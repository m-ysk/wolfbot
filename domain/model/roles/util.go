package roles

func AvailableRoleNames() string {
	var result string
	for i, v := range roleDefinitions {
		if i != 0 {
			result += "\n"
		}
		result += v.Name + "（" + v.Abbr.String() + "）"
	}
	return result
}
