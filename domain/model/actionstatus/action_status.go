package actionstatus

import (
	"database/sql"
	"errors"
)

type ActionStatus string

const (
	Unacted ActionStatus = "Unacted"
	Acted   ActionStatus = "Acted"
)

var stringToStatus = map[string]ActionStatus{
	string(Unacted): Unacted,
	string(Acted):   Acted,
}

var statusToStringForHuman = map[ActionStatus]string{
	Unacted: "未行動",
	Acted:   "行動済み",
}

var (
	ErrorInvalidActionStatus = errors.New("invalid_action_status")
)

func New(str string) (ActionStatus, error) {
	if v, ok := stringToStatus[str]; ok {
		return v, nil
	}
	return "", ErrorInvalidActionStatus
}

func Must(str string) ActionStatus {
	v, err := New(str)
	if err != nil {
		panic(err)
	}
	return v
}

func (s ActionStatus) String() string {
	return string(s)
}

func (s ActionStatus) NullString() sql.NullString {
	return sql.NullString{
		String: s.String(),
		Valid:  true,
	}
}

func (s ActionStatus) StringForHuman() string {
	if v, ok := statusToStringForHuman[s]; ok {
		return v
	}
	return ""
}
