package roles

// 勝敗判定時に狼としてカウントするかどうか
type WolfCountType string

const (
	// 狼としてカウントしない
	WolfCountTypeHuman = "Human"

	// 狼としてカウントする
	WolfCountTypeWolf = "Wolf"
)

func (t WolfCountType) WolfCountable() bool {
	return t == WolfCountTypeWolf
}
