package output

import (
	"fmt"
	"wolfbot/domain/model/gamestatus"
)

type VillageCheckStatus struct {
	VillageNotExist bool
	Status          gamestatus.GameStatus
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
%v`,
		o.Status.StringForHuman(),
	)
}

type VillageCreate struct{}

func (o VillageCreate) Reply() string {
	return "村を作成しました"
}

type VillageDelete struct{}

func (o VillageDelete) Reply() string {
	return "村を削除しました"
}

type VillageAddPlayer struct{}
