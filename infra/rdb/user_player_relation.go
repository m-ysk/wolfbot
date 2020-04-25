package rdb

import (
	"wolfbot/domain/model"
	"wolfbot/lib/unixtime"
)

type UserPlayerRelation struct {
	UserID     string `sql:"primary_key"`
	VillageID  string `sql:"primary_key"`
	PlayerName string `sql:"not null;default:''"`
	PlayerID   string `sql:"not null;default:''"`
	CreatedAt  int64  `sql:"not null;default:0"`
	UpdatedAt  int64  `sql:"not null;default:0"`
}

func NewUserPlayerRelation(relation model.UserPlayerRelation) UserPlayerRelation {
	return UserPlayerRelation{
		UserID:     relation.UserID.String(),
		VillageID:  relation.VillageID.String(),
		PlayerName: relation.PlayerName.String(),
		PlayerID:   relation.PlayerID.String(),
		CreatedAt:  relation.CreatedAt.Int64(),
		UpdatedAt:  relation.UpdatedAt.Int64(),
	}
}

func (r UserPlayerRelation) MustModel() model.UserPlayerRelation {
	return model.UserPlayerRelation{
		UserID:     model.UserID(r.UserID),
		VillageID:  model.VillageID(r.VillageID),
		PlayerName: model.MustPlayerName(r.PlayerName),
		PlayerID:   model.PlayerID(r.PlayerID),
		CreatedAt:  unixtime.UnixTime(r.CreatedAt),
		UpdatedAt:  unixtime.UnixTime(r.UpdatedAt),
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
