package day21keypadconundrum

import (
	"fmt"

	calculator "github.com/ewoutquax/advent-of-code-2024/internal/day-21-keypad-conundrum/services/calculator"
	"github.com/ewoutquax/advent-of-code-2024/pkg/register"
	"github.com/ewoutquax/advent-of-code-2024/pkg/utils"
)

const (
	Day string = "21"
)

func solvePart1(inputFile string) {
	codes := utils.ReadFileAsLines(inputFile)

	fmt.Printf("Result of day-%s / part-1: %d\n", Day, calculator.SumComplexities(codes, 2))
}

func solvePart2(inputFile string) {
	codes := utils.ReadFileAsLines(inputFile)

	fmt.Printf("Result of day-%s / part-2: %d\n", Day, calculator.SumComplexities(codes, 25))
}

func init() {
	register.Day(Day+"a", solvePart1)
	register.Day(Day+"b", solvePart2)
}
