package bacnet

import (
	"encoding/binary"
	"errors"

	"github.com/absmach/bacnet/internal"
)

// NPDU Network Protocol Data Unit netwrok layer data packet.
type NPDU struct {
	Version     uint8 // Always one.
	Control     NPDUControlInformation
	DNET        uint16
	DLEN        uint8
	DADR        []byte
	Destination Address
	SNET        uint16
	SLEN        uint8
	SADR        []byte
	Source      Address
	MessageType byte
	HopCount    byte
	VendorID    uint16
}

// NPDUControlInformation a bit array to define network control information.
type NPDUControlInformation struct {
	control internal.BitArray
}

// NewNPDUControlInformation creates a new bit array for network control info.
func NewNPDUControlInformation() *NPDUControlInformation {
	return &NPDUControlInformation{
		control: *internal.NewBitArray(8),
	}
}

// IsNetworkLayerMessage returns wether a message is a network layer message.
func (nci *NPDUControlInformation) IsNetworkLayerMessage() bool {
	val, err := nci.control.Get(0)
	if err != nil {
		return false
	}
	return val
}

// SetNetworkLayerMessage sets the value for the netwrok layer message bit.
func (nci *NPDUControlInformation) SetNetworkLayerMessage(a bool) {
	nci.control.Set(0, a)
}

// IsDestinationSpecifier returns based on the npdu control bit bit if it is a message specifier.
func (nci *NPDUControlInformation) IsDestinationSpecifier() bool {
	val, err := nci.control.Get(2)
	if err != nil {
		return false
	}
	return val
}

func (nci *NPDUControlInformation) SetDestinationSpecifier(a bool) {
	nci.control.Set(2, a)
}

func (nci *NPDUControlInformation) IsSourceSpecifier() bool {
	val, err := nci.control.Get(4)
	if err != nil {
		return false
	}
	return val
}

func (nci *NPDUControlInformation) SetSourceSpecifier(a bool) {
	nci.control.Set(4, a)
}

func (nci *NPDUControlInformation) IsDataExpectingReply() bool {
	val, err := nci.control.Get(5)
	if err != nil {
		return false
	}
	return val
}

func (nci *NPDUControlInformation) SetDataExpectingReply(a bool) {
	nci.control.Set(5, a)
}

// NetworkPriority returns the network priority based on the network control information bits.
func (nci *NPDUControlInformation) NetworkPriority() (NetworkPriority, error) {
	control7, err := nci.control.Get(7)
	if err != nil {
		return 0, err
	}
	control6, err := nci.control.Get(6)
	if err != nil {
		return 0, err
	}
	switch {
	case !control6 && !control7:
		return NormalMessage, nil
	case !control6 && control7:
		return UrgentMessage, nil
	case control6 && !control7:
		return CriticalEquipmentMessage, nil
	case control6 && control7:
		return LifeSafetyMessage, nil
	default:
		return 0, errors.New("invalid network priority")
	}
}

// SetNetworkPriority sets the network control information based on network priority set.
func (nci *NPDUControlInformation) SetNetworkPriority(a NetworkPriority) error {
	switch a {
	case NormalMessage:
		nci.control.Set(6, false)
		nci.control.Set(7, false)
	case UrgentMessage:
		nci.control.Set(6, false)
		nci.control.Set(7, true)
	case CriticalEquipmentMessage:
		nci.control.Set(6, true)
		nci.control.Set(7, false)
	case LifeSafetyMessage:
		nci.control.Set(6, true)
		nci.control.Set(7, true)
	default:
		return errors.New("invalid network priority")
	}
	return nil
}

// Encode encodes network control information.
func (nci *NPDUControlInformation) Encode() ([]byte, error) {
	b, err := nci.control.ToByte()
	if err != nil {
		return []byte{}, err
	}
	return []byte{b}, nil
}

// Decode decodes network control information from a byte buffer.
func (nci *NPDUControlInformation) Decode(buffer []byte, offset int) int {
	if offset < len(buffer) {
		nci.control = *internal.NewBitArrayFromByte(buffer[offset])
		return 1
	}
	return 0
}

