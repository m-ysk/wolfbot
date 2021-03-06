package service

import (
	"errors"
	"fmt"
	"strconv"
	"time"
	"wolfbot/domain/interfaces"
	"wolfbot/domain/model"
	"wolfbot/domain/model/debug"
	"wolfbot/domain/model/gamestatus"
	"wolfbot/domain/model/judge"
	"wolfbot/domain/model/roles"
	"wolfbot/domain/output"
	"wolfbot/lib/errorwr"
)

type VillageService struct {
	villageRepository            interfaces.VillageRepository
	playerRepository             interfaces.PlayerRepository
	userPlayerRelationRepository interfaces.UserPlayerRelationRepository
	gameRepository               interfaces.GameRepository
	uuidGenerator                interfaces.UUIDGenerator
	randomGenerator              interfaces.RandomGenerator
}

func NewVillageService(
	villageRepository interfaces.VillageRepository,
	playerRepository interfaces.PlayerRepository,
	userPlayerRelationRepository interfaces.UserPlayerRelationRepository,
	gameRepository interfaces.GameRepository,
	uuidGenerator interfaces.UUIDGenerator,
	randomGenerator interfaces.RandomGenerator,
) VillageService {
	return VillageService{
		villageRepository:            villageRepository,
		playerRepository:             playerRepository,
		userPlayerRelationRepository: userPlayerRelationRepository,
		gameRepository:               gameRepository,
		uuidGenerator:                uuidGenerator,
		randomGenerator:              randomGenerator,
	}
}

func (s VillageService) CheckState(
	id model.VillageID,
) (output.VillageCheckState, error) {
	village, err := s.villageRepository.FindByID(id)
	if err != nil {
		if model.IsErrorNotFound(err) {
			return output.VillageCheckState{
				VillageNotExist: true,
			}, nil
		}

		return output.VillageCheckState{}, err
	}

	players, err := s.playerRepository.FindByVillageID(id)
	if err != nil {
		return output.VillageCheckState{}, err
	}

	return output.VillageCheckState{
		VillageNotExist: false,
		Village:         village,
		Players:         players,
	}, nil
}

func (s VillageService) Create(
	id model.VillageID,
) (output.VillageCreate, error) {
	village := model.NewVillage(id, debug.Normal)

	if err := s.villageRepository.Create(village); err != nil {
		return output.VillageCreate{}, err
	}

	return output.VillageCreate{}, nil
}

func (s VillageService) CreateForDebug(
	id model.VillageID,
) (output.VillageCreate, error) {
	village := model.NewVillage(id, debug.Debug)

	if err := s.villageRepository.Create(village); err != nil {
		return output.VillageCreate{}, err
	}

	return output.VillageCreate{}, nil
}

func (s VillageService) Delete(id model.VillageID) (output.VillageDelete, error) {
	if err := s.villageRepository.Delete(id); err != nil {
		return output.VillageDelete{}, err
	}

	return output.VillageDelete{}, nil
}

func (s VillageService) AddPlayer(
	villageID model.VillageID,
	userID model.UserID,
	name string,
) (output.VillageAddPlayer, error) {
	village, err := s.villageRepository.FindByID(villageID)
	if err != nil {
		return output.VillageAddPlayer{}, err
	}

	if village.Status != gamestatus.RecruitingPlayers {
		return output.VillageAddPlayer{}, ErrorCommandUnauthorized
	}

	playerName, err := model.NewPlayerName(name)
	if err != nil {
		return output.VillageAddPlayer{}, err
	}

	relations, err := s.userPlayerRelationRepository.FindByUserID(userID)
	if err != nil {
		return output.VillageAddPlayer{}, err
	}

	// 同一Group内で同じUserが既にPlayer登録されている場合はエラー
	if _, ok := relations.FindByVillageID(villageID); !village.IsDebug() && ok {
		return output.VillageAddPlayer{}, ErrorDuplicatedPlayerInGroup
	}

	// 同一Userが同じPlayerNameで別の村に参加している場合はエラー
	if _, ok := relations.FindByPlayerName(playerName); ok {
		return output.VillageAddPlayer{}, ErrorDuplicatedPlayerNameInSameUser
	}

	playerID := model.PlayerID(s.uuidGenerator.Generate())

	player := model.NewPlayer(
		playerID,
		villageID,
		model.PlayerName(name),
	)

	newRelation := model.NewUserPlayerRelation(
		userID,
		villageID,
		playerName,
		playerID,
	)

	if err := s.playerRepository.Create(player, newRelation); err != nil {
		return output.VillageAddPlayer{}, err
	}

	return output.VillageAddPlayer{
		PlayerName: playerName,
	}, nil
}

