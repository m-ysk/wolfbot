package regulation

type FirstDayDivination string

const (
	// 人狼ではない人物をランダムで通知
	FirstDayDivinationRandomWhite = "RandomWhite"

	// 通常の占いを行う
	FirstDayDivinationNormalDivination = "NormalDivination"

	// 初日占いなし
	FirstDayDivinationOmit = "Omit"
)
