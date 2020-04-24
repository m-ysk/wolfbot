package output

import (
	"fmt"
	"wolfbot/domain/model/gamestatus"
)

type VillageCheckStatus struct {
	VillageNotExist bool
	Status          gamestatus.GameStatus
}

func (o VillageCheckStatus) StringForHuman() string {
	if o.VillageNotExist {
		return `○現在の状況
村が作成されていません

○ヘルプ
このグループで人狼ゲームを行う場合は、
＠村作成
と入力してください`
	}

	return fmt.Sprintf(`○現在の状況
%v`,
		o.Status.StringForHuman(),
	)
}
