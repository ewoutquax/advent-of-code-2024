package day03mullitover

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/ewoutquax/advent-of-code-2024/pkg/register"
	"github.com/ewoutquax/advent-of-code-2024/pkg/utils"
)

const (
	Day string = "03"
)

type Operation struct {
	Left  int
	Right int
}

func (o Operation) Resolve() int {
	return o.Left * o.Right
}

func SumValidOperations(operations []Operation) int {
	var sum int = 0

	for _, op := range operations {
		sum += op.Resolve()
	}

	return sum
}

func ParseInput(lines []string) []Operation {
	var operations []Operation = make([]Operation, 0)

	ex := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	for _, line := range lines {
		matches := ex.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			operations = append(operations, Operation{
				Left:  convAtoi(string(match[1])),
				Right: convAtoi(string(match[2])),
			})
		}
	}

	return operations
}

func PreParseInput(lines []string) []string {
	var validLines []string = make([]string, 0)

	parts := strings.Split(strings.Join(lines, ""), "do()")
	for _, part := range parts {
		validParts := strings.Split(part, "don't()")
		validLines = append(validLines, validParts[0])
	}

	return validLines
}

func convAtoi(s string) int {
	nr, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return nr
}

func solvePart1(inputFile string) {
	lines := utils.ReadFileAsLines(inputFile)
	operations := ParseInput(lines)

	fmt.Printf("Result of day-%s / part-1: %d\n", Day, SumValidOperations(operations))
}

func solvePart2(inputFile string) {
	lines := utils.ReadFileAsLines(inputFile)
	validLines := PreParseInput(lines)
	operations := ParseInput(validLines)

	fmt.Printf("Result of day-%s / part-2: %d\n", Day, SumValidOperations(operations))
}

func init() {
	register.Day(Day+"a", solvePart1)
	register.Day(Day+"b", solvePart2)
}
