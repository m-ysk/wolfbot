package regulation

import (
	"database/sql"
	"encoding/json"
)

type Regulation struct {
	TieVote            TieVote
	FirstDayDivination FirstDayDivination
	ContinuousGuard    ContinuousGuard
	SelfGuard          SelfGuard
}

func NewByDefault() Regulation {
	return Regulation{
		TieVote:            TieVoteRevoting,
		FirstDayDivination: FirstDayDivinationRandomWhite,
		ContinuousGuard:    ContinuousGuardDisabled,
		SelfGuard:          SelfGuardDisabled,
	}
}

func (r Regulation) MustNullString() sql.NullString {
	b, err := json.Marshal(&r)
	if err != nil {
		panic(err)
	}

	return sql.NullString{
		String: string(b),
		Valid:  true,
	}
}
