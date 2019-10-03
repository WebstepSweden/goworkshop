package goworkshop

import "testing"

func TestSumTo100(t *testing.T) {
	want := 100 * 101 / 2 // this is Gauss solution, in case you're wondering
	got := SumTo100()
	if got != want {
		t.Fatalf(`Fail! Wanted '%d', got '%d'`, want, got)
	}
}

func TestLoopWords(t *testing.T) {
	want := "never gonna give You up "
	got := LoopWords()
	if got != want {
		t.Fatalf(`Fail! Wanted '%s', got '%s'`, want, got)
	}
}
