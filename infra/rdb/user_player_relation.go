package rdb

import (
	"database/sql"
	"time"
	"wolfbot/domain/model"
	"wolfbot/lib/optlock"
	"wolfbot/lib/unixtime"
)

type UserPlayerRelation struct {
	UserID     sql.NullString `sql:"primary_key;type:varchar;size:255"`
	PlayerName sql.NullString `sql:"primary_key;type:varchar;size:255"`
	VillageID  sql.NullString `sql:"type:varchar;size:255;not null;default:''"`
	PlayerID   sql.NullString `sql:"type:varchar;size:255;not null;default:''"`
	Version    sql.NullInt64  `sql:"not null;default:0"`
	CreatedAt  time.Time      `sql:"not null"`
	UpdatedAt  time.Time      `sql:"not null"`
}

func NewUserPlayerRelation(relation model.UserPlayerRelation) UserPlayerRelation {
	return UserPlayerRelation{
		UserID:     relation.UserID.NullString(),
		VillageID:  relation.VillageID.NullString(),
		PlayerName: relation.PlayerName.NullString(),
		PlayerID:   relation.PlayerID.NullString(),
		Version:    relation.Version.NullInt64WithIncrement(),
		CreatedAt:  time.Unix(relation.CreatedAt.Int64(), 0),
		UpdatedAt:  time.Unix(relation.UpdatedAt.Int64(), 0),
	}
}

func (r UserPlayerRelation) MustModel() model.UserPlayerRelation {
	return model.UserPlayerRelation{
		UserID:     model.UserID(r.UserID.String),
		VillageID:  model.VillageID(r.VillageID.String),
		PlayerName: model.MustPlayerName(r.PlayerName.String),
		PlayerID:   model.PlayerID(r.PlayerID.String),
		Version:    optlock.Version(r.Version.Int64),
		CreatedAt:  unixtime.UnixTime(r.CreatedAt.Unix()),
		UpdatedAt:  unixtime.UnixTime(r.UpdatedAt.Unix()),
	}
}

func (r UserPlayerRelation) CurrentVersion() sql.NullInt64 {
	return sql.NullInt64{
		Int64: r.Version.Int64 - 1,
		Valid: true,
	}
}

type UserPlayerRelations []UserPlayerRelation

func (rs UserPlayerRelations) MustModel() model.UserPlayerRelations {
	var result model.UserPlayerRelations
	for _, v := range rs {
		result = append(result, v.MustModel())
	}
	return result
}
