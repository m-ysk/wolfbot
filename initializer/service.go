package initializer

import "wolfbot/domain/service"

type Service struct {
	VillageService service.VillageService
}

func InitService(infra Infra) Service {
	villageService := service.NewVillageService(
		infra.VillageRepository,
		infra.PlayerRepository,
	)

	return Service{
		VillageService: villageService,
	}
}
