package handler

import (
	"strings"
	"wolfbot/domain/model"
	"wolfbot/domain/service"
)

type MessageHandler struct {
	villageService service.VillageService
}

func NewMessageHandler(
	villageService service.VillageService,
) MessageHandler {
	return MessageHandler{
		villageService: villageService,
	}
}

func (h MessageHandler) HandleGroupMessage(
	message string,
	userID model.PlayerID,
	groupID model.VillageID,
) (Output, error) {
	command := parseGroupMessage(message, userID, groupID)

	switch command.Action {
	case actionNone:
		return NoReplyOutput{}, nil

	case actionCheckGroupState:
		return h.villageService.CheckStatus(groupID)

	case actionCreateVillage:
		return h.villageService.Create(groupID)

	case actionDeleteVillage:
		return h.villageService.Delete(groupID)
	}

	panic("unreachable")
}

func parseGroupMessage(
	message string,
	userID model.PlayerID,
	groupID model.VillageID,
) command {
	replacedMsg := strings.ReplaceAll(message, "＠", "@")

	if replacedMsg == "@" {
		return command{
			Action:  actionCheckGroupState,
			Target:  "",
			UserID:  userID,
			GroupID: groupID,
		}
	}

	splitMsg := strings.Split(replacedMsg, "@")

	if len(splitMsg) != 2 {
		return newActionNoneCommand()
	}

	return newGroupCommand(
		splitMsg[1],
		splitMsg[0],
		userID,
		groupID,
	)
}

type Output interface {
	Reply() string
}

type NoReplyOutput struct{}

func (o NoReplyOutput) Reply() string {
	return ""
}
