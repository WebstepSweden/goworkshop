package goworkshop

import "testing"

// *** Variables ***
func TestVariables(t *testing.T) {
	CreateVariables()
}

// *** For loops ***
func TestLooping(t *testing.T) {
	LoopNumbers()
	//LoopStrings()
	//LoopRunes()
}

// *** Collections ***
func TestCollections(t *testing.T) {
	Slices()
	//Maps()
}
