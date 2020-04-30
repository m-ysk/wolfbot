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

func newUserCommand(
	action string,
	target string,
) command {
	if v, ok := userActionMap[action]; ok {
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
	actionNone                        action = "None"
	actionCheckGroupState             action = "CheckGroupState"
	actionCreateVillage               action = "CreateVillage"
	actionCreateVillageForDebug       action = "CreateVillageForDebug"
	actionDeleteVillage               action = "DeleteVillage"
	actionJoinVillage                 action = "JoinVillage"
	actionAddPlayersForDebug          action = "AddPlayersForDebug"
	actionFinishRecruiting            action = "FinishRecruiting"
	actionConfigureCasting            action = "ConfigureCasting"
	actionFinishConfiguringRegulation action = "FinishConfiguringRegulation"
	actionStartGame                   action = "GameStart"
	actionFinishVoting                action = "FinishVoting"
	actionConfirm                     action = "Confirm"
	actionReject                      action = "Reject"
)

const (
	actionCheckUserState action = "CheckUserState"
	actionVote           action = "Vote"
	actionBite           action = "Bite"
)

var groupActionMap = map[string]action{
	"確認":      actionCheckGroupState,
	"村作成":     actionCreateVillage,
	"デバッグ村作成": actionCreateVillageForDebug,
	"村削除":     actionDeleteVillage,
	"参加":      actionJoinVillage,
	"デバッグ参加":  actionAddPlayersForDebug,
	"募集終了":    actionFinishRecruiting,
	"配役設定":    actionConfigureCasting,
	"設定終了":    actionFinishConfiguringRegulation,
	"村開始":     actionStartGame,
	"投票終了":    actionFinishVoting,
	"はい":      actionConfirm,
	"いいえ":     actionReject,
}

var userActionMap = map[string]action{
	"確認": actionCheckUserState,
	"投票": actionVote,
	"噛む": actionBite,
}
