package service

import (
	"errors"
	"wolfbot/lib/errorwr"
)

var (
	ErrorCommandUnauthorized = errorwr.New(
		errors.New("command_unauthorized"),
		"現在はこのコマンドを実行できません",
	)

	ErrorRoleCommanndUnauthorized = errorwr.New(
		errors.New("role_command_unauthorized"),
		"あなたはこのコマンドを実行できません",
	)

	ErrorDeadPlayerCommandUnauthorized = errorwr.New(
		errors.New("player_dead"),
		"あなたは死亡しています。死亡したプレイヤーは一切行動できません",
	)

	ErrorDuplicatedPlayerInGroup = errorwr.New(
		errors.New("duplicated_player_in_group"),
		"あなたは既にこの村に参加しています",
	)

	ErrorDuplicatedPlayerNameInSameUser = errorwr.New(
		errors.New("duplicated_player_name_in_same_user"),
		"あなたは同じプレイヤー名で他の村に参加しています。複数の村に同時に参加する場合、異なるプレイヤー名を使用してください。",
	)

	ErrorInvalidCallToDebugFunction = errorwr.New(
		errors.New("invalid_call_to_debug_function"),
		"現在はデバッグモードではありません",
	)

	ErrorInvalidTargetPlayerName = errorwr.New(
		errors.New("invalid_target_player_name"),
		"対象プレイヤー名が誤っています。対象プレイヤー名を確認の上、もう一度入力してください",
	)
)
