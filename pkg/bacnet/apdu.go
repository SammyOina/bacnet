package bacnet

import "errors"

// errUnknownPDU is an error returned when an unknown PDU type is encountered.
var errUnknownPDU = errors.New("unkown PDU type")

// BACnetConfirmedServiceChoice represents BACnet confirmed service choices.
type BACnetConfirmedServiceChoice int

// BACnetUnconfirmedServiceChoice represents BACnet unconfirmed service choices.
type BACnetUnconfirmedServiceChoice int

const (
	segmentedMessageMask          = 0x10
	moreFollowsMask               = 0x20
	segmentedResponseAcceptedMask = 0x40
)

const (
	// Alarm and Event Services
	AcknowledgeAlarm                 BACnetConfirmedServiceChoice = 0
	ConfirmedCovNotification         BACnetConfirmedServiceChoice = 1
	ConfirmedCovNotificationMultiple BACnetConfirmedServiceChoice = 31
	ConfirmedEventNotification       BACnetConfirmedServiceChoice = 2
	GetAlarmSummary                  BACnetConfirmedServiceChoice = 3
	GetEnrollmentSummary             BACnetConfirmedServiceChoice = 4
	GetEventInformation              BACnetConfirmedServiceChoice = 29
	LifeSafetyOperation              BACnetConfirmedServiceChoice = 27
	SubscribeCov                     BACnetConfirmedServiceChoice = 5
	SubscribeCovProperty             BACnetConfirmedServiceChoice = 28
	SubscribeCovPropertyMultiple     BACnetConfirmedServiceChoice = 30
	// File Access Services
	AtomicReadFile  BACnetConfirmedServiceChoice = 6
	AtomicWriteFile BACnetConfirmedServiceChoice = 7
	// Object Access Services
	AddListElement    BACnetConfirmedServiceChoice = 8
	RemoveListElement BACnetConfirmedServiceChoice = 9
	CreateObject      BACnetConfirmedServiceChoice = 10
	DeleteObject      BACnetConfirmedServiceChoice = 11
	ReadProperty      BACnetConfirmedServiceChoice = 12
	// SERVICE_CONFIRMED_READ_CONDITIONAL = 13 removed
	ReadPropertyMultiple  BACnetConfirmedServiceChoice = 14
	ReadRange             BACnetConfirmedServiceChoice = 26
	WriteProperty         BACnetConfirmedServiceChoice = 15
	WritePropertyMultiple BACnetConfirmedServiceChoice = 16
	// Remote Device Management Services
	DeviceCommunicationControl BACnetConfirmedServiceChoice = 17
	ConfirmedPrivateTransfer   BACnetConfirmedServiceChoice = 18
	ConfirmedTextMessage       BACnetConfirmedServiceChoice = 19
	ReinitializeDevice         BACnetConfirmedServiceChoice = 20
	// Virtual Terminal Services
	VTOpen  BACnetConfirmedServiceChoice = 21
	VTClose BACnetConfirmedServiceChoice = 22
	VTData  BACnetConfirmedServiceChoice = 23
	// Security Services
	// SERVICE_CONFIRMED_AUTHENTICATE = 24 removed
	// SERVICE_CONFIRMED_REQUEST_KEY = 25 removed
)

const (
	ServiceChoiceIAm BACnetUnconfirmedServiceChoice = iota
	ServiceChoiceIHave
	UnconfirmedCovNotification
	UnconfirmedEVENTNotification
	UnconfirmedPrivateTransfer
	UnconfirmedTextMessage
	TimeSynchronization
	ServiceChoiceWhoHas
	ServiceChoiceWhoIs
	UTCTimeSynchronization
	WriteGroup
	UnconfirmedCovNotificationMultiple
	ServiceChoiceWhoAmI
	ServiceChoiceYouAre
)

// APDU Application Protocol Data Unit.
type APDU struct {
	PduType                   PduTypes
	SegmentedMessage          bool
	MoreFollows               bool
	SegmentedResponseAccepted bool
	MaxSegmentsAccepted       MaxSegments
	MaxApduLengthAccepted     MaxAPDU
	SequenceNumber            byte
	InvokeID                  byte
	RequenceNumber            byte
	ProposedWindowSize        byte
	ServiceChoice             byte
}

// Encode encodes APDU data to []byte.
func (a APDU) Encode() ([]byte, error) {
	buffer := make([]byte, 0)

	tmp := byte(a.PduType)
	if a.SegmentedMessage {
		tmp |= segmentedMessageMask
	}
	if a.MoreFollows {
		tmp |= moreFollowsMask
	}
	if a.SegmentedResponseAccepted {
		tmp |= segmentedResponseAcceptedMask
	}

	switch a.PduType {
	case PDUTypeConfirmedServiceRequest:
		buffer = append(buffer, tmp)
		buffer = append(buffer, byte(a.MaxSegmentsAccepted)|byte(a.MaxApduLengthAccepted))
		buffer = append(buffer, a.InvokeID)
		if a.SegmentedMessage {
			buffer = append(buffer, a.SequenceNumber, a.ProposedWindowSize)
		}
	case PDUTypeUnconfirmedServiceRequest:
		buffer = append(buffer, tmp)
	case PDUTypeSimpleAck:
		buffer = append(buffer, tmp)
		buffer = append(buffer, a.InvokeID)
	case PDUTypeComplexAck:
		buffer = append(buffer, a.InvokeID)
		if a.SegmentedMessage {
			buffer = append(buffer, a.SequenceNumber, a.ProposedWindowSize)
		}
	default:
		return []byte{}, errUnknownPDU
	}

	buffer = append(buffer, a.ServiceChoice)
	return buffer, nil
}

// Decode decodes []byte to APDU data.
func (a *APDU) Decode(buffer []byte, offset int) (int, error) {
	length := 0
	a.PduType = PduTypes(buffer[offset])
	tmp := byte(buffer[offset])
	length++

	switch a.PduType {
	case PDUTypeConfirmedServiceRequest:
		a.SegmentedMessage = tmp&segmentedMessageMask != 0
		a.MoreFollows = tmp&moreFollowsMask != 0
		a.SegmentedResponseAccepted = tmp&segmentedResponseAcceptedMask != 0
		a.MaxSegmentsAccepted = MaxSegments(buffer[offset+length] & 0xF0)
		a.MaxApduLengthAccepted = MaxAPDU(buffer[offset+length] & 0x0F)
		length++
		a.InvokeID = buffer[offset+length]
		length++
		if a.SegmentedMessage {
			a.SequenceNumber = buffer[offset+length]
			length++
			a.ProposedWindowSize = buffer[offset+length]
			length++
		}
		a.ServiceChoice = buffer[offset+length]
		length++
	case PDUTypeUnconfirmedServiceRequest:
		a.ServiceChoice = buffer[offset+length]
		length++
	case PDUTypeSimpleAck:
		a.InvokeID = buffer[offset+length]
		length++
		a.ServiceChoice = buffer[offset+length]
		length++
	case PDUTypeComplexAck:
		a.SegmentedMessage = tmp&segmentedMessageMask != 0
		a.InvokeID = buffer[offset+length]
		length++
		a.ServiceChoice = buffer[offset+length]
		length++
		if a.SegmentedMessage {
			a.SequenceNumber = buffer[offset+length]
			length++
			a.ProposedWindowSize = buffer[offset+length]
			length++
		}
	default:
		return -1, errUnknownPDU
	}

	return length, nil
}
