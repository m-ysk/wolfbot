package service

import (
	"errors"
	"strconv"
	"time"
	"wolfbot/domain/interfaces"
	"wolfbot/domain/model"
	"wolfbot/domain/model/debug"
	"wolfbot/domain/model/errorwr"
	"wolfbot/domain/model/gamestatus"
	"wolfbot/domain/output"
)

type VillageService struct {
	villageRepository            interfaces.VillageRepository
	playerRepository             interfaces.PlayerRepository
	userPlayerRelationRepository interfaces.UserPlayerRelationRepository
	gameRepository               interfaces.GameRepository
	uuidGenerator                interfaces.UUIDGenerator
}

func NewVillageService(
	villageRepository interfaces.VillageRepository,
	playerRepository interfaces.PlayerRepository,
	userPlayerRelationRepository interfaces.UserPlayerRelationRepository,
	gameRepository interfaces.GameRepository,
	uuidGenerator interfaces.UUIDGenerator,
) VillageService {
	return VillageService{
		villageRepository:            villageRepository,
		playerRepository:             playerRepository,
		userPlayerRelationRepository: userPlayerRelationRepository,
		gameRepository:               gameRepository,
		uuidGenerator:                uuidGenerator,
	}
}

func (s VillageService) CheckStatus(
	id model.VillageID,
) (output.VillageCheckStatus, error) {
	village, err := s.villageRepository.FindByID(id)
	if err != nil {
		if model.IsErrorNotFound(err) {
			return output.VillageCheckStatus{
				VillageNotExist: true,
			}, nil
		}

		return output.VillageCheckStatus{}, err
	}

	players, err := s.playerRepository.FindByVillageID(id)
	if err != nil {
		return output.VillageCheckStatus{}, err
	}

	return output.VillageCheckStatus{
		VillageNotExist: false,
		Village:         village,
		Players:         players,
	}, nil
}

func (s VillageService) Create(
	id model.VillageID,
) (output.VillageCreate, error) {
	village := model.NewVillage(id, debug.Normal)

	if err := s.villageRepository.Create(village); err != nil {
		return output.VillageCreate{}, err
	}

	return output.VillageCreate{}, nil
}

func (s VillageService) CreateForDebug(
	id model.VillageID,
) (output.VillageCreate, error) {
	village := model.NewVillage(id, debug.Debug)

	if err := s.villageRepository.Create(village); err != nil {
		return output.VillageCreate{}, err
	}

	return output.VillageCreate{}, nil
}

func (s VillageService) Delete(id model.VillageID) (output.VillageDelete, error) {
	if err := s.villageRepository.Delete(id); err != nil {
		return output.VillageDelete{}, err
	}

	return output.VillageDelete{}, nil
}

func (s VillageService) AddPlayer(
	villageID model.VillageID,
	userID model.UserID,
	name string,
) (output.VillageAddPlayer, error) {
	village, err := s.villageRepository.FindByID(villageID)
	if err != nil {
		return output.VillageAddPlayer{}, err
	}

	if village.Status != gamestatus.RecruitingPlayers {
		return output.VillageAddPlayer{}, ErrorCommandUnauthorized
	}

	playerName, err := model.NewPlayerName(name)
	if err != nil {
		return output.VillageAddPlayer{}, err
	}

	relations, err := s.userPlayerRelationRepository.FindByUserID(userID)
	if err != nil {
		return output.VillageAddPlayer{}, err
	}

	// 同一Group内で同じUserが既にPlayer登録されている場合はエラー
	if _, ok := relations.FindByVillageID(villageID); !village.IsDebug() && ok {
		return output.VillageAddPlayer{}, ErrorDuplicatedPlayerInGroup
	}

	// 同一Userが同じPlayerNameで別の村に参加している場合はエラー
	if _, ok := relations.FindByPlayerName(playerName); ok {
		return output.VillageAddPlayer{}, ErrorDuplicatedPlayerNameInSameUser
	}

	playerID := model.PlayerID(s.uuidGenerator.Generate())

	player := model.NewPlayer(
		playerID,
		villageID,
		model.PlayerName(name),
	)

	newRelation := model.NewUserPlayerRelation(
		userID,
		villageID,
		playerName,
		playerID,
	)

	if err := s.playerRepository.Create(player, newRelation); err != nil {
		return output.VillageAddPlayer{}, err
	}

	return output.VillageAddPlayer{
		PlayerName: playerName,
	}, nil
}

func (s VillageService) AddPlayersForDebug(
	villageID model.VillageID,
	userID model.UserID,
	number string,
) (output.VillageAddPlayersForDebug, error) {
	num, err := strconv.Atoi(number)
	if err != nil {
		return output.VillageAddPlayersForDebug{}, err
	}

	names := []string{"あ", "い", "う", "え", "お", "か", "き", "く", "け", "こ", "さ", "し", "す", "せ", "そ"}

	for i := 0; i < num; i++ {
		if _, err := s.AddPlayer(villageID, userID, names[i]); err != nil {
			return output.VillageAddPlayersForDebug{}, err
		}
		time.Sleep(time.Millisecond * 200)
	}

	return output.VillageAddPlayersForDebug{Number: num}, nil
}

func (s VillageService) FinishRecruiting(
	villageID model.VillageID,
) (output.VillageFinishRecruiting, error) {
	game, err := s.gameRepository.FindByVillageID(villageID)
	if err != nil {
		return output.VillageFinishRecruiting{}, err
	}

	if game.Village.Status != gamestatus.RecruitingPlayers {
		return output.VillageFinishRecruiting{}, ErrorCommandUnauthorized
	}

	if game.Players.Count() < 3 {
		return output.VillageFinishRecruiting{}, errorwr.New(
			errors.New("insufficient_player_count"),
			"ゲームの開始には3人以上のプレイヤーの参加が必要です",
		)
	}

	game.Village.UpdateStatus(gamestatus.ConfiguringCasting)

	if err := s.gameRepository.Update(game); err != nil {
		return output.VillageFinishRecruiting{}, err
	}

	return output.VillageFinishRecruiting{Game: game}, nil
}
