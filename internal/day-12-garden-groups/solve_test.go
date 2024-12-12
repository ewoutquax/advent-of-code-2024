package day12gardengroups_test

import (
	"testing"

	. "github.com/ewoutquax/advent-of-code-2024/internal/day-12-garden-groups"
	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	assert := assert.New(t)

	garden := ParseInput(testInput())

	assert.IsType(Garden{}, garden)
	assert.Len(garden.Locations, 16)
	assert.Len(garden.Regions, 5)
	assert.Len(garden.Regions[1], 4)
	assert.Len(garden.Regions[4], 1) // This is for plant 'D'
}

func TestCalculatePerimiters(t *testing.T) {
	garden := ParseInput(testInput())

	assert.Equal(t, 10, CalculatePerimiter(Region(1), garden))
}

func TestCalculateSides(t *testing.T) {
	garden := ParseInput(testInput())

	// Test the region of plant 'C'
	assert.Equal(t, 8, CalculateSides(Region(3), garden))

	// Test the region of plant 'A'
	assert.Equal(t, 4, CalculateSides(Region(1), garden))
}

func TestCalculatePrice(t *testing.T) {
	garden := ParseInput(testInput())

	assert.Equal(t, 40, CalculatePriceByPerimiter(Region(1), garden))
}

func TestSumPricesByPerimiter(t *testing.T) {
	garden := ParseInput(testInput())
	assert.Equal(t, 140, SumPricesByPerimiter(garden))

	gardenMedium := ParseInput(testInputMedium())
	assert.Equal(t, 772, SumPricesByPerimiter(gardenMedium))

	gardenLarge := ParseInput(testInputLarge())
	assert.Equal(t, 1930, SumPricesByPerimiter(gardenLarge))
}

func TestSumPricesBySides(t *testing.T) {
	garden := ParseInput(testInput())
	assert.Equal(t, 80, SumPricesBySides(garden))

	gardenMedium := ParseInput(testInputMedium())
	assert.Equal(t, 436, SumPricesBySides(gardenMedium))

	gardenLarge := ParseInput(testInputLarge())
	assert.Equal(t, 1206, SumPricesBySides(gardenLarge))

	gardenComplex := ParseInput(testInputComplex())
	assert.Equal(t, 236, SumPricesBySides(gardenComplex))
}

func testInput() []string {
	return []string{
		"AAAA",
		"BBCD",
		"BBCC",
		"EEEC",
	}
}

func testInputMedium() []string {
	return []string{
		"OOOOO",
		"OXOXO",
		"OOOOO",
		"OXOXO",
		"OOOOO",
	}
}

func testInputLarge() []string {
	return []string{
		"RRRRIICCFF",
		"RRRRIICCCF",
		"VVRRRCCFFF",
		"VVRCCCJFFF",
		"VVVVCJJCFE",
		"VVIVCCJJEE",
		"VVIIICJJEE",
		"MIIIIIJJEE",
		"MIIISIJEEE",
		"MMMISSJEEE",
	}
}

func testInputComplex() []string {
	return []string{
		"EEEEE",
		"EXXXX",
		"EEEEE",
		"EXXXX",
		"EEEEE",
	}
}
