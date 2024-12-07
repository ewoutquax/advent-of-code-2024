package day07bridgerepair_test

import (
	"fmt"
	"reflect"
	"testing"

	. "github.com/ewoutquax/advent-of-code-2024/internal/day-07-bridge-repair"
	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	assert := assert.New(t)

	equations := ParseInput(testInput())

	assert.Equal("day07bridgerepair.Equations", fmt.Sprintf("%s", reflect.TypeOf(equations)))
	assert.Len(equations, 9)
	assert.Equal(190, equations[0].Answer)
	assert.Len(equations[0].Numbers, 2)
	assert.Equal(10, equations[0].Numbers[0])
	assert.Equal(19, equations[0].Numbers[1])
}

func TestEquationIsValid(t *testing.T) {
	equations := ParseInput(testInput())

	testCases := map[*Equation]bool{
		&equations[0]: true,
		&equations[1]: true,
		&equations[2]: false,
		&equations[3]: false,
		&equations[4]: false,
		&equations[5]: false,
		&equations[6]: false,
		&equations[7]: false,
		&equations[8]: true,
	}

	for equationInput, expectedResult := range testCases {
		assert.Equal(t, expectedResult, equationInput.IsValid(), fmt.Sprintf(equationInput.ToS()))
	}
}

func TestSumValidEquations(t *testing.T) {
	equations := ParseInput(testInput())

	assert.Equal(t, 3749, SumValidEquations(equations))
}

func TestSumValidEquationsWithConcatenator(t *testing.T) {
	equations := ParseInput(testInput())
	MaxOperators = 3

	assert.Equal(t, 11387, SumValidEquations(equations))
}

func testInput() []string {
	return []string{
		"190: 10 19",
		"3267: 81 40 27",
		"83: 17 5",
		"156: 15 6",
		"7290: 6 8 6 15",
		"161011: 16 10 13",
		"192: 17 8 14",
		"21037: 9 7 18 13",
		"292: 11 6 16 20",
	}
}
