package handler

import (
	"testing"
	"wolfbot/domain/model"
	"wolfbot/domain/model/judge"
	"wolfbot/domain/output"
	"wolfbot/domain/service"
	"wolfbot/infra/randgen"
	"wolfbot/infra/rdb"
	"wolfbot/infra/uuidgen"

	"github.com/google/go-cmp/cmp"
)

func InitMessageHandler() MessageHandler {
	var villages []model.Village
	var players model.Players
	var relations model.UserPlayerRelations

	villageRepo := rdb.VillageRepositoryMock{
		Villages: &villages,
	}
	playerRepo := rdb.PlayerRepositoryMock{
		Players:   &players,
		Relations: &relations,
	}
	relationRepo := rdb.UserPlayerRelationRepositoryMock{
		Relations: &relations,
	}
	gameRepo := rdb.GameRepositoryMock{
		Villages: &villages,
		Players:  &players,
	}

	villageService := service.NewVillageService(
		villageRepo,
		playerRepo,
		relationRepo,
		gameRepo,
		uuidgen.NewUUIDGenerator(),
		randgen.RandomGeneratorMock{},
	)

	playerService := service.NewPlayerService(
		playerRepo,
		gameRepo,
		randgen.RandomGeneratorMock{},
	)

	relationService := service.NewUserPlayerRelationService(
		relationRepo,
	)

	return MessageHandler{
		villageService:            villageService,
		playerService:             playerService,
		userPlayerRelationService: relationService,
	}
}

type TestCommand struct {
	// message
	m string
	// userID
	u model.UserID
	// groupID
	g string
}

var villager2Wolf1GameStart = []TestCommand{
	{
		m: "＠村作成",
		u: "a",
		g: "g",
	},
	{
		m: "あ＠参加",
		u: "a",
		g: "g",
	},
	{
		m: "い＠参加",
		u: "i",
		g: "g",
	},
	{
		m: "う＠参加",
		u: "u",
		g: "g",
	},
	{
		m: "＠募集終了",
		u: "a",
		g: "g",
	},
	{
		m: "狼＠配役設定",
		u: "a",
		g: "g",
	},
	{
		m: "＠はい",
		u: "a",
		g: "g",
	},
	{
		m: "＠設定終了",
		u: "a",
		g: "g",
	},
	{
		m: "＠確認",
		u: "a",
	},
	{
		m: "＠確認",
		u: "i",
	},
	{
		m: "＠確認",
		u: "u",
	},
	{
		m: "＠村開始",
		u: "a",
		g: "g",
	},
}

var villager4Wolf2GameStart = []TestCommand{
	{
		m: "＠村作成",
		u: "a",
		g: "g",
	},
	{
		m: "あ＠参加",
		u: "a",
		g: "g",
	},
	{
		m: "い＠参加",
		u: "i",
		g: "g",
	},
	{
		m: "う＠参加",
		u: "u",
		g: "g",
	},
	{
		m: "え＠参加",
		u: "e",
		g: "g",
	},
	{
		m: "お＠参加",
		u: "o",
		g: "g",
	},
	{
		m: "か＠参加",
		u: "ka",
		g: "g",
	},
	{
		m: "＠募集終了",
		u: "a",
		g: "g",
	},
	{
		m: "狼狼＠配役設定",
		u: "a",
		g: "g",
	},
	{
		m: "＠はい",
		u: "a",
		g: "g",
	},
	{
		m: "＠設定終了",
		u: "a",
		g: "g",
	},
	{
		m: "＠確認",
		u: "a",
	},
	{
		m: "＠確認",
		u: "i",
	},
	{
		m: "＠確認",
		u: "u",
	},
	{
		m: "＠確認",
		u: "e",
	},
	{
		m: "＠確認",
		u: "o",
	},
	{
		m: "＠確認",
		u: "ka",
	},
	{
		m: "＠村開始",
		u: "a",
		g: "g",
	},
}

