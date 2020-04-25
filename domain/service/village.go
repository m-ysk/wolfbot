package service

import (
	"wolfbot/domain/interfaces"
	"wolfbot/domain/model"
	"wolfbot/domain/model/gamestatus"
	"wolfbot/domain/output"
)

type VillageService struct {
	villageRepository interfaces.VillageRepository
	playerRepository  interfaces.PlayerRepository
}

func NewVillageService(
	villageRepository interfaces.VillageRepository,
	playerRepository interfaces.PlayerRepository,
) VillageService {
	return VillageService{
		villageRepository: villageRepository,
		playerRepository:  playerRepository,
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
	groupID model.GroupID,
	userID model.UserID,
	name string,
) (output.VillageAddPlayer, error) {
	village, err := s.villageRepository.FindByID(groupID.VillageID())
	if err != nil {
		return output.VillageAddPlayer{}, err
	}

	if village.Status != gamestatus.RecruitingPlayers {
		return output.VillageAddPlayer{}, ErrorCommandUnauthorized
	}

	player := model.NewPlayer(model.PlayerID(userID), groupID.VillageID(), model.PlayerName(name))

	if err := s.playerRepository.Create(player); err != nil {
		return output.VillageAddPlayer{}, err
	}

	return output.VillageAddPlayer{}, nil
}
