package initializer

import "wolfbot/domain/service"

type Service struct {
	VillageService            service.VillageService
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

	userPlayerRelationService := service.NewUserPlayerRelationService(
		infra.UserPlayerRelationRepository,
	)

	return Service{
		VillageService:            villageService,
		UserPlayerRelationService: userPlayerRelationService,
	}
}
