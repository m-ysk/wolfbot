package model

import (
	"database/sql"
	"time"
	"wolfbot/domain/model/actionstatus"
	"wolfbot/domain/model/lifestatus"
	"wolfbot/domain/model/roles"
	"wolfbot/domain/model/votestatus"
	"wolfbot/lib/optlock"
)

type Player struct {
	ID           PlayerID
	VillageID    VillageID
	Name         PlayerName
	LifeStatus   lifestatus.LifeStatus
	Role         roles.Role
	ActionStatus actionstatus.ActionStatus
	ActTo        PlayerID
	VoteStatus   votestatus.VoteStatus
	VoteTo       PlayerID
	Version      optlock.Version
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func NewPlayer(
	id PlayerID,
	villageID VillageID,
	name PlayerName,
) Player {
	return Player{
		ID:           id,
		VillageID:    villageID,
		Name:         name,
		LifeStatus:   lifestatus.Alive,
		Role:         roles.Must(roles.Unassigned.String()),
		ActionStatus: actionstatus.Unacted,
		VoteStatus:   votestatus.Unvoted,
		Version:      0,
	}
}

func (p *Player) Alive() bool {
	if p == nil {
		return false
	}
	return p.LifeStatus == lifestatus.Alive
}

func (p *Player) Acted() bool {
	if p == nil {
		return false
	}
	return p.ActionStatus == actionstatus.Acted
}

func (p *Player) Act(to PlayerID) {
	if p == nil {
		return
	}
	p.ActionStatus = actionstatus.Acted
	p.ActTo = to
}

func (p *Player) Unact() {
	if p == nil {
		return
	}
	p.ActionStatus = actionstatus.Unacted
	p.ActTo = ""
}

func (p *Player) Vote(target PlayerID) {
	if p == nil {
		return
	}
	p.VoteStatus = votestatus.Voted
	p.VoteTo = target
}

func (p *Player) Unvote() {
	if p == nil {
		return
	}
	p.VoteStatus = votestatus.Unvoted
	p.VoteTo = ""
}

func (p *Player) Kill() {
	if p == nil {
		return
	}
	p.LifeStatus = lifestatus.Dead
}

type Players []Player

func (ps Players) Bite(target PlayerID) {
	for i, v := range ps {
		if v.Role.Bitable() {
			ps[i].Act(target)
		}
	}
}

func (ps Players) FilterUnacted() Players {
	var result Players
	for _, v := range ps {
		if !v.Acted() {
			result = append(result, v)
		}
	}
	return result
}

func (ps Players) FilterBitable() Players {
	var result Players
	for _, v := range ps {
		if v.Role.Bitable() {
			result = append(result, v)
		}
	}
	return result
}

func (ps Players) Names() PlayerNames {
	var result PlayerNames
	for _, v := range ps {
		result = append(result, v.Name)
	}
	return result
}

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

func (ps Players) CountAliveForJudge() int {
	count := 0
	for _, v := range ps {
		if v.LifeStatus == lifestatus.Alive && !v.Role.UncountableForJudge() {
			count++
		}
	}
	return count
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

func (ps Players) CountWolf() int {
	count := 0
	for _, v := range ps {
		if v.Role.WolfCountable() {
			count++
		}
	}
	return count
}

func (ps Players) CountUnacted() int {
	count := 0
	for _, v := range ps {
		if v.ActionStatus == actionstatus.Unacted {
			count++
		}
	}
	return count
}

func (ps Players) CountUnvoted() int {
	count := 0
	for _, v := range ps {
		if v.VoteStatus == votestatus.Unvoted {
			count++
		}
	}
	return count
}

func (ps Players) FindByID(id PlayerID) (Player, bool) {
	for _, v := range ps {
		if v.ID == id {
			return v, true
		}
	}
	return Player{}, false
}

func (ps Players) FindByName(name PlayerName) (Player, bool) {
	for _, v := range ps {
		if v.Name == name {
			return v, true
		}
	}
	return Player{}, false
}

func (ps Players) UpdatePlayer(player Player) {
	for i, v := range ps {
		if v.ID == player.ID {
			ps[i] = player
		}
	}
}

type PlayerID string

type PlayerIDs []PlayerID

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

type PlayerNames []PlayerName

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
