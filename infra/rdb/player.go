package rdb

import (
	"database/sql"
	"wolfbot/domain/model"
	"wolfbot/domain/model/lifestatus"
	"wolfbot/domain/model/role"
	"wolfbot/lib/optlock"
	"wolfbot/lib/unixtime"
)

type Player struct {
	ID         sql.NullString `sql:"primary_key;not null"`
	VillageID  sql.NullString `sql:"not null;default:''"`
	Name       sql.NullString `sql:"not null;default:''"`
	LifeStatus sql.NullString `sql:"not null;default:''"`
	Role       sql.NullString `sql:"not null;default:''"`
	Version    sql.NullInt64  `sql:"not null;default:0"`
	CreatedAt  sql.NullInt64  `sql:"not null;default:0"`
	UpdatedAt  sql.NullInt64  `sql:"not null;default:0"`
}

func NewPlayer(player model.Player) Player {
	return Player{
		ID:         player.ID.NullString(),
		VillageID:  player.VillageID.NullString(),
		Name:       player.Name.NullString(),
		LifeStatus: player.LifeStatus.NullString(),
		Role:       player.Role.NullString(),
		Version:    player.Version.NullInt64WithIncrement(),
		CreatedAt:  player.CreatedAt.NullInt64(),
		UpdatedAt:  player.UpdatedAt.NullInt64(),
	}
}

func (p Player) MustModel() model.Player {
	return model.Player{
		ID:         model.PlayerID(p.ID.String),
		VillageID:  model.VillageID(p.VillageID.String),
		Name:       model.PlayerName(p.Name.String),
		LifeStatus: lifestatus.Must(p.LifeStatus.String),
		Role:       role.Must(p.Role.String),
		Version:    optlock.Version(p.Version.Int64),
		CreatedAt:  unixtime.UnixTime(p.CreatedAt.Int64),
		UpdatedAt:  unixtime.UnixTime(p.UpdatedAt.Int64),
	}
}

func (p Player) CurrentVersion() sql.NullInt64 {
	return sql.NullInt64{
		Int64: p.Version.Int64 - 1,
		Valid: true,
	}
}

type Players []Player

func NewPlayers(players model.Players) Players {
	var result Players
	for _, v := range players {
		result = append(result, NewPlayer(v))
	}
	return result
}

func (ps Players) MustModel() model.Players {
	var result model.Players
	for _, v := range ps {
		result = append(result, v.MustModel())
	}
	return result
}
