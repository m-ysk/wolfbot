package role

import "errors"

type ID string

const (
	Unassigned ID = "Unassigned"
	Villager   ID = "Villager"
	Wolf       ID = "Wolf"
)

var (
	ErrorInvalidRoleID = errors.New("invalid_role_id")
)

func (id ID) String() string {
	return string(id)
}