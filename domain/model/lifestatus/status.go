package lifestatus

import "errors"

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

func (s LifeStatus) String() string {
	return string(s)
}

func (s LifeStatus) StringForHuman() string {
	if v, ok := statusToStringForHuman[s]; ok {
		return v
	}
	return ""
}
