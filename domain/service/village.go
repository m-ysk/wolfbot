package service

import (
	"wolfbot/domain/interfaces"
	"wolfbot/domain/model"
)

type VillageService struct {
	villageRepository interfaces.VillageRepository
}

func NewVillageService(
	villageRepository interfaces.VillageRepository,
) VillageService {
	return VillageService{
		villageRepository: villageRepository,
	}
}

func (s VillageService) Create(id model.GroupID) error {
	village := model.NewVillage(id)

	if err := s.villageRepository.Create(village); err != nil {
		return err
	}

	return nil
}

func (s VillageService) Delete(id model.GroupID) error {
	return s.villageRepository.Delete(id)
}
