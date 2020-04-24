package gamestatus

import "errors"

type GameStatus string

const (
	RecruitingPlayers GameStatus = "RecruitingPlayers"
)

var stringToStatus = map[string]GameStatus{
	string(RecruitingPlayers): RecruitingPlayers,
}

var statusToStringForHuman = map[GameStatus]string{
	RecruitingPlayers: "参加者募集中",
}

var (
	ErrorInvalidGameStatus = errors.New("invalid_game_status")
)

func New(str string) (GameStatus, error) {
	if v, ok := stringToStatus[str]; ok {
		return v, nil
	}
	return "", ErrorInvalidGameStatus
}

func (s GameStatus) String() string {
	return string(s)
}

func (s GameStatus) StringForHuman() string {
	if v, ok := statusToStringForHuman[s]; ok {
		return v
	}
	return ""
}
