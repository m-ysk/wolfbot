package rdb

import "wolfbot/domain/model"

func FromGameModel(game model.Game) (Village, Players) {
	return NewVillage(game.Village), NewPlayers(game.Players)
}

func MustGameModel(village Village, players Players) model.Game {
	return model.Game{
		Village: village.MustModel(),
		Players: players.MustModel(),
	}
}
