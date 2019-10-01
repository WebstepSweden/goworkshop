// Package goworkshop includes, not too surprisingly, a workshop about the Go programing language
package goworkshop

import (
	"fmt"
	"strings"
)

//                          `..------..`
//                     `-:////::::::::/+//:. .-.
//               :///o/+/:---:/::::/:-..-://o+//+/
//              +/ods//-.`    `+::/.:`    `+:/do:+-
//              +//y::/+Mh.    ::/-dNs     +::/o++`
//               -o/::+-:-    .+++o-.    `:/:::o-
//                o:::://:--:/+yNMNo+/:://::::://
//               `o::::::::::+:.::-.:+::::::::::+
//               `o:::::::::::+/`/`o/:::::::::::o                   .-:-.
//                o::::::::::::/:+:/::::::::::::o`               /hNMMNMMNo
//                +:::::::::::::::::::::::::::::o`             `dMNo.   .+`     .::-`
//                +:::::::::::::::::::::::::::::o`             yMM-           omMmdNMm/
//               `+/::::::::::::::::::::::::::::o-`            NMm    .dddd+ sMN-   +MM/
//             :/:+/::::::::::::::::::::::::::::o-:+`          dMN.   `/+MMs NMh    `MMs
//             :/:o:::::::::::::::::::::::::::::++:/           -NMm:    .MMs hMN`   /MM/
//                +:::::::::::::::::::::::::::::/:              .yNMNmdmMMN/ `yMNhydMN+
//                +::::::::::::::::::::::::::::://                 -/++/:.     `:++/-
//                o::::::::::::::::::::::::::::://
//               `o:::::::::::::::::::::::::::::+-
//                o:::::::::::::::::::::::::::::o`
//                //:::::::::::::::::::::::::::/+
//                 `o:::::::::::::::::::::::::::o`
//                  .+/:::::::::::::::::::::::/+`
//                  -+/+::::::::::::::::::/+//:`
//                 //.-/+//+++///////+++//://-/+
//                 ./:-     ``.----..``     `::`

// *********************
// *** Go or Golang? ***
// *********************
//
// Welcome to the Go workshop. I'll try to take you through some of what I learned when
// going through this strange language and see if I can make you interested as well.
//
// But first, you can sometimes see people refer to this language as Go and sometimes as
// Golang. So which one is correct? Well, the name of the language is Go. However, googling
// for something called Go is quite difficult, so many sites about Go are called Golang,
// that is, the domain name is golang.org, the twitter handle is #golang, the subreddit name
// is /r/golang etc. So use what you want, and let's go with the workshop.

// Go resources:
// * https://golang.org/
// * https://play.golang.org/
// * https://tour.golang.org/
// * https://twitter.com/search?q=%23golang
// * https://www.reddit.com/r/golang/

// *** Setting up the development environment ***
//
// If you haven't already, install Go using the instructions found in https://golang.org/doc/install
//
// Your Go environment should look like this:
// Go
// |- bin
// |- pkg
// |- src
//    |- github.com
//       |- goworkshop
//
// You should also have two environment variables:
// GOROOT - pointing to wherever you installed Go
// GOPATH - pointing to where your Go code is, that is, the root folder of the structure above

// Go dev env resources:
// * https://golang.org/doc/install
// * https://medium.com/rungo/working-in-go-workspace-3b0576e0534a
// * https://www.callicoder.com/golang-installation-setup-gopath-workspace/

// *****************
// *** Exporting ***
// *****************
//
// Defining identifiers in Go (such as functions and fields), you can choose to export them or keep them local.
// Exporting means they will be available to other packages for import. So writing things like strings.ToLower("HELLO")
// means I'm using the exported function ToLower() from package strings.
// And the secret to exporting stuff from your package is this:
// .........they need to start with a capital letter.
// That's it! So in these cases:

func mySecretLittleFunction() {}

// UseMeImExported is an example of an exported function.
func UseMeImExported() {}

// only the second function will be exported.
// Also, exporting means you have to document stuff, to let potential users know what your functions do.
// If you don't, you're going to be haunted by the go fmt tool!

// Exporting resources:
// * https://tour.golang.org/basics/3
// * https://medium.com/golangspec/exported-identifiers-in-go-518e93cc98af

