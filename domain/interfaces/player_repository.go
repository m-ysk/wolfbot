package interfaces

import "wolfbot/domain/model"

type PlayerRepository interface {
	Create(player model.Player) error
	Delete(id model.PlayerID) error
	FindByID(id model.PlayerID) (model.Player, error)
}
