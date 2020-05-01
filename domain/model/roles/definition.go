package roles

// 勝敗判定で狼としてカウントされる役職
var wolfCountables = IDs{
	Wolf,
}

// 勝敗判定時にプレイヤー総数にカウントしない役職
var uncountablesForJudge = IDs{}

// 夜に能力を必ず実行しなければならない役職
var mustActs = IDs{
	Wolf,
}

// 「噛む」コマンドを実行できる役職
var bitables = IDs{
	Wolf,
}

var roleUnassigned = Role{
	ID:   Unassigned,
	Name: "未割当",
	Abbr: "未",
}

var roleDefinitions = Roles{
	{
		ID:   Villager,
		Name: "村人",
		Abbr: "村",
	},
	{
		ID:   Wolf,
		Name: "人狼",
		Abbr: "狼",
	},
}
