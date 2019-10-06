package goworkshop

var createFractionTestCases = []struct {
	description string
	numerator   int
	denominator int
	expected    Fraction
}{
	{
		"simple fraction, 2/3",
		2, 3,
		Fraction{2, 3},
	},
	{
		"fraction with gcd, 4/6 -> 2/3",
		4, 6,
		Fraction{2, 3},
	},
	{
		"fraction > 1, 3/2",
		3, 2,
		Fraction{3, 2},
	},
	{
		"fraction > 1 with gcd, 200/50 -> 4/1",
		200, 50,
		Fraction{4, 1},
	},
}

var addFractionsTestCases = []struct {
	description  string
	numerator1   int
	denominator1 int
	numerator2   int
	denominator2 int
	expected     Fraction
}{
	{
		"simple addition, 1/2 + 2/3 = 5/6",
		1, 2,
		1, 3,
		Fraction{5, 6},
	},
	{
		"addition with gcd, 1/2 + 1/4 = 3/4",
		1, 2,
		1, 4,
		Fraction{3, 4},
	},
	{
		"simple addition with sum > 1, 2/3 + 2/3 = 4/3",
		2, 3,
		2, 3,
		Fraction{4, 3},
	},
	{
		"simple addition with sum > 1 and gcd, 1/2 + 3/4 = 4/3",
		1, 2,
		3, 4,
		Fraction{5, 4},
	},
}

var printFractionsTestCases = []struct {
	description string
	fractions   []*Fraction
	expected    string
}{
	{
		"print one fraction",
		[]*Fraction{
			{1, 2},
		},
		"Fractions: 1/2",
	},
	{
		"print two fractions",
		[]*Fraction{
			{1, 2},
			{2, 3},
		},
		"Fractions: 1/2, 2/3",
	},
	{
		"print three fractions",
		[]*Fraction{
			{1, 2},
			{2, 3},
			{3, 4},
		},
		"Fractions: 1/2, 2/3, 3/4",
	},
}
