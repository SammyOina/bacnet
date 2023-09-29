package bacnet

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/absmach/bacnet/internal"
	"github.com/absmach/bacnet/pkg/encoding"
)

type ReadPropertyRequest struct {
	ObjectIdentifier   *ObjectIdentifier
	PropertyIdentifier interface{}
	PropertyArrayIndex uint32
}

func (r ReadPropertyRequest) Encode() []byte {
	var ret []byte

	if r.ObjectIdentifier != nil {
		ret = append(ret, r.ObjectIdentifier.EncodeContext(0)...)
	}

	propID := r.PropertyIdentifier.(encoding.PropertyIdentifier)

	ret = append(ret, encoding.EncodeContextEnumerated(1, uint32(propID))...)

	if r.PropertyArrayIndex != 0 {
		ret = append(ret, encoding.EncodeContextUnsigned(2, r.PropertyArrayIndex)...)
	}

	return ret
}

func (r *ReadPropertyRequest) Decode(buffer []byte, offset, apduLen int) (int, error) {
	leng := 0

	// objectIdentifier
	if encoding.IsContextTag(buffer, offset+leng, 0) {
		leng1, _, lenValue, err := encoding.DecodeTagNumberAndValue(buffer, offset+leng)
		if err != nil {
			return -1, err
		}
		leng += leng1
		r.ObjectIdentifier = &ObjectIdentifier{}
		leng1, err = r.ObjectIdentifier.Decode(buffer, offset+leng, int(lenValue))
		if err != nil {
			return -1, err
		}
		leng += leng1
	} else {
		return -1, errInvalidTagNumber
	}

	// propertyIdentifier
	if encoding.IsContextTag(buffer, offset+leng, 1) {
		leng1, _, lenValue, err := encoding.DecodeTagNumberAndValue(buffer, offset+leng)
		if err != nil {
			return -1, err
		}
		leng += leng1
		propID := encoding.PropertyList
		leng1, r.PropertyIdentifier, err = encoding.DecodeEnumerated(buffer, offset+leng, lenValue, nil, &propID)
		if err != nil {
			return -1, err
		}
		leng += leng1
	} else {
		return -1, errInvalidTagNumber
	}

	// propertyArrayIndex (optional)
	if leng < apduLen && encoding.IsContextTag(buffer, offset+leng, 2) {
		leng1, _, lenValue, err := encoding.DecodeTagNumberAndValue(buffer, offset+leng)
		if err != nil {
			return -1, err
		}
		leng += leng1
		leng1, r.PropertyArrayIndex, err = encoding.DecodeUnsigned(buffer, offset+leng, int(lenValue))
		if err != nil {
			return -1, err
		}

		leng += leng1
	}

	return leng, nil
}

type ReadPropertyACK struct {
	ObjectIdentifier   ObjectIdentifier
	PropertyIdentifier interface{}
	PropertyArrayIndex uint32
	PropertyValue      []Value
}

func (r *ReadPropertyACK) Decode(buffer []byte, offset, apduLen int) (int, error) {
	leng := 0

	// 0 object_identifier
	if encoding.IsContextTag(buffer, offset+leng, 0) {
		r.ObjectIdentifier = ObjectIdentifier{}
		leng1, err := r.ObjectIdentifier.DecodeContext(buffer, offset+leng, apduLen-leng, 0)
		if err != nil {
			return -1, err
		}
		leng += leng1
	} else {
		return -1, errors.New("decoding error for object_identifier")
	}

	// 2 propertyidentifier
	if encoding.IsContextTag(buffer, offset+leng, 1) {
		leng1, _, lenValue, err := encoding.DecodeTagNumberAndValue(buffer, offset+leng)
		if err != nil {
			return -1, err
		}
		leng += leng1
		PropertyID := encoding.PropertyList
		leng1, propID, err := encoding.DecodeEnumerated(buffer, offset+leng, lenValue, nil, &PropertyID)
		if err != nil {
			return -1, err
		}
		r.PropertyIdentifier = propID
		leng += leng1
	} else {
		return -1, errors.New("decoding error for property_identifier")
	}

	// 2 property_array_index
	if encoding.IsContextTag(buffer, offset+leng, 2) {
		leng1, _, lenValue, err := encoding.DecodeTagNumberAndValue(buffer, offset+leng)
		if err != nil {
			return -1, err
		}
		leng += leng1
		leng1, r.PropertyArrayIndex, err = encoding.DecodeUnsigned(buffer, offset+leng, int(lenValue))
		if err != nil {
			return -1, err
		}
		leng += leng1
	}

	// tag 3 property-value
	if encoding.IsOpeningTagNumber(buffer, offset+leng, 3) {
		leng++
		r.PropertyValue = make([]Value, 0)
		for !encoding.IsClosingTagNumber(buffer, offset+leng, 3) && leng < apduLen {
			bValue := Value{}
			propId := r.PropertyIdentifier.(encoding.PropertyIdentifier)
			leng1, err := bValue.Decode(buffer, offset+leng, apduLen-leng, &r.ObjectIdentifier.Type, &propId)
			if err != nil {
				return -1, err
			}
			leng += leng1
			r.PropertyValue = append(r.PropertyValue, bValue)
		}
		if encoding.IsClosingTagNumber(buffer, offset+leng, 3) {
			leng++
		} else {
			return -1, errors.New("decoding error for property_value")
		}
	} else {
		return -1, errors.New("decoding error for property_value")
	}

	return leng, nil
}

type Value struct {
	Tag   *ApplicationTags
	Value interface{}
}

