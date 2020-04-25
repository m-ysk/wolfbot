package rdb

import (
	"errors"
	"wolfbot/domain/interfaces"
	"wolfbot/domain/model"

	"github.com/jinzhu/gorm"
)

type playerRepository struct {
	db *gorm.DB
}

var _ interfaces.PlayerRepository = playerRepository{}

func NewPlayerRepository(db *gorm.DB) playerRepository {
	return playerRepository{db: db}
}

func (r playerRepository) Create(player model.Player) error {
	p := NewPlayer(player)
	if err := r.db.Create(&p).Error; err != nil {
		return err
	}
	return nil
}

func (r playerRepository) Delete(id model.PlayerID) error {
	if err := r.db.Delete(&Player{
		ID: id.String(),
	}).Error; err != nil {
		return err
	}
	return nil
}

func (r playerRepository) FindByID(
	id model.PlayerID,
) (model.Player, error) {
	var p Player

	result := r.db.Where(&Player{
		ID: id.String(),
	}).First(&p)
	if result.RecordNotFound() {
		return model.Player{}, NewErrorNotFound(
			errors.New("player not found: id: " + id.String()),
		)
	}
	if err := result.Error; err != nil {
		return model.Player{}, err
	}

	return p.Model()
}
