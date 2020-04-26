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
	GameRepository               interfaces.GameRepository
	UUIDGenerator                interfaces.UUIDGenerator
}

func InitInfra(db *gorm.DB) Infra {
	villageRepository := rdb.NewVillageRepository(db)
	playerRepository := rdb.NewPlayerRepository(db)
	userPlayerRelationRepository := rdb.NewUserPlayerRelationRepository(db)
	gameRepository := rdb.NewGameRepository(db)

	uuidGenerator := uuidgen.NewUUIDGenerator()

	return Infra{
		VillageRepository:            villageRepository,
		PlayerRepository:             playerRepository,
		UserPlayerRelationRepository: userPlayerRelationRepository,
		GameRepository:               gameRepository,
		UUIDGenerator:                uuidGenerator,
	}
}
