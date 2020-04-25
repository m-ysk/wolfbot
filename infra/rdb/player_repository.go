package rdb

import (
	"errors"
	"log"
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

func (repo playerRepository) Create(
	player model.Player,
	relation model.UserPlayerRelation,
) error {
	tx := repo.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			Rollback(tx)
		}
	}()

	p := NewPlayer(player)
	if err := tx.Create(&p).Error; err != nil {
		log.Println(err)
		Rollback(tx)
		return err
	}

	r := NewUserPlayerRelation(relation)
	if err := tx.Create(&r).Error; err != nil {
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

func (repo playerRepository) Delete(id model.PlayerID) error {
	if err := repo.db.Delete(&Player{
		ID: id.String(),
	}).Error; err != nil {
		return err
	}
	return nil
}

func (repo playerRepository) FindByID(
	id model.PlayerID,
) (model.Player, error) {
	var p Player

	result := repo.db.Where(&Player{
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
