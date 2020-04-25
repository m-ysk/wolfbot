package handler

import (
	"testing"
	"wolfbot/domain/model"

	"github.com/google/go-cmp/cmp"
)

func Test_parseGroupMessage(t *testing.T) {
	tests := []struct {
		message string
		userID  model.UserID
		groupID model.VillageID
		want    command
	}{
		{
			message: "@村作成",
			userID:  "user",
			groupID: "group",
			want: command{
				Action: actionCreateVillage,
				Target: "",
			},
		},
		{
			message: "＠村作成",
			userID:  "user",
			groupID: "group",
			want: command{
				Action: actionCreateVillage,
				Target: "",
			},
		},
	}

	for _, tt := range tests {
		if got := parseGroupMessage(tt.message); !cmp.Equal(tt.want, got) {
			t.Error(cmp.Diff(tt.want, got))
		}
	}
}
