package regulation

type TieVote string

const (
	// 再投票
	TieVoteRevoting TieVote = "Revoting"

	// 最多得票者をランダムで処刑
	TieVoteRandomExecution TieVote = "RandomExecution"
)
