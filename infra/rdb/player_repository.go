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

func (repo playerRepository) Update(player model.Player) error {
	p := NewPlayer(player)

	result := repo.db.Model(&Player{}).Where(map[string]interface{}{
		"id":      p.ID.String,
		"version": p.CurrentVersion().Int64,
	}).Omit("id").Updates(&p)
	if err := result.Error; err != nil {
		return err
	}
	if result.RowsAffected == 0 {
		return errors.New("failed to update player: id: " + p.ID.String)
	}

	return nil
}

func (repo playerRepository) Delete(id model.PlayerID) error {
	if err := repo.db.Delete(&Player{
		ID: id.NullString(),
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
		ID: id.NullString(),
	}).First(&p)
	if result.RecordNotFound() {
		return model.Player{}, NewErrorNotFound(
			errors.New("player not found: id: " + id.String()),
		)
	}
	if err := result.Error; err != nil {
		return model.Player{}, err
	}

	return p.MustModel(), nil
}

func (repo playerRepository) FindByVillageID(
	villageID model.VillageID,
) (model.Players, error) {
	var ps Players
	if err := repo.db.Where(Player{
		VillageID: villageID.NullString(),
	}).Find(&ps).Error; err != nil {
		return nil, err
	}
	return ps.MustModel(), nil
}
