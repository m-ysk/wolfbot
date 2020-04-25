package service

import (
	"wolfbot/domain/interfaces"
	"wolfbot/domain/model"
	"wolfbot/domain/model/gamestatus"
	"wolfbot/domain/output"
)

type VillageService struct {
	villageRepository            interfaces.VillageRepository
	playerRepository             interfaces.PlayerRepository
	userPlayerRelationRepository interfaces.UserPlayerRelationRepository
	uuidGenerator                interfaces.UUIDGenerator
}

func NewVillageService(
	villageRepository interfaces.VillageRepository,
	playerRepository interfaces.PlayerRepository,
	userPlayerRelationRepository interfaces.UserPlayerRelationRepository,
	uuidGenerator interfaces.UUIDGenerator,
) VillageService {
	return VillageService{
		villageRepository:            villageRepository,
		playerRepository:             playerRepository,
		userPlayerRelationRepository: userPlayerRelationRepository,
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
				Status:          "",
			}, nil
		}

		return output.VillageCheckStatus{}, err
	}

	return output.VillageCheckStatus{
		VillageNotExist: false,
		Status:          village.Status,
	}, nil
}

func (s VillageService) Create(
	id model.VillageID,
) (output.VillageCreate, error) {
	village := model.NewVillage(id)

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
	if _, ok := relations.FindByVillageID(villageID); ok {
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

	return output.VillageAddPlayer{}, nil
}
