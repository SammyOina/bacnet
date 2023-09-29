package bacnet

import (
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/absmach/bacnet/pkg/transport"
)

var (
	errUnsupportedTransport = errors.New("unsupported transport")
	errInvalidMessageLength = errors.New("invalid message length")
	errUnsupportedFunction  = errors.New("unsupported BVLC function")
)

//go:generate stringer -type=BVLCFunctions
type BVLCFunctions int

const (
	BVLCResult BVLCFunctions = iota
	BVLCWriteBroadcastDistributionTable
	BVLCReadBroadcastDistTable
	BVLCReadBroadcastDistTableAck
	BVLCForwardedNPDU
	BVLCRegisterForeignDevice
	BVLCReadForeignDeviceTable
	BVLCReadForeignDeviceTableAck
	BVLCDeleteForeignDeviceTableENTRY
	BVLCDistributeBroadcastToNetwok
	BVLCOriginalUnicastNPDU
	BVLCOriginalBroadcastNPDU
)

type MaxAPDU int

const (
	MaxAPDU50 MaxAPDU = iota
	MaxAPDU128
	MaxAPDU206
	MaxAPDU480
	MaxAPDU1024
	MaxAPDU1476
)

// BVLL for BACnet/IPv4
const BVLLTypeBACnetIP = 0x81

// BACnet Virtual Link Control
type BVLC struct {
	BVLLTypeBACnetIP byte
	BVLCHeaderLength byte
	BVLCMaxAPDU      MaxAPDU
}

// NewBVLC creates a new BVLC with the given transport.
func NewBVLC(transprt transport.Transport) (*BVLC, error) {
	bvlc := &BVLC{
		BVLLTypeBACnetIP: BVLLTypeBACnetIP,
		BVLCHeaderLength: 4,
	}

	if transprt == transport.IP {
		bvlc.BVLCMaxAPDU = MaxAPDU1476
	} else {
		return nil, errUnsupportedTransport
	}

	return bvlc, nil
}

// Decode decodes incoming buffer.
// TODO support other functions.
func (bvlc *BVLC) Decode(buffer []byte, offset int) (int, BVLCFunctions, uint16, error) {
	msgType := buffer[0]
	function := BVLCFunctions(buffer[1])
	msgLength := binary.BigEndian.Uint16(buffer[2:4])

	if msgType != bvlc.BVLLTypeBACnetIP || msgLength != uint16(len(buffer)) {
		fmt.Println(msgType, bvlc.BVLLTypeBACnetIP, msgLength, uint16(len(buffer)))
		return 0, 0, 0, errUnsupportedTransport
	}

	switch function {
	case BVLCResult:
		return 4, function, msgLength, nil
	case BVLCOriginalUnicastNPDU:
		return 4, function, msgLength, nil
	case BVLCOriginalBroadcastNPDU:
		return 4, function, msgLength, nil
	case BVLCForwardedNPDU:
		// Handle this case
	case BVLCDistributeBroadcastToNetwok:
		// Handle this case
	case BVLCRegisterForeignDevice:
		// Handle this case
	case BVLCReadForeignDeviceTable:
		// Handle this case
	case BVLCDeleteForeignDeviceTableENTRY:
		// Handle this case
	case BVLCReadBroadcastDistTable:
		// Handle this case
	case BVLCWriteBroadcastDistributionTable:
		// Handle this case
	default:
		return -1, 0, 0, errUnsupportedFunction
	}

	return 0, 0, 0, errors.New("todo")
}

// Encode encodes the bvlc header to []byte.
func (bvlc BVLC) Encode(function BVLCFunctions, msgLength uint16) []byte {
	b := make([]byte, 4)
	b[0] = bvlc.BVLLTypeBACnetIP
	b[1] = byte(function)
	binary.BigEndian.PutUint16(b[2:4], msgLength)
	return b
}
