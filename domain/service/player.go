package service

import (
	"errors"
	"fmt"
	"wolfbot/domain/interfaces"
	"wolfbot/domain/model"
	"wolfbot/domain/model/gamestatus"
	"wolfbot/domain/model/regulation"
	"wolfbot/domain/model/roles"
	"wolfbot/domain/output"
	"wolfbot/lib/errorwr"
)

type PlayerService struct {
	playerRepository interfaces.PlayerRepository
	gameRepository   interfaces.GameRepository
	randomGenerator  interfaces.RandomGenerator
}

func NewPlayerService(
	playerRepository interfaces.PlayerRepository,
	gameRepository interfaces.GameRepository,
	randomGenerator interfaces.RandomGenerator,
) PlayerService {
	return PlayerService{
		playerRepository: playerRepository,
		gameRepository:   gameRepository,
		randomGenerator:  randomGenerator,
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

	if !player.Alive() {
		return output.PlayerCheckStateDead{}, nil
	}

	if game.Village.Status == gamestatus.Daytime {
		voteTo, _ := game.Players.FindByID(player.VoteTo)

		return output.PlayerCheckStateInDaytime{
			Role:       player.Role,
			VoteStatus: player.VoteStatus,
			VoteTo:     voteTo.Name,
		}, nil
	}

	switch player.Role.ID {
	case roles.Wolf:
		return s.checkStateForWolf(game, player)

	case roles.Diviner:
		return s.checkStateForDiviner(game, player)

	default:
		return s.checkState(game, player)
	}
}

func (s PlayerService) checkState(
	game model.Game,
	player model.Player,
) (output.Interface, error) {
	switch game.Village.Status {
	case gamestatus.CheckingRole:
		player.Act("")
		s.playerRepository.Update(player)

		return output.PlayerCheckStateInCheckingRole{
			Role: player.Role,
		}, nil

	default:
		return output.PlayerCheckState{}, nil
	}
}

func (s PlayerService) checkStateForWolf(
	game model.Game,
	player model.Player,
) (output.Interface, error) {
	var otherWolfNames []string
	for _, v := range game.Players {
		if v.Role.WolfCountable() && v.Name != player.Name {
			otherWolfNames = append(otherWolfNames, v.Name.String())
		}
	}

	switch game.Village.Status {
	case gamestatus.CheckingRole:
		player.Act("")
		s.playerRepository.Update(player)

		return output.PlayerCheckStateInCheckinRoleForWolf{
			Role:           player.Role,
			OtherWolfNames: otherWolfNames,
		}, nil

	case gamestatus.Nighttime:
		target, _ := game.Players.FindByID(player.ActTo)

		return output.PlayerCheckStateInNighttimeForWolf{
			Role:         player.Role,
			ActionStatus: player.ActionStatus,
			ActTo:        target.Name,
		}, nil
	}

	return nil, fmt.Errorf("unreachable in checkStateForWolf: %+v", game)
}

func (s PlayerService) checkStateForDiviner(
	game model.Game,
	player model.Player,
) (output.Interface, error) {
	switch game.Village.Status {
	case gamestatus.CheckingRole:
		switch game.Village.Regulation.FirstDayDivination {
		case regulation.FirstDayDivinationRandomWhite:
			if player.Acted() {
				target, _ := game.Players.FindByID(player.ActTo)

				return output.PlayerCheckStateInCheckingRoleForDivinerRandomWhite{
					Role:      player.Role,
					WhiteName: target.Name,
				}, nil
			}

			result := game.RandomWhite(player.ID, s.randomGenerator.Intn)
			if err := s.playerRepository.Update(result.Diviner); err != nil {
				return nil, err
			}

			return output.PlayerCheckStateInCheckingRoleForDivinerRandomWhite{
				Role:      player.Role,
				WhiteName: result.WhiteName,
			}, nil
		}

	case gamestatus.Nighttime:
		if player.Acted() {
			target, _ := game.Players.FindByID(player.ActTo)

			return output.PlayerCheckStateInNighttimeForDiviner{
				Role:    player.Role,
				Divined: true,
				Target:  target.Name,
				IsWolf:  target.Role.Black(),
			}, nil
		}

		return output.PlayerCheckStateInNighttimeForDiviner{
			Role:    player.Role,
			Divined: false,
			Target:  "",
			IsWolf:  false,
		}, nil
	}

	return nil, fmt.Errorf("unreachable in checkStateForDiviner: %+v", game)
}

func (s PlayerService) Vote(
	playerID model.PlayerID,
	villageID model.VillageID,
	target string,
) (output.PlayerVote, error) {
	game, err := s.gameRepository.FindByVillageID(villageID)
	if err != nil {
		return output.PlayerVote{}, err
	}

	if game.Village.Status != gamestatus.Daytime {
		return output.PlayerVote{}, ErrorCommandUnauthorized
	}

	player, _ := game.Players.FindByID(playerID)
	if !player.Alive() {
		return output.PlayerVote{}, ErrorDeadPlayerCommandUnauthorized
	}

	targetPlayer, ok := game.Players.FindByName(model.PlayerName(target))
	if !ok {
		return output.PlayerVote{}, ErrorInvalidTargetPlayerName
	}
	if !targetPlayer.Alive() {
		return output.PlayerVote{}, ErrorDeadTargetPlayerName
	}

	player.Vote(targetPlayer.ID)

	if err := s.playerRepository.Update(player); err != nil {
		return output.PlayerVote{}, err
	}

	return output.PlayerVote{
		Target: targetPlayer.Name,
	}, nil
}

func (s PlayerService) Bite(
	playerID model.PlayerID,
	villageID model.VillageID,
	target string,
) (output.PlayerBite, error) {
	result, err := s.validateRoleCommandWithTarget(playerID, villageID, model.PlayerName(target))
	if err != nil {
		return output.PlayerBite{}, err
	}

	if !result.Actor.Role.Bitable() {
		return output.PlayerBite{}, ErrorRoleCommanndUnauthorized
	}

	bitables := result.Game.Players.FilterBitable()
	bitables.Bite(result.Target.ID)

	if err := s.playerRepository.UpdateAll(bitables); err != nil {
		return output.PlayerBite{}, err
	}

	return output.PlayerBite{
		Target: result.Target.Name,
	}, nil
}

func (s PlayerService) Divine(
	playerID model.PlayerID,
	villageID model.VillageID,
	target string,
) (output.PlayerDivine, error) {
	validated, err := s.validateRoleCommandWithTarget(playerID, villageID, model.PlayerName(target))
	if err != nil {
		return output.PlayerDivine{}, err
	}

	player := validated.Actor

	if player.Role.ID != roles.Diviner {
		return output.PlayerDivine{}, ErrorRoleCommanndUnauthorized
	}

	if player.Acted() {
		return output.PlayerDivine{}, errorwr.New(
			errors.New("cannot repeat divine"),
			"あなたは既に今晩の占いを実行しています。占い結果をもう一度確認する場合は、\n＠確認\nと入力してください",
		)
	}

	player.Act(validated.Target.ID)

	if err := s.playerRepository.Update(player); err != nil {
		return output.PlayerDivine{}, err
	}

	return output.PlayerDivine{
		Target: validated.Target.Name,
		IsWolf: validated.Target.Role.Black(),
	}, nil
}

type validateRoleCommandWithTargetResult struct {
	Game   model.Game
	Actor  model.Player
	Target model.Player
}

func (s PlayerService) validateRoleCommandWithTarget(
	playerID model.PlayerID,
	villageID model.VillageID,
	target model.PlayerName,
) (validateRoleCommandWithTargetResult, error) {
	game, err := s.gameRepository.FindByVillageID(villageID)
	if err != nil {
		return validateRoleCommandWithTargetResult{}, err
	}

	if game.Village.Status != gamestatus.Nighttime {
		return validateRoleCommandWithTargetResult{}, ErrorCommandUnauthorized
	}

	player, _ := game.Players.FindByID(playerID)
	if !player.Alive() {
		return validateRoleCommandWithTargetResult{}, ErrorDeadPlayerCommandUnauthorized
	}

	targetPlayer, ok := game.Players.FindByName(target)
	if !ok {
		return validateRoleCommandWithTargetResult{}, ErrorInvalidTargetPlayerName
	}
	if !targetPlayer.Alive() {
		return validateRoleCommandWithTargetResult{}, ErrorDeadTargetPlayerName
	}

	return validateRoleCommandWithTargetResult{
		Game:   game,
		Actor:  player,
		Target: targetPlayer,
	}, nil
}
