package rdb

import (
	"wolfbot/domain/model"
	"wolfbot/domain/model/gamestatus"
	"wolfbot/lib/unixtime"
)

type Village struct {
	ID        string `sql:"primary_key;not null"`
	Status    string `sql:"not null;default:''"`
	CreatedAt int64  `sql:"not null;default:0"`
	UpdatedAt int64  `sql:"not null;default:0"`
}

func NewVillage(village model.Village) Village {
	return Village{
		ID:        village.ID.String(),
		Status:    village.Status.String(),
		CreatedAt: village.CreatedAt.Int64(),
		UpdatedAt: village.UpdatedAt.Int64(),
	}
}

func (v Village) MustModel() model.Village {
	return model.Village{
		ID:        model.VillageID(v.ID),
		Status:    gamestatus.Must(v.Status),
		CreatedAt: unixtime.UnixTime(v.CreatedAt),
		UpdatedAt: unixtime.UnixTime(v.UpdatedAt),
	}
}
