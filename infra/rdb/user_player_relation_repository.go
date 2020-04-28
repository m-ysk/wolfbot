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

func (repo userPlayerRelationRepository) Delete(userID model.UserID, villageID model.VillageID) error {
	if err := repo.db.Delete(&UserPlayerRelation{
		UserID:    userID.NullString(),
		VillageID: villageID.NullString(),
	}).Error; err != nil {
		return err
	}
	return nil
}

func (repo userPlayerRelationRepository) FindByUserIDAndVillageID(
	userID model.UserID,
	villageID model.VillageID,
) (model.UserPlayerRelation, error) {
	var r UserPlayerRelation

	result := repo.db.Where(&UserPlayerRelation{
		UserID:    userID.NullString(),
		VillageID: villageID.NullString(),
	}).First(&r)
	if result.RecordNotFound() {
		return model.UserPlayerRelation{}, NewErrorNotFound(
			errors.New(
				fmt.Sprintf(
					"user_player_relation_not_found: user_id: %v, group_id: %v",
					userID.String(),
					villageID.String(),
				),
			),
		)
	}
	if err := result.Error; err != nil {
		return model.UserPlayerRelation{}, err
	}

	return r.MustModel(), nil
}

func (repo userPlayerRelationRepository) FindByUserIDAndPlayerName(
	userID model.UserID,
	playerName model.PlayerName,
) (model.UserPlayerRelation, error) {
	var r UserPlayerRelation

	result := repo.db.Where(&UserPlayerRelation{
		UserID:     userID.NullString(),
		PlayerName: playerName.NullString(),
	}).First(&r)
	if result.RecordNotFound() {
		return model.UserPlayerRelation{}, NewErrorNotFound(
			errors.New(
				fmt.Sprintf(
					"user_player_relation_not_found: user_id: %v, player_name: %v",
					userID.String(),
					playerName.String(),
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
		UserID: userID.NullString(),
	}).Find(&rs)
	if err := result.Error; err != nil {
		return nil, err
	}

	return rs.MustModel(), nil
}
