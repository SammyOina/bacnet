package bacnet

import (
	"time"

	"github.com/absmach/bacnet/pkg/encoding"
)

type DateTime struct {
	DateTime time.Time
}

func (dt *DateTime) Decode(buf []byte, offset int) int {
	len, date := encoding.DecodeApplicationDate(buf, offset)
	len1, ttime := encoding.DecodeApplicationTime(buf, offset+len)
	dt.DateTime = ttime.AddDate(date.Year(), int(date.Month()), date.Day())
	return len + len1
}

func (dt DateTime) Encode() []byte {
	return append(encoding.EncodeApplicationDate(dt.DateTime), encoding.EncodeApplicationTime(dt.DateTime)...)
}
