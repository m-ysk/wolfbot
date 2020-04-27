package model

import (
	"wolfbot/domain/model/gamestatus"
	"wolfbot/domain/model/role"
	"wolfbot/lib/randutil"
)

type Game struct {
	Village Village
	Players Players
}

func (g *Game) Start() {
	g.assignRole()
	g.Village.Status = gamestatus.CheckingRole
}

func (g *Game) assignRole() {
	if g == nil {
		return
	}

	roleIDs := g.Village.Casting.RoleIDs()

	shuffledInts := randutil.GenerateShuffledPermutation(
		g.Players.Count(),
	)

	for i, v := range shuffledInts {
		newRole := role.Must(roleIDs[v].String())
		g.Players[i].Role = newRole
	}
}
