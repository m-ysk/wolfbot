package model

import "database/sql"

type UserID string

func (id UserID) String() string {
	return string(id)
}

func (id UserID) NullString() sql.NullString {
	return sql.NullString{
		String: id.String(),
		Valid:  true,
	}
}
