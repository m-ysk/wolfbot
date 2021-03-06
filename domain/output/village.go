package output

import (
	"fmt"
	"wolfbot/domain/model"
	"wolfbot/domain/model/gamestatus"
	"wolfbot/domain/model/judge"
	"wolfbot/domain/model/roles"
)

type VillageCheckState struct {
	VillageNotExist bool
	Village         model.Village
	Players         model.Players
}

func (o VillageCheckState) Reply() string {
	if o.VillageNotExist {
		return `○現在の状況
村が作成されていません

○ガイド
このグループで人狼ゲームを行う場合は、
＠村作成
と入力してください`
	}

	guide := "特になし"
	if o.Village.Status == gamestatus.Daytime {
		if unvoted := len(o.Players.FilterUnvoted()); unvoted == 0 {
			guide = `すべてのプレイヤーが投票を終了しています。
投票を終了して処刑を実行する場合は、
＠投票終了
と発言してください`
		} else {
			guide = fmt.Sprintf("投票を行っていないプレイヤーが%v人います。", unvoted)
		}
	}

	return fmt.Sprintf(`○現在の状況
%v

○ガイド
%v

○参加者
%v`,
		o.Village.Status.StringForHuman(),
		guide,
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
		roles.AvailableRoleNames(),
	)
}

