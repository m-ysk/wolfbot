package votestatus

import (
	"database/sql"
	"errors"
)

type VoteStatus string

const (
	Unvoted VoteStatus = "Unvoted"
	Voted   VoteStatus = "Voted"
)

var stringToStatus = map[string]VoteStatus{
	string(Unvoted): Unvoted,
	string(Voted):   Voted,
}

var statusToStringForHuman = map[VoteStatus]string{
	Unvoted: "未行動",
	Voted:   "行動済み",
}

var (
	ErrorInvalidVoteStatus = errors.New("invalid_vote_status")
)

func New(str string) (VoteStatus, error) {
	if v, ok := stringToStatus[str]; ok {
		return v, nil
	}
	return "", ErrorInvalidVoteStatus
}

func Must(str string) VoteStatus {
	v, err := New(str)
	if err != nil {
		panic(err)
	}
	return v
}

func (s VoteStatus) String() string {
	return string(s)
}

func (s VoteStatus) NullString() sql.NullString {
	return sql.NullString{
		String: s.String(),
		Valid:  true,
	}
}

func (s VoteStatus) StringForHuman() string {
	if v, ok := statusToStringForHuman[s]; ok {
		return v
	}
	return ""
}
