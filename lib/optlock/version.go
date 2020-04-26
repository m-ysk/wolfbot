package optlock

import "database/sql"

type Version int64

func (v Version) Int64() int64 {
	return int64(v)
}

func (v Version) NullInt64WithIncrement() sql.NullInt64 {
	return sql.NullInt64{
		Int64: v.Int64() + 1,
		Valid: true,
	}
}
