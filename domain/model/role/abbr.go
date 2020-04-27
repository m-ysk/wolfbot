package role

import "errors"

type Abbr string

var (
	ErrorInvalidRoleAbbr = errors.New("invalid_role_abbr")
)

func (a Abbr) String() string {
	return string(a)
}
