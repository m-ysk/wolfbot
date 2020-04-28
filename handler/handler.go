package handler

import (
	"strings"
	"wolfbot/domain/model"
	"wolfbot/domain/service"
)

type MessageHandler struct {
	villageService            service.VillageService
	playerService             service.PlayerService
	userPlayerRelationService service.UserPlayerRelationService
}

func NewMessageHandler(
	villageService service.VillageService,
	playerService service.PlayerService,
	userPlayerRelationService service.UserPlayerRelationService,
) MessageHandler {
	return MessageHandler{
		villageService:            villageService,
		playerService:             playerService,
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

	case actionCreateVillageForDebug:
		return h.villageService.CreateForDebug(villageID)

	case actionDeleteVillage:
		return h.villageService.Delete(villageID)

	case actionJoinVillage:
		return h.villageService.AddPlayer(villageID, userID, cmd.Target)

	case actionAddPlayersForDebug:
		return h.villageService.AddPlayersForDebug(villageID, userID, cmd.Target)

	case actionFinishRecruiting:
		return h.villageService.FinishRecruiting(villageID)

	case actionConfigureCasting:
		return h.villageService.ConfigureCasting(villageID, cmd.Target)

	case actionFinishConfiguringRegulation:
		return h.villageService.FinishConfiguringRegulation(villageID)

	case actionStartGame:
		return h.villageService.StartGame(villageID)

	case actionConfirm:
		return h.villageService.Confirm(villageID)

	case actionReject:
		return h.villageService.Reject(villageID)
	}

	_, err := h.userPlayerRelationService.GetPlayerIDByUserIDAndVillageID(
		userID,
		villageID,
	)
	if err != nil {
		return NoReplyOutput{}, err
	}

	switch cmd.Action {
	}

	panic("unreachable")
}

func (h MessageHandler) HandleUserMessage(
	message string,
	userID model.UserID,
) (Output, error) {
	cmd, playerName := parseUserMessage(message)

	switch cmd.Action {
	case actionNone:
		return NoReplyOutput{}, nil
	}

	var relation model.UserPlayerRelation
	var err error
	if playerName == "" {
		relation, err = h.userPlayerRelationService.GetOneOrErrByUserID(userID)
	} else {
		relation, err = h.userPlayerRelationService.GetByUserIDAndPlayerName(
			userID,
			model.PlayerName(playerName),
		)
	}
	if err != nil {
		return NoReplyOutput{}, err
	}

	playerID := relation.PlayerID
	villageID := relation.VillageID

	switch cmd.Action {
	case actionCheckUserState:
		return h.playerService.CheckState(playerID, villageID)
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

func parseUserMessage(
	message string,
) (command, string) {
	replacedMsg := strings.ReplaceAll(message, "＠", "@")
	replacedMsg = strings.ReplaceAll(replacedMsg, "／", "/")

	splitSlash := strings.Split(replacedMsg, "/")

	var msg, userName string
	switch len(splitSlash) {
	case 1:
		msg = splitSlash[0]
	case 2:
		msg = splitSlash[0]
		userName = splitSlash[1]
	default:
		return newActionNoneCommand(), ""
	}

	if msg == "@" {
		return command{
			Action: actionCheckUserState,
			Target: "",
		}, userName
	}

	splitMsg := strings.Split(msg, "@")

	if len(splitMsg) != 2 {
		return newActionNoneCommand(), ""
	}

	return newUserCommand(
		splitMsg[1],
		splitMsg[0],
	), userName
}

type Output interface {
	Reply() string
}

type NoReplyOutput struct{}

func (o NoReplyOutput) Reply() string {
	return ""
}
