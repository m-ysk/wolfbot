package randgen

type RandomGeneratorMock struct{}

func (g RandomGeneratorMock) Intn(upper int) int {
	return 0
}

func (g RandomGeneratorMock) GenerateShuffledPermutation(upper int) []int {
	var result []int
	for i := 0; i < upper; i++ {
		result = append(result, i)
	}
	return result
}
