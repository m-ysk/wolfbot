package handler

import (
	"strings"
	"wolfbot/domain/model"
	"wolfbot/domain/service"
)

type MessageHandler struct {
	villageService            service.VillageService
	userPlayerRelationService service.UserPlayerRelationService
}

func NewMessageHandler(
	villageService service.VillageService,
	userPlayerRelationService service.UserPlayerRelationService,
) MessageHandler {
	return MessageHandler{
		villageService:            villageService,
		userPlayerRelationService: userPlayerRelationService,
	}
}

func (h MessageHandler) HandleGroupMessage(
	message string,
	userID model.UserID,
	groupID string,
) (Output, error) {
	cmd := parseGroupMessage(message)

	villageID := model.VillageID(groupID)

	switch cmd.Action {
	case actionNone:
		return NoReplyOutput{}, nil

	case actionCheckGroupState:
		return h.villageService.CheckStatus(villageID)

	case actionCreateVillage:
		return h.villageService.Create(villageID)

	case actionDeleteVillage:
		return h.villageService.Delete(villageID)

	case actionJoinVillage:
		return h.villageService.AddPlayer(villageID, userID, cmd.Target)
	}

	_, err := h.userPlayerRelationService.GetPlayerIDByUserIDAndVillageID(
		userID,
		villageID,
	)
	if err != nil {
		return NoReplyOutput{}, err
	}

	panic("unreachable")
}

func parseGroupMessage(
	message string,
) command {
	replacedMsg := strings.ReplaceAll(message, "＠", "@")

	if replacedMsg == "@" {
		return command{
			Action: actionCheckGroupState,
			Target: "",
		}
	}

	splitMsg := strings.Split(replacedMsg, "@")

	if len(splitMsg) != 2 {
		return newActionNoneCommand()
	}

	return newGroupCommand(
		splitMsg[1],
		splitMsg[0],
	)
}

type Output interface {
	Reply() string
}

type NoReplyOutput struct{}

func (o NoReplyOutput) Reply() string {
	return ""
}
