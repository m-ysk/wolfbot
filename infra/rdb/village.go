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

func (v Village) Model() (model.Village, error) {
	gameStatus, err := gamestatus.New(v.Status)
	if err != nil {
		return model.Village{}, err
	}

	return model.Village{
		ID:        model.GroupID(v.ID),
		Status:    gameStatus,
		CreatedAt: unixtime.UnixTime(v.CreatedAt),
		UpdatedAt: unixtime.UnixTime(v.UpdatedAt),
	}, nil
}