func (bv *Value) Decode(buffer []byte, offset, apduLen int, objType *encoding.ObjectType, propID *encoding.PropertyIdentifier) (int, error) {
	length := 0
	var err error

	if !encoding.IsContextSpecific(buffer[offset]) {
		tagLen, tagNumber, lenValueType, err := encoding.DecodeTagNumberAndValue(buffer, offset)
		if err != nil {
			return -1, err
		}
		if tagLen > 0 {
			ttag := ApplicationTags(tagNumber)
			bv.Tag = &ttag
			length += tagLen

			decodeLen := 0

			switch *bv.Tag {
			case Null:
				bv.Value = nil
				decodeLen = 0
			case Boolean:
				if lenValueType > 0 {
					bv.Value = true
				} else {
					bv.Value = false
				}
			case UnsignedInt:
				if *propID == encoding.RoutingTable {
					bv.Tag = nil
					bv.Value = &RouterEntry{}
					length--
					decodeLen, err = bv.Value.(*RouterEntry).Decode(buffer, offset+length, apduLen)
					if err != nil {
						return -1, err
					}
				} else if *propID == encoding.ActiveVtSessions {
					bv.Tag = nil
					bv.Value = &VTSession{}
					length--
					decodeLen = bv.Value.(*VTSession).Decode(buffer, offset+length, apduLen)
				} else if *propID == encoding.ThreatLevel || *propID == encoding.ThreatAuthority {
					bv.Tag = nil
					bv.Value = &AccessThreatLevel{}
					length--
					decodeLen = bv.Value.(*AccessThreatLevel).Decode(buffer, offset+length, apduLen)
				} else {
					var uintVal uint32
					decodeLen, uintVal, err = encoding.DecodeUnsigned(buffer, offset+length, int(lenValueType))
					if err != nil {
						return -1, err
					}
					bv.Value = uintVal
				}
			case SignedInt:
				var intValue int
				decodeLen, intValue = encoding.DecodeSigned(buffer, offset+length, int(lenValueType))
				bv.Value = intValue
			case Real:
				var floatValue float32
				decodeLen, floatValue = encoding.DecodeRealSafe(buffer, offset+length, int(lenValueType))
				bv.Value = floatValue
			case Double:
				var doubleValue float64
				decodeLen, doubleValue = encoding.DecodeDoubleSafe(buffer, offset+length, int(lenValueType))
				bv.Value = doubleValue
			case OctetString:
				var octetValue []byte
				decodeLen, octetValue = encoding.DecodeOctetString(buffer, offset+length, int(lenValueType))
				bv.Value = octetValue
			case CharacterString:
				var stringValue string
				decodeLen, stringValue = encoding.DecodeCharacterString(buffer, offset+length, apduLen, int(lenValueType))
				bv.Value = stringValue
			case ApplicationTagsBitString:
				switch *propID {
				case encoding.RecipientList:
					bv.Tag = nil
					bv.Value = &Destination{}
					length--
					decodeLen, err = bv.Value.(*Destination).Decode(buffer, offset+length, apduLen)
					if err != nil {
						return -1, err
					}
				case encoding.StatusFlags:
					bv.Tag = nil
					bitValue := &StatusFlags{}
					decodeLen = bitValue.Decode(buffer, offset, int(lenValueType))
					bv.Value = bitValue
				case encoding.EventEnable, encoding.AckedTransitions:
					bv.Tag = nil
					bitValue := &EventTransitionBits{}
					decodeLen = bitValue.Decode(buffer, offset, int(lenValueType))
					bv.Value = bitValue
				case encoding.LimitEnable:
					bv.Tag = nil
					bitValue := &LimitEnable{}
					decodeLen = bitValue.Decode(buffer, offset, int(lenValueType))
					bv.Value = bitValue
				case encoding.ProtocolObjectTypesSupported:
					bv.Tag = nil
					bitValue := &ObjectTypesSupported{}
					decodeLen = bitValue.Decode(buffer, offset, int(lenValueType))
					bv.Value = bitValue
				case encoding.ProtocolServicesSupported:
					bv.Tag = nil
					bitValue := &ServicesSupported{}
					decodeLen = bitValue.Decode(buffer, offset, int(lenValueType))
					bv.Value = bitValue
				default:
					bitValue := &BitString{}
					decodeLen = bitValue.Decode(buffer, offset, int(lenValueType))
					bv.Value = bitValue
				}
			case Enumerated:
				decodeLen, bv.Value, err = encoding.DecodeEnumerated(buffer, offset+length, lenValueType, objType, propID)
				if err != nil {
					return -1, err
				}
			case Date:
				switch *propID {
				case encoding.EffectivePeriod:
					bv.Tag = nil
					bv.Value = &DateRange{}
					length--
					decodeLen, err = bv.Value.(*DateRange).Decode(buffer, offset+length, apduLen)
					if err != nil {
						return -1, err
					}
				case encoding.MinimumValueTimestamp,
					encoding.MaximumValueTimestamp,
					encoding.ChangeOfStateTime,
					encoding.TimeOfStateCountReset,
					encoding.TimeOfActiveTimeReset,
					encoding.ModificationDate,
					encoding.UpdateTime,
					encoding.CountChangeTime,
					encoding.StartTime,
					encoding.StopTime,
					encoding.LastCredentialAddedTime,
					encoding.LastCredentialRemovedTime,
					encoding.ActivationTime,
					encoding.ExpiryTime,
					encoding.LastUseTime,
					encoding.TimeOfStrikeCountReset,
					encoding.ValueChangeTime:
					bv.Tag = nil
					bv.Value = &DateTime{}
					length--
					decodeLen = bv.Value.(*DateTime).Decode(buffer, offset+length)
				default:
					decodeLen, bv.Value = encoding.DecodeDateSafe(buffer, offset+length, int(lenValueType))
				}
				if (*objType == encoding.DateTimeValue || *objType == encoding.TimePatternValue) && (*propID == encoding.PresentValue || *propID == encoding.RelinquishDefault) {
					decodeLen, bv.Value = encoding.DecodeDateSafe(buffer, offset+length, int(lenValueType))
				}
			case Time:
				decodeLen, bv.Value = encoding.DecodeBACnetTimeSafe(buffer, offset+length, int(lenValueType))
			case BACnetObjectIdentifier:
				if *propID == encoding.LastKeyServer ||
					*propID == encoding.ManualSlaveAddressBinding ||
					*propID == encoding.SlaveAddressBinding ||
					*propID == encoding.DeviceAddressBinding {
					bv.Tag = nil
					bv.Value = &AddressBinding{}
					length--
					decodeLen, err = bv.Value.(*AddressBinding).Decode(buffer, offset+length, apduLen)
					if err != nil {
						return -1, err
					}
				} else {
					var objectType encoding.ObjectType
					var instance uint32
					decodeLen, objectType, instance, err = encoding.DecodeObjectIDSafe(buffer, offset+length, lenValueType)
					if err != nil {
						return -1, err
					}
					bv.Value = ObjectIdentifier{Type: objectType, Instance: ObjectInstance(instance)}
				}
			default:
				log.Println("Unhandled tag:", bv.Tag)
				length = apduLen
			}

			if decodeLen < 0 {
				return -1, fmt.Errorf("no tags decoded")
			}
			length += decodeLen
		}
	} else {
		switch *propID {
		case encoding.BacnetIpGlobalAddress, encoding.FdBbmdAddress:
			bv.Value = &HostNPort{}
			length1, err := bv.Value.(*HostNPort).Decode(buffer, offset+length, apduLen-length)
			if err != nil {
				return -1, err
			}
			length += length1
		case encoding.UtcTimeSynchronizationRecipients,
			encoding.RestartNotificationRecipients,
			encoding.TimeSynchronizationRecipients,
			encoding.CovuRecipients:
			bv.Value = &Recipient{}
			length1, err := bv.Value.(*Recipient).Decode(buffer, offset+length, apduLen-length)
			if err != nil {
				return -1, err
			}
			length += length1
		case encoding.KeySets:
			bv.Value = &SecurityKeySet{}
			length1, err := bv.Value.(*SecurityKeySet).Decode(buffer, offset+length, apduLen-length)
			if err != nil {
				return -1, err
			}
			length += length1
		case encoding.EventTimeStamps,
			encoding.LastCommandTime,
			encoding.CommandTimeArray,
			encoding.LastRestoreTime,
			encoding.TimeOfDeviceRestart,
			encoding.AccessEventTime,
			encoding.UpdateTime:
			bv.Value = &TimeStamp{}
			length1, err := bv.Value.(*TimeStamp).Decode(buffer, offset+length, apduLen-length)
			if err != nil {
				return -1, err
			}
			length += length1
		case encoding.ListOfGroupMembers:
			bv.Value = &ReadAccessSpecification{}
			length, err = bv.Value.(*ReadAccessSpecification).Decode(buffer, offset+length, apduLen-length)
			if err != nil {
				return -1, err
			}
		case encoding.ListOfObjectPropertyReferences:
			bv.Value = &DeviceObjectPropertyReference{}
			length, err = bv.Value.(*DeviceObjectPropertyReference).Decode(buffer, offset+length, apduLen-length)
			if err != nil {
				return -1, err
			}
		case encoding.MemberOf,
			encoding.ZoneMembers,
			encoding.DoorMembers,
			encoding.SubordinateList,
			encoding.Represents,
			encoding.AccessEventCredential,
			encoding.AccessDoors,
			encoding.ZoneTo,
			encoding.ZoneFrom,
			encoding.CredentialsInZone,
			encoding.LastCredentialAdded,
			encoding.LastCredentialRemoved,
			encoding.EntryPoints,
			encoding.ExitPoints,
			encoding.Members,
			encoding.Credentials,
			encoding.Accompaniment,
			encoding.BelongsTo,
			encoding.LastAccessPoint,
			encoding.EnergyMeterRef:
			bv.Value = &DeviceObjectReference{}
			length1, err := bv.Value.(*DeviceObjectReference).Decode(buffer, offset+length, apduLen-length)
			if err != nil {
				return -1, err
			}
			length += length1
		case encoding.EventAlgorithmInhibitRef,
			encoding.InputReference,
			encoding.ManipulatedVariableReference,
			encoding.ControlledVariableReference:
			bv.Value = &ObjectPropertyReference{}
			length1, err := bv.Value.(*ObjectPropertyReference).Decode(buffer, offset+length, apduLen-length)
			if err != nil {
				return -1, err
			}
			length += length1
		case encoding.LoggingRecord:
			bv.Value = &AccumulatorRecord{}
			length, err = bv.Value.(*AccumulatorRecord).Decode(buffer, offset+length, apduLen-length)
			if err != nil {
				return -1, err
			}
		case encoding.PropertyIdentifierAction:
			bv.Value = &ActionList{}
			length1, err := bv.Value.(*ActionList).Decode(buffer, offset+length, apduLen-length)
			if err != nil {
				return -1, err
			}
			length += length1
		case encoding.Scale:
			bv.Value = &Scale{}
			length1, err := bv.Value.(*Scale).Decode(buffer, offset+length, apduLen-length)
			if err != nil {
				return -1, err
			}
			length += length1
		case encoding.LightingCommand:
			bv.Value = &LightingCommand{}
			length1, err := bv.Value.(*LightingCommand).Decode(buffer, offset+length, apduLen-length)
			if err != nil {
				return -1, err
			}
			length += length1
		case encoding.Prescale:
			bv.Value = &Prescale{}
			length1, err := bv.Value.(*Prescale).Decode(buffer, offset+length, apduLen-length)
			if err != nil {
				return -1, err
			}
			length += length1
		case encoding.RequestedShedLevel,
			encoding.ExpectedShedLevel,
			encoding.ActualShedLevel:
			bv.Value = &ShedLevel{}
			length1, err := bv.Value.(*ShedLevel).Decode(buffer, offset+length, apduLen-length)
			if err != nil {
				return -1, err
			}
			length += length1
		case encoding.LogBuffer:
			switch *objType {
			case encoding.TrendLog:
				bv.Value = &LogRecord{}
				length1, err := bv.Value.(*LogRecord).Decode(buffer, offset+length, apduLen-length, nil, nil)
				if err != nil {
					return -1, err
				}
				length += length1
			case encoding.EventLog:
				bv.Value = &EventLogRecord{}
				length1, err := bv.Value.(*EventLogRecord).Decode(buffer, offset+length, apduLen-length)
				if err != nil {
					return -1, err
				}
				length += length1
			}
		case encoding.DateList:
			bv.Value = &CalendarEntry{}
			length1, err := bv.Value.(*CalendarEntry).Decode(buffer, offset+length, apduLen-length)
			if err != nil {
				return -1, err
			}
			length += length1
		case encoding.PresentValue:
			switch *objType {
			case encoding.Group:
				bv.Value = &ReadAccessResult{}
				length1, err := bv.Value.(*ReadAccessResult).Decode(buffer, offset+length, apduLen-length)
				if err != nil {
					return -1, err
				}
				length += length1
			case encoding.Channel:
				bv.Value = &ChannelValue{}
				length += bv.Value.(*ChannelValue).Decode(buffer, offset+length, apduLen-length)
			case encoding.GlobalGroup:
				bv.Value = &PropertyAccessResult{}
				length += bv.Value.(*PropertyAccessResult).Decode(buffer, offset+length, apduLen-length)
			case encoding.CredentialDataInput:
				bv.Value = &AuthenticationFactor{}
				length += bv.Value.(*AuthenticationFactor).Decode(buffer, offset+length, apduLen-length)
			}
		case encoding.NegativeAccessRules,
			encoding.PositiveAccessRules:
			bv.Value = &AccessRule{}
			length += bv.Value.(*AccessRule).Decode(buffer, offset+length, apduLen-length)
		case encoding.Tags:
			bv.Value = &NameValue{}
			length1, err := bv.Value.(*NameValue).Decode(buffer, offset+length, apduLen-length)
			if err != nil {
				return -1, err
			}
			length += length1
		case encoding.SubordinateTags:
			bv.Value = &NameValueCollection{}
			length1, err := bv.Value.(*NameValueCollection).Decode(buffer, offset+length, apduLen-length)
			if err != nil {
				return -1, err
			}
			length += length1
		case encoding.NetworkAccessSecurityPolicies:
			bv.Value = &NetworkSecurityPolicy{}
			length1, err := bv.Value.(*NetworkSecurityPolicy).Decode(buffer, offset+length, apduLen-length)
			if err != nil {
				return -1, err
			}
			length += length1
		case encoding.PortFilter:
			bv.Value = &PortPermission{}
			length1, err := bv.Value.(*PortPermission).Decode(buffer, offset+length, apduLen-length)
			if err != nil {
				return -1, err
			}
			length += length1
		case encoding.PriorityArray:
			bv.Value = &PriorityArray{}
			length += bv.Value.(*PriorityArray).Decode(buffer, offset+length, apduLen-length)
		case encoding.ProcessIdentifierFilter:
			bv.Value = &ProcessIdSelection{}
			length += bv.Value.(*ProcessIdSelection).Decode(buffer, offset+length, apduLen-length)
		case encoding.SetpointReference:
			bv.Value = &SetpointReference{}
			length += bv.Value.(*SetpointReference).Decode(buffer, offset+length, apduLen-length)
		case encoding.ExceptionSchedule:
			bv.Value = &SpecialEvent{}
			length += bv.Value.(*SpecialEvent).Decode(buffer, offset+length, apduLen-length)
		case encoding.StateChangeValues:
			bv.Value = &TimerStateChangeValue{}
			length += bv.Value.(*TimerStateChangeValue).Decode(buffer, offset+length, apduLen-length)
		case encoding.ValueSource, encoding.ValueSourceArray:
			bv.Value = &ValueSource{}
			length += bv.Value.(*ValueSource).Decode(buffer, offset+length, apduLen-length)
		case encoding.VirtualMacAddressTable:
			bv.Value = &VMACEntry{}
			length += bv.Value.(*VMACEntry).Decode(buffer, offset+length, apduLen-length)
		case encoding.AssignedAccessRights:
			bv.Value = &AssignedAccessRights{}
			length += bv.Value.(*AssignedAccessRights).Decode(buffer, offset+length, apduLen-length)
		case encoding.AssignedLandingCalls:
			bv.Value = &AssignedLandingCalls{}
			length += bv.Value.(*AssignedLandingCalls).Decode(buffer, offset+length, apduLen-length)
		case encoding.AccessEventAuthenticationFactor:
			bv.Value = &AuthenticationFactor{}
			length += bv.Value.(*AuthenticationFactor).Decode(buffer, offset+length, apduLen-length)
		case encoding.SupportedFormats:
			bv.Value = &AuthenticationFactorFormat{}
			length += bv.Value.(*AuthenticationFactorFormat).Decode(buffer, offset+length, apduLen-length)
		case encoding.AuthenticationPolicyList:
			bv.Value = &AuthenticationPolicy{}
			length += bv.Value.(*AuthenticationPolicy).Decode(buffer, offset+length, apduLen-length)
		case encoding.ActiveCovSubscriptions:
			bv.Value = &COVSubscription{}
			length += bv.Value.(*COVSubscription).Decode(buffer, offset+length, apduLen-length)
		case encoding.AuthenticationFactors:
			bv.Value = &CredentialAuthenticationFactor{}
			length += bv.Value.(*CredentialAuthenticationFactor).Decode(buffer, offset+length, apduLen-length)
		case encoding.WeeklySchedule:
			bv.Value = &DailySchedule{}
			length += bv.Value.(*DailySchedule).Decode(buffer, offset+length, apduLen-length)
		case encoding.SubscribedRecipients:
			bv.Value = &EventNotificationSubscription{}
			length += bv.Value.(*EventNotificationSubscription).Decode(buffer, offset+length, apduLen-length)
		case encoding.EventParameters:
			bv.Value = &EventParameter{}
			length += bv.Value.(*EventParameter).Decode(buffer, offset+length, apduLen-length)
		case encoding.FaultParameters:
			bv.Value = &FaultParameter{}
			length += bv.Value.(*FaultParameter).Decode(buffer, offset+length, apduLen-length)
		default:
			bv.Value = nil
		}
	}
	return length, nil
}

func (bv *Value) Encode() []byte {
	if bv.Tag == nil {
		return nil
	} else {
		switch *bv.Tag {
		case Boolean:
			return encoding.EncodeApplicationBoolean(bv.Value.(bool))
		case UnsignedInt:
			return encoding.EncodeApplicationUnsigned(bv.Value.(uint32))
		case SignedInt:
			return encoding.EncodeApplicationSigned(bv.Value.(int32))
		case Real:
			return encoding.EncodeApplicationReal(bv.Value.(float32))
		case Double:
			return encoding.EncodeApplicationDouble(bv.Value.(float64))
		case OctetString:
			return encoding.EncodeApplicationOctetString(bv.Value.([]byte), 0, len(bv.Value.([]byte)))
		case CharacterString:
			return encoding.EncodeApplicationCharacterString(bv.Value.(string))
		case ApplicationTagsBitString:
			return encoding.EncodeApplicationBitString(bv.Value)
		case Enumerated:
			return encoding.EncodeApplicationEnumerated(bv.Value.(uint32))
		case Date:
			return encoding.EncodeApplicationDate(bv.Value.(time.Time))
		case Time:
			return encoding.EncodeApplicationTime(bv.Value.(time.Time))
		case BACnetObjectIdentifier:
			return bv.Value.(*ObjectIdentifier).EncodeApp()
		default:
			switch bv.Value.(type) {
			case int:
				return encoding.EncodeApplicationEnumerated(uint32(bv.Value.(int)))
			}
			log.Printf("Unsupported BACnetApplicationTag: %v", bv.Tag)
			return nil
		}
	}
}

// BACnetRouterEntryStatus is an enumeration for the status of BACnetRouterEntry.
type RouterEntryStatus int

const (
	Available RouterEntryStatus = iota
	BACnetRouterEntryStatusBusy
	Disconnected
)

// BACnetRouterEntry represents a BACnet router entry.
type RouterEntry struct {
	NetworkNumber    uint32
	MACAddress       []byte
	Status           RouterEntryStatus
	PerformanceIndex uint32
}

