package day04ceressearch

import (
	"fmt"
	"strings"

	"github.com/ewoutquax/advent-of-code-2024/pkg/register"
	"github.com/ewoutquax/advent-of-code-2024/pkg/utils"
)

const (
	Day string = "04"
)

type Location struct {
	X int
	Y int
}

type Puzzle struct {
	Letters map[Location]string
	MaxX    int
	MaxY    int
}

type Vector [2]int

type SecondVector struct {
	OffsetLocation [2]int
	Vector
}

type VectorWithSecondVector struct {
	Vector
	SecondVector
}

func CountOccurances(puzzle Puzzle, searchWord string) int {
	var count int = 0

	for x := 0; x < puzzle.MaxX; x++ {
		for y := 0; y < puzzle.MaxY; y++ {
			for _, vector := range directions() {
				matchFound := true
				for offset := 0; matchFound && offset < len(searchWord); offset++ {
					searchLoc := Location{
						X: x + offset*vector[0],
						Y: y + offset*vector[1],
					}
					if char, exists := puzzle.Letters[searchLoc]; !(exists && char == string(searchWord[offset])) {
						matchFound = false
					}
				}
				if matchFound {
					count++
				}
			}
		}
	}

	return count
}

func CountXmasOccurances(puzzle Puzzle, searchWord string) int {
	var count int = 0
	var matchFound bool

	for x := 0; x < puzzle.MaxX; x++ {
		for y := 0; y < puzzle.MaxY; y++ {
			for _, vectorWithSecondVector := range xMasDirections() {
				matchFound = true
				for offset := 0; matchFound && offset < len(searchWord); offset++ {
					searchLoc := Location{
						X: x + offset*vectorWithSecondVector.Vector[0],
						Y: y + offset*vectorWithSecondVector.Vector[1],
					}
					secondSearchLoc := Location{
						X: x + (len(searchWord)-1)*vectorWithSecondVector.OffsetLocation[0] + offset*vectorWithSecondVector.SecondVector.Vector[0],
						Y: y + (len(searchWord)-1)*vectorWithSecondVector.OffsetLocation[1] + offset*vectorWithSecondVector.SecondVector.Vector[1],
					}
					firstChar, firstExists := puzzle.Letters[searchLoc]
					secondChar, secondExists := puzzle.Letters[secondSearchLoc]
					if !(firstExists && secondExists && firstChar == secondChar && firstChar == string(searchWord[offset])) {
						matchFound = false
					}
				}
				if matchFound {
					count++
				}
			}
		}
	}

	return count
}

func ParseInput(lines []string) Puzzle {
	var puzzle Puzzle = Puzzle{
		Letters: make(map[Location]string, len(lines)*len(lines[0])),
		MaxX:    len(lines[0]),
		MaxY:    len(lines),
	}

	for y, line := range lines {
		for x, char := range strings.Split(line, "") {
			loc := Location{X: x, Y: y}
			puzzle.Letters[loc] = char
		}
	}

	return puzzle
}

func directions() [8]Vector {
	return [8]Vector{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}
}

func xMasDirections() [4]VectorWithSecondVector {
	return [4]VectorWithSecondVector{
		{Vector: [2]int{1, 1}, SecondVector: SecondVector{OffsetLocation: [2]int{1, 0}, Vector: [2]int{-1, 1}}},
		{Vector: [2]int{-1, 1}, SecondVector: SecondVector{OffsetLocation: [2]int{0, 1}, Vector: [2]int{-1, -1}}},
		{Vector: [2]int{-1, -1}, SecondVector: SecondVector{OffsetLocation: [2]int{-1, 0}, Vector: [2]int{1, -1}}},
		{Vector: [2]int{1, -1}, SecondVector: SecondVector{OffsetLocation: [2]int{0, -1}, Vector: [2]int{1, 1}}},
	}
}

func solvePart1(inputFile string) {
	lines := utils.ReadFileAsLines(inputFile)
	puzzle := ParseInput(lines)

	fmt.Printf("Result of day-%s / part-1: %d\n", Day, CountOccurances(puzzle, "XMAS"))
}

func solvePart2(inputFile string) {
	lines := utils.ReadFileAsLines(inputFile)
	puzzle := ParseInput(lines)

	fmt.Printf("Result of day-%s / part-2: %d\n", Day, CountXmasOccurances(puzzle, "MAS"))
}

func init() {
	register.Day(Day+"a", solvePart1)
	register.Day(Day+"b", solvePart2)
}
