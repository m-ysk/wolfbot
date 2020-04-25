package rdb

import (
	"errors"
	"fmt"
	"wolfbot/domain/interfaces"
	"wolfbot/domain/model"

	"github.com/jinzhu/gorm"
)

type userPlayerRelationRepository struct {
	db *gorm.DB
}

var _ interfaces.UserPlayerRelationRepository = userPlayerRelationRepository{}

func NewUserPlayerRelationRepository(db *gorm.DB) userPlayerRelationRepository {
	return userPlayerRelationRepository{db: db}
}

func (repo userPlayerRelationRepository) Create(relation model.UserPlayerRelation) error {
	r := NewUserPlayerRelation(relation)
	if err := repo.db.Create(&r).Error; err != nil {
		return err
	}
	return nil
}

func (repo userPlayerRelationRepository) Delete(userID model.UserID, groupID model.GroupID) error {
	if err := repo.db.Delete(&UserPlayerRelation{
		UserID:  userID.String(),
		GroupID: groupID.String(),
	}).Error; err != nil {
		return err
	}
	return nil
}

func (repo userPlayerRelationRepository) FindByUserIDAndGroupID(
	userID model.UserID,
	groupID model.GroupID,
) (model.UserPlayerRelation, error) {
	var r UserPlayerRelation

	result := repo.db.Where(&UserPlayerRelation{
		UserID:  userID.String(),
		GroupID: groupID.String(),
	}).First(&r)
	if result.RecordNotFound() {
		return model.UserPlayerRelation{}, NewErrorNotFound(
			errors.New(
				fmt.Sprintf(
					"user_player_relation_not_found: user_id: %v, group_id: %v",
					userID.String(),
					groupID.String(),
				),
			),
		)
	}
	if err := result.Error; err != nil {
		return model.UserPlayerRelation{}, err
	}

	return r.MustModel(), nil
}

func (repo userPlayerRelationRepository) FindByUserID(
	userID model.UserID,
) (model.UserPlayerRelations, error) {
	var rs UserPlayerRelations

	result := repo.db.Where(&UserPlayerRelation{
		UserID: userID.String(),
	}).Find(&rs)
	if err := result.Error; err != nil {
		return nil, err
	}

	return rs.MustModel(), nil
}
