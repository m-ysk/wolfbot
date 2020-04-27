package rdb

import (
	"database/sql"
	"encoding/json"
	"time"
	"wolfbot/domain/model"
	"wolfbot/domain/model/debug"
	"wolfbot/domain/model/gamestatus"
	"wolfbot/domain/model/regulation"
	"wolfbot/domain/model/roles"
	"wolfbot/lib/optlock"
)

type Village struct {
	ID         sql.NullString `sql:"primary_key;type:varchar;size:255;not null"`
	Status     sql.NullString `sql:"type:varchar;size:255;not null;default:''"`
	Casting    sql.NullString `sql:"type:varchar;size:5000;not null;default:''"`
	Regulation sql.NullString `sql:"type:varchar;size:5000;not null;default:''"`
	Debug      sql.NullString `sql:"type:varchar;size:255;not null;default:''"`
	Version    sql.NullInt64  `sql:"not null;default:0"`
	CreatedAt  time.Time      `sql:"not null"`
	UpdatedAt  time.Time      `sql:"not null"`
}

func MustVillage(village model.Village) Village {
	return Village{
		ID:         village.ID.NullString(),
		Status:     village.Status.NullString(),
		Casting:    village.Casting.MustNullString(),
		Regulation: village.Regulation.MustNullString(),
		Debug:      village.Debug.NullString(),
		Version:    village.Version.NullInt64WithIncrement(),
		CreatedAt:  village.CreatedAt,
		UpdatedAt:  village.UpdatedAt,
	}
}

func (v Village) MustModel() model.Village {
	var casting roles.Casting
	if c := v.Casting.String; c != "" {
		if err := json.Unmarshal([]byte(c), &casting); err != nil {
			panic(err)
		}
	} else {
		casting = make(roles.Casting)
	}

	var reg regulation.Regulation
	if r := v.Regulation.String; r != "" {
		if err := json.Unmarshal([]byte(r), &reg); err != nil {
			panic(err)
		}
	} else {
		reg = regulation.NewByDefault()
	}

	return model.Village{
		ID:         model.VillageID(v.ID.String),
		Status:     gamestatus.Must(v.Status.String),
		Casting:    casting,
		Regulation: reg,
		Debug:      debug.Must(v.Debug.String),
		Version:    optlock.Version(v.Version.Int64),
		CreatedAt:  v.CreatedAt,
		UpdatedAt:  v.UpdatedAt,
	}
}

func (v Village) CurrentVersion() sql.NullInt64 {
	return sql.NullInt64{
		Int64: v.Version.Int64 - 1,
		Valid: true,
	}
}
