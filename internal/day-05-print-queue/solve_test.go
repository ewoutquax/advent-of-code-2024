package day05printqueue_test

import (
	"fmt"
	"testing"

	. "github.com/ewoutquax/advent-of-code-2024/internal/day-05-print-queue"
	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	assert := assert.New(t)

	manual := ParseInput(testInput())

	assert.IsType(Manual{}, manual)
	assert.Len(manual.Orders, 21)
	assert.Len(manual.Updates, 6)
	assert.Len(manual.Updates[0], 5)
}

func TestUpdateIsValid(t *testing.T) {
	manual := ParseInput(testInput())

	testCases := map[*Update]bool{
		&manual.Updates[0]: true,
		&manual.Updates[1]: true,
		&manual.Updates[2]: true,
		&manual.Updates[3]: false,
		&manual.Updates[4]: false,
		&manual.Updates[5]: false,
	}
	for update, expectedResult := range testCases {
		assert.Equal(t, expectedResult, update.IsValid(manual.Orders), fmt.Sprintf("Update: %v", update))
	}
}

func TestSumMiddlePages(t *testing.T) {
	manual := ParseInput(testInput())

	validUpdates := []Update{
		manual.Updates[0],
		manual.Updates[1],
		manual.Updates[2],
	}

	assert.Equal(t, 143, SumMiddlePages(validUpdates))
}

func TestSumMiddlePagesOfValidReports(t *testing.T) {
	assert.Equal(t, 143, SumMiddlePagesOfValidReports(testInput()))
}

func TestOrderUpdates(t *testing.T) {
	manual := ParseInput(testInput())

	update := Update{
		75, 97, 47, 61, 53,
	}

	orderedUpdate := OrderUpdate(update, manual.Orders)

	assert.Len(t, orderedUpdate, 5)
	assert.Equal(t, PageNr(97), orderedUpdate[0])
	assert.Equal(t, PageNr(75), orderedUpdate[1])
	assert.Equal(t, PageNr(47), orderedUpdate[2])
	assert.Equal(t, PageNr(61), orderedUpdate[3])
	assert.Equal(t, PageNr(53), orderedUpdate[4])
}

func TestSumMiddlePagesOfCorrectedInvalidReports(t *testing.T) {
	assert.Equal(t, 123, SumMiddlePagesOfCorrectedInvalidReports(testInput()))
}

func testInput() [][]string {
	return [][]string{
		{
			"47|53",
			"97|13",
			"97|61",
			"97|47",
			"75|29",
			"61|13",
			"75|53",
			"29|13",
			"97|29",
			"53|29",
			"61|53",
			"97|53",
			"61|29",
			"47|13",
			"75|47",
			"97|75",
			"47|61",
			"75|61",
			"47|29",
			"75|13",
			"53|13",
		}, {
			"75,47,61,53,29",
			"97,61,53,29,13",
			"75,29,13",
			"75,97,47,61,53",
			"61,13,29",
			"97,13,75,29,47",
		},
	}
}
