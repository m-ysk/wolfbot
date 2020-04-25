package model

import (
	"wolfbot/lib/unixtime"
)

type UserPlayerRelation struct {
	UserID     UserID
	GroupID    GroupID
	PlayerName PlayerName
	PlayerID   PlayerID
	CreatedAt  unixtime.UnixTime
	UpdatedAt  unixtime.UnixTime
}

type UserPlayerRelations []UserPlayerRelation
