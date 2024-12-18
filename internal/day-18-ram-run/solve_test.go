package day18ramrun_test

import (
	"image"
	"testing"

	. "github.com/ewoutquax/advent-of-code-2024/internal/day-18-ram-run"
	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	assert := assert.New(t)

	universe := BuildUniverse(
		WithMaxValues(6, 6),
		WithCorruptSpaces(testInput()[:12]),
	)

	assert.IsType(Universe{}, universe)
	assert.Len(universe.CorruptSpaces, 12)
	assert.True(universe.CorruptSpaces[image.Pt(5, 4)])
}

func TestFindPaths(t *testing.T) {
	universe := BuildUniverse(
		WithMaxValues(6, 6),
		WithCorruptSpaces(testInput()[:12]),
	)

	assert.Equal(t, 22, FindMinSteps(universe))
}

func TestFirstFullBlockingBte(t *testing.T) {
	universe := BuildUniverse(
		WithMaxValues(6, 6),
	)

	byte := FindFirstBlockingByte(universe, testInput())
	assert.Equal(t, "6,1", byte)
}

func testInput() []string {
	return []string{
		"5,4",
		"4,2",
		"4,5",
		"3,0",
		"2,1",
		"6,3",
		"2,4",
		"1,5",
		"0,6",
		"3,3",
		"2,6",
		"5,1",
		"1,2",
		"5,5",
		"2,5",
		"6,5",
		"1,4",
		"0,4",
		"6,4",
		"1,1",
		"6,1",
		"1,0",
		"0,5",
		"1,6",
		"2,0",
	}
}
