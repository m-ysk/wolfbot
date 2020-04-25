package rdb

import "wolfbot/domain/model"

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
