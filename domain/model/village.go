package model

import (
	"wolfbot/domain/model/gamestatus"
	"wolfbot/lib/unixtime"
)

type Village struct {
	ID        VillageID
	Status    gamestatus.GameStatus
	CreatedAt unixtime.UnixTime
	UpdatedAt unixtime.UnixTime
}

func NewVillage(id VillageID) Village {
	return Village{
		ID:        id,
		Status:    gamestatus.RecruitingPlayers,
		CreatedAt: unixtime.Now(),
		UpdatedAt: unixtime.Now(),
	}
}

func (v *Village) UpdateStatus(status gamestatus.GameStatus) {
	if v == nil {
		return
	}
	v.Status = status
	v.updateTimestamp()
}

func (v *Village) updateTimestamp() {
	if v == nil {
		return
	}
	v.UpdatedAt = unixtime.Now()
}

type VillageID string

func (id VillageID) String() string {
	return string(id)
}

type VillageStatus string

func (s VillageStatus) String() string {
	return string(s)
}
