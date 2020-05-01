package rdb

import (
	"errors"
	"wolfbot/domain/interfaces"
	"wolfbot/domain/model"
)

type VillageRepositoryMock struct {
	Villages *[]model.Village
}

var _ interfaces.VillageRepository = VillageRepositoryMock{}

func (repo VillageRepositoryMock) Create(village model.Village) error {
	for _, v := range *repo.Villages {
		if v.ID == village.ID {
			return errors.New("duplicated_village_id")
		}
	}
	*repo.Villages = append(*repo.Villages, village)
	return nil
}

func (repo VillageRepositoryMock) Delete(id model.VillageID) error {
	var newVillages []model.Village
	for _, v := range *repo.Villages {
		if v.ID != id {
			newVillages = append(newVillages, v)
		}
	}
	*repo.Villages = newVillages
	return nil
}

func (repo VillageRepositoryMock) FindByID(id model.VillageID) (model.Village, error) {
	for _, v := range *repo.Villages {
		if v.ID == id {
			return v, nil
		}
	}
	return model.Village{}, NewErrorNotFound(errors.New("village not found: id:" + id.String()))
}
