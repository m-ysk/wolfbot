package service

import (
	"errors"
	"wolfbot/domain/model/errorwr"
)

var (
	ErrorCommandUnauthorized = errorwr.New(
		errors.New("command_unauthorized"),
		"現在はこのコマンドを実行できません",
	)
)
