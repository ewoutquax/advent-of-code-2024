package day11plutonianpebbles_test

import (
	"fmt"
	"testing"

	. "github.com/ewoutquax/advent-of-code-2024/internal/day-11-plutonian-pebbles"
	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	assert := assert.New(t)

	rocks := ParseInput(testInput())

	assert.Equal("[]day11plutonianpebbles.Rock", fmt.Sprintf("%T", rocks))
	assert.Len(rocks, 2)
	assert.Equal(Rock(125), rocks[0])
	assert.Equal(Rock(17), rocks[len(rocks)-1])
}

func TestRockNextRocks(t *testing.T) {
	testCases := map[Rock][]Rock{
		0:    {1},
		10:   {1, 0},
		1024: {10, 24},
		8000: {80, 0},
		8:    {16192},
	}

	for inputRock, expectedOutput := range testCases {
		actualOutput := inputRock.NextRocks()
		assert.Equal(t, expectedOutput, actualOutput)
	}
}

func TestCountRocksAfterBlinks(t *testing.T) {
	testCases := map[int]int{
		6:  22,
		25: 55312,
	}

	rocks := ParseInput(testInput())
	for blinks, expectedCount := range testCases {
		actualCount := CountRocksAfterBlinks(rocks, blinks)
		assert.Equal(t, expectedCount, actualCount)
	}
}

func testInput() string {
	return "125 17"
}
