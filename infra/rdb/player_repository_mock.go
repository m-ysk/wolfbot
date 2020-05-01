package rdb

import (
	"errors"
	"fmt"
	"wolfbot/domain/interfaces"
	"wolfbot/domain/model"
)

type PlayerRepositoryMock struct {
	Players   *model.Players
	Relations *model.UserPlayerRelations
}

var _ interfaces.PlayerRepository = PlayerRepositoryMock{}

func (repo PlayerRepositoryMock) Create(
	player model.Player,
	relation model.UserPlayerRelation,
) error {
	for _, v := range *repo.Players {
		if v.ID == player.ID {
			return errors.New("duplicated player id: " + player.ID.String())
		}
	}
	*repo.Players = append(*repo.Players, player)

	for _, v := range *repo.Relations {
		if v.UserID == relation.UserID && v.PlayerName == relation.PlayerName {
			return fmt.Errorf("duplicated primary key: %+v", relation)
		}
	}
	*repo.Relations = append(*repo.Relations, relation)

	return nil
}

func (repo PlayerRepositoryMock) Update(player model.Player) error {
	updated := false
	for i, v := range *repo.Players {
		if v.ID == player.ID {
			(*repo.Players)[i] = player
			updated = true
		}
	}
	if !updated {
		return fmt.Errorf("failed to update player: %+v", player)
	}
	return nil
}

func (repo PlayerRepositoryMock) UpdateAll(players model.Players) error {
	for _, player := range players {
		updated := false
		for i, v := range *repo.Players {
			if v.ID == player.ID {
				(*repo.Players)[i] = player
				updated = true
			}
		}
		if !updated {
			return fmt.Errorf("failed to update player: %+v", player)
		}
	}
	return nil
}

func (repo PlayerRepositoryMock) FindByID(id model.PlayerID) (model.Player, error) {
	for _, v := range *repo.Players {
		if v.ID == id {
			return v, nil
		}
	}
	return model.Player{}, NewErrorNotFound(
		errors.New("player not found: id: " + id.String()),
	)
}

func (repo PlayerRepositoryMock) FindByVillageID(
	villageID model.VillageID,
) (model.Players, error) {
	var result model.Players
	for _, v := range *repo.Players {
		if v.VillageID == villageID {
			result = append(result, v)
		}
	}
	return result, nil
}