func (s VillageService) AddPlayersForDebug(
	villageID model.VillageID,
	userID model.UserID,
	number string,
) (output.VillageAddPlayersForDebug, error) {
	num, err := strconv.Atoi(number)
	if err != nil {
		return output.VillageAddPlayersForDebug{}, err
	}

	village, err := s.villageRepository.FindByID(villageID)
	if err != nil {
		return output.VillageAddPlayersForDebug{}, err
	}

	if !village.IsDebug() {
		return output.VillageAddPlayersForDebug{}, ErrorInvalidCallToDebugFunction
	}

	names := []string{"あ", "い", "う", "え", "お", "か", "き", "く", "け", "こ", "さ", "し", "す", "せ", "そ"}

	for i := 0; i < num; i++ {
		if _, err := s.AddPlayer(villageID, userID, names[i]); err != nil {
			return output.VillageAddPlayersForDebug{}, err
		}
		time.Sleep(time.Millisecond * 200)
	}

	return output.VillageAddPlayersForDebug{Number: num}, nil
}

func (s VillageService) FinishRecruiting(
	villageID model.VillageID,
) (output.VillageFinishRecruiting, error) {
	game, err := s.gameRepository.FindByVillageID(villageID)
	if err != nil {
		return output.VillageFinishRecruiting{}, err
	}

	if game.Village.Status != gamestatus.RecruitingPlayers {
		return output.VillageFinishRecruiting{}, ErrorCommandUnauthorized
	}

	if len(game.Players) < 3 {
		return output.VillageFinishRecruiting{}, errorwr.New(
			errors.New("insufficient_player_count"),
			"ゲームの開始には3人以上のプレイヤーの参加が必要です",
		)
	}

	game.Village.UpdateStatus(gamestatus.ConfiguringCasting)

	if err := s.gameRepository.Update(game); err != nil {
		return output.VillageFinishRecruiting{}, err
	}

	return output.VillageFinishRecruiting{Game: game}, nil
}

func (s VillageService) ConfigureCasting(
	villageID model.VillageID,
	castingStr string,
) (output.VillageConfigureCasting, error) {
	game, err := s.gameRepository.FindByVillageID(villageID)
	if err != nil {
		return output.VillageConfigureCasting{}, err
	}

	if game.Village.Status != gamestatus.ConfiguringCasting {
		return output.VillageConfigureCasting{}, ErrorCommandUnauthorized
	}

	casting, err := roles.ParseAndValidateCasting(
		castingStr,
		len(game.Players),
	)
	if err != nil {
		return output.VillageConfigureCasting{}, err
	}

	game.Village.UpdateCasting(casting)
	game.Village.UpdateStatus(gamestatus.ConfirmingCasting)

	if err := s.gameRepository.Update(game); err != nil {
		return output.VillageConfigureCasting{}, err
	}

	return output.VillageConfigureCasting{
		Casting: casting,
	}, err
}

