package gamestatus

import (
	"database/sql"
	"errors"
)

type GameStatus string

const (
	RecruitingPlayers  GameStatus = "RecruitingPlayers"
	ConfiguringCasting GameStatus = "ConfiguringCasting"
)

var stringToStatus = map[string]GameStatus{
	string(RecruitingPlayers):  RecruitingPlayers,
	string(ConfiguringCasting): ConfiguringCasting,
}

var statusToStringForHuman = map[GameStatus]string{
	RecruitingPlayers:  "参加者募集中",
	ConfiguringCasting: "配役設定中",
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

func Must(str string) GameStatus {
	v, err := New(str)
	if err != nil {
		panic(err)
	}
	return v
}

func (s GameStatus) String() string {
	return string(s)
}

func (s GameStatus) NullString() sql.NullString {
	return sql.NullString{
		String: s.String(),
		Valid:  true,
	}
}

func (s GameStatus) StringForHuman() string {
	if v, ok := statusToStringForHuman[s]; ok {
		return v
	}
	return ""
}
