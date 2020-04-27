package randutil

import "math/rand"

func GenerateShuffledPermutation(upper int) []int {
	var result []int
	for i := 0; i < upper; i++ {
		result = append(result, i)
	}

	ShuffleInts(result)

	return result
}

func ShuffleInts(data []int) {
	n := len(data)
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		data[i], data[j] = data[j], data[i]
	}
}

type ShuffleData interface {
	Len() int
	Swap(i, j int)
}

func Shuffle(data ShuffleData) {
	n := data.Len()
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		data.Swap(i, j)
	}
}
