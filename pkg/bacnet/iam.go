package bacnet

import (
	"errors"

	"github.com/absmach/bacnet/pkg/encoding"
)

type IAmRequest struct {
	IamDeviceIdentifier   ObjectIdentifier
	MaxAPDULengthAccepted uint32
	SegmentationSupported interface{}
	VendorID              uint32
}

func (iam *IAmRequest) Decode(buffer []byte, offset int) (int, error) {
	leng := 0
	iam.IamDeviceIdentifier = ObjectIdentifier{}
	// OBJECT ID - object id
	leng1, tagNumber, lenValue, err := encoding.DecodeTagNumberAndValue(buffer, offset+leng)
	if err != nil {
		return -1, err
	}
	leng += leng1

	if tagNumber != byte(BACnetObjectIdentifier) {
		return -1, errors.New("Invalid tag number")
	}

	leng1, err = iam.IamDeviceIdentifier.Decode(buffer, offset+leng, int(lenValue))
	if err != nil {
		return -1, err
	}
	leng += leng1

	if iam.IamDeviceIdentifier.Type != encoding.ObjectTypeDevice {
		return -1, errors.New("Got Iam from no device")
	}

	// MAX APDU - unsigned
	leng1, tagNumber, lenValue, err = encoding.DecodeTagNumberAndValue(buffer, offset+leng)
	if err != nil {
		return -1, err
	}
	leng += leng1
	if tagNumber != byte(UnsignedInt) {
		return -1, errors.New("Invalid tag number")
	}

	leng1, decodedValue, err := encoding.DecodeUnsigned(buffer, offset+leng, int(lenValue))
	if err != nil {
		return -1, err
	}
	leng += leng1
	iam.MaxAPDULengthAccepted = decodedValue

	// Segmentation - enumerated
	leng1, tagNumber, lenValue, err = encoding.DecodeTagNumberAndValue(buffer, offset+leng)
	if err != nil {
		return -1, err
	}
	leng += leng1
	if tagNumber != byte(Enumerated) {
		return -1, errors.New("Invalid tag number")
	}
	propID := encoding.SegmentationSupported
	leng1, segmentationSupported, err := encoding.DecodeEnumerated(buffer, offset+leng, lenValue, nil, &propID)
	if err != nil {
		return -1, err
	}
	leng += leng1
	iam.SegmentationSupported = segmentationSupported

	// Vendor ID - unsigned16
	leng1, tagNumber, lenValue, err = encoding.DecodeTagNumberAndValue(buffer, offset+leng)
	if err != nil {
		return -1, err
	}
	leng += leng1

	if tagNumber != byte(UnsignedInt) {
		return -1, errors.New("Invalid tag number")
	}

	leng1, decodedValue, err = encoding.DecodeUnsigned(buffer, offset+leng, int(lenValue))
	if err != nil {
		return -1, err
	}

	leng += leng1
	if decodedValue > 0xFFFF {
		return -1, errors.New("Value exceeds 0xFFFF")
	}
	iam.VendorID = decodedValue

	return leng, nil
}

func (iam IAmRequest) Encode() []byte {
	tmp := iam.IamDeviceIdentifier.Encode()
	propID := iam.SegmentationSupported.(encoding.PropertyIdentifier)
	return append(append(append(append(encoding.EncodeTag(encoding.BACnetApplicationTag(BACnetObjectIdentifier), false, len(tmp)), tmp...), encoding.EncodeApplicationUnsigned(iam.MaxAPDULengthAccepted)...), encoding.EncodeApplicationEnumerated(uint32(propID))...), encoding.EncodeApplicationUnsigned(iam.VendorID)...)
}

type YouAreRequest struct {
	VendorID         uint32
	ModelName        string
	SerialNumber     string
	DeviceIdentifier ObjectIdentifier
	DeviceMACAddress []byte
}

func (youAre *YouAreRequest) Decode(buffer []byte, offset int, apduLen int) (int, error) {
	leng := 0

	leng1, tagNumber, lenValue, err := encoding.DecodeTagNumberAndValue(buffer, offset+leng)
	if err != nil {
		return -1, err
	}
	leng += leng1
	if tagNumber == byte(UnsignedInt) {
		leng1, decodedValue, err := encoding.DecodeUnsigned(buffer, offset+leng, int(lenValue))
		if err != nil {
			return -1, err
		}
		leng += leng1
		youAre.VendorID = decodedValue
	} else {
		return -1, errors.New("Invalid tag number")
	}

	leng1, tagNumber, lenValue, err = encoding.DecodeTagNumberAndValue(buffer, offset+leng)
	if err != nil {
		return -1, err
	}
	leng += leng1
	if tagNumber == byte(CharacterString) {
		leng1, decodedValue := encoding.DecodeCharacterString(buffer, offset+leng, apduLen-leng, int(lenValue))

		leng += leng1
		youAre.ModelName = decodedValue
	} else {
		return -1, errors.New("Invalid tag number")
	}

	leng1, tagNumber, lenValue, err = encoding.DecodeTagNumberAndValue(buffer, offset+leng)
	if err != nil {
		return -1, err
	}
	leng += leng1
	if tagNumber == byte(CharacterString) {
		leng1, decodedValue := encoding.DecodeCharacterString(buffer, offset+leng, apduLen-leng, int(lenValue))
		leng += leng1
		youAre.SerialNumber = decodedValue
	} else {
		return -1, errors.New("Invalid tag number")
	}

	if leng < apduLen {
		leng1, tagNumber, lenValue, err = encoding.DecodeTagNumberAndValue(buffer, offset+leng)
		if err != nil {
			return -1, err
		}
		if tagNumber == byte(BACnetObjectIdentifier) {
			leng += leng1
			youAre.DeviceIdentifier = ObjectIdentifier{}
			leng1, err = youAre.DeviceIdentifier.Decode(buffer, offset+leng, int(lenValue))
			if err != nil {
				return -1, err
			}
			leng += leng1
		}
	}

	if leng < apduLen {
		leng1, tagNumber, lenValue, err = encoding.DecodeTagNumberAndValue(buffer, offset+leng)
		if err != nil {
			return -1, err
		}
		if tagNumber == byte(OctetString) {
			leng += leng1
			leng1, decodedValue := encoding.DecodeOctetString(buffer, offset+leng, int(lenValue))
			leng += leng1
			youAre.DeviceMACAddress = decodedValue
		}
	}

	return leng, nil
}

func (youAre YouAreRequest) Encode() []byte {
	buffer := append(append(append([]byte{}, encoding.EncodeApplicationUnsigned(youAre.VendorID)...),
		encoding.EncodeApplicationCharacterString(youAre.ModelName)...),
		encoding.EncodeApplicationCharacterString(youAre.SerialNumber)...)

	if youAre.DeviceIdentifier != (ObjectIdentifier{}) {
		buffer = append(buffer, youAre.DeviceIdentifier.EncodeApp()...)
	}

	if len(youAre.DeviceMACAddress) > 0 {
		buffer = append(buffer, encoding.EncodeApplicationOctetString(youAre.DeviceMACAddress, 0, len(youAre.DeviceMACAddress))...)
	}

	return buffer
}
