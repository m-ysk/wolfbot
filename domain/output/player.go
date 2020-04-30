package output

import (
	"fmt"
	"wolfbot/domain/model"
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

type PlayerCheckStateForWolf struct {
	Role           roles.Role
	OtherWolfNames []string
}

func (o PlayerCheckStateForWolf) Reply() string {
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
