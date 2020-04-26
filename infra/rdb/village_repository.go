package rdb

import (
	"errors"
	"log"
	"wolfbot/domain/interfaces"
	"wolfbot/domain/model"

	"github.com/jinzhu/gorm"
)

type villageRepository struct {
	db *gorm.DB
}

var _ interfaces.VillageRepository = villageRepository{}

func NewVillageRepository(db *gorm.DB) villageRepository {
	return villageRepository{db: db}
}

func (repo villageRepository) Create(village model.Village) error {
	v := NewVillage(village)
	if err := repo.db.Create(&v).Error; err != nil {
		return err
	}
	return nil
}

func (repo villageRepository) Delete(id model.VillageID) error {
	tx := repo.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			Rollback(tx)
		}
	}()

	if err := tx.Delete(&Village{
		ID: id.String(),
	}).Error; err != nil {
		log.Println(err)
		Rollback(tx)
		return err
	}

	if err := tx.Delete(&Player{
		VillageID: id.String(),
	}).Error; err != nil {
		log.Println(err)
		Rollback(tx)
		return err
	}

	if err := tx.Delete(&UserPlayerRelation{
		VillageID: id.String(),
	}).Error; err != nil {
		log.Println(err)
		Rollback(tx)
		return err
	}

	if err := tx.Commit().Error; err != nil {
		log.Println(err)
		Rollback(tx)
		return err
	}

	return nil
}

func (repo villageRepository) FindByID(
	id model.VillageID,
) (model.Village, error) {
	var v Village

	result := repo.db.Where(&Village{
		ID: id.String(),
	}).First(&v)
	if result.RecordNotFound() {
		return model.Village{}, NewErrorNotFound(
			errors.New("village not found: id: " + id.String()),
		)
	}
	if err := result.Error; err != nil {
		return model.Village{}, err
	}

	return v.MustModel(), nil
}