// Decode decodes a RouterEntry from an encoded byte buffer.
func (entry *RouterEntry) Decode(buffer []byte, offset, apduLen int) (int, error) {
	var length int

	// network_number
	length1, tagNumber, lenValue, err := encoding.DecodeTagNumberAndValue(buffer, offset)
	if err != nil {
		return -1, err
	}
	if tagNumber != byte(encoding.UnsignedInt) {
		return -1, errors.New("Error decoding network_number")
	}
	length += length1
	length1, entry.NetworkNumber, err = encoding.DecodeUnsigned(buffer, offset+length, int(lenValue))
	if err != nil {
		return -1, err
	}
	length += length1

	// mac_address
	length1, tagNumber, lenValue, err = encoding.DecodeTagNumberAndValue(buffer, offset+length)
	if err != nil {
		return -1, err
	}
	if tagNumber != byte(encoding.OctetString) {
		return -1, errors.New("Error decoding mac_address")
	}
	length += length1
	length1, entry.MACAddress = encoding.DecodeOctetString(buffer, offset+length, int(lenValue))
	length += length1

	// status
	length1, tagNumber, lenValue, err = encoding.DecodeTagNumberAndValue(buffer, offset+length)
	if err != nil {
		return -1, err
	}
	if tagNumber != byte(encoding.Enumerated) {
		return -1, errors.New("Error decoding status")
	}
	length += length1
	length1, Val, err := encoding.DecodeUnsigned(buffer, offset+length, int(lenValue))
	if err != nil {
		return -1, err
	}
	length += length1
	entry.Status = RouterEntryStatus(Val)

	// performance_index (optional)
	if offset < apduLen {
		length1, tagNumber, lenValue, err = encoding.DecodeTagNumberAndValue(buffer, offset+length)
		if err != nil {
			return -1, err
		}
		if tagNumber != byte(encoding.UnsignedInt) {
			length += length1
			length1, entry.PerformanceIndex, err = encoding.DecodeUnsigned(buffer, offset+length, int(lenValue))
			if err != nil {
				return -1, err
			}
			length += length1
		}
	}

	return length, nil
}

type VTSession struct {
	LocalVTSessionID  int
	RemoteVTSessionID int
	RemoteVTAddress   Address
}

// Decode method for BACnetVTSession
func (b *VTSession) Decode(buffer []byte, offset, apduLen int) int {
	// Implement decode logic here
	return -1
}

// BACnetAccessThreatLevel struct definition
type AccessThreatLevel struct {
	Value int
}

// decode method for BACnetAccessThreatLevel
func (b *AccessThreatLevel) Decode(buffer []byte, offset, apduLen int) int {
	// Implement decode logic here
	return -1
}

type Destination struct {
	ValidDays                   *DaysOfWeek
	FromTime                    time.Time
	ToTime                      time.Time
	Recipient                   *Recipient
	ProcessIdentifier           uint32
	IssueConfirmedNotifications bool
	Transitions                 *EventTransitionBits
}

func (b *Destination) Decode(buffer []byte, offset int, apduLen int) (int, error) {
	leng := 0

	leng1, tagNumber, lenValue, err := encoding.DecodeTagNumberAndValue(buffer, offset+leng)
	if err != nil {
		return -1, err
	}
	if tagNumber != byte(encoding.BitString) {
		return -1, errInvalidTagNumber
	}
	leng += leng1
	b.ValidDays = &DaysOfWeek{}

	leng1 = b.ValidDays.Decode(buffer, offset+leng, int(lenValue))

	if leng1 < 0 {
		return -1, errInvalidMessageLength
	}
	leng += leng1

	leng1, tagNumber, lenValue, err = encoding.DecodeTagNumberAndValue(buffer, offset+leng)
	if err != nil {
		return -1, err
	}
	if tagNumber != byte(encoding.Time) {
		return -1, errInvalidTagNumber
	}
	leng += leng1
	leng1, b.FromTime = encoding.DecodeBACnetTimeSafe(buffer, offset+leng, int(lenValue))

	leng += leng1

	leng1, tagNumber, lenValue, err = encoding.DecodeTagNumberAndValue(buffer, offset+leng)
	if err != nil {
		return -1, err
	}
	if tagNumber != byte(encoding.Time) {
		return -1, errInvalidTagNumber
	}
	leng += leng1
	leng1, b.ToTime = encoding.DecodeBACnetTimeSafe(buffer, offset+leng, int(lenValue))

	leng += leng1

	b.Recipient = &Recipient{}
	leng1, err = b.Recipient.Decode(buffer, offset+leng, apduLen-leng)
	if err != nil {
		return -1, err
	}

	if leng1 < 0 {
		return -1, errInvalidMessageLength
	}
	leng += leng1

	leng1, tagNumber, lenValue, err = encoding.DecodeTagNumberAndValue(buffer, offset+leng)
	if err != nil {
		return -1, err
	}
	if tagNumber != byte(encoding.UnsignedInt) {
		return -1, errInvalidTagNumber
	}
	leng += leng1
	leng1, b.ProcessIdentifier, err = encoding.DecodeUnsigned(buffer, offset+leng, int(lenValue))
	if err != nil {
		return -1, err
	}

	leng += leng1

	leng1, tagNumber, lenValue, err = encoding.DecodeTagNumberAndValue(buffer, offset+leng)
	if err != nil {
		return -1, err
	}
	if tagNumber != byte(encoding.Boolean) {
		return -1, errInvalidTagNumber
	}
	leng += leng1
	if lenValue > 0 {
		b.IssueConfirmedNotifications = true
	} else {
		b.IssueConfirmedNotifications = false
	}

	leng1, tagNumber, lenValue, err = encoding.DecodeTagNumberAndValue(buffer, offset+leng)
	if err != nil {
		return -1, err
	}
	if tagNumber != byte(encoding.BitString) {
		return -1, errInvalidTagNumber
	}
	leng += leng1

	b.Transitions = &EventTransitionBits{}
	leng1 = b.Transitions.Decode(buffer, offset+leng, int(lenValue))
	if leng1 < 0 {
		return -1, errInvalidMessageLength
	}
	leng += leng1

	return leng, nil
}

type DaysOfWeek struct {
	unusedBits byte
	bitString  BitString
	monday     bool
	tuesday    bool
	wednesday  bool
	thursday   bool
	friday     bool
	saturday   bool
	sunday     bool
}

func NewBACnetDaysOfWeek() *DaysOfWeek {
	return &DaysOfWeek{
		unusedBits: 1,
		bitString:  *NewBACnetBitString(1, *internal.NewBitArray(8)),
	}
}

func (d *DaysOfWeek) Decode(buffer []byte, offset int, apduLen int) int {
	d.bitString = BitString{}
	return d.bitString.Decode(buffer, offset, apduLen)
}

func (d *DaysOfWeek) SetDay(day int, value bool) error {
	if day < 0 || day > 6 {
		return fmt.Errorf("Day index out of range")
	}
	d.bitString.Value.Set(day, value)
	return nil
}

func (d *DaysOfWeek) GetDay(day int) (bool, error) {
	if day < 0 || day > 6 {
		return false, fmt.Errorf("Day index out of range")
	}
	return d.bitString.Value.Get(day)
}

type BitString struct {
	UnusedBits byte
	Value      internal.BitArray
}

func NewBACnetBitString(unusedBits byte, value internal.BitArray) *BitString {
	return &BitString{
		UnusedBits: unusedBits,
		Value:      value,
	}
}

func (b *BitString) Decode(buffer []byte, offset, apduLen int) int {
	leng := 0
	if apduLen > 0 {
		b.UnusedBits = buffer[offset]
		leng += 1
		b.Value = *internal.NewBitArray((apduLen - 1) * 8)
		bit := 0
		for i := 1; i < apduLen; i++ {
			for i2 := 0; i2 < 8; i2++ {
				b.Value.Set(bit, buffer[offset+i]&(1<<(7-i2)) != 0)
				bit++
			}
		}
	}
	return apduLen
}

type Recipient struct {
	Value interface{}
}

func (br *Recipient) Decode(buffer []byte, offset, apduLen int) (int, error) {
	leng := 0
	leng1, tagNumber, lenValue, err := encoding.DecodeTagNumberAndValue(buffer, offset+leng)
	if err != nil {
		return -1, err
	}

	if tagNumber == 0 {
		// device_identifier
		leng += leng1
		br.Value = &ObjectIdentifier{}
		leng1, err = br.Value.(*ObjectIdentifier).Decode(buffer, offset+leng, int(lenValue))
		if err != nil {
			return -1, err
		}
		leng += leng1
	} else if tagNumber == 1 {
		// address
		br.Value = &Address{}
		leng1, err = br.Value.(*Address).Decode(buffer, offset+leng, int(lenValue))
		if err != nil {
			return -1, err
		}
		leng += leng1
	} else {
		return -1, errInvalidTagNumber
	}

	return leng, nil
}

// BACnetEventTransitionBits represents a BACnet event transition bits structure.
type EventTransitionBits struct {
	UnusedBits uint8
	BitString  *BitString
}

// NewBACnetEventTransitionBits creates a new BACnet event transition bits instance.
func NewBACnetEventTransitionBits() *EventTransitionBits {
	return &EventTransitionBits{
		UnusedBits: 5,
		BitString:  NewBACnetBitString(5, *internal.NewBitArray(5)),
	}
}

// Decode decodes the bit string from a buffer.
func (e *EventTransitionBits) Decode(buffer []byte, offset, apduLen int) int {
	bitString := NewBACnetBitString(0, internal.BitArray{})
	decodedLen := bitString.Decode(buffer, offset, apduLen)

	e.BitString = bitString
	return decodedLen
}

// ToOffNormal returns the value of ToOffNormal property.
func (e *EventTransitionBits) ToOffNormal() (bool, error) {
	return e.BitString.Value.Get(0)
}

// SetToOffNormal sets the value of ToOffNormal property.
func (e *EventTransitionBits) SetToOffNormal(a bool) {
	e.BitString.Value.Set(0, a)
}

// ToFault returns the value of ToFault property.
func (e *EventTransitionBits) ToFault() (bool, error) {
	return e.BitString.Value.Get(1)
}

// SetToFault sets the value of ToFault property.
func (e *EventTransitionBits) SetToFault(a bool) {
	e.BitString.Value.Set(1, a)
}

// ToNormal returns the value of ToNormal property.
func (e *EventTransitionBits) ToNormal() (bool, error) {
	return e.BitString.Value.Get(2)
}

// SetToNormal sets the value of ToNormal property.
func (e *EventTransitionBits) SetToNormal(a bool) {
	e.BitString.Value.Set(2, a)
}

// BACnetStatusFlags represents a BACnet status flags.
type StatusFlags struct {
	unusedbits   int
	bitstring    BitString
	inalarm      bool
	fault        bool
	overridden   bool
	outofservice bool
}

// NewBACnetStatusFlags creates a new BACnetStatusFlags instance.
func NewBACnetStatusFlags() *StatusFlags {
	return &StatusFlags{
		unusedbits:   4,
		bitstring:    *NewBACnetBitString(4, *internal.NewBitArrayFromByte(0x00)),
		inalarm:      false,
		fault:        false,
		overridden:   false,
		outofservice: false,
	}
}

// decode decodes BACnetStatusFlags from a buffer.
func (s *StatusFlags) Decode(buffer []byte, offset, apduLen int) int {
	s.bitstring = *NewBACnetBitString(byte(s.unusedbits), *internal.NewBitArrayFromByte(0x00))
	return s.bitstring.Decode(buffer, offset, apduLen)
}

// InAlarm returns the inalarm property.
func (s *StatusFlags) InAlarm() (bool, error) {
	return s.bitstring.Value.Get(0)
}

// SetInAlarm sets the inalarm property.
func (s *StatusFlags) SetInAlarm(a bool) {
	s.bitstring.Value.Set(0, a)
}

// Fault returns the fault property.
func (s *StatusFlags) Fault() (bool, error) {
	return s.bitstring.Value.Get(1)
}

// SetFault sets the fault property.
func (s *StatusFlags) SetFault(a bool) {
	s.bitstring.Value.Set(1, a)
}

// Overridden returns the overridden property.
func (s *StatusFlags) Overridden() (bool, error) {
	return s.bitstring.Value.Get(2)
}

// SetOverridden sets the overridden property.
func (s *StatusFlags) SetOverridden(a bool) {
	s.bitstring.Value.Set(2, a)
}

// OutOfService returns the outofservice property.
func (s *StatusFlags) OutOfService() (bool, error) {
	return s.bitstring.Value.Get(3)
}

// SetOutOfService sets the outofservice property.
func (s *StatusFlags) SetOutOfService(a bool) {
	s.bitstring.Value.Set(3, a)
}

type LimitEnable struct {
	unusedBits      uint8
	bitString       BitString
	lowLimitEnable  bool
	highLimitEnable bool
}

func NewBACnetLimitEnable() *LimitEnable {
	return &LimitEnable{
		unusedBits:      6,
		bitString:       *NewBACnetBitString(6, *internal.NewBitArrayFromByte(0x00)),
		lowLimitEnable:  false,
		highLimitEnable: false,
	}
}

func (b *LimitEnable) Decode(buffer []byte, offset, apduLen int) int {
	b.bitString = *NewBACnetBitString(0, *internal.NewBitArrayFromByte(0x00))
	return b.bitString.Decode(buffer, offset, apduLen)
}

func (b *LimitEnable) LowLimitEnable() (bool, error) {
	return b.bitString.Value.Get(0)
}

