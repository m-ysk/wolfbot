package unixtime

import (
	"database/sql"
	"time"
)

type UnixTime int64

func Now() UnixTime {
	return UnixTime(time.Now().Unix())
}

func (ut UnixTime) Int64() int64 {
	return int64(ut)
}

func (ut UnixTime) NullInt64() sql.NullInt64 {
	return sql.NullInt64{
		Int64: ut.Int64(),
		Valid: true,
	}
}