type VillageConfigureCasting struct {
	Casting roles.Casting
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

type VillageFinishConfiguringRegulation struct{}

func (o VillageFinishConfiguringRegulation) Reply() string {
	return fmt.Sprintf(`各プレイヤーに役職を割り振りました。
各プレイヤーは、【わたしへの個別トーク】にて、
＠確認
と入力して役職を確認してください。

全員が役職を確認したら、【このグループ】にて、
＠村開始
と発言してください`)
}

type VillageStartGame struct {
	WolfCount int
}

func (o VillageStartGame) Reply() string {
	return fmt.Sprintf(`ある朝のこと、村の村長が無残に喰い殺された姿で発見されました。
どうやら、村人の中に恐ろしい人狼が%v匹、紛れ込んでいるようです。
村人に成りすましている人狼を突き止め、処刑しましょう。
処刑対象は、村人の投票により決定します。
村で決めた期限までに、処刑すべき人物の名前を投票してください。
投票は、【わたしへの個別トーク】にて、

（投票先プレイヤー名）＠投票

と発言してください。

投票の締切時間になったら、【このグループ】にて、

＠投票終了

と発言してください。`,
		o.WolfCount,
	)
}

type VillageFinishVoting struct {
	UnvotedCount int
}

func (o VillageFinishVoting) Reply() string {
	var notice string
	if o.UnvotedCount == 0 {
		notice = "すべてのプレイヤーが投票を終了しています。"
	} else {
		notice = fmt.Sprintf(
			"まだ投票を行っていないプレイヤーが%v人います。",
			o.UnvotedCount,
		)
	}
	return fmt.Sprintf(`%v

本当に投票を終了して処刑を実行してもよろしいですか？

投票を終了して処刑を実行する場合は、
＠はい

キャンセルする場合は、
＠いいえ

と発言してください`,
		notice,
	)
}

type VillageFinishNighttime struct{}

func (o VillageFinishNighttime) Reply() string {
	return `能力の実行が必要なすべてのプレイヤーが能力の実行を完了しています。

本当に夜の時間を終了して次の日に進んでもよろしいですか？

夜の時間を終了する場合は、
＠はい

キャンセルする場合は、
＠いいえ

と発言してください`
}

type VillageConfirmCasting struct{}

func (o VillageConfirmCasting) Reply() string {
	return `配役を設定しました。
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

と入力してください。`
}

type VillageConfirmFinishVotingLynched struct {
	LynchedPlayer model.PlayerName
	VoteDetail    model.VoteDetail
}

func (o VillageConfirmFinishVotingLynched) Reply() string {
	return fmt.Sprintf(`投票の結果、%vさんが処刑されました。

○投票結果
%v

処刑を行ったにも関わらず、恐ろしい夜はやってきます。

能力を持っているプレイヤーは、村で決めた時間までに、役職を実行してください。

役職の実行方法が分からないプレイヤーは、【わたしへの個別トーク】にて、
＠確認
と入力してください。

村で決めた時間になったら、【このグループ】にて、
＠夜明け
と入力してください`,
		o.LynchedPlayer.String(),
		o.VoteDetail.StringForHuman(),
	)
}

type VillageConfirmFinishVotingRevoting struct {
	VoteDetail model.VoteDetail
}

func (o VillageConfirmFinishVotingRevoting) Reply() string {
	return fmt.Sprintf(`投票結果が同数のため、再投票を行います。

各プレイヤーは、もう一度、【わたしへの個別トーク】にて
（投票先プレイヤー名）＠投票
と入力して再投票を行ってください。
（再投票を行わなかったプレイヤーは、前回と同じ投票先に投票したものとみなされます）

全員の再投票が終わったら、【このグループ】にて、
＠投票終了
と発言してください。

○投票結果
%v`,
		o.VoteDetail.StringForHuman(),
	)
}

type VillageConfirmFinishVotingJudged struct {
	Judge         judge.Judge
	LynchedPlayer model.PlayerName
	VoteDetail    model.VoteDetail
}

func (o VillageConfirmFinishVotingJudged) Reply() string {
	return fmt.Sprintf(`投票の結果、%vさんが処刑されました。

○投票結果
%v

%v`,
		o.LynchedPlayer.String(),
		o.VoteDetail.StringForHuman(),
		judgeResultMessage(o.Judge),
	)
}

type VillageConfirmFinishNighttimeJudged struct {
	Judge   judge.Judge
	Victims model.PlayerNames
}

func (o VillageConfirmFinishNighttimeJudged) Reply() string {
	return fmt.Sprintf(`%v

%v`,
		victimsMessage(o.Victims),
		judgeResultMessage(o.Judge),
	)
}

type VillageConfirmFinishNighttimeOngoing struct {
	Victims model.PlayerNames
}

func (o VillageConfirmFinishNighttimeOngoing) Reply() string {
	return fmt.Sprintf(`%v

今日も議論を行い、村で決めた期限までに、処刑すべき人物の名前を投票してください。

投票は、【わたしへの個別トーク】にて、

（投票先プレイヤー名）＠投票

と入力してください。

投票の締切時間になったら、【このグループ】にて、

＠投票終了

と入力してください。`,
		victimsMessage(o.Victims),
	)
}

type VillageRejectConfirmCasting struct {
	PrevStatus gamestatus.GameStatus
}

func (o VillageRejectConfirmCasting) Reply() string {
	return fmt.Sprintf(`配役の設定をキャンセルしました。

もう一度配役を設定します。

設定可能な役職は、
%v
です。

設定したい配役を以下のように発言してください。
（残った人数には自動的に村人が設定されます）

○占い師1人、霊能者1人、狩人1人、人狼2人、残りを村人に設定する場合の例

占霊狩狼狼＠配役設定`,
		roles.AvailableRoleNames())
}

type VillageRejectFinishVoting struct{}

func (o VillageRejectFinishVoting) Reply() string {
	return `投票の終了をキャンセルしました。
投票がまだのプレイヤーは、投票を行ってください。

投票は、【わたしへの個別トーク】にて、

（投票先プレイヤー名）＠投票

と発言してください。

投票の締切時間になったら、【このグループ】にて、

＠投票終了

と発言してください。`
}

type VillageRejectFinishNighttime struct{}

func (o VillageRejectFinishNighttime) Reply() string {
	return `夜時間の終了をキャンセルしました。

夜時間を終了し、次の日に進む場合は、【このグループ】にて、
＠夜明け
と入力してください`
}

func judgeResultMessage(result judge.Judge) string {
	switch result {
	case judge.Villagers:
		return `……村に平和が訪れました！
【村人側の勝利】です。
もう一度ゲームを行う場合は、
＠村作成
と入力してください。`

	case judge.Wolves:
		return `……村は人狼に乗っ取られてしまいました。
【人狼側の勝利】です。
もう一度ゲームを行う場合は、
＠村作成
と入力してください。`

	case judge.Foxes:
		return `……なんと、村は妖狐に乗っ取られてしまいました！
【妖狐の勝利】です。
もう一度ゲームを行う場合は、
＠村作成
と入力してください。`
	}

	// Unreachable
	return ""
}

func victimsMessage(victims model.PlayerNames) string {
	if len(victims) == 0 {
		return "夜が明けました……幸いなことに、昨夜は犠牲者がいなかったようです。"
	}

	var names string
	for _, v := range victims {
		if names != "" {
			names += "、"
		}
		names += "「" + v.String() + "」さん"
	}

	return fmt.Sprintf("夜が明けました……%vが無残な姿で発見されました", names)
}
