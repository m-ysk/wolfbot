package roles

// 「噛む」コマンドを実行できる役職
var bitableRoles = IDs{
	Wolf,
}

var roleUnassigned = Role{
	ID:            Unassigned,
	Name:          "未割当",
	Abbr:          "未",
	WolfCountType: WolfCountTypeHuman,
}

var roleDefinitions = Roles{
	{
		ID:            Villager,
		Name:          "村人",
		Abbr:          "村",
		WolfCountType: WolfCountTypeHuman,
	},
	{
		ID:            Wolf,
		Name:          "人狼",
		Abbr:          "狼",
		WolfCountType: WolfCountTypeWolf,
	},
}
