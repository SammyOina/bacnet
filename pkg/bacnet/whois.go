package bacnet

import (
	"fmt"

	"github.com/absmach/bacnet/pkg/encoding"
)

type WhoIs struct {
	HighLimit, LowLimit *uint32
}

func (w *WhoIs) Decode(buf []byte, offset, apduLen int) (int, error) {
	if apduLen <= 0 {
		return 0, fmt.Errorf("invalid apdu length")
	}
	length, tagNum, lenVal, err := encoding.DecodeTagNumberAndValue(buf, offset)
	if err != nil {
		return -1, err
	}
	if tagNum != 0 {
		return -1, fmt.Errorf("invalid tag number")
	}
	if apduLen > length {
		len1, decVal, err := encoding.DecodeUnsigned(buf, offset+length, int(lenVal))
		if err != nil {
			return -1, err
		}
		length += len1

		if decVal <= encoding.MaxInstance {
			w.LowLimit = &decVal
		}

		if apduLen > length {
			len1, tagNum, lenVal, err := encoding.DecodeTagNumberAndValue(buf, offset+length)
			if err != nil {
				return -1, err
			}
			length += len1
			if tagNum != 1 {
				return -1, fmt.Errorf("invalid tag number")
			}
			if apduLen > length {
				len1, decVal, err := encoding.DecodeUnsigned(buf, offset+length, int(lenVal))
				if err != nil {
					return -1, err
				}
				length += len1
				if decVal <= encoding.MaxInstance {
					w.HighLimit = &decVal
				}
			} else {
				return -1, fmt.Errorf("apdu lenth greater than message lenth")
			}
		} else {
			return -1, fmt.Errorf("apdu lenth greater than message lenth")
		}
	} else {
		return -1, fmt.Errorf("apdu lenth greater than message lenth")
	}
	return length, nil
}

func (w WhoIs) Encode() []byte {
	var res []byte
	if w.LowLimit != nil && *w.LowLimit <= encoding.MaxInstance {
		res = append(res, encoding.EncodeContextUnsigned(0, *w.LowLimit)...)
	}
	if w.HighLimit != nil && *w.HighLimit <= encoding.MaxInstance {
		res = append(res, encoding.EncodeContextUnsigned(1, *w.HighLimit)...)
	}
	return res
}