func (b *LimitEnable) SetLowLimitEnable(a bool) {
	b.bitString.Value.Set(0, a)
}

func (b *LimitEnable) HighLimitEnable() (bool, error) {
	return b.bitString.Value.Get(1)
}

func (b *LimitEnable) SetHighLimitEnable(a bool) {
	b.bitString.Value.Set(1, a)
}

type ObjectTypesSupported struct {
	unusedbits uint8
	bitstring  BitString
}

type ObjectTypesSupportedProperty int

const (
	AnalogInput ObjectTypesSupportedProperty = iota
	AanalofOutput
	AnalogValue
	BinaryInput
	BinaryOutput
	// TODO Add other properties here
)

func NewBACnetObjectTypesSupported() *ObjectTypesSupported {
	return &ObjectTypesSupported{
		unusedbits: 3,
		bitstring:  *NewBACnetBitString(3, *internal.NewBitArray(64)),
	}
}

func (b *ObjectTypesSupported) Set(property ObjectTypesSupportedProperty, value bool) {
	b.bitstring.Value.Set(int(property), value)
}

func (b *ObjectTypesSupported) Get(property ObjectTypesSupportedProperty) (bool, error) {
	return b.bitstring.Value.Get(int(property))
}

func (b *ObjectTypesSupported) Decode(buf []byte, offset, apduLen int) int {
	return b.bitstring.Decode(buf, offset, apduLen)
}

type ServicesSupported struct {
	unusedbits uint8
	bitstring  BitString
}

type ServicesSupportedProperty int

// TODO check index sequence
const (
	acknowledgeAlarm ServicesSupportedProperty = iota
	confirmedCOVNotification
	confirmedCOVNotificationMultiple
	confirmedEventNotification
	getAlarmSummary
	getEnrollmentSummary
	getEventInformation
	lifeSafetyOperation
	subscribeCOV
	subscribeCOVProperty
	subscribeCOVPropertyMultiple
	atomicReadFile
	atomicWriteFile
	addListElement
	removeListElement
	createObject
	deleteObject
	readProperty
	readPropertyMultiple
	readRange
	writeGroup
	writeProperty
	writePropertyMultiple
	deviceCommunicationControl
	confirmedPrivateTransfer
	confirmedTextMessage
	reinitializeDevice
	vtOpen
	vtClose
	vtData
	whoAmI
	youAre
	iAm
	iHave
	unconfirmedCOVNotification
	unconfirmedCOVNotificationMultiple
	unconfirmedEventNotification
	unconfirmedPrivateTransfer
	unconfirmedTextMessage
	timeSynchronization
	utcTimeSynchronization
	whoHas
	whoIs
)

func NewBACnetServicesSupported() *ServicesSupported {
	return &ServicesSupported{
		unusedbits: 7,
		bitstring:  *NewBACnetBitString(7, *internal.NewBitArrayFromByte(0x00000000000000)),
	}
}

func (b *ServicesSupported) Set(property ServicesSupportedProperty, value bool) {
	b.bitstring.Value.Set(int(property), value)
}

func (b *ServicesSupported) Get(property ServicesSupportedProperty) (bool, error) {
	return b.bitstring.Value.Get(int(property))
}

func (b *ServicesSupported) Decode(buf []byte, offset, apduLen int) int {
	return b.bitstring.Decode(buf, offset, apduLen)
}

// BACnetDateRange is a struct representing a date range in BACnet.
type DateRange struct {
	StartDate time.Time
	EndDate   time.Time
}

// Decode decodes a BACnetDateRange from the given buffer, offset, and length.
func (dr *DateRange) Decode(buffer []byte, offset, apduLen int) (int, error) {
	var leng int

	leng1, tagNumber, lenValue, err := encoding.DecodeTagNumberAndValue(buffer, offset+leng)
	if err != nil {
		return -1, err
	}
	if tagNumber == byte(Date) {
		leng += leng1
		leng1, startDate := encoding.DecodeDateSafe(buffer, offset+leng, int(lenValue))
		dr.StartDate = startDate
		leng += leng1
	} else {
		return -1, fmt.Errorf("Unexpected tag number: %v", tagNumber)
	}

	leng1, tagNumber, lenValue, err = encoding.DecodeTagNumberAndValue(buffer, offset+leng)
	if err != nil {
		return -1, err
	}
	if tagNumber == byte(Date) {
		leng += leng1
		leng1, endDate := encoding.DecodeDateSafe(buffer, offset+leng, int(lenValue))

		dr.EndDate = endDate
		leng += leng1
	} else {
		return -1, fmt.Errorf("Unexpected tag number: %v", tagNumber)
	}

	return leng, nil
}

type AddressBinding struct {
	DeviceIdentifier ObjectIdentifier
	DeviceAddress    Address
}

func (binding *AddressBinding) Decode(buffer []byte, offset int, apduLen int) (int, error) {
	length := 0
	length1, tagNumber, lenValue, err := encoding.DecodeTagNumberAndValue(buffer, offset+length)
	if err != nil {
		return -1, err
	}

	// device_identifier
	if tagNumber == byte(BACnetObjectIdentifier) {
		length += length1
		binding.DeviceIdentifier = ObjectIdentifier{}
		leng1, err := binding.DeviceIdentifier.Decode(buffer, offset+length, int(lenValue))
		if err != nil {
			return -1, err
		}
		length += leng1
	} else {
		return -1, errInvalidTagNumber
	}

	length1, tagNumber, lenValue, err = encoding.DecodeTagNumberAndValue(buffer, offset+length)
	if err != nil {
		return -1, err
	}

	if tagNumber == byte(UnsignedInt) {
		binding.DeviceAddress = Address{}
		leng1, err := binding.DeviceAddress.Decode(buffer, offset+length, int(lenValue))
		if err != nil {
			return -1, err
		}
		length += leng1
	} else {
		return -1, errInvalidTagNumber
	}

	return length, nil
}

type HostNPort struct {
	Host *HostAddress
	Port uint32
}

func (b *HostNPort) Decode(buffer []byte, offset, apduLen int) (int, error) {
	leng := 0

	if !encoding.IsOpeningTagNumber(buffer, offset+leng, 0) {
		return -1, errors.New("Invalid opening tag")
	}
	leng++
	b.Host = &HostAddress{}
	hostLen, err := b.Host.Decode(buffer, offset+leng, apduLen-leng)
	if err != nil {
		return -1, err
	}
	leng += hostLen
	leng++

	if encoding.IsContextTag(buffer, offset+leng, 1) {
		leng1, tagNumber, lenValue, err := encoding.DecodeTagNumberAndValue(buffer, offset+leng)
		if err != nil {
			return -1, err
		}
		leng += leng1
		if tagNumber == 1 {
			leng1, b.Port, err = encoding.DecodeUnsigned(buffer, offset+leng, int(lenValue))
			if err != nil {
				return -1, err
			}
			leng += leng1
		} else {
			return -1, errors.New("Invalid tag number")
		}
	} else {
		return -1, errors.New("Invalid context tag")
	}

	return leng, nil
}

type HostAddress struct {
	Value interface{}
}

func (b *HostAddress) Decode(buffer []byte, offset, apduLen int) (int, error) {
	leng := 0
	leng1, tagNumber, lenValue, err := encoding.DecodeTagNumberAndValue(buffer, offset+leng)
	if err != nil {
		return -1, err
	}

	switch tagNumber {
	case byte(Null):
		leng += leng1
		b.Value = nil
	case byte(OctetString):
		leng += leng1
		leng1, octetString := encoding.DecodeOctetString(buffer, offset+leng, int(lenValue))
		b.Value = octetString
		leng += leng1
	case byte(CharacterString):
		leng += leng1
		leng1, characterString := encoding.DecodeCharacterString(buffer, offset+leng, apduLen-leng, int(lenValue))
		b.Value = characterString
		leng += leng1
	default:
		return -1, errors.New("Invalid tag number")
	}

	return leng, nil
}

type SecurityKeySet struct {
	KeyRevision    uint32
	ActivationTime *DateTime
	ExpirationTime *DateTime
	KeyIDs         []*KeyIdentifier
}

func (b *SecurityKeySet) Decode(buffer []byte, offset, apduLen int) (int, error) {
	leng := 0

	// key_revision
	if encoding.IsContextTag(buffer, offset+leng, 0) {
		leng1, tagNumber, lenValue, err := encoding.DecodeTagNumberAndValue(buffer, offset+leng)
		if err != nil {
			return -1, err
		}
		leng += leng1
		if tagNumber == 0 {
			leng1, b.KeyRevision, err = encoding.DecodeUnsigned(buffer, offset+leng, int(lenValue))
			if err != nil {
				return -1, err
			}
			leng += leng1
		} else {
			return -1, errors.New("Invalid tag number")
		}
	} else {
		return -1, errors.New("Invalid context tag")
	}

	// activation_time
	if encoding.IsOpeningTagNumber(buffer, offset+leng, 1) {
		leng++
		b.ActivationTime = &DateTime{}
		leng1 := b.ActivationTime.Decode(buffer, offset+leng)
		leng += leng1
	} else {
		return -1, errors.New("Invalid opening tag for activation_time")
	}

	// expiration_time
	if encoding.IsOpeningTagNumber(buffer, offset+leng, 2) {
		leng++
		b.ExpirationTime = &DateTime{}
		leng1 := b.ExpirationTime.Decode(buffer, offset+leng)
		leng += leng1
	} else {
		return -1, errors.New("Invalid opening tag for expiration_time")
	}

	b.KeyIDs = make([]*KeyIdentifier, 0)
	if encoding.IsOpeningTagNumber(buffer, offset+leng, 3) && leng < apduLen {
		leng++
		for !encoding.IsClosingTagNumber(buffer, offset+leng, 3) {
			bValue := &KeyIdentifier{}
			leng1, err := bValue.Decode(buffer, offset+leng, apduLen-leng)
			if err != nil {
				return -1, err
			}
			b.KeyIDs = append(b.KeyIDs, bValue)
			leng += leng1
		}
		leng++
	} else {
		return -1, errors.New("Invalid opening tag for key_ids or unexpected end of data")
	}

	return leng, nil
}

type KeyIdentifier struct {
	Algorithm uint32
	KeyID     uint32
}

func (b *KeyIdentifier) Decode(buffer []byte, offset, apduLen int) (int, error) {
	leng := 0

	// algorithm
	if encoding.IsContextTag(buffer, offset+leng, 0) {
		leng1, tagNumber, lenValue, err := encoding.DecodeTagNumberAndValue(buffer, offset+leng)
		if err != nil {
			return -1, err
		}
		leng += leng1
		if tagNumber == 0 {
			leng1, b.Algorithm, err = encoding.DecodeUnsigned(buffer, offset+leng, int(lenValue))
			if err != nil {
				return -1, err
			}
			leng += leng1
		} else {
			return -1, errors.New("Invalid tag number for algorithm")
		}
	} else {
		return -1, errors.New("Invalid context tag for algorithm")
	}

	// key_id
	if encoding.IsContextTag(buffer, offset+leng, 1) {
		leng1, tagNumber, lenValue, err := encoding.DecodeTagNumberAndValue(buffer, offset+leng)
		if err != nil {
			return -1, err
		}
		leng += leng1
		if tagNumber == 1 {
			leng1, b.KeyID, err = encoding.DecodeUnsigned(buffer, offset+leng, int(lenValue))
			if err != nil {
				return -1, err
			}
			leng += leng1
		} else {
			return -1, errors.New("Invalid tag number for key_id")
		}
	} else {
		return -1, errors.New("Invalid context tag for key_id")
	}

	return leng, nil
}

type TimeStamp struct {
	Value interface{}
}

func (b *TimeStamp) Decode(buffer []byte, offset, apduLen int) (int, error) {
	leng := 0
	if encoding.IsContextTag(buffer, offset+leng, 2) {
		// BACnetDateTime
		leng1, tagNumber, _, err := encoding.DecodeTagNumberAndValue(buffer, offset+leng)
		if err != nil {
			return -1, err
		}
		leng += leng1
		if tagNumber == 2 {
			b.Value = &DateTime{}
			leng1 := b.Value.(*DateTime).Decode(buffer, offset+leng)
			leng += leng1
		} else {
			return -1, errors.New("Invalid tag number for BACnetDateTime")
		}
	} else if encoding.IsContextTag(buffer, offset+leng, 1) {
		// sequence number
		leng1, tagNumber, lenValue, err := encoding.DecodeTagNumberAndValue(buffer, offset+leng)
		if err != nil {
			return -1, err
		}
		leng += leng1
		if tagNumber == 1 {
			leng1, seqNum, err := encoding.DecodeUnsigned(buffer, offset+leng, int(lenValue))
			if err != nil {
				return -1, err
			}
			b.Value = seqNum
			leng += leng1
		} else {
			return -1, errors.New("Invalid tag number for sequence number")
		}
	} else if encoding.IsContextTag(buffer, offset+leng, 0) {
		// time
		leng1, tagNumber, lenValue, err := encoding.DecodeTagNumberAndValue(buffer, offset+leng)
		if err != nil {
			return -1, err
		}
		leng += leng1
		if tagNumber == 0 {
			leng1, bacnetTime := encoding.DecodeBACnetTimeSafe(buffer, offset+leng, int(lenValue))
			b.Value = bacnetTime
			leng += leng1
		} else {
			return -1, errors.New("Invalid tag number for time")
		}
	} else {
		return -1, errors.New("Invalid context tag")
	}

	return leng, nil
}

