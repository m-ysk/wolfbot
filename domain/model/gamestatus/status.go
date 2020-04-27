package gamestatus

import (
	"database/sql"
	"errors"
)

type GameStatus string

const (
	// 参加者募集中
	RecruitingPlayers GameStatus = "RecruitingPlayers"

	// 配役設定中
	ConfiguringCasting GameStatus = "ConfiguringCasting"

	// 配役設定結果の確認中
	ConfirmingCasting GameStatus = "ConfirmingCasting"
)

var stringToStatus = map[string]GameStatus{
	string(RecruitingPlayers):  RecruitingPlayers,
	string(ConfiguringCasting): ConfiguringCasting,
	string(ConfirmingCasting):  ConfirmingCasting,
}

var statusToStringForHuman = map[GameStatus]string{
	RecruitingPlayers:  "参加者募集中",
	ConfiguringCasting: "配役設定中",
	ConfirmingCasting:  "配役設定結果の確認中",
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
