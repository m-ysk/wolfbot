package role

import "database/sql"

type Role struct {
	ID            ID
	Name          string
	Abbr          Abbr
	WolfCountType WolfCountType
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

func NewFromAbbr(abbr Abbr) (Role, error) {
	for _, v := range roleDefinitions {
		if v.Abbr == abbr {
			return v, nil
		}
	}

	return Role{}, ErrorInvalidRoleAbbr
}

func (r Role) String() string {
	return string(r.ID)
}

func (r Role) NullString() sql.NullString {
	return sql.NullString{
		String: r.String(),
		Valid:  true,
	}
}
