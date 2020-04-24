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
		fmt.Sprintf("村が作成されていません")
	}

	return fmt.Sprintf(`
○現在の状況
%v
`,
		o.Status.StringForHuman(),
	)
}
