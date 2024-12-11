package day10hoofit_test

import (
	"testing"

	. "github.com/ewoutquax/advent-of-code-2024/internal/day-10-hoof-it"
	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	assert := assert.New(t)

	topographicMap := ParseInput(testInput())

	assert.IsType(TopographicMap{}, topographicMap)
	assert.Len(topographicMap.Locations, 64)
	assert.Len(topographicMap.Trailheads, 9)
	assert.Equal(8, topographicMap.Locations[Coordinate{0, 0}].Height)
	assert.Len(topographicMap.Trailheads[Coordinate{2, 0}].Ends, 0)
	assert.Equal(0, topographicMap.Trailheads[Coordinate{2, 0}].Rating)
}

func TestFindTrailheadsEnds(t *testing.T) {
	topographicMap := ParseInput(testInput())

	for _, th := range topographicMap.Trailheads {
		if th.Coordinate.X != 2 || th.Coordinate.Y != 0 {
			delete(topographicMap.Trailheads, th.Coordinate)
		}
	}

	FindTrailheadsEnds(topographicMap)

	assert.Len(t, topographicMap.Trailheads[Coordinate{2, 0}].Ends, 5)
	assert.Equal(t, topographicMap.Trailheads[Coordinate{2, 0}].Rating, 20)
}

func TestSumTrailheadScores(t *testing.T) {
	topographicMap := ParseInput(testInput())

	FindTrailheadsEnds(topographicMap)

	assert.Equal(t, 36, SumTrailheadScores(topographicMap))
}

func TestSumTrailheadRating(t *testing.T) {
	topographicMap := ParseInput(testInput())

	FindTrailheadsEnds(topographicMap)

	assert.Equal(t, 81, SumTrailheadRatings(topographicMap))
}

func testInput() []string {
	return []string{
		"89010123",
		"78121874",
		"87430965",
		"96549874",
		"45678903",
		"32019012",
		"01329801",
		"10456732",
	}
}
