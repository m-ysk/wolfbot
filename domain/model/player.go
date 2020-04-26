package model

import (
	"database/sql"
	"wolfbot/domain/model/lifestatus"
	"wolfbot/domain/model/role"
	"wolfbot/lib/optlock"
	"wolfbot/lib/unixtime"
)

type Player struct {
	ID         PlayerID
	VillageID  VillageID
	Name       PlayerName
	LifeStatus lifestatus.LifeStatus
	Role       role.Role
	Version    optlock.Version
	CreatedAt  unixtime.UnixTime
	UpdatedAt  unixtime.UnixTime
}

func NewPlayer(
	id PlayerID,
	villageID VillageID,
	name PlayerName,
) Player {
	return Player{
		ID:         id,
		VillageID:  villageID,
		Name:       name,
		LifeStatus: lifestatus.Alive,
		Role:       role.Must(role.Unassigned.String()),
		Version:    0,
		CreatedAt:  unixtime.Now(),
		UpdatedAt:  unixtime.Now(),
	}
}

type Players []Player

func (ps Players) NamesForHuman() string {
	var result string
	for i, v := range ps {
		if i != 0 {
			result += "\n"
		}
		result += v.Name.String()
	}
	return result
}

func (ps Players) Count() int {
	return len(ps)
}

type PlayerID string

func (id PlayerID) String() string {
	return string(id)
}

func (id PlayerID) NullString() sql.NullString {
	return sql.NullString{
		String: id.String(),
		Valid:  true,
	}
}

type PlayerName string

func NewPlayerName(name string) (PlayerName, error) {
	return PlayerName(name), nil
}

func MustPlayerName(name string) PlayerName {
	n, err := NewPlayerName(name)
	if err != nil {
		panic(err)
	}
	return n
}

func (n PlayerName) String() string {
	return string(n)
}

func (n PlayerName) NullString() sql.NullString {
	return sql.NullString{
		String: n.String(),
		Valid:  true,
	}
}
