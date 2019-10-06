// Package goworkshop includes, not too surprisingly, a workshop about the Go programing language
package goworkshop

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strings"
	"testing"
	"time"
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
//

//    ____                                     _
//  / ___| ___      __ _  ___      __ _  ___ | |
// | |  _ / _ \    / _` |/ _ \    / _` |/ _ \| |
// | |_| | (_) |  | (_| | (_) |  | (_| | (_) |_|
//  \____|\___/    \__, |\___/    \__, |\___/(_)
//                 |___/          |___/

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
// * https://godoc.org/
// * https://play.golang.org/
// * https://tour.golang.org/
// * https://twitter.com/search?q=%23golang
// * https://www.reddit.com/r/golang/

//  ____
// |  _ \   ___ __   __    ___  _ __ __   __
// | | | | / _ \\ \ / /   / _ \| '_ \\ \ / /
// | |_| ||  __/ \ V /   |  __/| | | |\ V /
// |____/  \___|  \_/     \___||_| |_| \_/

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

//  _____                            _    _
// | ____|__  __ _ __    ___   _ __ | |_ (_) _ __    __ _
// |  _|  \ \/ /| '_ \  / _ \ | '__|| __|| || '_ \  / _` |
// | |___  >  < | |_) || (_) || |   | |_ | || | | || (_| |
// |_____|/_/\_\| .__/  \___/ |_|    \__||_||_| |_| \__, |
//              |_|                                 |___/

// Defining identifiers in Go (such as functions and fields), you can choose to export them
// or keep them local.
//
// Exporting means they will be available to other packages for import. So writing things like:
// strings.ToLower("HELLO")
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

//                 __             _      ___
//   __ _   ___   / _| _ __ ___  | |_   ( _ )
//  / _` | / _ \ | |_ | '_ ` _ \ | __|  / _ \/\
// | (_| || (_) ||  _|| | | | | || |_  | (_>  <
//  \__, | \___/ |_|  |_| |_| |_| \__|  \___/\/
//  |___/         _  _         _
//   __ _   ___  | |(_) _ __  | |_
//  / _` | / _ \ | || || '_ \ | __|
// | (_| || (_) || || || | | || |_
//  \__, | \___/ |_||_||_| |_| \__|
//  |___/

// Go is quite opinionated about the coding style and defines some rigid rules to control the looks of
// your code. In order to keep things nice and clean, there are two tools you need to use, either using
// the terminal or integrating them in your IDE. These are:
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

// __     __            _         _      _
// \ \   / /__ _  _ __ (_)  __ _ | |__  | |  ___  ___
//  \ \ / // _` || '__|| | / _` || '_ \ | | / _ \/ __|
//   \ V /| (_| || |   | || (_| || |_) || ||  __/\__ \
//    \_/  \__,_||_|   |_| \__,_||_.__/ |_| \___||___/

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

//  _____                     _    _
// |  ___|_   _  _ __    ___ | |_ (_)  ___   _ __   ___
// | |_  | | | || '_ \  / __|| __|| | / _ \ | '_ \ / __|
// |  _| | |_| || | | || (__ | |_ | || (_) || | | |\__ \
// |_|    \__,_||_| |_| \___| \__||_| \___/ |_| |_||___/

// Functions have the following syntax:
// func functionName(arg1Name arg1Type, arg2Name arg2Type, ...etc) returnType {
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

// or even like this:

func returningTuple(s string) (int, int, int) {
	return 1, 2, 3
}

// Function resources:
// * https://tour.golang.org/basics/4
// * https://www.callicoder.com/golang-functions/

//  _____               _
// |  ___|___   _ __   | |  ___    ___   _ __   ___
// | |_  / _ \ | '__|  | | / _ \  / _ \ | '_ \ / __|
// |  _|| (_) || |     | || (_) || (_) || |_) |\__ \
// |_|   \___/ |_|     |_| \___/  \___/ | .__/ |___/
//                                      |_|

// For loops are the only way you can loop through stuff in Go. There are different ways to use them though:

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

//   ____        _  _              _    _
//  / ___| ___  | || |  ___   ___ | |_ (_)  ___   _ __   ___
// | |    / _ \ | || | / _ \ / __|| __|| | / _ \ | '_ \ / __|
// | |___| (_) || || ||  __/| (__ | |_ | || (_) || | | |\__ \
//  \____|\___/ |_||_| \___| \___| \__||_| \___/ |_| |_||___/

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

// Collections resources:
// * http://www.golangbootcamp.com/book/collection_types
// * https://www.golang-book.com/books/intro/6

//  ____   _                       _
// / ___| | |_  _ __  _   _   ___ | |_  ___
// \___ \ | __|| '__|| | | | / __|| __|/ __|
//  ___) || |_ | |   | |_| || (__ | |_ \__ \
// |____/  \__||_|    \__,_| \___| \__||___/

// Structs are objects that group some fields together. Here's one struct:

// URL represents a url with a protocol, a domain and a path
type URL struct {
	protocol string
	domain   string
	port     int
}

// They are used like this:

// CreateStruct creates a URL struct and prints it
func CreateStruct() {
	url := URL{
		"http",
		"google.se",
		80,
	}

	fmt.Printf("URL struct: %+v\n", url)
}

// Structs resources:
// * https://gobyexample.com/structs
// * https://medium.com/rungo/structures-in-go-76377cc106a2

