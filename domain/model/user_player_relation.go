package model

import (
	"time"
	"wolfbot/lib/optlock"
)

type UserPlayerRelation struct {
	UserID     UserID
	VillageID  VillageID
	PlayerName PlayerName
	PlayerID   PlayerID
	Version    optlock.Version
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func NewUserPlayerRelation(
	userID UserID,
	villageID VillageID,
	playerName PlayerName,
	playerID PlayerID,
) UserPlayerRelation {
	return UserPlayerRelation{
		UserID:     userID,
		VillageID:  villageID,
		PlayerName: playerName,
		PlayerID:   playerID,
		Version:    0,
	}
}

type UserPlayerRelations []UserPlayerRelation

func (rs UserPlayerRelations) FindByVillageID(
	villageID VillageID,
) (UserPlayerRelation, bool) {
	for _, v := range rs {
		if v.VillageID == villageID {
			return v, true
		}
	}
	return UserPlayerRelation{}, false
}

func (rs UserPlayerRelations) FindByPlayerName(
	name PlayerName,
) (UserPlayerRelation, bool) {
	for _, v := range rs {
		if v.PlayerName == name {
			return v, true
		}
	}
	return UserPlayerRelation{}, false
}
