package output

import (
	"fmt"
	"wolfbot/domain/model"
	"wolfbot/domain/model/roles"
)

type PlayerCheckState struct{}

func (o PlayerCheckState) Reply() string {
	return "現在はゲームの開始前です。グループトーク内でゲームの設定を行ってください"
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
