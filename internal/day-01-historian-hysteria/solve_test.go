package day01historianhysteria_test

import (
	"testing"

	. "github.com/ewoutquax/advent-of-code-2024/internal/day-01-historian-hysteria"
	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	assert := assert.New(t)

	lists := ParseInput(testInput())

	assert.IsType(Lists{}, lists)

	assert.Len(lists.Left, 6)
	assert.Equal(3, lists.Left[0])
	assert.Equal(3, lists.Left[len(lists.Left)-1])

	assert.Len(lists.Right, 6)
	assert.Equal(4, lists.Right[0])
	assert.Equal(3, lists.Right[len(lists.Right)-1])
}

func TestSumDistanceBetweenSmallest(t *testing.T) {
	lists := ParseInput(testInput())
	assert.Equal(t, 11, SumDistanceBetweenSmallest(lists))
}

func TestSumSimilarityScore(t *testing.T) {
	lists := ParseInput(testInput())
	assert.Equal(t, 31, SimilarityScore(lists))
}

func testInput() []string {
	return []string{
		"3   4",
		"4   3",
		"2   5",
		"1   3",
		"3   9",
		"3   3",
	}
}
