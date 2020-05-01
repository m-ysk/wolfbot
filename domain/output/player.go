package output

import (
	"fmt"
	"wolfbot/domain/model"
	"wolfbot/domain/model/actionstatus"
	"wolfbot/domain/model/roles"
	"wolfbot/domain/model/votestatus"
)

type PlayerCheckState struct{}

func (o PlayerCheckState) Reply() string {
	return "現在はゲームの開始前です。グループトーク内でゲームの設定を行ってください"
}

type PlayerCheckStateDead struct{}

func (o PlayerCheckStateDead) Reply() string {
	return "残念ながらあなたは死んでしまいました。死亡したプレイヤーは一切行動できません。"
}

type PlayerCheckStateInCheckingRole struct {
	Role roles.Role
}

func (o PlayerCheckStateInCheckingRole) Reply() string {
	return fmt.Sprintf(`○あなたの役職
%v`,
		o.Role.Name,
	)
}

type PlayerCheckStateInDaytime struct {
	Role       roles.Role
	VoteStatus votestatus.VoteStatus
	VoteTo     model.PlayerName
}

func (o PlayerCheckStateInDaytime) Reply() string {
	var vote string
	if o.VoteStatus == votestatus.Unvoted {
		vote = `まだ本日の投票を行っていません。

このトーク内にて、
（投票先プレイヤー名）＠投票
と発言して投票してください`
	} else {
		vote = fmt.Sprintf(`あなたは「%v」さんに投票済みです。

投票先を変更する場合は、もう一度、
（投票先プレイヤー名）＠投票
と発言して再投票してください`,
			o.VoteTo.String(),
		)
	}

	return fmt.Sprintf(`○あなたの役職
%v

○本日の投票先
%v`,
		o.Role.Name,
		vote,
	)
}

type PlayerCheckStateInCheckinRoleForWolf struct {
	Role           roles.Role
	OtherWolfNames []string
}

func (o PlayerCheckStateInCheckinRoleForWolf) Reply() string {
	var otherWolves string
	if len(o.OtherWolfNames) == 0 {
		otherWolves = "人狼はあなた1人です"
	} else {
		for i, v := range o.OtherWolfNames {
			if i != 0 {
				otherWolves += "\n"
			}
			otherWolves += v
		}
	}

	return fmt.Sprintf(`○あなたの役職
%v

○仲間の人狼
%v`,
		o.Role.Name,
		otherWolves,
	)
}

type PlayerCheckStateInNighttimeForWolf struct {
	Role         roles.Role
	ActionStatus actionstatus.ActionStatus
	ActTo        model.PlayerName
}

func (o PlayerCheckStateInNighttimeForWolf) Reply() string {
	var bite string
	if o.ActionStatus == actionstatus.Unacted {
		bite = `まだ本日の襲撃先を指定していません。

このトークにて、
（襲撃先プレイヤー名）＠噛む
と発言して襲撃先を指定してください。

※人狼が複数名いる場合、代表者1人のみ襲撃先の設定を行ってください`
	} else {
		bite = fmt.Sprintf(`%v

※襲撃機を変更する場合、このトークにて
（襲撃先プレイヤー名）＠噛む
と発言して襲撃先を指定してください。`,
			o.ActTo,
		)
	}

	return fmt.Sprintf(`○あなたの役職
%v

○本日の襲撃先
%v`,
		o.Role.Name,
		bite,
	)
}

type PlayerCheckStateInCheckingRoleForDivinerRandomWhite struct {
	Role      roles.Role
	WhiteName model.PlayerName
}

func (o PlayerCheckStateInCheckingRoleForDivinerRandomWhite) Reply() string {
	return fmt.Sprintf(`○あなたの役職
%v

○占い結果
%vさんは人狼ではありません
（妖狐が含まれる配役の場合、妖狐でもありません）`,
		o.Role.Name,
		o.WhiteName.String(),
	)
}

type PlayerCheckStateInNighttimeForDiviner struct {
	Role    roles.Role
	Divined bool
	Target  model.PlayerName
	IsWolf  bool
}

func (o PlayerCheckStateInNighttimeForDiviner) Reply() string {
	var divineResult string
	if o.Divined {
		var wolfOrHuman string
		if o.IsWolf {
			wolfOrHuman = "人狼"
		} else {
			wolfOrHuman = "人狼ではありません"
		}
		divineResult = fmt.Sprintf(
			`あなたは「%v」さんを占い、結果は、%vでした`,
			o.Target.String(),
			wolfOrHuman,
		)
	} else {
		divineResult = `まだ本日の占いを実行していません。
占いを実行するには、このトーク内にて、

（対象のプレイヤー名）＠占う

と発言してください`
	}

	return fmt.Sprintf(`○あなたの役職
%v

○占い結果
%v`,
		o.Role.Name,
		divineResult,
	)
}

type PlayerVote struct {
	Target model.PlayerName
}

func (o PlayerVote) Reply() string {
	return fmt.Sprintf(`「%v」さんに投票しました。

投票先を変更する場合は、もう一度、
（投票先プレイヤー名）＠投票
と入力してください`,
		o.Target.String(),
	)
}

type PlayerBite struct {
	Target model.PlayerName
}

func (o PlayerBite) Reply() string {
	return fmt.Sprintf(`「%v」さんを今晩の襲撃先に設定しました。

襲撃先を変更する場合は、もう一度、
（襲撃先プレイヤー名）＠噛む
と入力してください`,
		o.Target,
	)
}

type PlayerDivine struct {
	Target model.PlayerName
	IsWolf bool
}

func (o PlayerDivine) Reply() string {
	var result string
	if o.IsWolf {
		result = "人狼"
	} else {
		result = "人狼ではありません"
	}

	return fmt.Sprintf(
		"占いの結果、「%v」さんは、%vでした。",
		o.Target,
		result,
	)
}
