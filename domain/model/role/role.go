package role

type Role struct {
	ID   ID
	Name string
	Abbr string
}

type Roles []Role

func New(id string) (Role, error) {
	if id == Unassigned.String() {
		return roleUnassigned, nil
	}

	for _, v := range roleDefinitions {
		if v.ID.String() == id {
			return v, nil
		}
	}

	return Role{}, ErrorInvalidRoleID
}

func Must(str string) Role {
	v, err := New(str)
	if err != nil {
		panic(err)
	}
	return v
}

func (r Role) String() string {
	return string(r.ID)
}
