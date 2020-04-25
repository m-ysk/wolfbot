package role

import "errors"

type Role string

const (
	Unassigned Role = "Unassigned"
	Villager   Role = "Villager"
	Wolf       Role = "Wolf"
)

var stringToRole = map[string]Role{
	string(Unassigned): Unassigned,
	string(Villager):   Villager,
	string(Wolf):       Wolf,
}

var roleToStringForHuman = map[Role]string{
	Unassigned: "未割当",
	Villager:   "村人",
	Wolf:       "人狼",
}

var (
	ErrorInvalidRole = errors.New("invalid_role")
)

func New(str string) (Role, error) {
	if v, ok := stringToRole[str]; ok {
		return v, nil
	}
	return "", ErrorInvalidRole
}

func Must(str string) Role {
	v, err := New(str)
	if err != nil {
		panic(err)
	}
	return v
}

func (r Role) String() string {
	return string(r)
}

func (r Role) StringForHuman() string {
	if v, ok := roleToStringForHuman[r]; ok {
		return v
	}
	return ""
}
