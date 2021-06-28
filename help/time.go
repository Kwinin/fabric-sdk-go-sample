package help

import "time"

const TokenExpire = time.Hour * 12

func FormatTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func ParseTime(value string) (time.Time, error) {
	return time.ParseInLocation("2006-01-02 15:04:05", value, time.Local)
}

func FormatTimeWithInt(sec int) string {
	return time.Unix(int64(sec), 0).Format("2006-01-02 15:04:05")
}
