package interfaces

import "wolfbot/domain/model"

type VillageRepository interface {
	Create(village model.Village) error
	Delete(id model.VillageID) error
	FindByID(id model.VillageID) (model.Village, error)
}
