package day06guardgallivant_test

import (
	"testing"

	. "github.com/ewoutquax/advent-of-code-2024/internal/day-06-guard-gallivant"
	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	assert := assert.New(t)

	universe := ParseInput(testInput())

	assert.IsType(Universe{}, universe)
	assert.Len(universe.Blocks, 8)
	assert.Equal(9, universe.MaxX)
	assert.Equal(9, universe.MaxY)

	assert.Equal(4, universe.Guard.X)
	assert.Equal(6, universe.Guard.Y)
	assert.Equal(DirectionUp, universe.Guard.Direction)

	assert.Len(universe.VisitedLocations, 1)
}

func TestMoveGuard(t *testing.T) {
	assert := assert.New(t)

	universe := ParseInput(testInput())

	MoveGuard(&universe)

	assert.Equal(DirectionUp, universe.Guard.Direction)
	assert.Equal(4, universe.Guard.Location.X)
	assert.Equal(5, universe.Guard.Location.Y)

	assert.Len(universe.VisitedLocations, 2)
}

func TestTurnGuard(t *testing.T) {
	assert := assert.New(t)

	universe := ParseInput(testInput())

	for i := 0; i < 6; i++ {
		MoveGuard(&universe)
	}

	assert.Equal(DirectionRight, universe.Guard.Direction)
	assert.Equal(4, universe.Guard.Location.X)
	assert.Equal(1, universe.Guard.Location.Y)

	assert.Len(universe.VisitedLocations, 6)
}

func TestCountDistinctLocations(t *testing.T) {
	assert := assert.New(t)

	universe := ParseInput(testInput())

	MoveGuardOffMap(&universe)

	// The one location outside the map is also counted here
	assert.Len(universe.VisitedLocations, 42)
}

func TestDetectLooping(t *testing.T) {
	universe := ParseInput(testInput())

	universe.Blocks[Location{3, 6}] = true
	assert.True(t, GuardIsLooping(&universe))
	universe.Reset()
	delete(universe.Blocks, Location{3, 6})

	universe.Blocks[Location{6, 7}] = true
	assert.True(t, GuardIsLooping(&universe))
	universe.Reset()
	delete(universe.Blocks, Location{6, 7})

	universe.Blocks[Location{7, 7}] = true
	assert.True(t, GuardIsLooping(&universe))
	universe.Reset()
	delete(universe.Blocks, Location{7, 7})

	universe.Reset()
	assert.False(t, GuardIsLooping(&universe))
}

func TestCountLoopingBlocks(t *testing.T) {
	universe := ParseInput(testInput())

	assert.Equal(t, 6, CountLoopingBlocks(universe))
}

func testInput() []string {
	return []string{
		"....#.....",
		".........#",
		"..........",
		"..#.......",
		".......#..",
		"..........",
		".#..^.....",
		"........#.",
		"#.........",
		"......#...",
	}
}