// ************************
// *** gofmt and golint ***
// ************************
//
// Go is quite opinionated about the coding style and defines some rigid rules to control the looks of your code.
// In order to keep things nice and clean, there are two tools you need to use, either using the terminal or
// integrating them in your IDE. These are:
// * gofmt - a tool that formats your code according to the Go formatting rules
// * golint - a tool that lints your code and prints out style mistakes

// To see them in action, uncomment this line:

//  func     Not_well_formatted() {println("hello")}

// then run:
// gofmt -w workshop.go
// golint workshop.go

// gofmt and golint resources:
// * https://blog.golang.org/go-fmt-your-code
// * https://github.com/golang/lint

// *** Variables ***
// There are different ways to create a variable in Go. We usually use the first two forms.

// CreateVariables shows how to create variables in different ways
func CreateVariables() {

	// Create a variable, then assign a value to it. Use keyword 'var' and the '=' symbol.
	var x string
	x = "Hello"

	// Create a variable and assign it a value directly. Use the ':=' symbol, no need for 'var'.
	y := "World"

	// Create a pointer to a newly allocated variable.
	z := new(string)
	*z = "Wazzup?"

	println(x)
	println(y)
	println(*z)
}

// Variable resources:
// * https://gobyexample.com/variables
// * https://www.golang-book.com/books/intro/4

// *****************
// *** Functions ***
// *****************
//
// Functions have the following syntax:
// func functionsName(arg1Name arg1Type, arg2Name arg2Type, ...etc,) returnType {
// 	   implementation
// }
//
// Functions can look like this:

func nonReturningFunc(arg int) {
}

// or like this:

func returningFunc(arg1 string, arg2 int) float32 {
	return 0
}

// Function resources:
// * https://tour.golang.org/basics/4
// * https://www.callicoder.com/golang-functions/

// *****************
// *** For loops ***
// *****************
//
// For loops are the only way you can loop through stuff in Go. There are different ways to use it though:

// LoopNumbers loops through numbers 1 to 10 and prints all the numbers that are divisible by 3
func LoopNumbers() {
	for i := 1; i < 10; i++ {
		if i%3 == 0 {
			print(i)
		}
	}
	println()
}

// LoopStrings loops through names and find those who contain "ar"
func LoopStrings() {
	var namesWithAr []string
	names := []string{"Barbara", "Cindy", "Laura", "Carina"}
	for _, name := range names {
		if strings.Contains(name, "ar") {
			namesWithAr = append(namesWithAr, name)
		}
	}
	fmt.Printf("names with ar: %v", namesWithAr)
	println()
}

// LoopRunes loops through letters and creates a new string with every forth letter
func LoopRunes() {
	var letters = ""
	for index, letter := range "abcdefghijklmnopqrstuvwxyz" {
		if index%4 == 0 {
			letters += string(letter)
		}
	}
	fmt.Printf("letters: %v", letters)
	println()
}

// For loops resources:
// * https://tour.golang.org/flowcontrol/1
// * https://golangbot.com/loops/

// *******************
// *** Collections ***
// *******************
//
// Let's take a look at two common collection: slices (flexible arrays) and maps.
// You can create them using the := symbol or by using make:

// Slices shows ways to create slices
func Slices() {
	numbers := []string{"One", "Two", "Tree"}

	primes := make([]int, 4)
	primes[0] = 2
	primes[1] = 3
	primes[2] = 5
	primes[3] = 7

	fmt.Printf("numbers: %v\n", numbers)
	fmt.Printf("primes: %v\n", primes)
}

// Maps shows ways to create maps
func Maps() {
	danishNumbers := map[string]int{
		"ti":        10,
		"tyve":      20,
		"tredive":   30,
		"fyrre":     40,
		"halvtreds": 50,
	}

	inchesToCm := make(map[int]float32)
	inchesToCm[1] = 2.54
	inchesToCm[2] = 5.08
	inchesToCm[3] = 7.62
	inchesToCm[4] = 10.16

	fmt.Printf("danishNumbers: %v\n", danishNumbers)
	fmt.Printf("inchesToCm: %v\n", inchesToCm)
}

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

/*
Things not in Go
*/
