package encoding

import "time"

func DecodeApplicationDate(buf []byte, offset int) (int, time.Time) {
	len, tagNum := decodeTagNumber(buf, offset)
	if tagNum == byte(Date) {
		len1, date := decodeDate(buf, offset+len)
		return len + len1, date
	}
	return -1, time.Time{}
}

func DecodeApplicationTime(buf []byte, offset int) (int, time.Time) {
	len, tagNum := decodeTagNumber(buf, offset)
	if tagNum == byte(Time) {
		len1, btime := decodeBACnetTime(buf, offset+len)
		return len + len1, btime
	}
	return -1, time.Time{}
}

func decodeDate(buf []byte, offset int) (int, time.Time) {
	year := buf[offset]
	month := buf[offset+1]
	day := buf[offset+2]
	wday := buf[offset+3]
	if month == 0xFF && day == 0xFF && wday == 0xFF && year == 0xFF {
		return 4, time.Time{}
	}
	//BACnet year starts in 1900
	return 4, time.Date(int(year)+1900, time.Month(month), int(day), 0, 0, 0, 0, nil)
}

func DecodeDateSafe(buf []byte, offset, lenVal int) (int, time.Time) {
	if lenVal != 4 {
		return lenVal, time.Time{}
	}
	return decodeDate(buf, offset)
}

func decodeBACnetTime(buf []byte, offset int) (int, time.Time) {
	hour := buf[offset]
	min := buf[offset+1]
	sec := buf[offset+2]
	hundredths := buf[offset+3]
	if hour == 0xFF && min == 0xFF && sec == 0xFF && hundredths == 0xFF {
		return 4, time.Time{}
	}
	if hundredths > 100 {
		hundredths = 0
	}
	return 4, time.Date(0, 0, 0, int(hour), int(min), int(sec), int(hundredths*10), nil)
}

func decodeBACnetTimeSafe(buf []byte, offset int, lenVal int) (int, time.Time) {
	if lenVal != 4 {
		return lenVal, time.Time{}
	}
	return decodeBACnetTime(buf, offset)
}

func EncodeApplicationDate(date time.Time) []byte {
	return append(EncodeTag(Date, false, 4), encodeBacnetDate(date)...)
}
func EncodeApplicationTime(date time.Time) []byte {
	return append(EncodeTag(Time, false, 4), encodeBacnetTime(date)...)
}

func encodeBacnetDate(date time.Time) []byte {
	data := make([]byte, 4)
	data[0] = byte(date.Year() - 1900)
	data[1] = byte(date.Month())
	data[2] = byte(date.Day())
	data[3] = byte(date.Weekday())
	return data
}

func encodeBacnetTime(date time.Time) []byte {
	data := make([]byte, 4)
	data[0] = byte(date.Hour())
	data[1] = byte(date.Minute())
	data[2] = byte(date.Second())
	data[3] = byte(date.Nanosecond() / 10)
	return data
}
