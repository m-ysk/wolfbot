package interfaces

import (
	"wolfbot/domain/model"
)

type UserPlayerRelationRepository interface {
	Create(relation model.UserPlayerRelation) error
	Delete(userID model.UserID, villageID model.VillageID) error
	FindByUserIDAndVillageID(
		userID model.UserID,
		villageID model.VillageID,
	) (model.UserPlayerRelation, error)
	FindByUserID(userID model.UserID) (model.UserPlayerRelations, error)
}