func (s VillageService) FinishConfiguringRegulation(
	villageID model.VillageID,
) (output.VillageFinishConfiguringRegulation, error) {
	game, err := s.gameRepository.FindByVillageID(villageID)
	if err != nil {
		return output.VillageFinishConfiguringRegulation{}, err
	}

	if game.Village.Status != gamestatus.ConfiguringRegulation {
		return output.VillageFinishConfiguringRegulation{}, ErrorCommandUnauthorized
	}

	game.AssignRole(s.randomGenerator.GenerateShuffledPermutation)
	game.Village.UpdateStatus(gamestatus.CheckingRole)

	if err := s.gameRepository.Update(game); err != nil {
		return output.VillageFinishConfiguringRegulation{}, err
	}

	return output.VillageFinishConfiguringRegulation{}, nil
}

func (s VillageService) StartGame(
	villageID model.VillageID,
) (output.VillageStartGame, error) {
	game, err := s.gameRepository.FindByVillageID(villageID)
	if err != nil {
		return output.VillageStartGame{}, err
	}

	if game.Village.Status != gamestatus.CheckingRole {
		return output.VillageStartGame{}, ErrorCommandUnauthorized
	}

	if unacted := game.Players.CountUnacted(); unacted != 0 {
		return output.VillageStartGame{}, errorwr.New(
			errors.New("unacted_player_remaining"),
			fmt.Sprintf(
				"まだ役職を確認していないプレイヤーが%v人います。全員が役職を確認するまでゲームを開始できません",
				unacted,
			),
		)
	}

	game.Start()

	if err := s.gameRepository.Update(game); err != nil {
		return output.VillageStartGame{}, err
	}

	return output.VillageStartGame{
		WolfCount: game.Players.CountAliveWolf(),
	}, nil
}

func (s VillageService) FinishVoting(
	villageID model.VillageID,
) (output.VillageFinishVoting, error) {
	game, err := s.gameRepository.FindByVillageID(villageID)
	if err != nil {
		return output.VillageFinishVoting{}, err
	}

	if game.Village.Status != gamestatus.Daytime {
		return output.VillageFinishVoting{}, ErrorCommandUnauthorized
	}

	game.Village.UpdateStatus(gamestatus.ConfirmingFinishDaytime)

	if err := s.gameRepository.Update(game); err != nil {
		return output.VillageFinishVoting{}, err
	}

	return output.VillageFinishVoting{
		UnvotedCount: game.Players.CountUnvoted(),
	}, nil
}

func (s VillageService) FinishNighttime(
	villageID model.VillageID,
) (output.VillageFinishNighttime, error) {
	game, err := s.gameRepository.FindByVillageID(villageID)
	if err != nil {
		return output.VillageFinishNighttime{}, err
	}

	if game.Village.Status != gamestatus.Nighttime {
		return output.VillageFinishNighttime{}, ErrorCommandUnauthorized
	}

	if unacted := game.Players.FilterMustAct().FilterUnacted(); len(unacted) > 0 {
		return output.VillageFinishNighttime{}, errorwr.New(
			errors.New("unacted_player_exists"),
			"まだ能力を実行していないプレイヤーが存在します。すべてのプレイヤーが能力を実行するまで次の日に進むことはできません",
		)
	}

	game.Village.UpdateStatus(gamestatus.ConfirmingFinishNighttime)

	if err := s.gameRepository.Update(game); err != nil {
		return output.VillageFinishNighttime{}, err
	}

	return output.VillageFinishNighttime{}, nil
}

func (s VillageService) Confirm(
	villageID model.VillageID,
) (output.Interface, error) {
	game, err := s.gameRepository.FindByVillageID(villageID)
	if err != nil {
		return nil, err
	}

	switch game.Village.Status {
	case gamestatus.ConfirmingCasting:
		return s.confirmCasting(game)

	case gamestatus.ConfirmingFinishDaytime:
		return s.confirmFinishVoting(game)

	case gamestatus.ConfirmingFinishNighttime:
		return s.confirmFinishNighttime(game)

	default:
		return nil, ErrorCommandUnauthorized
	}
}

