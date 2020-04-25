package interfaces

import (
	"wolfbot/domain/model"
)

type UserPlayerRelationRepository interface {
	Create(relation model.UserPlayerRelation) error
	Delete(userID model.UserID, groupID model.GroupID) error
	FindByUserIDAndGroupID(
		userID model.UserID,
		groupID model.GroupID,
	) (model.UserPlayerRelation, error)
	FindByUserID(userID model.UserID) (model.UserPlayerRelations, error)
}
