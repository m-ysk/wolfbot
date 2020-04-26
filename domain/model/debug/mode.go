package debug

import (
	"database/sql"
	"errors"
)

type Mode string

const (
	Normal Mode = "Normal"
	Debug  Mode = "Debug"
)

var stringToMode = map[string]Mode{
	string(Normal): Normal,
	string(Debug):  Debug,
}

var ErrorInvalidDebugMode = errors.New("invalid_debug_mode")

func New(str string) (Mode, error) {
	if v, ok := stringToMode[str]; ok {
		return v, nil
	}
	return "", ErrorInvalidDebugMode
}

func Must(str string) Mode {
	v, err := New(str)
	if err != nil {
		panic(err)
	}
	return v
}

func (m Mode) String() string {
	return string(m)
}

func (m Mode) NullString() sql.NullString {
	return sql.NullString{
		String: m.String(),
		Valid:  true,
	}
}
