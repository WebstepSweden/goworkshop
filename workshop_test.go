package goworkshop

import "testing"

func TestVariables(t *testing.T) {
	CreateVariables()
}

func TestLooping(t *testing.T) {
	LoopNumbers()
	//LoopStrings()
	//LoopRunes()
}

func TestCollections(t *testing.T) {
	Slices()
	//Maps()
}

func TestStructs(t *testing.T) {
	CreateStruct()
}

func TestStringsAndFmt(b *testing.T) {
	UseStrings("Spam")
	//UseFmt()
}

func TestPointers(t *testing.T) {
	UsePointers()
}

func TestDivisionRunner(t *testing.T) {
	DivisionRunner(2, 3)
	DivisionRunner(5, 0)
}

func TestGetANumber_(t *testing.T) {
	want := 8
	got := GetANumber()
	if got != want {
		t.Fatalf(`Fail! Wanted '%v', got '%v'`, want, got)
	}
}

func BenchmarkGetANumber_(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetANumber()
	}
}