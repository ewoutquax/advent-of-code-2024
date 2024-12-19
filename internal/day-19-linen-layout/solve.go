package day19linenlayout

import (
	"fmt"
	"strings"

	"github.com/ewoutquax/advent-of-code-2024/pkg/register"
	"github.com/ewoutquax/advent-of-code-2024/pkg/utils"
)

const (
	Day string = "19"
)

type Pattern string
type Towel string

var cache map[Towel]int

func (towel Towel) CountCombinations(patterns []Pattern) int {
	if towel == "" {
		return 1
	}

	if cache == nil {
		cache = make(map[Towel]int)
	}

	if result, exists := cache[towel]; exists {
		return result
	}

	count := 0
	for _, pattern := range patterns {
		if strings.HasPrefix(string(towel), string(pattern)) {
			subTowel := towel[len(pattern):]
			count += subTowel.CountCombinations(patterns)
		}
	}
	cache[towel] = count

	return count
}

type Universe struct {
	Patterns []Pattern
	Towels   []Towel
}

func SumPossibleCombinations(u Universe) int {
	var sum int = 0

	for _, towel := range u.Towels {
		sum += towel.CountCombinations(u.Patterns)
	}
	return sum
}

func CountPossibleTowels(u Universe) int {
	var sum int = 0

	for _, towel := range u.Towels {
		if towel.CountCombinations(u.Patterns) > 0 {
			sum++
		}
	}

	return sum
}

func ParseInput(blocks [][]string) Universe {
	patterns := strings.Split(blocks[0][0], ", ")
	towels := blocks[1]

	var universe = Universe{
		Patterns: make([]Pattern, 0, len(patterns)),
		Towels:   make([]Towel, 0, len(towels)),
	}

	for _, pattern := range patterns {
		universe.Patterns = append(universe.Patterns, Pattern(pattern))
	}

	for _, towel := range towels {
		universe.Towels = append(universe.Towels, Towel(towel))
	}

	return universe
}

func solvePart1(inputFile string) {
	lines := utils.ReadFileAsBlocks(inputFile)
	universe := ParseInput(lines)

	fmt.Printf("Result of day-%s / part-1: %d\n", Day, CountPossibleTowels(universe))
}

func solvePart2(inputFile string) {
	blocks := utils.ReadFileAsBlocks(inputFile)
	universe := ParseInput(blocks)

	fmt.Printf("Result of day-%s / part-2: %d\n", Day, SumPossibleCombinations(universe))
}

func init() {
	register.Day(Day+"a", solvePart1)
	register.Day(Day+"b", solvePart2)
}
