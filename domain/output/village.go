package output

import (
	"fmt"
	"wolfbot/domain/model"
	"wolfbot/domain/model/role"
)

type VillageCheckStatus struct {
	VillageNotExist bool
	Village         model.Village
	Players         model.Players
}

func (o VillageCheckStatus) Reply() string {
	if o.VillageNotExist {
		return `○現在の状況
村が作成されていません

○ガイド
このグループで人狼ゲームを行う場合は、
＠村作成
と入力してください`
	}

	return fmt.Sprintf(`○現在の状況
%v

○参加者
%v`,
		o.Village.Status.StringForHuman(),
		o.Players.NamesForHuman(),
	)
}

type VillageCreate struct{}

func (o VillageCreate) Reply() string {
	return `村を作成しました。
この村に参加するプレイヤーは、
（プレイヤー名）＠参加
と発言してください`
}

type VillageDelete struct{}

func (o VillageDelete) Reply() string {
	return "村を削除しました"
}

type VillageAddPlayer struct {
	PlayerName model.PlayerName
}

func (o VillageAddPlayer) Reply() string {
	return fmt.Sprintf(`プレイヤー名：%v
で村に参加しました`,
		o.PlayerName.String())
}

type VillageFinishRecruiting struct {
	Game model.Game
}

func (o VillageFinishRecruiting) Reply() string {
	return fmt.Sprintf(`参加者の募集を締め切りました。

○参加者
%v

続いて、配役を設定します。

設定可能な役職は、
%v
です。`,
		o.Game.Players.NamesForHuman(),
		role.AvailableRoleNames(),
	)
}
