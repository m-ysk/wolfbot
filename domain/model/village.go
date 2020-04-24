package model

import (
	"errors"
	"wolfbot/domain/model/gamestatus"
	"wolfbot/lib/unixtime"
)

type Village struct {
	ID        GroupID
	Status    gamestatus.GameStatus
	CreatedAt unixtime.UnixTime
	UpdatedAt unixtime.UnixTime
}

var (
	ErrorVillageNotFound = errors.New("village_not_found")
)

func IsVillageNotFound(err error) bool {
	return err == ErrorVillageNotFound
}

func NewVillage(id GroupID) Village {
	return Village{
		ID:        id,
		Status:    gamestatus.RecruitingPlayers,
		CreatedAt: unixtime.Now(),
		UpdatedAt: unixtime.Now(),
	}
}

type GroupID string

func (id GroupID) String() string {
	return string(id)
}

type VillageStatus string

func (s VillageStatus) String() string {
	return string(s)
}
