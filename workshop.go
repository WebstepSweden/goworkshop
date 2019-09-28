// Package goworkshop includes, not too surprisingly, a workshop about Go
package goworkshop

import (
	"fmt"
	"strings"
)

/*
Go/Golang?
*/

/*
Setting up development environment
*/

/*
Exporting
*/

/*
go fmt and golint
*/

/*
Things not in Go
*/

/*
Variables
*/

/*
For loops
*/

// Looping shows how to use loops in Go
func Looping() {

	// loop through number 1 to 10 and prints all the numbers that are divisible by 3
	for i := 1; i < 10; i++ {
		if i%3 == 0 {
			println(i)
		}
	}

	// loops through names and find those who contains "ar"
	var namesWithAr []string
	names := []string{"Barbara", "Cindy", "Laura", "Carina"}
	for _, name := range names {
		if strings.Contains(name, "ar") {
			namesWithAr = append(namesWithAr, name)
		}
	}
	fmt.Printf("%v", namesWithAr)
}

/*
Collections
*/

/*
Structs
*/

/*
strings and fmt
*/

/*
Pointers
*/

/*
Errors
*/

/*
Testing
*/
