package unixtime

import "time"

type UnixTime int64

func Now() UnixTime {
	return UnixTime(time.Now().Unix())
}

func (ut UnixTime) Int64() int64 {
	return int64(ut)
}
