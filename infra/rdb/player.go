package rdb

import (
	"wolfbot/domain/model"
	"wolfbot/domain/model/lifestatus"
	"wolfbot/domain/model/role"
	"wolfbot/lib/unixtime"
)

type Player struct {
	ID         string `sql:"primary_key;not null"`
	VillageID  string `sql:"not null;default:''"`
	Name       string `sql:"not null;default:''"`
	LifeStatus string `sql:"not null;default:''"`
	Role       string `sql:"not null;default:''"`
	CreatedAt  int64  `sql:"not null;default:0"`
	UpdatedAt  int64  `sql:"not null;default:0"`
}

func NewPlayer(player model.Player) Player {
	return Player{
		ID:         player.ID.String(),
		VillageID:  player.VillageID.String(),
		Name:       player.Name.String(),
		LifeStatus: player.LifeStatus.String(),
		Role:       player.Role.String(),
		CreatedAt:  player.CreatedAt.Int64(),
		UpdatedAt:  player.UpdatedAt.Int64(),
	}
}

func (p Player) Model() (model.Player, error) {
	lifeStatus, err := lifestatus.New(p.LifeStatus)
	if err != nil {
		return model.Player{}, err
	}

	roleVal, err := role.New(p.Role)
	if err != nil {
		return model.Player{}, err
	}

	return model.Player{
		ID:         model.PlayerID(p.ID),
		VillageID:  model.VillageID(p.VillageID),
		Name:       model.PlayerName(p.Name),
		LifeStatus: lifeStatus,
		Role:       roleVal,
		CreatedAt:  unixtime.UnixTime(p.CreatedAt),
		UpdatedAt:  unixtime.UnixTime(p.UpdatedAt),
	}, nil
}
