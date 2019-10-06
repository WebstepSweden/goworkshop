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
