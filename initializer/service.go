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
	)

	playerService := service.NewPlayerService(
		infra.PlayerRepository,
		infra.GameRepository,
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
