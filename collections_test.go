package goworkshop

import (
	"math/rand"
	"testing"
)

func TestSlices(t *testing.T) {
	sliceOfSevens := CreateSliceWithSevens(1000)

	for i := 0; i < 10; i++ {
		r := rand.Intn(1000)
		want := 7
		got := sliceOfSevens[r]

		if got != want {
			t.Fatalf(`Fail! Wanted '%d', '%d'`, want, got)
		}
	}
}

func TestMaps(t *testing.T) {
	squareMap := CreateMapWithSquares(1000)

	for i := 0; i < 10; i++ {
		r := rand.Intn(1000)
		want := r * r
		got := squareMap[r]

		if got != want {
			t.Fatalf(`Fail! Wanted '%d', '%d'`, want, got)
		}
	}
}
