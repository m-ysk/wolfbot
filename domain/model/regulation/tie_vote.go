package regulation

type TieVote string

const (
	// 再投票
	TieVoteRevoting TieVote = "Revoting"

	// 引き分け
	TieVoteDraw TieVote = "Draw"
)
