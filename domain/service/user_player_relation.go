package service

import (
	"wolfbot/domain/interfaces"
	"wolfbot/domain/model"
)

type UserPlayerRelationService struct {
	userPlayerRelationRepository interfaces.UserPlayerRelationRepository
}

func NewUserPlayerRelationService(
	userPlayerRelationRepository interfaces.UserPlayerRelationRepository,
) UserPlayerRelationService {
	return UserPlayerRelationService{
		userPlayerRelationRepository: userPlayerRelationRepository,
	}
}

func (s UserPlayerRelationService) GetPlayerIDByUserIDAndVillageID(
	userID model.UserID,
	villageID model.VillageID,
) (model.PlayerID, error) {
	relation, err := s.userPlayerRelationRepository.FindByUserIDAndVillageID(userID, villageID)
	if err != nil {
		return "", err
	}
	return relation.PlayerID, nil
}