func (b *TimeStamp) Encode() []byte {
	// Implement the encoding logic as needed for your specific application.
	return nil
}

func (b *TimeStamp) EncodeContext(tagNumber encoding.BACnetApplicationTag) []byte {
	tmp := b.Encode()
	return append(encoding.EncodeTag(tagNumber, true, len(tmp)), tmp...)
}

// ReadAccessSpecification represents a BACnet Read Access Specification.
type ReadAccessSpecification struct {
	ObjectIdentifier         ObjectIdentifier
	ListOfPropertyReferences []PropertyReference
}

// Decode decodes the ReadAccessSpecification from the buffer.
func (ras *ReadAccessSpecification) Decode(buffer []byte, offset, apduLen int) (int, error) {
	leng := 0

	// ObjectIdentifier
	ras.ObjectIdentifier = ObjectIdentifier{}
	leng1, err := ras.ObjectIdentifier.DecodeContext(buffer, offset+leng, apduLen-leng, 0)
	if err != nil {
		return -1, err
	}
	leng += leng1

	// ListOfPropertyReferences
	if buffer[offset+leng] == 0x30 { // Check for opening tag (0x30)
		leng++

		ras.ListOfPropertyReferences = make([]PropertyReference, 0)

		for apduLen-leng > 1 && buffer[offset+leng] != 0x00 { // Check for closing tag (0x00)
			bValue := PropertyReference{}
			leng1, err := bValue.Decode(buffer, offset+leng, apduLen-leng)
			if err != nil {
				return -1, err
			}
			leng += leng1

			ras.ListOfPropertyReferences = append(ras.ListOfPropertyReferences, bValue)
		}
	} else {
		return -1, errors.New("Invalid opening tag for ListOfPropertyReferences")
	}

	leng++

	return leng, nil
}

// PropertyReference represents a BACnet property reference.
type PropertyReference struct {
	PropertyIdentifier interface{}
	PropertyArrayIndex uint32
}

// Decode decodes the PropertyReference from the buffer.
func (ref *PropertyReference) Decode(buffer []byte, offset, apduLen int) (int, error) {
	leng := 0

	// propertyIdentifier
	if encoding.IsContextTag(buffer, offset+leng, 0) {
		leng1, _, lenValue, err := encoding.DecodeTagNumberAndValue(buffer, offset+leng)
		if err != nil {
			return -1, err
		}
		leng += leng1
		propID := encoding.PropertyList
		leng1, ref.PropertyIdentifier, err = encoding.DecodeEnumerated(buffer, offset+leng, lenValue, nil, &propID)
		if err != nil {
			return -1, err
		}
		leng += leng1
	} else {
		return -1, errors.New("Missing context tag for PropertyIdentifier")
	}

	if leng < apduLen {
		if encoding.IsContextTag(buffer, offset+leng, 1) && !encoding.IsClosingTagNumber(buffer, offset+leng, 1) {
			leng1, _, lenValue, err := encoding.DecodeTagNumberAndValue(buffer, offset+leng)
			if err != nil {
				return -1, err
			}
			leng += leng1
			leng1, ref.PropertyArrayIndex, err = encoding.DecodeUnsigned(buffer, offset+leng, int(lenValue))
			if err != nil {
				return -1, err
			}
			leng += leng1
		}
	}

	return leng, nil
}

type DeviceObjectPropertyReference struct {
	ObjectIdentifier   ObjectIdentifier
	PropertyIdentifier interface{}
	PropertyArrayIndex uint32
	DeviceIdentifier   ObjectIdentifier
}

func (bdopr *DeviceObjectPropertyReference) Decode(buffer []byte, offset int, apduLen int) (int, error) {
	leng := 0

	// tag 0 objectidentifier
	bdopr.ObjectIdentifier = ObjectIdentifier{}
	leng1, err := bdopr.ObjectIdentifier.DecodeContext(buffer, offset+leng, apduLen-leng, 0)
	if err != nil {
		return -1, err
	}
	if leng1 < 0 {
		return -1, errors.New("failed to decode object identifier")
	}
	leng += leng1

	// tag 1 propertyidentifier
	if encoding.IsContextTag(buffer, offset+leng, 1) {
		leng1, _, lenValue, err := encoding.DecodeTagNumberAndValue(buffer, offset+leng)
		if err != nil {
			return -1, err
		}
		leng += leng1
		propID := encoding.PropertyList
		leng1, bdopr.PropertyIdentifier, err = encoding.DecodeEnumerated(buffer, offset+leng, lenValue, nil, &propID)
		if err != nil {
			return -1, err
		}
		leng += leng1
	} else {
		return -1, errors.New("Missing tag property Identifier")
	}

	if leng < apduLen {
		// tag 2 property-array-index optional
		if encoding.IsContextTag(buffer, offset+leng, 2) {
			leng1, _, lenValue, err := encoding.DecodeTagNumberAndValue(buffer, offset+leng)
			if err != nil {
				return -1, err
			}
			leng += leng1
			leng1, bdopr.PropertyArrayIndex, err = encoding.DecodeUnsigned(buffer, offset+leng, int(lenValue))
			if err != nil {
				return -1, err
			}
			leng += leng1
		}
	}

	if leng < apduLen {
		// tag 3 device-identifier optional
		bdopr.DeviceIdentifier = ObjectIdentifier{}
		leng1, err := bdopr.DeviceIdentifier.DecodeContext(buffer, offset+leng, apduLen-leng, 3)
		if err != nil {
			return -1, err
		}
		if leng1 < 0 {
			return -1, errors.New("failed to decode device identifier")
		}
		leng += leng1
	}

	return leng, nil
}

type DeviceObjectReference struct {
	DeviceIdentifier ObjectIdentifier
	ObjectIdentifier ObjectIdentifier
}

func (bdor *DeviceObjectReference) Decode(buffer []byte, offset int, apduLen int) (int, error) {
	leng := 0

	// tag 0 device-identifier optional
	bdor.DeviceIdentifier = ObjectIdentifier{}
	leng1, err := bdor.DeviceIdentifier.DecodeContext(buffer, offset+leng, apduLen-leng, 0)
	if err != nil {
		return -1, err
	}
	if leng1 > 0 {
		leng += leng1
	}

	// tag 1 objectidentifier
	bdor.ObjectIdentifier = ObjectIdentifier{}
	leng1, err = bdor.ObjectIdentifier.DecodeContext(buffer, offset+leng, apduLen-leng, 1)
	if err != nil {
		return -1, err
	}
	if leng1 < 0 {
		return -1, errInvalidMessageLength
	}
	leng += leng1

	return leng, nil
}

type ObjectPropertyReference struct {
	ObjectIdentifier   ObjectIdentifier
	PropertyIdentifier interface{}
	PropertyArrayIndex uint32
}

func (bopr *ObjectPropertyReference) Decode(buffer []byte, offset int, apduLen int) (int, error) {
	leng := 0

	// tag 0 objectidentifier
	bopr.ObjectIdentifier = ObjectIdentifier{}
	leng1, err := bopr.ObjectIdentifier.DecodeContext(buffer, offset+leng, apduLen-leng, 0)
	if err != nil {
		return -1, err
	}
	if leng1 < 0 {
		return -1, errInvalidMessageLength
	}
	leng += leng1

	// tag 1 propertyidentifier
	if encoding.IsContextTag(buffer, offset+leng, 1) {
		leng1, _, lenValue, err := encoding.DecodeTagNumberAndValue(buffer, offset+leng)
		if err != nil {
			return -1, err
		}
		leng += leng1
		propID := encoding.PropertyList
		leng1, bopr.PropertyIdentifier, err = encoding.DecodeEnumerated(buffer, offset+leng, lenValue, nil, &propID)
		if err != nil {
			return -1, err
		}
		leng += leng1
	} else {
		return -1, errInvalidTagNumber
	}

	if leng < apduLen {
		// tag 2 property-array-index optional
		if encoding.IsContextTag(buffer, offset+leng, 2) {
			leng1, _, lenValue, err := encoding.DecodeTagNumberAndValue(buffer, offset+leng)
			if err != nil {
				return -1, err
			}
			leng += leng1
			leng1, bopr.PropertyArrayIndex, err = encoding.DecodeUnsigned(buffer, offset+leng, int(lenValue))
			if err != nil {
				return -1, err
			}
			leng += leng1
		}
	}

	return leng, nil
}

type AccumulatorRecord struct {
	Timestamp         TimeStamp
	PresentValue      uint32
	AccumulatedValue  uint32
	AccumulatorStatus AccumulatorStatus
}

type AccumulatorStatus int

const (
	AccumulatorStatusNormal AccumulatorStatus = iota
	AccumulatorStatusStarting
	AccumulatorStatusRecovered
	AccumulatorStatusAbnormal
	AccumulatorStatusFailed
)

func (bar *AccumulatorRecord) Decode(buffer []byte, offset int, apduLen int) (int, error) {
	leng := 0

	// 0 timestamp
	if encoding.IsContextTag(buffer, offset+leng, 0) {
		leng1, _, lenValue, err := encoding.DecodeTagNumberAndValue(buffer, offset+leng)
		if err != nil {
			return -1, err
		}
		leng += leng1
		bar.Timestamp = TimeStamp{}
		leng1, err = bar.Timestamp.Decode(buffer, offset+leng, int(lenValue))
		if err != nil {
			return -1, err
		}
	} else {
		return -1, errors.New("Missing tag 0")
	}

	// 1 present-value
	if encoding.IsContextTag(buffer, offset+leng, 1) {
		leng1, _, lenValue, err := encoding.DecodeTagNumberAndValue(buffer, offset+leng)
		if err != nil {
			return -1, err
		}
		leng += leng1
		leng1, bar.PresentValue, err = encoding.DecodeUnsigned(buffer, offset+leng, int(lenValue))
		if err != nil {
			return -1, err
		}
		leng += leng1
	} else {
		return -1, errors.New("Missing tag 1")
	}

	// 2 accumulated-value
	if encoding.IsContextTag(buffer, offset+leng, 2) {
		leng1, _, lenValue, err := encoding.DecodeTagNumberAndValue(buffer, offset+leng)
		if err != nil {
			return -1, err
		}
		leng += leng1
		leng1, bar.AccumulatedValue, err = encoding.DecodeUnsigned(buffer, offset+leng, int(lenValue))
		if err != nil {
			return -1, err
		}
		leng += leng1
	} else {
		return -1, errors.New("Missing tag 2")
	}

	// 3 accumulator-status
	if encoding.IsContextTag(buffer, offset+leng, 3) {
		leng1, _, lenValue, err := encoding.DecodeTagNumberAndValue(buffer, offset+leng)
		if err != nil {
			return -1, err
		}
		leng += leng1
		leng1, statusValue, err := encoding.DecodeUnsigned(buffer, offset+leng, int(lenValue))
		if err != nil {
			return -1, err
		}
		bar.AccumulatorStatus = AccumulatorStatus(statusValue)
		leng += leng1
	} else {
		return -1, errors.New("Missing tag 3")
	}

	return leng, nil
}

type ActionList struct {
	Action []ActionCommand
}

func (bal *ActionList) Decode(buffer []byte, offset int, apduLen int) (int, error) {
	leng := 0

	// SEQUENCE OF BACnetActionCommand
	if encoding.IsOpeningTagNumber(buffer, offset+leng, 0) {
		leng += 1
		bal.Action = make([]ActionCommand, 0)
		for !encoding.IsClosingTagNumber(buffer, offset+leng, 0) {
			bac := ActionCommand{}
			leng1, err := bac.Decode(buffer, offset+leng, apduLen-leng)
			if err != nil {
				return -1, err
			}
			if leng1 < 0 {
				return -1, errInvalidMessageLength
			}
			leng += leng1
			bal.Action = append(bal.Action, bac)
		}
		leng += 1
	}

	return leng, nil
}

type ActionCommand struct {
	DeviceIdentifier   ObjectIdentifier
	ObjectIdentifier   ObjectIdentifier
	PropertyIdentifier interface{}
	PropertyArrayIndex int
	PropertyValue      []Value
	Priority           int
	PostDelay          int
	QuitOnFailure      bool
	WriteSuccessful    bool
}

