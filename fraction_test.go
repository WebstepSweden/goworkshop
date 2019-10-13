package goworkshop

import (
	"errors"
	"testing"
)

func TestCreateFraction(t *testing.T) {
	for _, testCase := range createFractionTestCases {
		fraction, err := CreateFraction(testCase.numerator, testCase.denominator)

		if fraction == nil {
			t.Fatalf(`Fail in test '%v'. Got nil`, testCase.description)
		}

		if err != nil {
			t.Fatalf(`Fail in test '%v'. Got error '%v'`, testCase.description, err)
		}

		want := testCase.expected
		got := *fraction

		if got != want {
			t.Fatalf(`Fail in test '%v'. Wanted '%v', got '%v'`, testCase.description, want, got)
		}
	}
}

func TestCreateFractionError(t *testing.T) {
	want := errors.New("denominator must not be zero")
	_, err := CreateFraction(2, 0)

	if err == nil {
		t.Fatalf(`Fail in test. Wanted error '%v', got nil`, want)
	}

	if err.Error() != want.Error() {
		t.Fatalf(`Fail in test. Wanted error '%v', got '%v'`, want, err)
	}
}

func TestAddTwoFractions(t *testing.T) {
	for _, testCase := range addFractionsTestCases {
		fraction1, err1 := CreateFraction(testCase.numerator1, testCase.denominator1)
		fraction2, err2 := CreateFraction(testCase.numerator2, testCase.denominator2)

		if fraction1 == nil || fraction2 == nil {
			t.Fatalf(`Fail in test '%v'. Got nil`, testCase.description)
		}

		if err1 != nil {
			t.Fatalf(`Fail in test '%v'. Got error '%v'`, testCase.description, err1)
		}

		if err2 != nil {
			t.Fatalf(`Fail in test '%v'. Got error '%v'`, testCase.description, err2)
		}

		want := testCase.expected
		result, err := AddFractions(fraction1, fraction2)

		if result == nil {
			t.Fatalf(`Fail in test '%v'. Got nil`, testCase.description)
		}

		if err != nil {
			t.Fatalf(`Fail in test '%v'. Got error '%v'`, testCase.description, err)
		}

		got := *result

		if got != want {
			t.Fatalf(`Fail in test '%v'. Wanted '%v', got '%v'`, testCase.description, want, got)
		}
	}
}

func TestPrintFractions(t *testing.T) {
	for _, testCase := range printFractionsTestCases {
		want := testCase.expected
		got := PrintFractions(testCase.fractions)

		if got != want {
			t.Fatalf(`Fail in test '%v'. Wanted '%v', got '%v'`, testCase.description, want, got)
		}
	}
}
