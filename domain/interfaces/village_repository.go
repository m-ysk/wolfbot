package interfaces

import "wolfbot/domain/model"

type VillageRepository interface {
	Create(village model.Village) error
	Delete(id model.GroupID) error
	FindByID(id model.GroupID) (model.Village, error)
}
