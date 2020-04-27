package roles

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCasting_RoleIDs(t *testing.T) {
	casting := Casting{
		Villager: 2,
		Wolf:     3,
	}

	want := IDs{
		Villager,
		Villager,
		Wolf,
		Wolf,
		Wolf,
	}

	if got := casting.RoleIDs(); !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
