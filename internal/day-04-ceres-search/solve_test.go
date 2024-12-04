package day04ceressearch_test

import (
	"testing"

	. "github.com/ewoutquax/advent-of-code-2024/internal/day-04-ceres-search"
	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	assert := assert.New(t)

	puzzle := ParseInput(testInput())

	assert.IsType(Puzzle{}, puzzle)
	assert.Len(puzzle.Letters, 100)
	assert.Equal(10, puzzle.MaxX)
	assert.Equal(10, puzzle.MaxY)
}

func TestCountOccurances(t *testing.T) {
	puzzle := ParseInput(testInput())

	assert.Equal(t, 18, CountOccurances(puzzle, "XMAS"))
}

func TestCountXmasOccurences(t *testing.T) {
	puzzle := ParseInput(testInput())

	assert.Equal(t, 9, CountXmasOccurances(puzzle, "MAS"))
}

func testInput() []string {
	return []string{
		"MMMSXXMASM",
		"MSAMXMSMSA",
		"AMXSXMAAMM",
		"MSAMASMSMX",
		"XMASAMXAMM",
		"XXAMMXXAMA",
		"SMSMSASXSS",
		"SAXAMASAAA",
		"MAMMMXMMMM",
		"MXMXAXMASX",
	}
}
