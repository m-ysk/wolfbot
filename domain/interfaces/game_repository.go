package interfaces

import "wolfbot/domain/model"

type GameRepository interface {
	Update(game model.Game) error
	FindByVillageID(villageID model.VillageID) (model.Game, error)
}
