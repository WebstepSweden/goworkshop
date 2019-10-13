package goworkshop

import (
	"errors"
	"fmt"
	"strings"
)

// Fraction contains a numerator and a denominator
type Fraction struct {
	numerator   int
	denominator int
}

// Gcd finds the greatest common denominator of two numbers
func Gcd(a int, b int) int {
	if b == 0 {
		return a
	}
	return Gcd(b, a%b)
}

// CreateFraction creates a Fraction given a numerator and a denominator and returns a pointer to it
// If the denominator is zero, it returns an error with text "denominator must not be zero"
// When creating the Fraction, this function uses the Gcd function to simplify the Fraction
func CreateFraction(num int, denom int) (*Fraction, error) {
	if denom == 0 {
		return nil, errors.New("denominator must not be zero")
	}
	gcd := Gcd(num, denom)
	return &Fraction{num / gcd, denom / gcd}, nil
}


// AddFractions adds two fraction pointers and return a pointer to the result which is a new Fraction
// If the result denominator is zero, it returns an error with text "denominator must not be zero"
// As a reminder, adding two fractions is calculated like this:
// a/x + b/y = (ay + bx)/xy
func AddFractions(f1 *Fraction, f2 *Fraction) (*Fraction, error) {
	num := f1.numerator*f2.denominator + f2.numerator*f1.denominator
	denom := f1.denominator * f2.denominator
	return CreateFraction(num, denom)
}

// PrintFractions returns a string representation of the given Fraction pointers
// examples:
// PrintFractions([]*Fraction{ {1, 2} }) => Fractions: 1/2
// PrintFractions([]*Fraction{ {1, 2}, {2, 3} }) => Fractions: 1/2, 2/3
func PrintFractions(fractions []*Fraction) string {
	result := "Fractions: "
	for _, f := range fractions {
		result += fmt.Sprintf("%v/%v, ", f.numerator, f.denominator)
	}
	return strings.TrimRight(result, ", ")
}
