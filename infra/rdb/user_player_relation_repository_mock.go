package rdb

import (
	"fmt"
	"wolfbot/domain/interfaces"
	"wolfbot/domain/model"
)

type UserPlayerRelationRepositoryMock struct {
	Relations *model.UserPlayerRelations
}

var _ interfaces.UserPlayerRelationRepository = UserPlayerRelationRepositoryMock{}

func (repo UserPlayerRelationRepositoryMock) Create(relation model.UserPlayerRelation) error {
	for _, v := range *repo.Relations {
		if v.UserID == relation.UserID && v.PlayerName == relation.PlayerName {
			return fmt.Errorf("duplicated primary key: %+v", relation)
		}
	}
	*repo.Relations = append(*repo.Relations, relation)
	return nil
}

func (repo UserPlayerRelationRepositoryMock) FindByUserIDAndVillageID(
	userID model.UserID,
	villageID model.VillageID,
) (model.UserPlayerRelation, error) {
	panic("not implemented yet!")
}

func (repo UserPlayerRelationRepositoryMock) FindByUserIDAndPlayerName(
	userID model.UserID,
	playerName model.PlayerName,
) (model.UserPlayerRelation, error) {
	for _, v := range *repo.Relations {
		if v.UserID == userID && v.PlayerName == playerName {
			return v, nil
		}
	}
	return model.UserPlayerRelation{}, fmt.Errorf(
		"relation not found: %v, %v",
		userID,
		playerName,
	)
}

func (repo UserPlayerRelationRepositoryMock) FindByUserID(
	userID model.UserID,
) (model.UserPlayerRelations, error) {
	var result model.UserPlayerRelations
	for _, v := range *repo.Relations {
		if v.UserID == userID {
			result = append(result, v)
		}
	}
	return result, nil
}
