package helper

import "time"

func GetNowTimestamp() *int64 {
	now := time.Now().Unix()
	return &now
}

func GetNowNanoTimestamp() *int64 {
	now := time.Now().UnixNano() / 1000000
	return &now
}

func GetNowDbDatetime() string {
	nowDate := time.Now()
	return nowDate.Format("2006-01-02 15:04:05")
}

func GetNowTimestampInt64() int64 {
	now := time.Now().Unix()
	return now
}

func GetNowTimestampInt32() int32 {
	now := time.Now().Unix()
	return int32(now)
}
