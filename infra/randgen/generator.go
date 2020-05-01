package randgen

import (
	"math/rand"
	"wolfbot/lib/randutil"
)

type RandomGenerator struct{}

func NewRandomGenerator() RandomGenerator {
	return RandomGenerator{}
}

func (g RandomGenerator) Intn(upper int) int {
	return rand.Intn(upper)
}

func (g RandomGenerator) GenerateShuffledPermutation(upper int) []int {
	return randutil.GenerateShuffledPermutation(upper)
}
