package output

import (
	"fmt"
	"wolfbot/domain/model"
	"wolfbot/domain/model/gamestatus"
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

type VillageAddPlayersForDebug struct {
	Number int
}

func (o VillageAddPlayersForDebug) Reply() string {
	return fmt.Sprintf("%v人の参加者を登録しました", o.Number)
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
です。

設定したい配役を以下のように発言してください。
（残った人数には自動的に村人が設定されます）

○占い師1人、霊能者1人、狩人1人、人狼2人、残りを村人に設定する場合の例

占霊狩狼狼＠配役設定`,
		o.Game.Players.NamesForHuman(),
		role.AvailableRoleNames(),
	)
}

type VillageConfigureCasting struct {
	Casting role.Casting
}

func (o VillageConfigureCasting) Reply() string {
	return fmt.Sprintf(`以下の配役を設定します。

○配役
%v

この配役で決定する場合は、
＠はい

配役の設定をやり直す場合は、
＠いいえ

と発言してください`,
		o.Casting.StringForHuman(),
	)
}

type VillageConfirm struct {
	PrevStatus gamestatus.GameStatus
}

func (o VillageConfirm) Reply() string {
	switch st := o.PrevStatus; st {
	case gamestatus.ConfirmingCasting:
		return fmt.Sprintf(`配役を設定しました。
続いて、ルールを設定します。

ルールはデフォルトで以下のように設定されています。
[1]投票結果が同数の場合：再投票
[2]初日の占い：人狼ではない人物をランダムで通知
[3]狩人の連続ガード：不可
[4]狩人の自分ガード：不可
ルールを変更する場合は、

（変更したい項目の番号）＠変更

と入力してください。
例えば、「狩人の連続ガード」を変更する場合のコマンドは、
3＠変更
となります。
このままの設定でゲームを開始する場合は、

＠設定終了

と入力してください。`)

	default:
		return ""
	}
}
