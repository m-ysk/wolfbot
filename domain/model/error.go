package model

type ErrorNotFound interface {
	IsErrorNotFound() bool
}

func IsErrorNotFound(err error) bool {
	if e, ok := err.(ErrorNotFound); ok {
		return e.IsErrorNotFound()
	}
	return false
}
