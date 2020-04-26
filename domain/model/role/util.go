package role

func AvailableRoleNames() string {
	var result string
	for i, v := range roleDefinitions {
		if i != 0 {
			result += "\n"
		}
		result += v.Name + "（" + v.Abbr + "）"
	}
	return result
}
