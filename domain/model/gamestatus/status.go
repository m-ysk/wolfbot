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

	// ルール設定中
	ConfiguringRegulation GameStatus = "ConfiguringRegulation"

	// 役職確認中
	CheckingRole GameStatus = "CheckingRole"

	// 昼
	Daytime GameStatus = "Daytime"

	// 昼を終了するか確認中
	ConfirmingFinishDaytime GameStatus = "ConfirmingFinishDaytime"
)

var stringToStatus = map[string]GameStatus{
	string(RecruitingPlayers):       RecruitingPlayers,
	string(ConfiguringCasting):      ConfiguringCasting,
	string(ConfirmingCasting):       ConfirmingCasting,
	string(ConfiguringRegulation):   ConfiguringRegulation,
	string(CheckingRole):            CheckingRole,
	string(Daytime):                 Daytime,
	string(ConfirmingFinishDaytime): ConfirmingFinishDaytime,
}

var statusToStringForHuman = map[GameStatus]string{
	RecruitingPlayers:       "参加者募集中",
	ConfiguringCasting:      "配役設定中",
	ConfirmingCasting:       "配役設定結果の確認中",
	ConfiguringRegulation:   "ルール設定中",
	CheckingRole:            "役職確認中",
	Daytime:                 "昼（議論と投票）",
	ConfirmingFinishDaytime: "昼の終了確認中",
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
