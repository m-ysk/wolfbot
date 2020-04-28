package service

import (
	"wolfbot/domain/interfaces"
	"wolfbot/domain/model"
	"wolfbot/domain/model/gamestatus"
	"wolfbot/domain/model/roles"
	"wolfbot/domain/output"
)

type PlayerService struct {
	playerRepository interfaces.PlayerRepository
	gameRepository   interfaces.GameRepository
}

func NewPlayerService(
	playerRepository interfaces.PlayerRepository,
	gameRepository interfaces.GameRepository,
) PlayerService {
	return PlayerService{
		playerRepository: playerRepository,
		gameRepository:   gameRepository,
	}
}

func (s PlayerService) CheckState(
	playerID model.PlayerID,
	villageID model.VillageID,
) (output.Interface, error) {
	game, err := s.gameRepository.FindByVillageID(villageID)
	if err != nil {
		return nil, err
	}

	player, _ := game.Players.FindByID(playerID)

	switch player.Role.ID {
	case roles.Wolf:
		return s.checkStateForWolf(game, player)

	default:
		return s.checkState(game, player)
	}
}

func (s PlayerService) checkState(
	game model.Game,
	player model.Player,
) (output.PlayerCheckState, error) {
	switch game.Village.Status {
	case gamestatus.CheckingRole:
		player.Acted()
		s.playerRepository.Update(player)

		return output.PlayerCheckState{
			Role: player.Role,
		}, nil
	}

	panic("unreachable")
}

func (s PlayerService) checkStateForWolf(
	game model.Game,
	player model.Player,
) (output.PlayerCheckStateForWolf, error) {
	var otherWolfNames []string
	for _, v := range game.Players {
		if v.Role.WolfCountType.WolfCountable() && v.Name != player.Name {
			otherWolfNames = append(otherWolfNames, v.Name.String())
		}
	}

	switch game.Village.Status {
	case gamestatus.CheckingRole:
		player.Acted()
		s.playerRepository.Update(player)

		return output.PlayerCheckStateForWolf{
			Role:           player.Role,
			OtherWolfNames: otherWolfNames,
		}, nil
	}

	panic("unreachable")
}
