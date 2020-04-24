package service

import (
	"wolfbot/domain/interfaces"
	"wolfbot/domain/model"
	"wolfbot/domain/output"
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

func (s VillageService) CheckStatus(
	id model.GroupID,
) (output.VillageCheckStatus, error) {
	village, err := s.villageRepository.FindByID(id)
	if err != nil {
		if model.IsVillageNotFound(err) {
			return output.VillageCheckStatus{
				VillageNotExist: true,
				Status:          "",
			}, nil
		}

		return output.VillageCheckStatus{}, err
	}

	return output.VillageCheckStatus{
		VillageNotExist: false,
		Status:          village.Status,
	}, nil
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
