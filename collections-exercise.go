package goworkshop

// CreateSliceWithSevens returns a slice of size max where every element is a seven
func CreateSliceWithSevens(max int) []int {
	sliceWithSevens := make([]int, max)
	for i := 0; i < max; i++ {
		sliceWithSevens[i] = 7
	}
	return sliceWithSevens
}

// CreateMapWithSquares returns a map with numbers and their squares
func CreateMapWithSquares(max int) map[int]int {
	squares := make(map[int]int)
	for i := 0; i < max; i++ {
		squares[i] = i * i
	}
	return squares
}
