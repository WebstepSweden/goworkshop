package goworkshop

// Coordinate is a struct that represents a geographic location with an x and y numbers (use type float32)
type Coordinate struct {
	x float32
	y float32
}

// PointerToCoordinate creates a coordinate using the given x and y and returns a pointer to it
func PointerToCoordinate(x float32, y float32) *Coordinate {
	coordinate := Coordinate{x, y}
	return &coordinate
}

// ChangeCoordinate changes a given coordinate pointer's x and y given new x and y
func ChangeCoordinate(coordinate *Coordinate, x float32, y float32) {
	coordinate.x = x
	coordinate.y = y
}
