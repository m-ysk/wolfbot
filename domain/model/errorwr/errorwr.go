package errorwr

type ErrorWithReply struct {
	Err     error
	Message string
}

func New(err error, msg string) ErrorWithReply {
	return ErrorWithReply{
		Err:     err,
		Message: msg,
	}
}

func (e ErrorWithReply) Error() string {
	return e.Err.Error()
}

func (e ErrorWithReply) Reply() string {
	return e.Message
}
