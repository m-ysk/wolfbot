package roles

import "errors"

type ID string

type IDs []ID

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
