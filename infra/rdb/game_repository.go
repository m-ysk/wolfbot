package rdb

import (
	"errors"
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

	result := tx.Where(&Village{
		ID:      village.ID,
		Version: village.CurrentVersion(),
	}).Updates(&village)
	if result.RowsAffected == 0 {
		Rollback(tx)
		return errors.New("failed to update village: id: " + village.ID.String)
	}
	if err := result.Error; err != nil {
		Rollback(tx)
		return err
	}

	for _, p := range players {
		result := tx.Where(&Village{
			ID:      p.ID,
			Version: p.CurrentVersion(),
		}).Updates(&p)
		if result.RowsAffected == 0 {
			Rollback(tx)
			return errors.New("failed to update player: id: " + p.ID.String)
		}
		if err := result.Error; err != nil {
			Rollback(tx)
			return err
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
