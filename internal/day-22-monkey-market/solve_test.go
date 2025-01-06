package day22monkeymarket_test

import (
	"fmt"
	"testing"

	. "github.com/ewoutquax/advent-of-code-2024/internal/day-22-monkey-market"
	"github.com/stretchr/testify/assert"
)

// func TestParseInput(t *testing.T) {
// 	t.Skip("Ff geduld nog")
//
// 	assert := assert.New(t)
//
// 	universe := ParseInput(testInput())
//
// 	assert.IsType(Universe{}, universe)
// }

func TestMixNumbers(t *testing.T) {
	assert.Equal(t, Number(37), Number(15).Mix(42))
}

func TestPruneSecret(t *testing.T) {
	assert.Equal(t, Number(16113920), PruneSecret(100000000))
}

func TestCalculateNextPrice(t *testing.T) {
	testCases := map[Number]Number{
		123:      15887950,
		15887950: 16495136,
		16495136: 527345,
		527345:   704524,
		704524:   1553684,
		1553684:  12683156,
		12683156: 11100544,
		11100544: 12249484,
		12249484: 7753432,
		7753432:  5908254,
	}

	for inputSecret, expectedResult := range testCases {
		assert.Equal(t, expectedResult, CalculateNextPrice(inputSecret))
	}
}

func TestCalculate2000thSecret(t *testing.T) {
	testCases := map[Number]Number{
		1:    8685429,
		10:   4700978,
		100:  15273692,
		2024: 8667524,
	}

	for inputSecret, expectedResult := range testCases {
		assert.Equal(t, expectedResult, Calculate2000thSecret(inputSecret))
	}
}

func TestSum2000thSecrets(t *testing.T) {
	numbers := ParseInput(testInput())

	result := Sum2000thSecret(numbers)

	assert.Equal(t, 37327623, result)
}

func TestCalculateMostBananas(t *testing.T) {
	numbers := ParseInput(testInput2())
	assert.Equal(t, 23, CalculateMostBananas(numbers))
}

func testInput() []string {
	return []string{
		"1",
		"10",
		"100",
		"2024",
	}
}

func testInput2() []string {
	return []string{
		"1",
		"2",
		"3",
		"2024",
	}
}
