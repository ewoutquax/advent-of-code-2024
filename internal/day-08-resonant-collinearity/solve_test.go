package day08resonantcollinearity_test

import (
	"fmt"
	"testing"

	. "github.com/ewoutquax/advent-of-code-2024/internal/day-08-resonant-collinearity"
	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	assert := assert.New(t)

	universe := ParseInput(testInput())

	assert.IsType(Universe{}, universe)
	assert.Equal(11, universe.MaxX)
	assert.Equal(11, universe.MaxY)
	assert.Len(universe.Signals, 2)
	assert.Len(universe.Signals["0"], 4)
	assert.Len(universe.Signals["A"], 3)
}

func TestAddAntinodesInEasyUniverse(t *testing.T) {
	assert := assert.New(t)

	universe := ParseInput(testInputEasy())

	AddAntinodes(&universe, false)

	assert.Len(universe.Signals, 1)
	assert.Len(universe.Antinodes, 2)

	antinodeLocs := make([]string, 0)
	for loc := range universe.Antinodes {
		antinodeLocs = append(antinodeLocs, loc.ToS())
	}

	assert.Contains(antinodeLocs, "[3, 1]")
	assert.Contains(antinodeLocs, "[6, 7]")
}

func TestAddAntinodesWithOffmapCheck(t *testing.T) {
	assert := assert.New(t)

	universe := ParseInput(testInputWithOffmapAntinodes())

	AddAntinodes(&universe, false)

	assert.Len(universe.Signals, 1)
	assert.Len(universe.Antinodes, 4)

	antinodeLocs := make([]string, 0)
	for loc := range universe.Antinodes {
		antinodeLocs = append(antinodeLocs, loc.ToS())
	}

	assert.Contains(antinodeLocs, "[3, 1]")
	assert.Contains(antinodeLocs, "[0, 2]")
	assert.Contains(antinodeLocs, "[2, 6]")
	assert.Contains(antinodeLocs, "[6, 7]")
}

func TestCountAntinodes(t *testing.T) {
	u := ParseInput(testInput())
	AddAntinodes(&u, false)

	for loc := range u.Antinodes {
		fmt.Printf("%s\n", loc.ToS())
	}

	assert.Equal(t, 14, CountAntinodes(u))
}

func TestCountAntinodesWithRepeat(t *testing.T) {
	u := ParseInput(testInput())
	AddAntinodes(&u, true)

	for loc := range u.Antinodes {
		fmt.Printf("%s\n", loc.ToS())
	}

	assert.Equal(t, 34, CountAntinodes(u))
}

func testInput() []string {
	return []string{
		"............",
		"........0...",
		".....0......",
		".......0....",
		"....0.......",
		"......A.....",
		"............",
		"............",
		"........A...",
		".........A..",
		"............",
		"............",
	}
}

func testInputEasy() []string {
	return []string{
		"..........",
		"..........",
		"..........",
		"....a.....",
		"..........",
		".....a....",
		"..........",
		"..........",
		"..........",
		"..........",
	}
}

func testInputWithOffmapAntinodes() []string {
	return []string{
		"..........",
		"..........",
		"..........",
		"....a.....",
		"........a.",
		".....a....",
		"..........",
		"..........",
		"..........",
		"..........",
	}
}
