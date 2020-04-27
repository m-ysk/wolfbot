package role

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"wolfbot/lib/errorwr"
)

type Casting map[ID]int

func ParseAndValidateCasting(str string, total int) (Casting, error) {
	casting, err := parseCasting(str, total)
	if err != nil {
		return nil, err
	}

	if err := validateCasting(casting); err != nil {
		return nil, err
	}

	return casting, nil
}

func parseCasting(str string, total int) (Casting, error) {
	result := make(Casting)
	assigned := 0
	for _, c := range str {
		r, err := NewFromAbbr(Abbr(c))
		if err != nil {
			return nil, err
		}
		assigned++
		result[r.ID]++
	}

	// 指定された役職の総数と参加者総数の差分は自動で村人に設定する
	result[Villager] += total - assigned

	return result, nil
}

var (
	ErrorWolfCountZero      = errors.New("wolf_count_zero")
	ErrorWolfCountExceeding = errors.New("wolf_count_exceeding")
)

func validateCasting(casting Casting) error {
	var total, wolf int
	for roleID, num := range casting {
		role, err := New(roleID.String())
		if err != nil {
			return err
		}

		total += num
		if role.WolfCountType.WolfCountable() {
			wolf += num
		}
	}

	if wolf == 0 {
		return errorwr.New(
			ErrorWolfCountZero,
			"人狼が設定されていません。最低でも1人の人狼またはその他の人狼系役職を設定してください",
		)
	}

	maxWolf := (total - 1) / 2
	if wolf > maxWolf {
		return errorwr.New(
			ErrorWolfCountExceeding,
			fmt.Sprintf("人狼の数が多すぎます。人狼系役職の合計は%v人以下に設定してください", maxWolf),
		)
	}

	return nil
}

// Slice は配役に含まれる役職のIDを人数分並べたスライスを返す
// 役職の並び順は順序保証される
func (c Casting) RoleIDs() IDs {
	var result IDs
	for _, role := range roleDefinitions {
		if num, ok := c[role.ID]; ok {
			for i := 0; i < num; i++ {
				result = append(result, role.ID)
			}
		}
	}
	return result
}

func (c Casting) MustNullString() sql.NullString {
	b, err := json.Marshal(&c)
	if err != nil {
		panic(err)
	}

	return sql.NullString{
		String: string(b),
		Valid:  true,
	}
}

func (c Casting) StringForHuman() string {
	var result string
	for _, role := range roleDefinitions {
		if num, ok := c[role.ID]; ok {
			if result != "" {
				result += "\n"
			}
			result += role.Name + ": " + strconv.Itoa(num) + "人"
		}
	}
	return result
}
