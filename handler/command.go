package handler

import (
	"errors"
)

type command struct {
	Action action
	Target string
}

var (
	ErrorInvalidAction = errors.New("invalid_action")
)

func newGroupCommand(
	action string,
	target string,
) command {
	if v, ok := groupActionMap[action]; ok {
		return command{
			Action: v,
			Target: target,
		}
	}

	return newActionNoneCommand()
}

func newActionNoneCommand() command {
	return command{
		Action: actionNone,
		Target: "",
	}
}

type action string

const (
	actionNone                  action = "None"
	actionCheckGroupState       action = "CheckGroupState"
	actionCreateVillage         action = "CreateVillage"
	actionCreateVillageForDebug action = "CreateVillageForDebug"
	actionDeleteVillage         action = "DeleteVillage"
	actionJoinVillage           action = "JoinVillage"
	actionAddPlayersForDebug    action = "AddPlayersForDebug"
	actionFinishRecruiting      action = "FinishRecruiting"
	actionConfigureCasting      action = "ConfigureCasting"
)

var groupActionMap = map[string]action{
	"村作成":     actionCreateVillage,
	"デバッグ村作成": actionCreateVillageForDebug,
	"村削除":     actionDeleteVillage,
	"参加":      actionJoinVillage,
	"デバッグ参加":  actionAddPlayersForDebug,
	"募集終了":    actionFinishRecruiting,
	"配役設定":    actionConfigureCasting,
}