func (bac *ActionCommand) Decode(buffer []byte, offset int, apduLen int) (int, error) {
	leng := 0

	// 0 device_identifier optional
	if encoding.IsContextTag(buffer, offset+leng, 0) {
		bac.DeviceIdentifier = ObjectIdentifier{}
		leng1, err := bac.DeviceIdentifier.DecodeContext(buffer, offset+leng, apduLen-leng, 0)
		if err != nil {
			return -1, err
		}
		leng += leng1
	}

	// 1 object_identifier
	if encoding.IsContextTag(buffer, offset+leng, 1) {
		bac.ObjectIdentifier = ObjectIdentifier{}
		leng1, err := bac.ObjectIdentifier.DecodeContext(buffer, offset+leng, apduLen-leng, 1)
		if err != nil {
			return -1, err
		}
		leng += leng1
	} else {
		return -1, errInvalidTagNumber
	}

	// 2 property_identifier
	if encoding.IsContextTag(buffer, offset+leng, 2) {
		leng1, _, lenValue, err := encoding.DecodeTagNumberAndValue(buffer, offset+leng)
		if err != nil {
			return -1, err
		}
		leng += leng1
		propID := encoding.PropertyList
		leng1, bac.PropertyIdentifier, err = encoding.DecodeEnumerated(buffer, offset+leng, lenValue, nil, &propID)
		if err != nil {
			return -1, err
		}
		if leng1 < 0 {
			return -1, errInvalidMessageLength
		}
		leng += leng1
	} else {
		return -1, errInvalidTagNumber
	}

	// 3 property_array_index
	if encoding.IsContextTag(buffer, offset+leng, 3) {
		leng1, _, lenValue, err := encoding.DecodeTagNumberAndValue(buffer, offset+leng)
		if err != nil {
			return -1, err
		}
		leng += leng1
		bac.PropertyArrayIndex, _, err = encoding.DecodeUnsigned(buffer, offset+leng, int(lenValue))
		if err != nil {
			return -1, err
		}
		leng += leng1
	}

	// tag 4 property-value
	if encoding.IsOpeningTagNumber(buffer, offset+leng, 4) {
		leng += 1
		bac.PropertyValue = []Value{}
		for !encoding.IsClosingTagNumber(buffer, offset+leng, 4) && leng < apduLen {
			bv := Value{}
			propID := bac.PropertyIdentifier.(encoding.PropertyIdentifier)
			leng1, _ := bv.Decode(buffer, offset+leng, apduLen-leng, &bac.ObjectIdentifier.Type, &propID)
			if leng1 < 0 {
				return -1, errInvalidMessageLength
			}
			leng += leng1
			bac.PropertyValue = append(bac.PropertyValue, bv)
		}
		if encoding.IsClosingTagNumber(buffer, offset+leng, 4) {
			leng += 1
		} else {
			return -1, errInvalidTagNumber
		}
	} else {
		return -1, errInvalidTagNumber
	}

	if leng < apduLen {
		// tag 5 priority optional
		if encoding.IsContextTag(buffer, offset+leng, 5) {
			leng1, _, lenValue, err := encoding.DecodeTagNumberAndValue(buffer, offset+leng)
			if err != nil {
				return -1, err
			}
			leng += leng1
			bac.Priority, _, err = encoding.DecodeUnsigned(buffer, offset+leng, int(lenValue))
			if err != nil {
				return -1, err
			}
			leng += leng1
		}
	}

	if leng < apduLen {
		// tag 6 post-delay optional
		if encoding.IsContextTag(buffer, offset+leng, 6) {
			leng1, _, lenValue, err := encoding.DecodeTagNumberAndValue(buffer, offset+leng)
			if err != nil {
				return -1, err
			}
			leng += leng1
			bac.PostDelay, _, err = encoding.DecodeUnsigned(buffer, offset+leng, int(lenValue))
			if err != nil {
				return -1, err
			}
			leng += leng1
		}
	}

	if leng < apduLen {
		// tag 7 quit-on-failure optional
		if encoding.IsContextTag(buffer, offset+leng, 7) {
			leng1, _, lenValue, err := encoding.DecodeTagNumberAndValue(buffer, offset+leng)
			if err != nil {
				return -1, err
			}
			leng += leng1
			uVal, _, err := encoding.DecodeUnsigned(buffer, offset+leng, int(lenValue))
			if err != nil {
				return -1, err
			}
			leng += leng1
			bac.QuitOnFailure = uVal > 0
		}
	}

	if leng < apduLen {
		// tag 8 write-successful optional
		if encoding.IsContextTag(buffer, offset+leng, 8) {
			leng1, _, lenValue, err := encoding.DecodeTagNumberAndValue(buffer, offset+leng)
			if err != nil {
				return -1, err
			}
			leng += leng1
			uVal, _, err := encoding.DecodeUnsigned(buffer, offset+leng, int(lenValue))
			if err != nil {
				return -1, err
			}
			leng += leng1
			bac.WriteSuccessful = uVal > 0
		}
	}

	return leng, nil
}

type Scale struct {
	Value interface{}
}

func (bs *Scale) Decode(buffer []byte, offset int, apduLen int) (int, error) {
	leng := 0

	if encoding.IsContextTag(buffer, offset+leng, 0) {
		// float-scale
		leng1, _, lenValue, err := encoding.DecodeTagNumberAndValue(buffer, offset+leng)
		if err != nil {
			return -1, err
		}
		leng += leng1
		leng1, bs.Value = encoding.DecodeRealSafe(buffer, offset+leng, int(lenValue))
		leng += leng1
	} else if encoding.IsContextTag(buffer, offset+leng, 1) {
		// integer-scale
		leng1, _, lenValue, err := encoding.DecodeTagNumberAndValue(buffer, offset+leng)
		if err != nil {
			return -1, err
		}
		leng += leng1
		leng1, bs.Value, err = encoding.DecodeUnsigned(buffer, offset+leng, int(lenValue))
		leng += leng1
	} else {
		return -1, errInvalidTagNumber
	}

	return leng, nil
}

type LightingCommand struct {
	Operation     LightingOperation
	TargetLevel   float32
	RampRate      float32
	StepIncrement float32
	FadeTime      uint32
	Priority      uint32
}

type LightingOperation uint32

const (
	LightingOperationUnknown LightingOperation = iota
	LightingOperationOff
	LightingOperationOn
	LightingOperationToggle
	LightingOperationDecrement
	LightingOperationIncrement
)

func (blc *LightingCommand) Decode(buffer []byte, offset int, apduLen int) (int, error) {
	leng := 0

	// operation
	if encoding.IsContextTag(buffer, offset+leng, 0) {
		leng1, _, lenValue, err := encoding.DecodeTagNumberAndValue(buffer, offset+leng)
		if err != nil {
			return -1, err
		}
		leng += leng1
		leng1, uVal, err := encoding.DecodeUnsigned(buffer, offset+leng, int(lenValue))
		if err != nil {
			return -1, err
		}
		leng += leng1
		blc.Operation = LightingOperation(uVal)
	} else {
		return -1, errInvalidTagNumber
	}

	if leng < apduLen {
		// target-level
		if encoding.IsContextTag(buffer, offset+leng, 1) {
			leng1, _, lenValue, err := encoding.DecodeTagNumberAndValue(buffer, offset+leng)
			if err != nil {
				return -1, err
			}
			leng += leng1
			leng1, blc.TargetLevel = encoding.DecodeRealSafe(buffer, offset+leng, int(lenValue))
			leng += leng1
		}
	}

	if leng < apduLen {
		// ramp-rate
		if encoding.IsContextTag(buffer, offset+leng, 2) {
			leng1, _, lenValue, err := encoding.DecodeTagNumberAndValue(buffer, offset+leng)
			if err != nil {
				return -1, err
			}
			leng += leng1
			leng1, blc.RampRate = encoding.DecodeRealSafe(buffer, offset+leng, int(lenValue))
			leng += leng1
		}
	}

	if leng < apduLen {
		// step-increment
		if encoding.IsContextTag(buffer, offset+leng, 3) {
			leng1, _, lenValue, err := encoding.DecodeTagNumberAndValue(buffer, offset+leng)
			if err != nil {
				return -1, err
			}
			leng += leng1
			leng1, blc.StepIncrement = encoding.DecodeRealSafe(buffer, offset+leng, int(lenValue))
			leng += leng1
		}
	}

	if leng < apduLen {
		// fade-time
		if encoding.IsContextTag(buffer, offset+leng, 4) {
			leng1, _, lenValue, err := encoding.DecodeTagNumberAndValue(buffer, offset+leng)
			if err != nil {
				return -1, err
			}
			leng += leng1
			leng1, blc.FadeTime, err = encoding.DecodeUnsigned(buffer, offset+leng, int(lenValue))
			if err != nil {
				return -1, err
			}
			leng += leng1
		}
	}

	if leng < apduLen {
		// priority
		if encoding.IsContextTag(buffer, offset+leng, 5) {
			leng1, _, lenValue, err := encoding.DecodeTagNumberAndValue(buffer, offset+leng)
			if err != nil {
				return -1, err
			}
			leng += leng1
			leng1, blc.Priority, err = encoding.DecodeUnsigned(buffer, offset+leng, int(lenValue))
			if err != nil {
				return -1, err
			}
			leng += leng1
		}
	}

	return leng, nil
}

type Prescale struct {
	Multiplier   uint32
	ModuloDivide uint32
}

func (bp *Prescale) Decode(buffer []byte, offset int, apduLen int) (int, error) {
	leng := 0

	// multiplier
	if encoding.IsContextTag(buffer, offset+leng, 0) {
		leng1, _, lenValue, err := encoding.DecodeTagNumberAndValue(buffer, offset+leng)
		if err != nil {
			return -1, err
		}
		leng += leng1
		leng1, bp.Multiplier, err = encoding.DecodeUnsigned(buffer, offset+leng, int(lenValue))
		if err != nil {
			return -1, err
		}
		leng += leng1
	} else {
		return -1, errInvalidTagNumber
	}

	// modulo_divide
	if encoding.IsContextTag(buffer, offset+leng, 1) {
		leng1, _, lenValue, err := encoding.DecodeTagNumberAndValue(buffer, offset+leng)
		if err != nil {
			return -1, err
		}
		leng += leng1
		leng1, bp.ModuloDivide, err = encoding.DecodeUnsigned(buffer, offset+leng, int(lenValue))
		if err != nil {
			return -1, err
		}
		leng += leng1
	} else {
		return -1, errInvalidTagNumber
	}

	return leng, nil
}

type ShedLevelChoice int

const (
	BACnetShedLevelChoicePercent ShedLevelChoice = iota
	BACnetShedLevelChoiceLevel
	BACnetShedLevelChoiceAmount
)

type ShedLevel struct {
	Choice ShedLevelChoice
	Value  interface{}
}

func (bsl *ShedLevel) Decode(buffer []byte, offset int, apduLen int) (int, error) {
	leng := 0

	if encoding.IsContextTag(buffer, offset+leng, 0) {
		// percent
		leng1, _, lenValue, err := encoding.DecodeTagNumberAndValue(buffer, offset+leng)
		if err != nil {
			return -1, err
		}
		leng += leng1
		leng1, bsl.Value, err = encoding.DecodeUnsigned(buffer, offset+leng, int(lenValue))
		if err != nil {
			return -1, err
		}
		leng += leng1
		bsl.Choice = BACnetShedLevelChoicePercent
	} else if encoding.IsContextTag(buffer, offset+leng, 1) {
		// level
		leng1, _, lenValue, err := encoding.DecodeTagNumberAndValue(buffer, offset+leng)
		if err != nil {
			return -1, err
		}
		leng += leng1
		leng1, bsl.Value, err = encoding.DecodeUnsigned(buffer, offset+leng, int(lenValue))
		if err != nil {
			return -1, err
		}
		leng += leng1
		bsl.Choice = BACnetShedLevelChoiceLevel
	} else if encoding.IsContextTag(buffer, offset+leng, 2) {
		// amount
		leng1, _, lenValue, err := encoding.DecodeTagNumberAndValue(buffer, offset+leng)
		if err != nil {
			return -1, err
		}
		leng += leng1
		leng1, bsl.Value = encoding.DecodeRealSafe(buffer, offset+leng, int(lenValue))
		leng += leng1
		bsl.Choice = BACnetShedLevelChoiceAmount
	} else {
		return -1, errInvalidTagNumber
	}

	return leng, nil
}

type LogRecordChoice int

const (
	BACnetLogRecordChoiceLogStatus LogRecordChoice = iota
	BACnetLogRecordChoiceBooleanValue
	BACnetLogRecordChoiceRealValue
	BACnetLogRecordChoiceEnumeratedValue
	BACnetLogRecordChoiceUnsignedValue
	BACnetLogRecordChoiceIntegerValue
	BACnetLogRecordChoiceBitstringValue
	BACnetLogRecordChoiceNullValue
	BACnetLogRecordChoiceFailure
	BACnetLogRecordChoiceTimeChange
	BACnetLogRecordChoiceAnyValue
)

type LogRecord struct {
	Timestamp   TimeStamp
	LogDatum    interface{}
	StatusFlags StatusFlags
}

