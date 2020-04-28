package interfaces

import "wolfbot/domain/model"

type PlayerRepository interface {
	Create(player model.Player, relation model.UserPlayerRelation) error
	Update(player model.Player) error
	Delete(id model.PlayerID) error
	FindByID(id model.PlayerID) (model.Player, error)
	FindByVillageID(villageID model.VillageID) (model.Players, error)
}
