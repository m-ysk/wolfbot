package model

import (
	"wolfbot/lib/optlock"
	"wolfbot/lib/unixtime"
)

type UserPlayerRelation struct {
	UserID     UserID
	VillageID  VillageID
	PlayerName PlayerName
	PlayerID   PlayerID
	Version    optlock.Version
	CreatedAt  unixtime.UnixTime
	UpdatedAt  unixtime.UnixTime
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
		CreatedAt:  unixtime.Now(),
		UpdatedAt:  unixtime.Now(),
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
