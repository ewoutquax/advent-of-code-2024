package day13clawcontraption

import (
	"errors"
	"fmt"
	. "image"
	"strings"

	"github.com/ewoutquax/advent-of-code-2024/pkg/register"
	"github.com/ewoutquax/advent-of-code-2024/pkg/utils"
)

const (
	Day string = "13"
)

type Vector Point
type Arcades []Arcade

type Arcade struct {
	ButtonA Vector
	ButtonB Vector
	Prize   Point
}

func SumCheapestSolutions(arcades []Arcade) int {
	var sum = 0

	for _, arcade := range arcades {
		minTokens, err := FindCheapestSolution(arcade)
		if err == nil {
			sum += minTokens
		}
	}

	return sum
}

func FindCheapestSolution(arcade Arcade) (int, error) {
	noemerRepeaterB := (arcade.ButtonA.X*arcade.Prize.Y - arcade.ButtonA.Y*arcade.Prize.X)
	delerRepeaterB := (arcade.ButtonA.X*arcade.ButtonB.Y - arcade.ButtonA.Y*arcade.ButtonB.X)

	if noemerRepeaterB%delerRepeaterB == 0 {
		repeaterB := noemerRepeaterB / delerRepeaterB

		restX := (arcade.Prize.X - arcade.ButtonB.X*repeaterB)
		if restX%arcade.ButtonA.X == 0 {
			repeaterA := restX / arcade.ButtonA.X
			return repeaterA*3 + repeaterB*1, nil
		}
	}

	return 0, errors.New("No solution found")
}

func ParseInput(blocks [][]string, offset int) []Arcade {
	var arcades = make([]Arcade, 0, len(blocks))

	for _, block := range blocks {
		arcades = append(arcades, parseBlock(block, offset))
	}

	return arcades
}

func parseBlock(lines []string, offset int) Arcade {
	parts := strings.Split(lines[2], ": ")
	pointParts := strings.Split(parts[1], ", ")
	pointXParts := strings.Split(pointParts[0], "=")
	pointYParts := strings.Split(pointParts[1], "=")

	return Arcade{
		ButtonA: parseButton(lines[0]),
		ButtonB: parseButton(lines[1]),
		Prize:   Pt(utils.ConvStrToI(pointXParts[1])+offset, utils.ConvStrToI(pointYParts[1])+offset),
	}
}

func parseButton(line string) Vector {
	parts := strings.Split(line, ": ")
	pointParts := strings.Split(parts[1], ", ")
	pointXParts := strings.Split(pointParts[0], "+")
	pointYParts := strings.Split(pointParts[1], "+")

	return Vector(Pt(utils.ConvStrToI(pointXParts[1]), utils.ConvStrToI(pointYParts[1])))
}

func max(nrs ...int) int {
	var lhs, rhs int

	if len(nrs) > 2 {
		rhs = max(nrs[1:]...)
	} else {
		rhs = nrs[1]
	}
	lhs = nrs[0]

	if lhs > rhs {
		return lhs
	}

	return rhs
}

func solvePart1(inputFile string) {
	blocks := utils.ReadFileAsBlocks(inputFile)
	arcades := ParseInput(blocks, 0)

	fmt.Printf("Result of day-%s / part-1: %d\n", Day, SumCheapestSolutions(arcades))
}

func solvePart2(inputFile string) {
	blocks := utils.ReadFileAsBlocks(inputFile)
	arcades := ParseInput(blocks, 10000000000000)

	fmt.Printf("Result of day-%s / part-2: %d\n", Day, SumCheapestSolutions(arcades))
}

func init() {
	register.Day(Day+"b", solvePart2)
	register.Day(Day+"a", solvePart1)
}
