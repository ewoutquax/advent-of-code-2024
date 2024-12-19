package day19linenlayout_test

import (
	"testing"

	. "github.com/ewoutquax/advent-of-code-2024/internal/day-19-linen-layout"
	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	assert := assert.New(t)

	universe := ParseInput(testInput())

	assert.IsType(Universe{}, universe)
	assert.Len(universe.Patterns, 8)
	assert.Len(universe.Towels, 8)
	assert.Equal(Pattern("r"), universe.Patterns[0])
	assert.Equal(Pattern("br"), universe.Patterns[len(universe.Patterns)-1])
	assert.Equal(Towel("brwrr"), universe.Towels[0])
	assert.Equal(Towel("bbrgwb"), universe.Towels[len(universe.Towels)-1])
}

func TestTowelIsPossible(t *testing.T) {
	universe := ParseInput(testInput())

	testCases := map[Towel]bool{
		universe.Towels[0]: true,
		universe.Towels[1]: true,
		universe.Towels[2]: true,
		universe.Towels[3]: true,
		universe.Towels[4]: false,
		universe.Towels[5]: true,
		universe.Towels[6]: true,
		universe.Towels[7]: false,
	}

	for inputTowel, expectedResult := range testCases {
		assert.Equal(t, expectedResult, inputTowel.CountCombinations(universe.Patterns) > 0)
	}
}

func TestCountPossibleTowels(t *testing.T) {
	universe := ParseInput(testInput())

	assert.Equal(t, 6, CountPossibleTowels(universe))
}

func TestCountPossibleCombinations(t *testing.T) {
	universe := ParseInput(testInput())

	testCases := map[Towel]int{
		"brwrr":  2,
		"bggr":   1,
		"gbbr":   4,
		"rrbgbr": 6,
		"ubwu":   0,
		"bbrgwb": 0,
	}

	for inputTowel, expectedResult := range testCases {
		assert.Equal(t, expectedResult, inputTowel.CountCombinations(universe.Patterns))
	}
}

func TestSumPossibleCombinations(t *testing.T) {
	universe := ParseInput(testInput())

	assert.Equal(t, 16, SumPossibleCombinations(universe))
}

func testInput() [][]string {
	return [][]string{
		{
			"r, wr, b, g, bwu, rb, gb, br",
		}, {
			"brwrr",
			"bggr",
			"gbbr",
			"rrbgbr",
			"ubwu",
			"bwurrg",
			"brgr",
			"bbrgwb",
		},
	}
}