//       _          _                        ___       __             _
//  ___ | |_  _ __ (_) _ __    __ _  ___    ( _ )     / _| _ __ ___  | |_
// / __|| __|| '__|| || '_ \  / _` |/ __|   / _ \/\  | |_ | '_ ` _ \ | __|
// \__ \| |_ | |   | || | | || (_| |\__ \  | (_>  <  |  _|| | | | | || |_
// |___/ \__||_|   |_||_| |_| \__, ||___/   \___/\/  |_|  |_| |_| |_| \__|
//                            |___/

// strings and fmt are two very useful packages that one uses quite often when dealing
// with strings. Here are some example usages:

// UseStrings shows some examples of using the strings package
func UseStrings(word string) {
	println(strings.ToLower(word))
	println(strings.ToUpper(word))
	println(strings.Contains(word, "am"))
	println(strings.Index(word, "am"))
	println(strings.Join([]string{word, word, word}, " & "))
	println(strings.Repeat(word, 7))
}

// UseFmt shows some examples of using the fmt package
func UseFmt() {
	fmt.Print(1, 2, 3)
	println()

	fmt.Printf("Pi: %v", math.Pi)
	println()

	waitForIt := "...wait for it..."
	fmt.Fprintf(os.Stdout, "It's going to be legen %v dary", waitForIt)
	println()
}

// strings and fmt resources:
// * https://golang.org/pkg/strings/
// * https://golang.org/pkg/fmt/

//  ____         _         _
// |  _ \  ___  (_) _ __  | |_  ___  _ __  ___
// | |_) |/ _ \ | || '_ \ | __|/ _ \| '__|/ __|
// |  __/| (_) || || | | || |_|  __/| |   \__ \
// |_|    \___/ |_||_| |_| \__|\___||_|   |___/

// Pointers are variables that contains a reference to other variables.
//
// To create a pointer, use the & symbol:
// pointerToSomething := &something
// => pointerToSomething will now contain the address of variable something
//
// To follow a pointer, use the * symbol:
// *pointerToSomething = "hi"
// => the data in variable something will now be overwritten

// UsePointers shows how to use pointers
func UsePointers() {
	something := "Hello"
	pointerToSomething := &something
	*pointerToSomething += " world"

	println(something)
}

// Pointer resources:
// * https://tour.golang.org/moretypes/1
// * https://medium.com/rungo/pointers-in-go-a789eafccd53

//  _____
// | ____| _ __  _ __  ___   _ __  ___
// |  _|  | '__|| '__|/ _ \ | '__|/ __|
// | |___ | |   | |  | (_) || |   \__ \
// |_____||_|   |_|   \___/ |_|   |___/

// Sometimes, a function wants to indicate to the caller whether everything went well or not.
// The idiomatic way to do so is to return two values: the return value and an error indicator.
// The caller would check if any errors were flagged, or else use the return value.

// Divide two numbers. If the denominator is zero, an error will be returned
func Divide(x float32, y float32) (float32, error) {
	if y == 0 {
		return 0, errors.New("division by zero")
	}
	return x / y, nil
}

// DivisionRunner runs divisions given two numbers
func DivisionRunner(x float32, y float32) {
	result, err := Divide(x, y)

	if err == nil {
		fmt.Printf("%v / %v = %v", x, y, result)
	} else {
		fmt.Printf("error: %v", err)
	}
	println()
}

// Errors resources:
// * https://gobyexample.com/errors
// * https://about.sourcegraph.com/go/gophercon-2019-handling-go-errors

//  _____           _    _
// |_   _|___  ___ | |_ (_) _ __    __ _
//   | | / _ \/ __|| __|| || '_ \  / _` |
//   | ||  __/\__ \| |_ | || | | || (_| |
//   |_| \___||___/ \__||_||_| |_| \__, |
//                                 |___/

// There is no library like Junit in Go. Instead we're using the testing package and if/else.

// GetANumber is the function we're testing. It returns a random number.
func GetANumber() int {
	rand.Seed(time.Now().UnixNano()) // use a different seeds every time
	return rand.Intn(100)            // generate a random number under 100
}

// TestGetANumber tests the GetANumber function
func TestGetANumber(t *testing.T) {
	want := 8
	got := GetANumber()
	if got != want {
		t.Fatalf(`Fail! Wanted '%v', got '%v'`, want, got)
	}
}

// Usually, we also add a benchmark test

// BenchmarkGetANumber tests function GetANumber's performance
func BenchmarkGetANumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetANumber()
	}
}

// The result might look like this:
//
// goos: darwin
// goarch: amd64
// pkg: exercism.io/github.com/webstep/goworkshop
// BenchmarkGetANumber_-12    	  163783	      7264 ns/op
// PASS
//
// This means that the GetANumber function ran 163783 times at an average speed of 7264 ns per loop.
// (The -12 thingy means it was run on 12 processors)
//
// Creating tests, you must follow these instructions:
// * Tests must be inside a file that ends with _test
// * Test functions must begin with func TestXxx(t *testing.T) - the first letter after Test must be capital
// * Benchmark functions must begin with func BenchmarkXxx(b *testing.B) - again, the first letter
//   after Benchmark must be capital
//
// Testing resources:
// * https://golang.org/pkg/testing/
// * https://marcofranssen.nl/test-and-benchmark-your-code-in-go/
