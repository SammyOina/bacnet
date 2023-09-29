package internal

import "errors"

var (
	errBitArrayLen = errors.New("bit array length must be 8 to convert to byte")
	errOutOfBounds = errors.New("index is out of the range for the bitarray")
)

// BitArray defines an array of bits.
type BitArray struct {
	bits []bool
}

// Creates a new bit array for the provided length.
func NewBitArray(length int) *BitArray {
	return &BitArray{
		bits: make([]bool, length),
	}
}

// Set sets the value of the bit array at the given index and value.
func (ba *BitArray) Set(index int, value bool) error {
	if index >= 0 && index < len(ba.bits) {
		ba.bits[index] = value
		return nil
	}
	return errOutOfBounds
}

// Get returns the the value of the bit array at the given index.
func (ba *BitArray) Get(index int) (bool, error) {
	if index >= 0 && index < len(ba.bits) {
		return ba.bits[index], nil
	}
	return false, errOutOfBounds
}

// ToByte converts bitarray to byte value.
// TODO return byte array.
func (ba *BitArray) ToByte() (byte, error) {
	// Ensure the length of the bit array is 8 to convert to a byte.
	if len(ba.bits) != 8 {
		return 0, errBitArrayLen
	}

	var byteValue byte
	for i := 0; i < 8; i++ {
		if ba.bits[i] {
			byteValue |= 1 << uint(7-i)
		}
	}

	return byteValue, nil
}

// NewBitArrayFromByte creates a new bit array from the given byte.
// TODO check length of incoming byte.
func NewBitArrayFromByte(byteValue byte) *BitArray {
	bitArray := NewBitArray(8)
	for j := 0; j < 8; j++ {
		bitArray.Set(j, byteValue&(1<<uint(7-j)) != 0)
	}
	return bitArray
}
