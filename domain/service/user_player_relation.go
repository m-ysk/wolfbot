package service

import (
	"errors"
	"wolfbot/domain/interfaces"
	"wolfbot/domain/model"
	"wolfbot/lib/errorwr"
)

type UserPlayerRelationService struct {
	userPlayerRelationRepository interfaces.UserPlayerRelationRepository
}

func NewUserPlayerRelationService(
	userPlayerRelationRepository interfaces.UserPlayerRelationRepository,
) UserPlayerRelationService {
	return UserPlayerRelationService{
		userPlayerRelationRepository: userPlayerRelationRepository,
	}
}

func (s UserPlayerRelationService) GetPlayerIDByUserIDAndVillageID(
	userID model.UserID,
	villageID model.VillageID,
) (model.PlayerID, error) {
	relation, err := s.userPlayerRelationRepository.FindByUserIDAndVillageID(userID, villageID)
	if err != nil {
		return "", err
	}
	return relation.PlayerID, nil
}

func (s UserPlayerRelationService) GetOneOrErrByUserID(
	userID model.UserID,
) (model.UserPlayerRelation, error) {
	relations, err := s.userPlayerRelationRepository.FindByUserID(userID)
	if err != nil {
		return model.UserPlayerRelation{}, err
	}

	if len(relations) == 0 {
		return model.UserPlayerRelation{}, errorwr.New(
			errors.New("user_not_joining_village"),
			"あなたは現在、人狼ゲームに参加していません",
		)
	}

	if len(relations) > 1 {
		return model.UserPlayerRelation{}, errorwr.New(
			errors.New("user_must_designate_player_name"),
			`あなたは複数の村に参加中です。
複数の村に同時に参加している場合、以下のように自分のプレイヤー名を指定してください。

【あなたの名前が「あーさー」で、「らんすろっと」さんに投票する場合の例】
らんすろっと＠投票／あーさー`,
		)
	}

	return relations[0], nil
}

func (s UserPlayerRelationService) GetByUserIDAndPlayerName(
	userID model.UserID,
	playerName model.PlayerName,
) (model.UserPlayerRelation, error) {
	return s.userPlayerRelationRepository.FindByUserIDAndPlayerName(userID, playerName)
}
