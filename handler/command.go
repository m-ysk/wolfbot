package handler

import (
	"errors"
	"wolfbot/domain/model"
)

type command struct {
	Action  action
	Target  string
	UserID  model.UserID
	GroupID model.GroupID
}

var (
	ErrorInvalidAction = errors.New("invalid_action")
)

func newGroupCommand(
	action string,
	target string,
	userID model.UserID,
	groupID model.GroupID,
) command {
	if v, ok := groupActionMap[action]; ok {
		return command{
			Action:  v,
			Target:  target,
			UserID:  userID,
			GroupID: groupID,
		}
	}

	return newActionNoneCommand()
}

func newActionNoneCommand() command {
	return command{
		Action:  actionNone,
		Target:  "",
		UserID:  "",
		GroupID: "",
	}
}

type action string

const (
	actionNone          action = "None"
	actionCreateVillage action = "CreateVillage"
	actionDeleteVillage action = "DeleteVillage"
)

var groupActionMap = map[string]action{
	"村作成": actionCreateVillage,
	"村削除": actionDeleteVillage,
}
