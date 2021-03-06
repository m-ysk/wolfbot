package rdb

import (
	"errors"
	"log"
	"wolfbot/domain/model"

	"github.com/jinzhu/gorm"
)

type gameRepository struct {
	db *gorm.DB
}

func NewGameRepository(db *gorm.DB) gameRepository {
	return gameRepository{db: db}
}

func (repo gameRepository) Update(game model.Game) error {
	village, players := FromGameModel(game)

	tx := repo.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			Rollback(tx)
		}
	}()

	result := tx.Model(&Village{}).Where(map[string]interface{}{
		"id":      village.ID.String,
		"version": village.CurrentVersion().Int64,
	}).Omit("id").Updates(&village)
	if err := result.Error; err != nil {
		Rollback(tx)
		return err
	}
	if result.RowsAffected == 0 {
		Rollback(tx)
		log.Println("failed to update village: id: " + village.ID.String)
		return ErrorConcurrentDBAccess
	}

	for _, p := range players {
		result := tx.Model(&Player{}).Where(map[string]interface{}{
			"id":      p.ID.String,
			"version": p.CurrentVersion().Int64,
		}).Omit("id").Updates(&p)
		if err := result.Error; err != nil {
			Rollback(tx)
			return err
		}
		if result.RowsAffected == 0 {
			Rollback(tx)
			log.Println("failed to update player: id: " + p.ID.String)
			return ErrorConcurrentDBAccess
		}
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

func (repo gameRepository) FindByVillageID(villageID model.VillageID) (model.Game, error) {
	var v Village
	result := repo.db.Where(&Village{
		ID: villageID.NullString(),
	}).First(&v)
	if result.RecordNotFound() {
		return model.Game{}, NewErrorNotFound(
			errors.New("village not found: id; " + villageID.String()),
		)
	}
	if err := result.Error; err != nil {
		return model.Game{}, err
	}

	var ps Players
	if err := repo.db.Where(&Player{
		VillageID: villageID.NullString(),
	}).Find(&ps).Error; err != nil {
		return model.Game{}, err
	}

	return MustGameModel(v, ps), nil
}
