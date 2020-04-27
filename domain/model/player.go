package model

import (
	"database/sql"
	"time"
	"wolfbot/domain/model/lifestatus"
	"wolfbot/domain/model/roles"
	"wolfbot/lib/optlock"
)

type Player struct {
	ID         PlayerID
	VillageID  VillageID
	Name       PlayerName
	LifeStatus lifestatus.LifeStatus
	Role       roles.Role
	Version    optlock.Version
	CreatedAt  time.Time
	UpdatedAt  time.Time
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
		Role:       roles.Must(roles.Unassigned.String()),
		Version:    0,
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

func (ps Players) CountRole(r roles.ID) int {
	count := 0
	for _, v := range ps {
		if v.Role.ID == r {
			count++
		}
	}
	return count
}

func (ps Players) WolfCount() int {
	count := 0
	for _, v := range ps {
		if v.Role.WolfCountType.WolfCountable() {
			count++
		}
	}
	return count
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