// NewNPDU creates a new NPDU.
func NewNPDU(destination *Address, source *Address, hopCount *uint8, vendorID *uint16) *NPDU {
	npdu := &NPDU{
		Version:     1,
		Control:     *NewNPDUControlInformation(),
		Destination: *destination,
		Source:      *source,
	}
	switch hopCount {
	case nil:
		npdu.HopCount = 255
	default:
		npdu.HopCount = *hopCount
	}
	if vendorID != nil {
		npdu.VendorID = *vendorID
	}

	if destination != nil && destination.NetworkNumber > 0 {
		npdu.Control.SetDestinationSpecifier(true)
		npdu.DNET = uint16(destination.NetworkNumber)
		npdu.DLEN = uint8(len(destination.MacAddress))
		npdu.DADR = destination.MacAddress
	}

	if source != nil && source.NetworkNumber > 0 && source.NetworkNumber < 0xFFFF {
		npdu.Control.SetSourceSpecifier(true)
		npdu.SNET = uint16(source.NetworkNumber)
		npdu.SLEN = uint8(len(source.MacAddress))
		npdu.SADR = source.MacAddress
	}

	return npdu
}

// Encode encodes the NPDU data to []byte.
func (npdu *NPDU) Encode() ([]byte, error) {
	buffer := make([]byte, 0)
	buffer = append(buffer, npdu.Version)
	ctrlBuf, err := npdu.Control.Encode()
	if err != nil {
		return buffer, err
	}
	buffer = append(buffer, ctrlBuf...)

	if npdu.Control.IsDestinationSpecifier() {
		buffer = append(buffer, uint8(npdu.DNET>>8), uint8(npdu.DNET&0xFF))
		if npdu.DNET == 0xFFFF {
			buffer = append(buffer, 0x00)
		} else {
			buffer = append(buffer, npdu.DLEN)
			buffer = append(buffer, npdu.DADR...)
		}
	}

	if npdu.Control.IsSourceSpecifier() {
		buffer = append(buffer, uint8(npdu.SNET>>8), uint8(npdu.SNET&0xFF))
		buffer = append(buffer, npdu.SLEN)
		buffer = append(buffer, npdu.SADR...)
	}

	if npdu.Control.IsDestinationSpecifier() {
		buffer = append(buffer, npdu.HopCount)
	}

	if npdu.Control.IsNetworkLayerMessage() {
		buffer = append(buffer, npdu.MessageType)
		if npdu.MessageType >= 0x80 && npdu.MessageType <= 0xFF {
			buffer = append(buffer, uint8(npdu.VendorID>>8), uint8(npdu.VendorID&0xFF))
		}
	}

	return buffer, nil
}

// Decode decodes []byte to NPDU.
func (npdu *NPDU) Decode(buffer []byte, offset int) int {
	length := 0
	version := buffer[offset] // always 1!!!!
	length++
	if version != npdu.Version {
		return -1
	}

	npdu.Control = *NewNPDUControlInformation()
	length += npdu.Control.Decode(buffer, offset+length)

	if npdu.Control.IsDestinationSpecifier() {
		npdu.DNET = binary.BigEndian.Uint16(buffer[offset+length : offset+length+2])
		length += 2
		npdu.DLEN = buffer[offset+length]
		length++
		npdu.DADR = buffer[offset+length : offset+length+int(npdu.DLEN)]
		length += int(npdu.DLEN)
		npdu.Destination = NewAddress(uint32(npdu.DNET), npdu.DADR, "", nil)
	}

	if npdu.Control.IsSourceSpecifier() {
		npdu.SNET = binary.BigEndian.Uint16(buffer[offset+length : offset+length+2])
		length += 2
		npdu.SLEN = buffer[offset+length]
		length++
		npdu.SADR = buffer[offset+length : offset+length+int(npdu.SLEN)]
		length += int(npdu.SLEN)
		npdu.Source = NewAddress(uint32(npdu.SNET), npdu.SADR, "", nil)
	}

	if npdu.Control.IsDestinationSpecifier() {
		npdu.HopCount = buffer[offset+length]
		length++
	}

	if npdu.Control.IsNetworkLayerMessage() {
		npdu.MessageType = buffer[offset+length]
		length++
		if npdu.MessageType >= 0x80 {
			npdu.VendorID = binary.BigEndian.Uint16(buffer[offset+length : offset+length+2])
			length += 2
		}
	}

	return length
}
