package model

import "wolfbot/lib/unixtime"

type Village struct {
	ID        GroupID
	Status    VillageStatus
	CreatedAt unixtime.UnixTime
	UpdatedAt unixtime.UnixTime
}

func NewVillage(id GroupID) Village {
	return Village{
		ID:        id,
		Status:    "",
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
