package rdb

import (
	"database/sql"
	"time"
	"wolfbot/domain/model"
	"wolfbot/domain/model/debug"
	"wolfbot/domain/model/gamestatus"
	"wolfbot/lib/optlock"
	"wolfbot/lib/unixtime"
)

type Village struct {
	ID        sql.NullString `sql:"primary_key;type:varchar;size:255;not null"`
	Status    sql.NullString `sql:"type:varchar;size:255;not null;default:''"`
	Debug     sql.NullString `sql:"type:varchar;size:255;not null;default:''"`
	Version   sql.NullInt64  `sql:"not null;default:0"`
	CreatedAt time.Time      `sql:"not null"`
	UpdatedAt time.Time      `sql:"not null"`
}

func NewVillage(village model.Village) Village {
	return Village{
		ID:        village.ID.NullString(),
		Status:    village.Status.NullString(),
		Debug:     village.Debug.NullString(),
		Version:   village.Version.NullInt64WithIncrement(),
		CreatedAt: time.Unix(village.CreatedAt.Int64(), 0),
		UpdatedAt: time.Unix(village.UpdatedAt.Int64(), 0),
	}
}

func (v Village) MustModel() model.Village {
	return model.Village{
		ID:        model.VillageID(v.ID.String),
		Status:    gamestatus.Must(v.Status.String),
		Debug:     debug.Must(v.Debug.String),
		Version:   optlock.Version(v.Version.Int64),
		CreatedAt: unixtime.UnixTime(v.CreatedAt.Unix()),
		UpdatedAt: unixtime.UnixTime(v.UpdatedAt.Unix()),
	}
}

func (v Village) CurrentVersion() sql.NullInt64 {
	return sql.NullInt64{
		Int64: v.Version.Int64 - 1,
		Valid: true,
	}
}