func (blr *LogRecord) Decode(buffer []byte, offset, apduLen int, objType *encoding.ObjectType, propID *encoding.PropertyIdentifier) (int, error) {
	leng := 0

	// timestamp
	if encoding.IsContextTag(buffer, offset+leng, 0) {
		leng1, _, lenValue, err := encoding.DecodeTagNumberAndValue(buffer, offset+leng)
		if err != nil {
			return -1, err
		}
		leng += leng1
		blr.Timestamp = TimeStamp{}
		leng1, err = blr.Timestamp.Decode(buffer, offset+leng, int(lenValue))
		if err != nil {
			return -1, err
		}
		leng += leng1
	} else {
		return -1, errInvalidTagNumber
	}

	if encoding.IsContextTag(buffer, offset+leng, 1) {
		leng1, tagNumber, lenValue, err := encoding.DecodeTagNumberAndValue(buffer, offset+leng)
		if err != nil {
			return -1, err
		}
		leng += leng1

		switch LogRecordChoice(tagNumber) {
		case BACnetLogRecordChoiceLogStatus:
			blr.LogDatum = &LogStatus{}
			leng += blr.LogDatum.(*LogStatus).Decode(buffer, offset+leng, int(lenValue))
		case BACnetLogRecordChoiceBooleanValue:
			blr.LogDatum = buffer[offset+leng] > 0
			leng++
		case BACnetLogRecordChoiceRealValue:
			leng1, logValue := encoding.DecodeRealSafe(buffer, offset+leng, int(lenValue))
			leng += leng1
			blr.LogDatum = logValue
		case BACnetLogRecordChoiceEnumeratedValue:
			leng1, logValue, err := encoding.DecodeEnumerated(buffer, offset+leng, lenValue, nil, nil)
			if err != nil {
				return -1, err
			}
			leng += leng1
			blr.LogDatum = logValue
		case BACnetLogRecordChoiceUnsignedValue:
			leng1, logValue, err := encoding.DecodeUnsigned(buffer, offset+leng, int(lenValue))
			if err != nil {
				return -1, err
			}
			leng += leng1
			blr.LogDatum = logValue
		case BACnetLogRecordChoiceIntegerValue:
			leng1, logValue := encoding.DecodeSigned(buffer, offset+leng, int(lenValue))
			leng += leng1
			blr.LogDatum = logValue
		case BACnetLogRecordChoiceBitstringValue:
			blr.LogDatum = &BitString{}
			leng += blr.LogDatum.(*BitString).Decode(buffer, offset+leng, int(lenValue))
		case BACnetLogRecordChoiceNullValue:
			blr.LogDatum = nil
			leng++
		case BACnetLogRecordChoiceFailure:
			blr.LogDatum = &Error{}
			leng1, err = blr.LogDatum.(*Error).Decode(buffer, offset+leng, apduLen-leng)
			if err != nil {
				return -1, err
			}
			leng += leng1
			if encoding.IsClosingTagNumber(buffer, offset+leng, byte(BACnetLogRecordChoiceFailure)) {
				leng++
			} else {
				return -1, errInvalidTagNumber
			}
		case BACnetLogRecordChoiceTimeChange:
			leng1, logValue := encoding.DecodeRealSafe(buffer, offset+leng, int(lenValue))
			leng += leng1
			blr.LogDatum = logValue
		case BACnetLogRecordChoiceAnyValue:
			blr.LogDatum = []Value{}
			for !encoding.IsClosingTagNumber(buffer, offset+leng, byte(BACnetLogRecordChoiceAnyValue)) && leng < apduLen {
				bValue := Value{}
				leng1, _ := bValue.Decode(buffer, offset+leng, apduLen-leng, objType, propID)
				if leng1 < 0 {
					return -1, errInvalidMessageLength
				}
				leng += leng1
				blr.LogDatum = append(blr.LogDatum.([]Value), bValue)
			}
			if encoding.IsClosingTagNumber(buffer, offset+leng, byte(BACnetLogRecordChoiceAnyValue)) {
				leng++
			} else {
				return -1, errInvalidTagNumber
			}
		default:
			return -1, errInvalidTagNumber
		}
	} else {
		return -1, errInvalidTagNumber
	}

	if leng < apduLen {
		// status-flags optional
		if encoding.IsContextTag(buffer, offset+leng, 2) {
			leng1, _, lenValue, err := encoding.DecodeTagNumberAndValue(buffer, offset+leng)
			if err != nil {
				return -1, err
			}
			leng += leng1
			blr.StatusFlags = StatusFlags{}
			leng += blr.StatusFlags.Decode(buffer, offset+leng, int(lenValue))
		}
	}

	return leng, nil
}

type LogStatus struct {
	UnusedBits uint8
	BitString  *BitString
}

func NewBACnetLogStatus() LogStatus {
	return LogStatus{
		UnusedBits: 5,
		BitString:  NewBACnetBitString(5, *internal.NewBitArrayFromByte(0x00)),
	}
}

func (bls *LogStatus) Decode(buffer []byte, offset, apduLen int) int {
	bls.BitString = NewBACnetBitString(5, *internal.NewBitArrayFromByte(0x00))
	return bls.BitString.Decode(buffer, offset, apduLen)
}

func (bls *LogStatus) SetLogDisabled(a bool) {
	bls.BitString.Value.Set(0, a)
}

func (bls *LogStatus) SetBufferPurged(a bool) {
	bls.BitString.Value.Set(1, a)
}

func (bls *LogStatus) SetLogInterrupted(a bool) {
	bls.BitString.Value.Set(2, a)
}

func (bls *LogStatus) LogDisabled() (bool, error) {
	return bls.BitString.Value.Get(0)
}

func (bls *LogStatus) BufferPurged() (bool, error) {
	return bls.BitString.Value.Get(1)
}

func (bls *LogStatus) LogInterrupted() (bool, error) {
	return bls.BitString.Value.Get(2)
}

type Error struct {
	ErrorClass ErrorClassEnum
	ErrorCode  ErrorCodeEnum
}

func (be *Error) Decode(buffer []byte, offset, apduLen int) (int, error) {
	leng := 0

	// Decode error_class
	leng1, tagNumber, lenValue, err := encoding.DecodeTagNumberAndValue(buffer, offset+leng)
	if err != nil {
		return -1, err
	}
	leng += leng1
	if tagNumber == byte(Enumerated) {
		leng1, eVal, err := encoding.DecodeEnumerated(buffer, offset+leng, lenValue, nil, nil)
		if err != nil {
			return -1, err
		}
		leng += leng1
		be.ErrorClass = ErrorClassEnum(eVal.(uint32))
	} else {
		return -1, errInvalidTagNumber
	}

	// Decode error_code
	leng1, tagNumber, lenValue, err = encoding.DecodeTagNumberAndValue(buffer, offset+leng)
	if err != nil {
		return -1, err
	}
	leng += leng1
	if tagNumber == byte(Enumerated) {
		leng1, eVal, err := encoding.DecodeEnumerated(buffer, offset+leng, lenValue, nil, nil)
		if err != nil {
			return -1, err
		}
		leng += leng1
		be.ErrorCode = ErrorCodeEnum(eVal.(uint32))
	} else {
		return -1, errInvalidTagNumber
	}

	return leng, nil
}

type CalendarEntry struct {
	Value interface{}
}

func (ce *CalendarEntry) Decode(buffer []byte, offset, apduLen int) (int, error) {
	leng := 0
	leng1, tagNumber, lenValue, err := encoding.DecodeTagNumberAndValue(buffer, offset+leng)
	if err != nil {
		return -1, err
	}
	leng += leng1
	if tagNumber == 0 {
		leng1, ce.Value = encoding.DecodeDateSafe(buffer, offset+leng, int(lenValue))
		leng += leng1
	} else if tagNumber == 1 {
		ce.Value = &DateRange{}
		leng1, err := ce.Value.(*DateRange).Decode(buffer, offset+leng, int(lenValue))
		if err != nil {
			return -1, err
		}
		leng += leng1
	} else if tagNumber == 2 {
		ce.Value = &WeekNDay{}
		leng += ce.Value.(*WeekNDay).Decode(buffer, offset+leng, int(lenValue))
	} else {
		return -1, errInvalidTagNumber
	}

	return leng, nil
}

type EventLogRecord struct {
	Timestamp DateTime
	LogDatum  interface{}
}

func (elr *EventLogRecord) Decode(buffer []byte, offset, apduLen int) (int, error) {
	leng := 0
	if encoding.IsContextTag(buffer, offset+leng, 0) {
		leng1, _, _, err := encoding.DecodeTagNumberAndValue(buffer, offset+leng)
		if err != nil {
			return -1, err
		}
		leng += leng1
		elr.Timestamp = DateTime{}
		leng += elr.Timestamp.Decode(buffer, offset+leng)
	} else {
		return -1, errInvalidTagNumber
	}

	return leng, nil
}

type WeekNDay struct {
	Month       int
	WeekOfMonth int
	DayOfWeek   int
}

func (wnd *WeekNDay) Decode(buffer []byte, offset, apduLen int) int {
	if apduLen >= 3 {
		wnd.Month = int(buffer[offset])
		wnd.WeekOfMonth = int(buffer[offset+1])
		wnd.DayOfWeek = int(buffer[offset+2])
	} else {
		return -1
	}
	return 3
}

type ReadAccessResultReadResult struct {
	PropertyIdentifier encoding.PropertyIdentifier
	PropertyArrayIndex uint32
	ReadResult         interface{} // Either BACnetValue or BACnetError
}

func (rarr *ReadAccessResultReadResult) Decode(buffer []byte, offset, apduLen int) (int, error) {
	leng := 0
	// 2 propertyidentifier
	if encoding.IsContextTag(buffer, offset+leng, 2) {
		leng1, _, lenValue, err := encoding.DecodeTagNumberAndValue(buffer, offset+leng)
		if err != nil {
			return -1, err
		}
		leng += leng1
		propID := encoding.PropertyList
		leng1, propertyIdentifier, err := encoding.DecodeEnumerated(buffer, offset+leng, lenValue, nil, &propID)
		if err != nil {
			return -1, err
		}
		leng += leng1
		rarr.PropertyIdentifier = encoding.PropertyIdentifier(propertyIdentifier.(uint32))
	} else {
		return -1, errInvalidTagNumber
	}

	// 3 property_array_index
	if encoding.IsContextTag(buffer, offset+leng, 3) {
		leng1, _, lenValue, err := encoding.DecodeTagNumberAndValue(buffer, offset+leng)
		if err != nil {
			return -1, err
		}
		leng += leng1
		leng1, propertyArrayIndex, err := encoding.DecodeUnsigned(buffer, offset+leng, int(lenValue))
		if err != nil {
			return -1, err
		}
		leng += leng1
		rarr.PropertyArrayIndex = propertyArrayIndex
	}

	if leng < apduLen {
		if encoding.IsOpeningTagNumber(buffer, offset+leng, 4) {
			rarr.ReadResult = &Value{}
			leng1, err := rarr.ReadResult.(*Value).Decode(buffer, offset+leng, apduLen-leng, nil, nil)
			if err != nil {
				return -1, err
			}
			leng += leng1
			if encoding.IsClosingTagNumber(buffer, offset+leng, 4) {
				leng += 1
			} else {
				return -1, errInvalidTagNumber
			}
		} else if encoding.IsOpeningTagNumber(buffer, offset+leng, 5) {
			rarr.ReadResult = &Error{}
			leng1, err := rarr.ReadResult.(*Error).Decode(buffer, offset+leng, apduLen-leng)
			if err != nil {
				return -1, err
			}
			leng += leng1
			if encoding.IsClosingTagNumber(buffer, offset+leng, 5) {
				leng += 1
			} else {
				return -1, errInvalidTagNumber
			}
		}
	}
	return leng, nil
}

type ReadAccessResult struct {
	ObjectIdentifier ObjectIdentifier
	ListOfResults    []ReadAccessResultReadResult
}

func (rar *ReadAccessResult) Decode(buffer []byte, offset, apduLen int) (int, error) {
	leng := 0
	// tag 0 objectidentifier
	rar.ObjectIdentifier = ObjectIdentifier{}
	if encoding.IsClosingTagNumber(buffer, offset+leng, 0) {
		leng1, err := rar.ObjectIdentifier.DecodeContext(buffer, offset+leng, apduLen-leng, 0)
		if err != nil {
			return -1, err
		}
		leng += leng1
	} else {
		return -1, errInvalidTagNumber
	}

	if encoding.IsOpeningTagNumber(buffer, offset+leng, 1) {
		leng += 1
		rar.ListOfResults = make([]ReadAccessResultReadResult, 0)

		for (apduLen-leng) > 1 && !encoding.IsClosingTagNumber(buffer, offset+leng, 1) {
			bValue := ReadAccessResultReadResult{}
			leng1, err := bValue.Decode(buffer, offset+leng, apduLen-leng)
			if err != nil {
				return -1, err
			}
			leng += leng1

			rar.ListOfResults = append(rar.ListOfResults, bValue)
		}

		if encoding.IsClosingTagNumber(buffer, offset+leng, 1) {
			leng += 1
		} else {
			return -1, errInvalidTagNumber
		}
	} else {
		return -1, errInvalidTagNumber
	}

	return leng, nil
}

type AccessRule struct {
	TimeRangeSpecifier TimeRangeSpecifierChoice
	TimeRange          DeviceObjectPropertyReference
	LocationSpecifier  LocationSpecifierChoice
	Location           DeviceObjectReference
	Enable             bool
}

