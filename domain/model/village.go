package model

import (
	"database/sql"
	"wolfbot/domain/model/debug"
	"wolfbot/domain/model/gamestatus"
	"wolfbot/lib/unixtime"
)

type Village struct {
	ID        VillageID
	Status    gamestatus.GameStatus
	Debug     debug.Mode
	CreatedAt unixtime.UnixTime
	UpdatedAt unixtime.UnixTime
}

func NewVillage(id VillageID, debug debug.Mode) Village {
	return Village{
		ID:        id,
		Status:    gamestatus.RecruitingPlayers,
		Debug:     debug,
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

func (v *Village) IsDebug() bool {
	if v == nil {
		return false
	}
	return v.Debug == debug.Debug
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

func (id VillageID) NullString() sql.NullString {
	return sql.NullString{
		String: id.String(),
		Valid:  true,
	}
}

type VillageStatus string

func (s VillageStatus) String() string {
	return string(s)
}

func (s VillageStatus) NullString() sql.NullString {
	return sql.NullString{
		String: s.String(),
		Valid:  true,
	}
}
