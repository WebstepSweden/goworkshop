package goworkshop

import "strings"

// SumTo100 loops through numbers 1 to 100, and returns their sum
func SumTo100() int {
	ret := 0
	for i := 1; i <= 100; i++ {
		ret += i
	}
	return ret
}

var sentence = "I have never thought I'm gonna need to give blood again. You saw me up there, right?"

// LoopWords loops through a sentence and returns a new string
// with every third word from the sentence (beginning with the third word).
// Have a look at the strings library!
func LoopWords() string {
	ret := ""
	for index, word := range strings.Split(sentence, " ") {
		if (index+1)%3 == 0 {
			ret += word + " "
		}
	}
	return ret
}
