package day03mullitover_test

import (
	"fmt"
	"testing"

	. "github.com/ewoutquax/advent-of-code-2024/internal/day-03-mull-it-over"
	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	assert := assert.New(t)

	operations := ParseInput(testInput())
	fmt.Printf("operations: %v\n", operations)

	assert.IsType([]Operation{}, operations)
	assert.Len(operations, 4)
}

func TestParseInputActive(t *testing.T) {
	assert := assert.New(t)

	operations := ParseInput(PreParseInput(testInput2()))
	fmt.Printf("operations: %v\n", operations)

	assert.IsType([]Operation{}, operations)
	assert.Len(operations, 2)
}

func TestSumValidOperations(t *testing.T) {
	operations := ParseInput(testInput())

	assert.Equal(t, 161, SumValidOperations(operations))
}

func TestSumValidOperationsActive(t *testing.T) {
	operations := ParseInput(PreParseInput(testInput2()))

	assert.Equal(t, 48, SumValidOperations(operations))
}

func testInput() []string {
	return []string{"xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"}
}

func testInput2() []string {
	return []string{
		"xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))",
	}
}
