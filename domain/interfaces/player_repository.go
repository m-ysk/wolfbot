package interfaces

import "wolfbot/domain/model"

type PlayerRepository interface {
	Create(player model.Player, relation model.UserPlayerRelation) error
	Update(player model.Player) error
	UpdateAll(players model.Players) error
	FindByID(id model.PlayerID) (model.Player, error)
	FindByVillageID(villageID model.VillageID) (model.Players, error)
}
