package initializer

import (
	"wolfbot/domain/interfaces"
	"wolfbot/infra/rdb"

	"github.com/jinzhu/gorm"
)

type Infra struct {
	VillageRepository interfaces.VillageRepository
}

func InitInfra(db *gorm.DB) Infra {
	villageRepository := rdb.NewVillageRepository(db)

	return Infra{
		VillageRepository: villageRepository,
	}
}