func (s VillageService) confirmCasting(
	game model.Game,
) (output.VillageConfirmCasting, error) {
	game.Village.UpdateStatus(gamestatus.ConfiguringRegulation)
	if err := s.gameRepository.Update(game); err != nil {
		return output.VillageConfirmCasting{}, err
	}
	return output.VillageConfirmCasting{}, nil
}

func (s VillageService) confirmFinishVoting(
	game model.Game,
) (output.Interface, error) {
	result, err := game.Lynch(s.randomGenerator.Intn)
	if err != nil {
		return nil, err
	}

	if result.Revoting {
		game.Village.UpdateStatus(gamestatus.Daytime)
		if err := s.gameRepository.Update(game); err != nil {
			return nil, err
		}

		return output.VillageConfirmFinishVotingRevoting{
			VoteDetail: game.VoteDetail(),
		}, nil
	}

	if judgeRes := game.Judge(); judgeRes != judge.Ongoing {
		if err := s.villageRepository.Delete(game.Village.ID); err != nil {
			return nil, err
		}

		return output.VillageConfirmFinishVotingJudged{
			Judge:         judgeRes,
			LynchedPlayer: result.LynchedPlayer.Name,
			VoteDetail:    game.VoteDetail(),
		}, nil
	}

	game.ProceedToNighttime()

	if err := s.gameRepository.Update(game); err != nil {
		return nil, err
	}

	return output.VillageConfirmFinishVotingLynched{
		LynchedPlayer: result.LynchedPlayer.Name,
		VoteDetail:    game.VoteDetail(),
	}, nil
}

func (s VillageService) confirmFinishNighttime(
	game model.Game,
) (output.Interface, error) {
	result := game.ExecuteRoleAction()

	if judgeRes := game.Judge(); judgeRes != judge.Ongoing {
		if err := s.villageRepository.Delete(game.Village.ID); err != nil {
			return nil, err
		}
		return output.VillageConfirmFinishNighttimeJudged{
			Judge:   judgeRes,
			Victims: result.Victims.Names(),
		}, nil
	}

	game.ProceedToNextDay()

	if err := s.gameRepository.Update(game); err != nil {
		return nil, err
	}

	return output.VillageConfirmFinishNighttimeOngoing{
		Victims: result.Victims.Names(),
	}, nil
}

func (s VillageService) Reject(
	villageID model.VillageID,
) (output.Interface, error) {
	game, err := s.gameRepository.FindByVillageID(villageID)
	if err != nil {
		return output.VillageRejectConfirmCasting{}, err
	}

	switch st := game.Village.Status; st {
	case gamestatus.ConfirmingCasting:
		return s.rejectConfirmCasting(game)

	case gamestatus.ConfirmingFinishDaytime:
		return s.rejectFinishVoting(game)

	case gamestatus.ConfirmingFinishNighttime:
		return s.rejectFinishNighttime(game)

	default:
		return output.VillageRejectConfirmCasting{}, ErrorCommandUnauthorized
	}
}

func (s VillageService) rejectConfirmCasting(
	game model.Game,
) (output.VillageRejectConfirmCasting, error) {
	game.Village.UpdateStatus(gamestatus.ConfiguringCasting)
	if err := s.gameRepository.Update(game); err != nil {
		return output.VillageRejectConfirmCasting{}, err
	}
	return output.VillageRejectConfirmCasting{}, nil
}

func (s VillageService) rejectFinishVoting(
	game model.Game,
) (output.VillageRejectFinishVoting, error) {
	game.Village.UpdateStatus(gamestatus.Daytime)
	if err := s.gameRepository.Update(game); err != nil {
		return output.VillageRejectFinishVoting{}, err
	}
	return output.VillageRejectFinishVoting{}, nil
}

func (s VillageService) rejectFinishNighttime(
	game model.Game,
) (output.VillageRejectFinishNighttime, error) {
	game.Village.UpdateStatus(gamestatus.Nighttime)
	if err := s.gameRepository.Update(game); err != nil {
		return output.VillageRejectFinishNighttime{}, err
	}
	return output.VillageRejectFinishNighttime{}, nil
}
