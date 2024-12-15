package day13clawcontraption_test

import (
	"fmt"
	"testing"

	. "github.com/ewoutquax/advent-of-code-2024/internal/day-13-claw-contraption"
	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	assert := assert.New(t)

	arcades := ParseInput(testInput(), 0)

	// assert.IsType([]Arcade, arcades)
	assert.Len(arcades, 4)
	assert.Equal(94, arcades[0].ButtonA.X)
	assert.Equal(71, arcades[len(arcades)-1].ButtonB.Y)
	assert.Equal(12748, arcades[1].Prize.X)
	assert.Equal(6450, arcades[2].Prize.Y)
}

func TestFindCheapestSolution(t *testing.T) {
	arcades := ParseInput(testInput(), 0)

	testCases := map[Arcade]int{
		arcades[0]: 280,
		arcades[2]: 200,
	}

	for input, expectedResult := range testCases {
		tokens, err := FindCheapestSolution(input)
		assert.Nil(t, err)
		assert.Equal(t, expectedResult, tokens)
	}

	impossibleArcades := []Arcade{arcades[1], arcades[3]}
	for _, impossible := range impossibleArcades {
		token, err := FindCheapestSolution(impossible)
		fmt.Printf("token: %v\n", token)
		assert.NotNil(t, err)
	}
}

func TestSumCheapestSolutions(t *testing.T) {
	arcades := ParseInput(testInput(), 0)

	assert.Equal(t, 480, SumCheapestSolutions(arcades))
}

func testInput() [][]string {
	return [][]string{
		{
			"Button A: X+94, Y+34",
			"Button B: X+22, Y+67",
			"Prize: X=8400, Y=5400",
		}, {
			"Button A: X+26, Y+66",
			"Button B: X+67, Y+21",
			"Prize: X=12748, Y=12176",
		}, {
			"Button A: X+17, Y+86",
			"Button B: X+84, Y+37",
			"Prize: X=7870, Y=6450",
		}, {
			"Button A: X+69, Y+23",
			"Button B: X+27, Y+71",
			"Prize: X=18641, Y=10279",
		},
	}
}
