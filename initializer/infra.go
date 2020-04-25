package initializer

import (
	"wolfbot/domain/interfaces"
	"wolfbot/infra/rdb"
	"wolfbot/infra/uuidgen"

	"github.com/jinzhu/gorm"
)

type Infra struct {
	VillageRepository            interfaces.VillageRepository
	PlayerRepository             interfaces.PlayerRepository
	UserPlayerRelationRepository interfaces.UserPlayerRelationRepository
	UUIDGenerator                interfaces.UUIDGenerator
}

func InitInfra(db *gorm.DB) Infra {
	villageRepository := rdb.NewVillageRepository(db)
	playerRepository := rdb.NewPlayerRepository(db)
	userPlayerRelationRepository := rdb.NewUserPlayerRelationRepository(db)

	uuidGenerator := uuidgen.NewUUIDGenerator()

	return Infra{
		VillageRepository:            villageRepository,
		PlayerRepository:             playerRepository,
		UserPlayerRelationRepository: userPlayerRelationRepository,
		UUIDGenerator:                uuidGenerator,
	}
}
