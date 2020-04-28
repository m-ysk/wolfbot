package rdb

import (
	"database/sql"
	"time"
	"wolfbot/domain/model"
	"wolfbot/domain/model/actionstatus"
	"wolfbot/domain/model/lifestatus"
	"wolfbot/domain/model/roles"
	"wolfbot/lib/optlock"
)

type Player struct {
	ID           sql.NullString `sql:"primary_key;type:varchar;size:255;not null"`
	VillageID    sql.NullString `sql:"type:varchar;size:255;not null;default:''"`
	Name         sql.NullString `sql:"type:varchar;size:255;not null;default:''"`
	LifeStatus   sql.NullString `sql:"type:varchar;size:255;not null;default:''"`
	Role         sql.NullString `sql:"type:varchar;size:255;not null;default:''"`
	ActionStatus sql.NullString `sql:"type:varchar;size:255;not null;default:''"`
	ActTo        sql.NullString `sql:"type:varchar;size:255;not null;default:''"`
	Version      sql.NullInt64  `sql:"not null;default:0"`
	CreatedAt    time.Time      `sql:"not null"`
	UpdatedAt    time.Time      `sql:"not null"`
}

func NewPlayer(player model.Player) Player {
	return Player{
		ID:           player.ID.NullString(),
		VillageID:    player.VillageID.NullString(),
		Name:         player.Name.NullString(),
		LifeStatus:   player.LifeStatus.NullString(),
		Role:         player.Role.NullString(),
		ActionStatus: player.ActionStatus.NullString(),
		ActTo:        player.ActTo.NullString(),
		Version:      player.Version.NullInt64WithIncrement(),
		CreatedAt:    player.CreatedAt,
		UpdatedAt:    player.UpdatedAt,
	}
}

func (p Player) MustModel() model.Player {
	return model.Player{
		ID:           model.PlayerID(p.ID.String),
		VillageID:    model.VillageID(p.VillageID.String),
		Name:         model.PlayerName(p.Name.String),
		LifeStatus:   lifestatus.Must(p.LifeStatus.String),
		Role:         roles.Must(p.Role.String),
		ActionStatus: actionstatus.Must(p.ActionStatus.String),
		ActTo:        model.PlayerID(p.ActTo.String),
		Version:      optlock.Version(p.Version.Int64),
		CreatedAt:    p.CreatedAt,
		UpdatedAt:    p.UpdatedAt,
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
