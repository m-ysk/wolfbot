package rdb

import (
	"errors"
	"wolfbot/domain/interfaces"
	"wolfbot/domain/model"
)

type GameRepositoryMock struct {
	Villages *[]model.Village
	Players  *model.Players
}

var _ interfaces.GameRepository = GameRepositoryMock{}

func (repo GameRepositoryMock) Update(game model.Game) error {
	villageUpdated := false
	for i, v := range *repo.Villages {
		if v.ID == game.Village.ID {
			(*repo.Villages)[i] = game.Village
			villageUpdated = true
		}
	}
	if !villageUpdated {
		return errors.New("failed to update village: id: " + game.Village.ID.String())
	}

	for _, player := range game.Players {
		playerUpdated := false
		for i, v := range *repo.Players {
			if v.ID == player.ID {
				(*repo.Players)[i] = player
				playerUpdated = true
			}
		}
		if !playerUpdated {
			return errors.New("failed to update player: id: " + player.ID.String())
		}
	}

	return nil
}

func (repo GameRepositoryMock) FindByVillageID(
	villageID model.VillageID,
) (model.Game, error) {
	var village model.Village
	villageFound := false
	for _, v := range *repo.Villages {
		if v.ID == villageID {
			village = v
			villageFound = true
		}
	}
	if !villageFound {
		return model.Game{}, NewErrorNotFound(
			errors.New("village not found: id: " + villageID.String()),
		)
	}

	var players model.Players
	for _, v := range *repo.Players {
		if v.VillageID == villageID {
			players = append(players, v)
		}
	}

	return model.Game{
		Village: village,
		Players: players,
	}, nil
}
