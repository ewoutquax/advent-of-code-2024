package day02rednosedreports

import (
	"fmt"
	"strings"

	types "github.com/ewoutquax/advent-of-code-2024/internal/day-02-red-nosed-reports/types"
	"github.com/ewoutquax/advent-of-code-2024/pkg/register"
	"github.com/ewoutquax/advent-of-code-2024/pkg/utils"
)

const (
	Day string = "02"
)

func SumValidReports(lines []string, dampeningLevel int) int {
	var sum int = 0

	reports := ParseInput(lines)
	for _, report := range reports {
		if report.IsSafe(dampeningLevel) {
			sum += 1
		}
	}

	return sum
}

func ParseInput(lines []string) []types.Report {
	var reports []types.Report = make([]types.Report, 0, len(lines))

	for _, line := range lines {
		numbers := strings.Split(line, " ")
		report := types.Report{
			Levels: make([]types.Level, 0, len(numbers)),
		}

		for _, nr := range numbers {
			report.Levels = append(report.Levels, types.ConvAtoLevel(nr))
		}

		reports = append(reports, report)
	}

	return reports
}

func solvePart1(inputFile string) {
	lines := utils.ReadFileAsLines(inputFile)

	fmt.Printf("Result of day-%s / part-1: %d\n", Day, SumValidReports(lines, 0))
}

func solvePart2(inputFile string) {
	lines := utils.ReadFileAsLines(inputFile)

	fmt.Printf("Result of day-%s / part-2: %d\n", Day, SumValidReports(lines, 1))
}

func init() {
	register.Day(Day+"a", solvePart1)
	register.Day(Day+"b", solvePart2)
}
