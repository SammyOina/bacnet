package encoding

import (
	"bytes"
	"encoding/binary"
	"math"
)

const (
	MaxObject         = 0x3FF
	InstanceBits      = 22
	MaxInstance       = 0x3FFFFF
	MaxBitstringBytes = 15
	ArrayAll          = 0xFFFFFFFF
	NoPriority        = 0
	MinPriority       = 1
	MaxPriority       = 16
	MaxUint8          = 0x100
	MaxUint16         = 0x10000
	MaxUint24         = 0x1000000
)

func EncodeUnsigned(value uint32) []byte {
	switch {
	case value < MaxUint8:
		buf := make([]byte, 1)
		buf[0] = uint8(value)
		return buf
	case value < MaxUint16:
		buf := make([]byte, 2)
		binary.BigEndian.PutUint16(buf, uint16(value))
		return buf
	case value < MaxUint24:
		buf := make([]byte, 3)
		buf[0] = byte((value & 0xff0000) >> 16)
		buf[1] = byte((value & 0x00ff00) >> 8)
		buf[2] = byte(value & 0x0000ff)
		return buf
	default:
		buf := make([]byte, 4)
		binary.BigEndian.PutUint32(buf, value)
		return buf
	}
}

func EncodeSigned(value int32) []byte {
	switch {
	case value < MaxUint8:
		buf := new(bytes.Buffer)
		binary.Write(buf, binary.BigEndian, uint8(value))
		return buf.Bytes()
	case value < MaxUint16:
		buf := new(bytes.Buffer)
		binary.Write(buf, binary.BigEndian, uint16(value))
		return buf.Bytes()
	case value < MaxUint24:
		buf := make([]byte, 3)
		buf[0] = byte((value & 0xff0000) >> 16)
		buf[1] = byte((value & 0x00ff00) >> 8)
		buf[2] = byte(value & 0x0000ff)
		return buf
	default:
		buf := new(bytes.Buffer)
		binary.Write(buf, binary.BigEndian, value)
		return buf.Bytes()
	}
}

func EncodeContextUnsigned(tagNum BACnetApplicationTag, val uint32) []byte {
	len := 0
	switch {
	case val < MaxUint8:
		len = 1
	case val < MaxUint16:
		len = 2
	case val < MaxUint24:
		len = 3
	default:
		len = 4
	}
	return append(EncodeTag(tagNum, true, len), EncodeUnsigned(val)...)
}

// EncodeApplicationUnsigned encodes an unsigned integer value as a BACnet application tag.
func EncodeApplicationUnsigned(value uint32) []byte {
	tmp := EncodeUnsigned(value)
	tag := EncodeTag(UnsignedInt, false, len(tmp))
	return append(tag, tmp...)
}

func EncodeApplicationEnumerated(value uint32) []byte {
	tmp := EncodeUnsigned(value)
	return append(EncodeTag(Enumerated, false, len(tmp)), tmp...)
}

func EncodeApplicationOctetString(octetString []byte, octetOffset, octetCount int) []byte {
	tag := EncodeTag(OctetString, false, octetCount)
	octetStringSegment := octetString[octetOffset : octetOffset+octetCount]
	return append(tag, octetStringSegment...)
}

func EncodeApplicationCharacterString(value string) []byte {
	tmp := encodeBACnetCharacterString(value)
	tag := EncodeTag(CharacterString, false, len(tmp))
	return append(tag, tmp...)
}

func encodeBACnetCharacterString(value string) []byte {
	encoding := []byte{byte(CharacterUTF8)}
	encodedValue := []byte(value)
	return append(encoding, encodedValue...)
}

func EncodeApplicationBoolean(val bool) []byte {
	if val {
		return EncodeTag(Boolean, false, 1)
	}
	return EncodeTag(Boolean, false, 0)
}

func EncodeApplicationSigned(val int32) []byte {
	tmp := EncodeSigned(val)
	return append(EncodeTag(SignedInt, false, len(tmp)), tmp...)
}

func EncodeApplicationReal(val float32) []byte {
	return append(EncodeTag(Real, false, 4), encodeBACnetReal(val)...)
}

func EncodeApplicationDouble(val float64) []byte {
	return append(EncodeTag(Double, false, 8), encodeBACnetDouble(val)...)
}

func encodeBACnetReal(value float32) []byte {
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, math.Float32bits(value))
	return buf
}

func encodeBACnetDouble(value float64) []byte {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, math.Float64bits(value))
	return buf
}

func EncodeApplicationBitString(val interface{}) []byte {
	// TODO
	return nil
}
