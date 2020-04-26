package rdb

import (
	"database/sql"
	"wolfbot/domain/model"
	"wolfbot/domain/model/debug"
	"wolfbot/domain/model/gamestatus"
	"wolfbot/lib/optlock"
	"wolfbot/lib/unixtime"
)

type Village struct {
	ID        sql.NullString `sql:"primary_key;not null"`
	Status    sql.NullString `sql:"not null;default:''"`
	Debug     sql.NullString `sql:"not null;default:''"`
	Version   sql.NullInt64  `sql:"not null;default:0"`
	CreatedAt sql.NullInt64  `sql:"not null;default:0"`
	UpdatedAt sql.NullInt64  `sql:"not null;default:0"`
}

func NewVillage(village model.Village) Village {
	return Village{
		ID:        village.ID.NullString(),
		Status:    village.Status.NullString(),
		Debug:     village.Debug.NullString(),
		Version:   village.Version.NullInt64WithIncrement(),
		CreatedAt: village.CreatedAt.NullInt64(),
		UpdatedAt: village.UpdatedAt.NullInt64(),
	}
}

func (v Village) MustModel() model.Village {
	return model.Village{
		ID:        model.VillageID(v.ID.String),
		Status:    gamestatus.Must(v.Status.String),
		Debug:     debug.Must(v.Debug.String),
		Version:   optlock.Version(v.Version.Int64),
		CreatedAt: unixtime.UnixTime(v.CreatedAt.Int64),
		UpdatedAt: unixtime.UnixTime(v.UpdatedAt.Int64),
	}
}
