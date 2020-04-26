package rdb

import (
	"database/sql"
	"wolfbot/domain/model"
	"wolfbot/lib/unixtime"
)

type UserPlayerRelation struct {
	UserID     sql.NullString `sql:"primary_key"`
	PlayerName sql.NullString `sql:"primary_key"`
	VillageID  sql.NullString `sql:"not null;default:''"`
	PlayerID   sql.NullString `sql:"not null;default:''"`
	CreatedAt  sql.NullInt64  `sql:"not null;default:0"`
	UpdatedAt  sql.NullInt64  `sql:"not null;default:0"`
}

func NewUserPlayerRelation(relation model.UserPlayerRelation) UserPlayerRelation {
	return UserPlayerRelation{
		UserID:     relation.UserID.NullString(),
		VillageID:  relation.VillageID.NullString(),
		PlayerName: relation.PlayerName.NullString(),
		PlayerID:   relation.PlayerID.NullString(),
		CreatedAt:  relation.CreatedAt.NullInt64(),
		UpdatedAt:  relation.UpdatedAt.NullInt64(),
	}
}

func (r UserPlayerRelation) MustModel() model.UserPlayerRelation {
	return model.UserPlayerRelation{
		UserID:     model.UserID(r.UserID.String),
		VillageID:  model.VillageID(r.VillageID.String),
		PlayerName: model.MustPlayerName(r.PlayerName.String),
		PlayerID:   model.PlayerID(r.PlayerID.String),
		CreatedAt:  unixtime.UnixTime(r.CreatedAt.Int64),
		UpdatedAt:  unixtime.UnixTime(r.UpdatedAt.Int64),
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
