package model

import (
	"testing"
	"wolfbot/domain/model/judge"
	"wolfbot/domain/model/lifestatus"
	"wolfbot/domain/model/roles"

	"github.com/google/go-cmp/cmp"
)

func TestGame_Judge(t *testing.T) {
	tests := []struct {
		game Game
		want judge.Judge
	}{
		{
			game: Game{
				Players: Players{
					{
						ID:         "a",
						LifeStatus: lifestatus.Alive,
						Role:       roles.Must(roles.Villager.String()),
					},
					{
						ID:         "b",
						LifeStatus: lifestatus.Alive,
						Role:       roles.Must(roles.Villager.String()),
					},
					{
						ID:         "c",
						LifeStatus: lifestatus.Alive,
						Role:       roles.Must(roles.Wolf.String()),
					},
				},
			},
			want: judge.Ongoing,
		},
		{
			game: Game{
				Players: Players{
					{
						ID:         "a",
						LifeStatus: lifestatus.Alive,
						Role:       roles.Must(roles.Villager.String()),
					},
					{
						ID:         "b",
						LifeStatus: lifestatus.Alive,
						Role:       roles.Must(roles.Villager.String()),
					},
					{
						ID:         "c",
						LifeStatus: lifestatus.Dead,
						Role:       roles.Must(roles.Wolf.String()),
					},
				},
			},
			want: judge.Villagers,
		},
		{
			game: Game{
				Players: Players{
					{
						ID:         "a",
						LifeStatus: lifestatus.Alive,
						Role:       roles.Must(roles.Villager.String()),
					},
					{
						ID:         "b",
						LifeStatus: lifestatus.Dead,
						Role:       roles.Must(roles.Villager.String()),
					},
					{
						ID:         "c",
						LifeStatus: lifestatus.Alive,
						Role:       roles.Must(roles.Wolf.String()),
					},
				},
			},
			want: judge.Wolves,
		},
	}

	for _, tt := range tests {
		if got := tt.game.Judge(); got != tt.want {
			t.Error(cmp.Diff(tt.want, got))
		}
	}
}
