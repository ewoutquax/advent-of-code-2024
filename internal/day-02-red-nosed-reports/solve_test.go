package day02rednosedreports_test

import (
	"fmt"
	"testing"

	. "github.com/ewoutquax/advent-of-code-2024/internal/day-02-red-nosed-reports"
	types "github.com/ewoutquax/advent-of-code-2024/internal/day-02-red-nosed-reports/types"
	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	assert := assert.New(t)

	reports := ParseInput(testInput())
	assert.IsType([]types.Report{}, reports)
	assert.Len(reports, 6)
	assert.Equal(types.Level(7), reports[0].Levels[0])
	assert.Equal(types.Level(1), reports[0].Levels[len(reports[0].Levels)-1])
	assert.Equal(types.Level(1), reports[len(reports)-1].Levels[0])
}

func TestReportIsSafe(t *testing.T) {
	reports := ParseInput(testInput())

	testCases := map[*types.Report]bool{
		&reports[0]: true,
		&reports[1]: false,
		&reports[2]: false,
		&reports[3]: false,
		&reports[4]: false,
		&reports[5]: true,
	}

	for inputReport, expectedResult := range testCases {
		assert.Equal(t, expectedResult, inputReport.IsSafe(0), fmt.Sprintf("report: %v\n", inputReport))
	}
}

func TestSumValidReports(t *testing.T) {
	assert.Equal(t, 2, SumValidReports(testInput(), 0))
}

func TestReportIsSafeWithDampening(t *testing.T) {
	reports := ParseInput(testInput())

	testCases := map[*types.Report]bool{
		&reports[0]: true,
		&reports[1]: false,
		&reports[2]: false,
		&reports[3]: true,
		&reports[4]: true,
		&reports[5]: true,
	}

	for inputReport, expectedResult := range testCases {
		assert.Equal(t, expectedResult, inputReport.IsSafe(1), fmt.Sprintf("report: %v\n", inputReport))
	}
}

func TestSumValidReportsWithDampening(t *testing.T) {
	assert.Equal(t, 4, SumValidReports(testInput(), 1))
}

func testInput() []string {
	return []string{
		"7 6 4 2 1",
		"1 2 7 8 9",
		"9 7 6 2 1",
		"1 3 2 4 5",
		"8 6 4 4 1",
		"1 3 6 7 9",
	}
}
