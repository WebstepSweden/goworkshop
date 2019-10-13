package goworkshop

// CreateSliceWithSevens returns a slice of size max (given as argument).
// Every element in the slice is a seven.
func CreateSliceWithSevens(max int) []int {
	sliceWithSevens := make([]int, max)
	for i := 0; i < max; i++ {
		sliceWithSevens[i] = 7
	}
	return sliceWithSevens
}

// CreateMapWithSquares returns a map of size max (given as argument).
// The map contains numbers (0 to max) as keys and their squares as values.
func CreateMapWithSquares(max int) map[int]int {
	squares := make(map[int]int)
	for i := 0; i < max; i++ {
		squares[i] = i * i
	}
	return squares
}
