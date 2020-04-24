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
	userID model.UserID,
	groupID model.GroupID,
) (string, error) {
	command := parseGroupMessage(message, userID, groupID)

	switch command.Action {
	case actionNone:
		return "", nil

	case actionCheckGroupState:
		output, err := h.villageService.CheckStatus(groupID)
		if err != nil {
			return "", err
		}
		return output.StringForHuman(), nil

	case actionCreateVillage:
		if err := h.villageService.Create(groupID); err != nil {
			return "", err
		}
		return "村を作成しました", nil

	case actionDeleteVillage:
		if err := h.villageService.Delete(groupID); err != nil {
			return "", err
		}
		return "村を削除しました", nil
	}

	panic("unreachable")
}

func parseGroupMessage(
	message string,
	userID model.UserID,
	groupID model.GroupID,
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