type TimeRangeSpecifierChoice int

const (
	Specified TimeRangeSpecifierChoice = iota
	Always
)

type LocationSpecifierChoice int

const (
	SpecifiedLocation LocationSpecifierChoice = iota
	All
)

func (bar *AccessRule) Decode(buffer []byte, offset, apduLen int) int {
	return -1
}

type NameValue struct {
	Name  string
	Value Value
}

func (bnv *NameValue) Decode(buffer []byte, offset, apduLen int) (int, error) {
	leng := 0

	// Name
	leng1, tagNumber, lenValue, err := encoding.DecodeTagNumberAndValue(buffer, offset+leng)
	if err != nil {
		return -1, err
	}
	if tagNumber != 0 {
		return -1, errInvalidTagNumber
	}
	leng += leng1
	leng1, bnv.Name = encoding.DecodeCharacterString(buffer, offset+leng, apduLen-leng, int(lenValue))
	leng += leng1

	// Decode value
	decodeLen := 0
	if leng < apduLen {
		leng1, tagNumber, lenValue, err := encoding.DecodeTagNumberAndValue(buffer, offset+leng)
		if err != nil {
			return -1, err
		}
		leng += leng1

		switch ApplicationTags(tagNumber) {
		case Null:
			bnv.Value = Value{Value: nil}
			decodeLen = 0
			// Fixme: fix null type nothing else to do, some Error occurs!!!!
		case Boolean:
			if lenValue > 0 {
				bnv.Value = Value{Value: true}
			} else {
				bnv.Value = Value{Value: false}
			}
		case UnsignedInt:
			decodeLen, bnv.Value.Value, err = encoding.DecodeUnsigned(buffer, offset+leng, int(lenValue))
			if err != nil {
				return -1, err
			}
		case SignedInt:
			decodeLen, bnv.Value.Value = encoding.DecodeSigned(buffer, offset+leng, int(lenValue))
		case Real:
			decodeLen, bnv.Value.Value = encoding.DecodeRealSafe(buffer, offset+leng, int(lenValue))
		case Double:
			decodeLen, bnv.Value.Value = encoding.DecodeDoubleSafe(buffer, offset+leng, int(lenValue))
		case OctetString:
			decodeLen, bnv.Value.Value = encoding.DecodeOctetString(buffer, offset+leng, int(lenValue))
		case CharacterString:
			decodeLen, bnv.Value.Value = encoding.DecodeCharacterString(buffer, offset+leng, apduLen-leng, int(lenValue))
		case ApplicationTagsBitString:
			bitValue := BitString{}
			decodeLen = bitValue.Decode(buffer, offset+leng, int(lenValue))
			bnv.Value.Value = bitValue
		case Enumerated:
			decodeLen, bnv.Value.Value, err = encoding.DecodeEnumerated(buffer, offset+leng, lenValue, nil, nil)
			if err != nil {
				return -1, err
			}
		case Date:
			decodeLen, dateValue := encoding.DecodeDateSafe(buffer, offset+leng, int(lenValue))

			if leng < apduLen {
				leng1, tagNumber, _, err := encoding.DecodeTagNumberAndValue(buffer, offset+leng+decodeLen)
				if err != nil {
					return -1, err
				}
				if tagNumber == byte(Time) {
					leng += leng1
					leng--
					bnv.Value.Value = &DateTime{}
					decodeLen = bnv.Value.Value.(*DateTime).Decode(buffer, offset+leng)
				}
			} else {
				bnv.Value.Value = dateValue
			}
		case Time:
			decodeLen, bnv.Value.Value = encoding.DecodeBACnetTimeSafe(buffer, offset+leng, int(lenValue))
		}

		if decodeLen < 0 {
			return -1, errInvalidMessageLength
		}
		leng += decodeLen
	}

	return leng, nil
}

type NameValueCollection struct {
	Members []NameValue
}

func (bnc *NameValueCollection) Decode(buffer []byte, offset, apduLen int) (int, error) {
	leng := 0

	// Check if it's an opening tag number
	if !encoding.IsOpeningTagNumber(buffer, offset+leng, 0) {
		return -1, errInvalidTagNumber
	}

	leng += 1
	bnc.Members = make([]NameValue, 0)

	for !encoding.IsClosingTagNumber(buffer, offset+leng, 0) {
		bValue := NameValue{}
		leng1, err := bValue.Decode(buffer, offset+leng, apduLen-leng)
		if err != nil {
			return -1, err
		}
		if leng1 < 0 {
			return -1, errInvalidMessageLength
		}
		leng += leng1
		bnc.Members = append(bnc.Members, bValue)
	}

	leng += 1
	return leng, nil
}

type SecurityPolicy int

type NetworkSecurityPolicy struct {
	PortID        int
	SecurityLevel SecurityPolicy
}

func (bns *NetworkSecurityPolicy) Decode(buffer []byte, offset, apduLen int) (int, error) {
	leng := 0

	// port_id
	leng1, tagNumber, lenValue, err := encoding.DecodeTagNumberAndValue(buffer, offset+leng)
	if err != nil {
		return -1, err
	}
	if tagNumber != 0 {
		return -1, errInvalidTagNumber
	}
	leng += leng1
	leng1, portID, err := encoding.DecodeUnsigned(buffer, offset+leng, int(lenValue))
	leng += leng1
	bns.PortID = int(portID)

	leng = 0
	// security_level
	leng1, tagNumber, lenValue, err = encoding.DecodeTagNumberAndValue(buffer, offset+leng)
	if err != nil {
		return -1, err
	}
	if tagNumber != 1 {
		return -1, errInvalidTagNumber
	}
	leng += leng1
	leng1, uVal, err := encoding.DecodeUnsigned(buffer, offset+leng, int(lenValue))
	if err != nil {
		return -1, err
	}
	leng += leng1
	bns.SecurityLevel = SecurityPolicy(uVal)

	return leng, nil
}

type PortPermission struct {
	PortID  int
	Enabled bool
}

func (bpp *PortPermission) Decode(buffer []byte, offset, apduLen int) (int, error) {
	leng := 0

	// port_id
	leng1, tagNumber, lenValue, err := encoding.DecodeTagNumberAndValue(buffer, offset+leng)
	if err != nil {
		return -1, err
	}
	if tagNumber != 0 {
		return -1, errInvalidTagNumber
	}
	leng += leng1
	leng1, portID, err := encoding.DecodeUnsigned(buffer, offset+leng, int(lenValue))
	if err != nil {
		return -1, err
	}
	leng += leng1
	bpp.PortID = int(portID)

	leng = 0
	// enabled
	leng1, tagNumber, lenValue, err = encoding.DecodeTagNumberAndValue(buffer, offset+leng)
	if err != nil {
		return -1, err
	}
	if tagNumber != 1 {
		return -1, errInvalidTagNumber
	}
	leng += leng1
	if lenValue > 0 {
		bpp.Enabled = true
	} else {
		bpp.Enabled = false
	}

	return leng, nil
}

type PriorityValue struct {
	Value interface{}
}

func (bpv *PriorityValue) Decode(buffer []byte, offset, apduLen int) int {
	// TODO Implement decoding logic for BACnetPriorityValue
	return -1
}

type PriorityArray struct {
	Value [16]PriorityValue
}

func (bpa *PriorityArray) Decode(buffer []byte, offset, apduLen int) int {
	leng := 0
	i := 0

	for leng < apduLen && i < 16 {
		bpa.Value[i] = PriorityValue{}
		leng += bpa.Value[i].Decode(buffer, offset+leng, apduLen-leng)
		i++
	}

	return leng
}

type ProcessIdSelection struct {
	Value interface{} // You can specify the type you expect here
}

func (bps *ProcessIdSelection) Decode(buffer []byte, offset, apduLen int) int {
	// TODO Implement decoding logic for BACnetProcessIdSelection
	return -1
}

type PropertyAccessResult struct {
	ObjectIdentifier   ObjectIdentifier
	PropertyIdentifier encoding.PropertyIdentifier
	PropertyArrayIndex int
	DeviceIdentifier   ObjectIdentifier
	AccessResult       interface{}
}

func (bpar *PropertyAccessResult) Decode(buffer []byte, offset, apduLen int) int {
	// TODO Implement decoding logic for BACnetPropertyAccessResult here
	return -1
}

type SetpointReference struct {
	Value interface{}
}

func (bsr *SetpointReference) Decode(buffer []byte, offset, apduLen int) int {
	// TODO Implement decoding logic for BACnetSetpointReference here
	return 0
}

type SpecialEvent struct {
	Period           interface{}
	ListOfTimeValues []interface{}
	EventPriority    int
}

func (bse *SpecialEvent) Decode(buffer []byte, offset, apduLen int) int {
	// TODO Implement decoding logic for BACnetSpecialEvent here
	return 0
}

type TimerStateChangeValue struct {
	Value interface{}
}

func (scv *TimerStateChangeValue) Decode([]byte, int, int) int {
	// TODO implement decoder
	return -1
}

type ValueSource struct {
	Value interface{}
}

func (scv *ValueSource) Decode([]byte, int, int) int {
	// TODO implement decoder
	return -1
}

type VMACEntry struct {
	VirtualMacAddress interface{}
	NativeMacAddress  interface{}
}

func (scv *VMACEntry) Decode([]byte, int, int) int {
	// TODO implement decoder
	return -1
}

type AssignedAccessRights struct {
	AssignedAccessRights DeviceObjectReference
	Enable               bool
}

func (scv *AssignedAccessRights) Decode([]byte, int, int) int {
	// TODO implement decoder
	return -1
}

type AssignedLandingCalls struct {
	LandingCalls []landingCall
}

type landingCall struct {
	FloorNumber int
	Direction   LiftCarDirection
}

func (balc *AssignedLandingCalls) Decode(buffer []byte, offset, apduLen int) int {
	// Implement decoding logic for BACnetAssignedLandingCalls
	return 0
}

type LiftCarDirection int

type AuthenticationFactor struct {
	FormatType  AuthenticationFactorType
	FormatClass int
	Value       []byte
}

func (baf *AuthenticationFactor) Decode(buffer []byte, offset, apduLen int) int {
	// Implement decoding logic for BACnetAuthenticationFactor
	return 0
}

type AuthenticationFactorType int

type AuthenticationFactorFormat struct {
	FormatType   AuthenticationFactorType
	VendorID     int
	VendorFormat int
}

func (baff *AuthenticationFactorFormat) Decode(buffer []byte, offset, apduLen int) int {
	// Implement decoding logic for AuthenticationFactorFormat
	return 0
}

type AuthenticationPolicy struct {
	Policies      []policy
	OrderEnforced bool
	Timeout       int
}

type policy struct {
	CredentialDataInput DeviceObjectReference
	Index               int
}

func (bap *AuthenticationPolicy) Decode(buffer []byte, offset, apduLen int) int {
	// Implement decoding logic for AuthenticationPolicy
	return 0
}

type BDTEntry struct {
	// Define BDTEntry structure here
}

type ChannelValue struct {
	Value interface{}
}

func (bcv *ChannelValue) Decode(buffer []byte, offset, apduLen int) int {
	// Implement decoding logic for BACnetChannelValue
	return 0
}

type COVSubscription struct {
	Recipient                   RecipientProcess
	MonitoredPropertyReference  ObjectPropertyReference
	IssueConfirmedNotifications bool
	TimeRemaining               int
	COVIncrement                float64
}

func (bcs *COVSubscription) Decode(buffer []byte, offset, apduLen int) int {
	// Implement decoding logic for BACnetCOVSubscription
	return 0
}

type AccessAuthenticationFactorDisable int

type CredentialAuthenticationFactor struct {
	Disable              AccessAuthenticationFactorDisable
	AuthenticationFactor AuthenticationFactor
}

func (bcaf *CredentialAuthenticationFactor) Decode(buffer []byte, offset, apduLen int) int {
	// Implement decoding logic for BACnetCredentialAuthenticationFactor
	return 0
}

type DailySchedule struct {
	DaySchedule []interface{}
}

func (bds *DailySchedule) Decode(buffer []byte, offset, apduLen int) int {
	// Implement decoding logic for BACnetDailySchedule
	return 0
}

type RecipientProcess struct {
	// Define BACnetRecipientProcess structure here
}

type EventNotificationSubscription struct {
	Recipient                   Recipient
	ProcessIdentifier           int
	IssueConfirmedNotifications bool
	TimeRemaining               int
}

func (bens *EventNotificationSubscription) Decode(buffer []byte, offset, apduLen int) int {
	// Implement decoding logic for BACnetEventNotificationSubscription
	return 0
}

type EventParameter struct {
	Value interface{}
}

func (bep *EventParameter) Decode(buffer []byte, offset, apduLen int) int {
	// Implement decoding logic for BACnetEventParameter
	return 0
}

type FaultParameter struct {
	Value interface{}
}

func (bfp *FaultParameter) Decode(buffer []byte, offset, apduLen int) int {
	// Implement decoding logic for BACnetFaultParameter
	return 0
}
