package initializer

import (
	"wolfbot/domain/interfaces"
	"wolfbot/infra/rdb"

	"github.com/jinzhu/gorm"
)

type Infra struct {
	VillageRepository interfaces.VillageRepository
	PlayerRepository  interfaces.PlayerRepository
}

func InitInfra(db *gorm.DB) Infra {
	villageRepository := rdb.NewVillageRepository(db)
	playerRepository := rdb.NewPlayerRepository(db)

	return Infra{
		VillageRepository: villageRepository,
		PlayerRepository:  playerRepository,
	}
}
