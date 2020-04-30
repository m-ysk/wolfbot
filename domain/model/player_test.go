package model

import (
	"testing"
	"wolfbot/domain/model/lifestatus"
	"wolfbot/domain/model/roles"
)

func TestPlayers_CountAliveForJudge(t *testing.T) {
	players := Players{
		{
			ID:         "a",
			LifeStatus: lifestatus.Alive,
			Role:       roles.Must(roles.Villager.String()),
		},
		{
			ID:         "b",
			LifeStatus: lifestatus.Alive,
			Role:       roles.Must(roles.Wolf.String()),
		},
		{
			ID:         "c",
			LifeStatus: lifestatus.Dead,
			Role:       roles.Must(roles.Villager.String()),
		},
	}

	if got := players.CountAliveForJudge(); got != 2 {
		t.Errorf("got %v, want %v", got, 2)
	}
}
