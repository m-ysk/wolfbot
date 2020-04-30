package rdb

import (
	"errors"
	"wolfbot/domain/model"
	"wolfbot/lib/errorwr"
)

type ErrorNotFound struct {
	Err error
}

var _ model.ErrorNotFound = ErrorNotFound{}

func NewErrorNotFound(err error) ErrorNotFound {
	return ErrorNotFound{
		Err: err,
	}
}

func (e ErrorNotFound) Error() string {
	return e.Err.Error()
}

func (e ErrorNotFound) IsErrorNotFound() bool {
	return true
}

var (
	ErrorConcurrentDBAccess = errorwr.New(
		errors.New("concurrent_db_access"),
		"複数プレイヤーの同時アクセスによるエラーが発生しました。少し時間を置いてから再度コマンドを実行してください",
	)
)
