package interfaces

type RandomGenerator interface {
	Intn(upper int) int
	GenerateShuffledPermutation(upper int) []int
}
