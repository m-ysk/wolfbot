package handler

import (
	"testing"
	"wolfbot/domain/model"

	"github.com/google/go-cmp/cmp"
)

func Test_parseGroupMessage(t *testing.T) {
	tests := []struct {
		message string
		userID  model.PlayerID
		groupID model.VillageID
		want    command
	}{
		{
			message: "@村作成",
			userID:  "user",
			groupID: "group",
			want: command{
				Action:  actionCreateVillage,
				Target:  "",
				UserID:  "user",
				GroupID: "group",
			},
		},
		{
			message: "＠村作成",
			userID:  "user",
			groupID: "group",
			want: command{
				Action:  actionCreateVillage,
				Target:  "",
				UserID:  "user",
				GroupID: "group",
			},
		},
	}

	for _, tt := range tests {
		if got := parseGroupMessage(tt.message, tt.userID, tt.groupID); !cmp.Equal(tt.want, got) {
			t.Error(cmp.Diff(tt.want, got))
		}
	}
}
