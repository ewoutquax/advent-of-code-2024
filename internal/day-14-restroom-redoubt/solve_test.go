package day14restroomredoubt_test

import (
	. "image"
	"testing"

	. "github.com/ewoutquax/advent-of-code-2024/internal/day-14-restroom-redoubt"
	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	assert := assert.New(t)

	bots := ParseInput(testInput())

	assert.IsType([]Bot{}, bots)
	assert.Equal(0, bots[0].Location.X)
	assert.Equal(4, bots[0].Location.Y)
	assert.Equal(-3, bots[len(bots)-1].Vector.X)
	assert.Equal(-3, bots[len(bots)-1].Vector.Y)
}

func TestMoveBot(t *testing.T) {
	bots := ParseInput([]string{"p=2,4 v=2,-3"})
	bot := bots[0]

	expectedLocations := []Point{
		// Pt(2, 4),
		Pt(4, 1),
		Pt(6, 5),
		Pt(8, 2),
		Pt(10, 6),
		Pt(1, 3),
	}

	for _, expectedLocation := range expectedLocations {
		bot.Move(1)
		assert.Equal(t, expectedLocation.String(), bot.Location.String())
	}
}

func TestLocationToQuadrant(t *testing.T) {
	testCases := map[Point]Quadrant{
		Pt(0, 2): 0,
		Pt(9, 0): 1,
		Pt(5, 4): -1,
		Pt(6, 6): 3,
	}

	for point, expectedQuadrant := range testCases {
		assert.Equal(t, expectedQuadrant, LocationToQuadrant(point))
	}
}

func TestCalculateSafetyFactorAfterMoves(t *testing.T) {
	bots := ParseInput(testInput())

	factor := CalculateSafetyFactorAfterMoves(bots, 100)

	assert.Equal(t, 12, factor)
}

func testInput() []string {
	return []string{
		"p=0,4 v=3,-3",
		"p=6,3 v=-1,-3",
		"p=10,3 v=-1,2",
		"p=2,0 v=2,-1",
		"p=0,0 v=1,3",
		"p=3,0 v=-2,-2",
		"p=7,6 v=-1,-3",
		"p=3,0 v=-1,-2",
		"p=9,3 v=2,3",
		"p=7,3 v=-1,2",
		"p=2,4 v=2,-3",
		"p=9,5 v=-3,-3",
	}
}
