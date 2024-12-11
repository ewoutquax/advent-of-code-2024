package day11plutonianpebbles_test

import (
	"fmt"
	"reflect"
	"testing"

	. "github.com/ewoutquax/advent-of-code-2024/internal/day-11-plutonian-pebbles"
	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	assert := assert.New(t)

	stones := ParseInput(testInput())

	assert.Equal("[]day11plutonianpebbles.Stone", fmt.Sprintf("%s", reflect.TypeOf(stones)))

	assert.Len(stones, 5)
	assert.Equal(Stone(0), stones[0])
	assert.Equal(Stone(999), stones[len(stones)-1])
}

func TestStoneResolve0to1(t *testing.T) {
	stones := ParseInput(testInput())

	assert.Equal(t, []Stone{1}, stones[0].Resolve())
}

func TestStoneResolve1to2024(t *testing.T) {
	stones := ParseInput(testInput())

	assert.Equal(t, []Stone{2024}, stones[1].Resolve())
}

func TestStoneResolve10To1And0(t *testing.T) {
	stones := ParseInput(testInput())

	assert.Equal(t, []Stone{1, 0}, stones[2].Resolve())
}

func TestBlinkOnce(t *testing.T) {
	stones := ParseInput(testInput())

	assert.Equal(t, 7, NrStonesAfterBlinks(stones, 1))
}

func TestBlink25Times(t *testing.T) {
	stones := ParseInput(testInput2())

	assert.Equal(t, 55312, NrStonesAfterBlinks(stones, 25))
}

func testInput() string {
	return "0 1 10 99 999"
}

func testInput2() string {
	return "125 17"
}
