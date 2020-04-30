package model

import (
	"errors"
	"wolfbot/domain/model/gamestatus"
	"wolfbot/domain/model/judge"
	"wolfbot/domain/model/regulation"
	"wolfbot/domain/model/roles"
	"wolfbot/domain/model/votestatus"
	"wolfbot/lib/randutil"
)

type Game struct {
	Village Village
	Players Players
}

func (g *Game) AssignRole() {
	if g == nil {
		return
	}

	roleIDs := g.Village.Casting.RoleIDs()

	shuffledInts := randutil.GenerateShuffledPermutation(
		len(g.Players),
	)

	for i, v := range shuffledInts {
		role := roles.Must(roleIDs[v].String())
		g.Players[i].Role = role
	}
}

func (g *Game) Start() {
	if g == nil {
		return
	}
	g.Village.UpdateStatus(gamestatus.Daytime)
	g.Village.Day = 1
}

func (g *Game) ProceedToNighttime() {
	if g == nil {
		return
	}

	for i, v := range g.Players {
		v.Unacted()
		g.Players[i] = v
	}

	g.Village.Status = gamestatus.Nighttime
}

func (g *Game) Judge() judge.Judge {
	if g == nil {
		return judge.Ongoing
	}

	aliveCountablePlayers := g.Players.CountAliveForJudge()
	wolves := g.Players.CountWolf()

	if wolves == 0 {
		return judge.Villagers
	}

	if wolves*2 >= aliveCountablePlayers {
		return judge.Wolves
	}

	return judge.Ongoing
}

type ExecutionResult struct {
	Revoting       bool
	ExecutedPlayer Player
}

// randomIntは、0以上、引数として与えた整数未満の整数をランダムに生成する関数
func (g *Game) Execute(randomInt func(int) int) (ExecutionResult, error) {
	if g == nil {
		return ExecutionResult{}, errors.New("nil_receiver")
	}

	voteResult := g.voteCounting()

	mostVoted := voteResult.mostVoted()

	if len(mostVoted) == 0 {
		return ExecutionResult{}, errors.New("execution_error")
	}

	if mv := len(mostVoted); mv > 1 {
		// 投票同数の場合の設定が再投票の場合
		if g.Village.Regulation.TieVote == regulation.TieVoteRevoting {
			return ExecutionResult{
				Revoting:       true,
				ExecutedPlayer: Player{},
			}, nil
		}

		// 投票同数の場合の設定が最多得票者をランダム処刑の場合
		executedPlayerID := mostVoted[randomInt(mv)]
		executedPlayer, _ := g.Players.FindByID(executedPlayerID)
		executedPlayer.Kill()
		g.Players.UpdatePlayer(executedPlayer)

		return ExecutionResult{
			Revoting:       false,
			ExecutedPlayer: executedPlayer,
		}, nil
	}

	executedPlayerID := mostVoted[0]
	executedPlayer, _ := g.Players.FindByID(executedPlayerID)
	executedPlayer.Kill()
	g.Players.UpdatePlayer(executedPlayer)

	return ExecutionResult{
		Revoting:       false,
		ExecutedPlayer: executedPlayer,
	}, nil
}

type VoteCountingResult map[PlayerID]int

func (g *Game) voteCounting() VoteCountingResult {
	if g == nil {
		return nil
	}

	result := make(VoteCountingResult)
	for _, voted := range g.Players {
		for _, voter := range g.Players {
			if voter.VoteStatus == votestatus.Voted && voter.VoteTo == voted.ID {
				result[voted.ID]++
			}
		}
	}

	return result
}

func (r VoteCountingResult) mostVoted() PlayerIDs {
	if r == nil {
		return nil
	}

	// Keyが得票数、Valueがその得票数を得たPlayerのIDのスライスであるmapを作る
	countToPlayerIDs := make(map[int]PlayerIDs)
	numOfPlayers := 0
	for k, v := range r {
		numOfPlayers++
		countToPlayerIDs[v] = append(countToPlayerIDs[v], k)
	}

	// 上で作ったmapを得票数が大きいほうから順に見ていく
	for i := numOfPlayers; i >= 0; i-- {
		if v, ok := countToPlayerIDs[i]; ok {
			return v
		}
	}

	return nil
}

type VoteDetail map[PlayerName]PlayerName

func (g *Game) VoteDetail() VoteDetail {
	if g == nil {
		return nil
	}

	result := make(VoteDetail)
	for _, voter := range g.Players {
		voted, _ := g.Players.FindByID(voter.VoteTo)
		result[voter.Name] = voted.Name
	}

	return result
}

func (d VoteDetail) StringForHuman() string {
	var result string
	for k, v := range d {
		if result != "" {
			result += "\n"
		}
		result += k.String() + "=>" + v.String()
	}
	return result
}