func TestMessageHandler_HandleMessage(t *testing.T) {
	tests := []struct {
		commands []TestCommand
		want     output.Interface
		wantErr  error
	}{
		// 1日目昼で村勝利
		{
			commands: append(villager2Wolf1GameStart, []TestCommand{
				{
					m: "う＠投票",
					u: "a",
				},
				{
					m: "う＠投票",
					u: "i",
				},
				{
					m: "あ＠投票",
					u: "u",
				},
				{
					m: "＠投票終了",
					u: "a",
					g: "g",
				},
				{
					m: "＠はい",
					u: "a",
					g: "g",
				},
			}...),
			want: output.VillageConfirmFinishVotingJudged{
				Judge:         judge.Villagers,
				LynchedPlayer: "う",
				VoteDetail: model.VoteDetail{
					"あ": "う",
					"い": "う",
					"う": "あ",
				},
			},
			wantErr: nil,
		},
		// 1日目昼で人狼勝利
		{
			commands: append(villager2Wolf1GameStart, []TestCommand{
				{
					m: "い＠投票",
					u: "a",
				},
				{
					m: "う＠投票",
					u: "i",
				},
				{
					m: "い＠投票",
					u: "u",
				},
				{
					m: "＠投票終了",
					u: "a",
					g: "g",
				},
				{
					m: "＠はい",
					u: "a",
					g: "g",
				},
			}...),
			want: output.VillageConfirmFinishVotingJudged{
				Judge:         judge.Wolves,
				LynchedPlayer: "い",
				VoteDetail: model.VoteDetail{
					"あ": "い",
					"い": "う",
					"う": "い",
				},
			},
			wantErr: nil,
		},
		// 2日目昼で村勝利
		{
			commands: append(villager4Wolf2GameStart, []TestCommand{
				{
					m: "か＠投票",
					u: "a",
				},
				{
					m: "＠投票終了",
					u: "a",
					g: "g",
				},
				{
					m: "＠はい",
					u: "a",
					g: "g",
				},
				{
					m: "あ＠噛む",
					u: "o",
				},
				{
					m: "＠夜明け",
					u: "i",
					g: "g",
				},
				{
					m: "＠はい",
					u: "i",
					g: "g",
				},
				{
					m: "お＠投票",
					u: "i",
				},
				{
					m: "＠投票終了",
					u: "i",
					g: "g",
				},
				{
					m: "＠はい",
					u: "i",
					g: "g",
				},
			}...),
			want: output.VillageConfirmFinishVotingJudged{
				Judge:         judge.Villagers,
				LynchedPlayer: "お",
				VoteDetail: model.VoteDetail{
					"あ": "",
					"い": "お",
					"う": "",
					"え": "",
					"お": "",
					"か": "",
				},
			},
			wantErr: nil,
		},
		// 2日目夜で狼勝利
		{
			commands: append(villager4Wolf2GameStart, []TestCommand{
				{
					m: "か＠投票",
					u: "a",
				},
				{
					m: "＠投票終了",
					u: "a",
					g: "g",
				},
				{
					m: "＠はい",
					u: "a",
					g: "g",
				},
				{
					m: "あ＠噛む",
					u: "o",
				},
				{
					m: "＠夜明け",
					u: "i",
					g: "g",
				},
				{
					m: "＠はい",
					u: "i",
					g: "g",
				},
				{
					m: "い＠投票",
					u: "u",
				},
				{
					m: "＠投票終了",
					u: "u",
					g: "g",
				},
				{
					m: "＠はい",
					u: "u",
					g: "g",
				},
				{
					m: "う＠噛む",
					u: "o",
				},
				{
					m: "＠夜明け",
					u: "u",
					g: "g",
				},
				{
					m: "＠はい",
					u: "u",
					g: "g",
				},
			}...),
			want: output.VillageConfirmFinishNighttimeJudged{
				Judge:   judge.Wolves,
				Victims: model.PlayerNames{"う"},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		h := InitMessageHandler()

		var out output.Interface
		var err error

		for _, c := range tt.commands {
			if c.g == "" {
				out, err = h.HandleUserMessage(c.m, c.u)
			} else {
				out, err = h.HandleGroupMessage(c.m, c.u, c.g)
			}

			if err != nil && !cmp.Equal(tt.wantErr, err) {
				t.Error(cmp.Diff(tt.wantErr, err))
			}
		}

		if !cmp.Equal(tt.want, out) {
			t.Error(cmp.Diff(tt.want, out))
		}
	}
}

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
