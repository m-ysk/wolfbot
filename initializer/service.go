package initializer

import "wolfbot/domain/service"

type Service struct {
	VillageService            service.VillageService
	PlayerService             service.PlayerService
	UserPlayerRelationService service.UserPlayerRelationService
}

func InitService(infra Infra) Service {
	villageService := service.NewVillageService(
		infra.VillageRepository,
		infra.PlayerRepository,
		infra.UserPlayerRelationRepository,
		infra.GameRepository,
		infra.UUIDGenerator,
		infra.RandomGenerator,
	)

	playerService := service.NewPlayerService(
		infra.PlayerRepository,
		infra.GameRepository,
		infra.RandomGenerator,
	)

	userPlayerRelationService := service.NewUserPlayerRelationService(
		infra.UserPlayerRelationRepository,
	)

	return Service{
		VillageService:            villageService,
		PlayerService:             playerService,
		UserPlayerRelationService: userPlayerRelationService,
	}
}
