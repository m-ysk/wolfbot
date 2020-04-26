package role

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
