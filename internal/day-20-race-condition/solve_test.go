package day20racecondition_test

import (
	"fmt"
	. "image"
	"testing"

	. "github.com/ewoutquax/advent-of-code-2024/internal/day-20-race-condition"
	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	assert := assert.New(t)

	track := ParseInput(testInput())

	assert.IsType(Track{}, track)
	assert.Len(track.Walls, 140)
	assert.True(track.Start.Eq(Pt(1, 3)))
	assert.True(track.Finish.Eq(Pt(5, 7)))
	assert.Equal(14, track.MaxX)
	assert.Equal(14, track.MaxY)
}

func TestFindPaths(t *testing.T) {
	track := ParseInput(testInput())

	solutionsByNrSteps2, _ := FindPaths(track, 0, 2)
	fmt.Printf("solutionsByNrSteps: %v\n", solutionsByNrSteps2)

	nrStepsSavedGrouped2 := map[int]int{
		2:  14,
		4:  14,
		6:  2,
		8:  4,
		10: 2,
		12: 3,
		20: 1,
		36: 1,
		38: 1,
		40: 1,
		64: 1,
	}

	for nrStepsSaved, expectedCount := range nrStepsSavedGrouped2 {
		assert.Equal(t, expectedCount, solutionsByNrSteps2[nrStepsSaved])
	}

	solutionsByNrSteps20, _ := FindPaths(track, 50, 20)
	fmt.Printf("solutionsByNrSteps: %v\n", solutionsByNrSteps20)

	nrStepsSavedGrouped20 := map[int]int{
		50: 32,
		52: 31,
		54: 29,
		56: 39,
		58: 25,
		60: 23,
		62: 20,
		64: 19,
		66: 12,
		68: 14,
		70: 12,
		72: 22,
		74: 4,
		76: 3,
	}

	for nrStepsSaved, expectedCount := range nrStepsSavedGrouped20 {
		assert.Equal(t, expectedCount, solutionsByNrSteps20[nrStepsSaved])
	}
}

func TestPathToKey(t *testing.T) {
	path1 := Path{
		Point:   Point{},
		NrSteps: 0,
	}
	assert.Equal(t, PathKey("(0,0)"), path1.ToKey())
}

func testInput() []string {
	return []string{
		"###############",
		"#...#...#.....#",
		"#.#.#.#.#.###.#",
		"#S#...#.#.#...#",
		"#######.#.#.###",
		"#######.#.#...#",
		"#######.#.###.#",
		"###..E#...#...#",
		"###.#######.###",
		"#...###...#...#",
		"#.#####.#.###.#",
		"#.#...#.#.#...#",
		"#.#.#.#.#.#.###",
		"#...#...#...###",
		"###############",
	}
}
