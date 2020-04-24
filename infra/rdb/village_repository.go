package rdb

import (
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

func (r villageRepository) Create(village model.Village) error {
	v := NewVillage(village)
	if err := r.db.Create(&v).Error; err != nil {
		return err
	}
	return nil
}

func (r villageRepository) Delete(id model.GroupID) error {
	if err := r.db.Delete(&Village{
		ID: id.String(),
	}).Error; err != nil {
		return err
	}
	return nil
}

func (r villageRepository) FindByID(
	id model.GroupID,
) (model.Village, error) {
	var v Village

	result := r.db.Where(&Village{
		ID: id.String(),
	}).First(&v)
	if result.RecordNotFound() {
		return model.Village{}, model.ErrorVillageNotFound
	}
	if err := result.Error; err != nil {
		return model.Village{}, err
	}

	return v.Model()
}
