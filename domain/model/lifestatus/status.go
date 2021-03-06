package lifestatus

import (
	"database/sql"
	"errors"
)

type LifeStatus string

const (
	Alive LifeStatus = "Alive"
	Dead  LifeStatus = "Dead"
)

var stringToStatus = map[string]LifeStatus{
	string(Alive): Alive,
	string(Dead):  Dead,
}

var statusToStringForHuman = map[LifeStatus]string{
	Alive: "生存",
	Dead:  "死亡",
}

var (
	ErrorInvalidLifeStatus = errors.New("invalid_life_status")
)

func New(str string) (LifeStatus, error) {
	if v, ok := stringToStatus[str]; ok {
		return v, nil
	}
	return "", ErrorInvalidLifeStatus
}

func Must(str string) LifeStatus {
	v, err := New(str)
	if err != nil {
		panic(err)
	}
	return v
}

func (s LifeStatus) String() string {
	return string(s)
}

func (s LifeStatus) NullString() sql.NullString {
	return sql.NullString{
		String: s.String(),
		Valid:  true,
	}
}

func (s LifeStatus) StringForHuman() string {
	if v, ok := statusToStringForHuman[s]; ok {
		return v
	}
	return ""
}
