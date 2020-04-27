package model

import (
	"database/sql"
	"time"
	"wolfbot/domain/model/debug"
	"wolfbot/domain/model/gamestatus"
	"wolfbot/domain/model/regulation"
	"wolfbot/domain/model/roles"
	"wolfbot/lib/optlock"
)

type Village struct {
	ID         VillageID
	Status     gamestatus.GameStatus
	Casting    roles.Casting
	Regulation regulation.Regulation
	Day        int
	Debug      debug.Mode
	Version    optlock.Version
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func NewVillage(id VillageID, debug debug.Mode) Village {
	return Village{
		ID:         id,
		Status:     gamestatus.RecruitingPlayers,
		Casting:    make(roles.Casting),
		Regulation: regulation.NewByDefault(),
		Day:        0,
		Debug:      debug,
		Version:    0,
	}
}

func (v *Village) UpdateStatus(status gamestatus.GameStatus) {
	if v == nil {
		return
	}
	v.Status = status
}

func (v *Village) UpdateCasting(casting roles.Casting) {
	if v == nil {
		return
	}
	v.Casting = casting
}

func (v *Village) IsDebug() bool {
	if v == nil {
		return false
	}
	return v.Debug == debug.Debug
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
