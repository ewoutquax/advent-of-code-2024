package day24crossedwires

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/ewoutquax/advent-of-code-2024/pkg/register"
	"github.com/ewoutquax/advent-of-code-2024/pkg/utils"
)

const (
	Day string = "24"
)

type (
	Name string
	Wire struct {
		Name
		Value      func() int
		isResolved bool
	}
	Wires    map[Name]*Wire
	Universe struct {
		Wires
	}
)

func (w *Wire) ResolveValue() int {
	if w.isResolved {
		value := w.Value()
		fmt.Printf("ResolveValue: wire '%s': return resolved value %d\n", w.Name, value)
		return value
	}

	value := w.Value()
	w.Value = func() int { return value }
	w.isResolved = true

	fmt.Printf("ResolveValue: wire '%s': return calculated value %d\n", w.Name, value)

	return value
}

func ConvertToDecimal(numbers []int) int64 {
	var binaryString string = ""

	for _, number := range numbers {
		binaryString = strconv.Itoa(number) + binaryString
	}

	fmt.Printf("binaryString: %v\n", binaryString)

	out, _ := (strconv.ParseInt(binaryString, 2, 64))
	return out
}

func ResolveZRegisters(u Universe) []int {
	var wireNames []string = make([]string, 0, len(u.Wires))
	for name := range u.Wires {
		if name[0] == 'z' {
			wireNames = append(wireNames, string(name))
		}
	}

	slices.Sort(wireNames)

	var out []int = make([]int, 0)

	for _, wireName := range wireNames {
		out = append(out, u.Wires[Name(wireName)].ResolveValue())
	}

	return out
}

func ParseInput(blocks [][]string) Universe {
	wires := parseWires(blocks[1])

	// Parse raw values
	for _, line := range blocks[0] {
		parts := strings.Split(line, ": ")

		wire := wires[Name(parts[0])]
		wire.Value = func() int { return utils.ConvStrToI(parts[1]) }
	}

	// Parse instructions
	for _, line := range blocks[1] {
		parts := strings.Split(line, " ")
		wireFrom1 := wires[Name(parts[0])]
		wireFrom2 := wires[Name(parts[2])]
		wireTo := wires[Name(parts[4])]

		switch parts[1] {
		case "AND":
			wireTo.Value = func() int { return wireFrom1.ResolveValue() & wireFrom2.ResolveValue() }
		case "OR":
			wireTo.Value = func() int { return wireFrom1.ResolveValue() | wireFrom2.ResolveValue() }
		case "XOR":
			wireTo.Value = func() int { return wireFrom1.ResolveValue() ^ wireFrom2.ResolveValue() }
		}
	}

	return Universe{
		Wires: wires,
	}
}

func parseWires(lines []string) Wires {
	var indexedWires = make(Wires, len(lines)*3)

	for _, line := range lines {
		parts := strings.Split(line, " ")

		indexedWires[Name(parts[0])] = &Wire{Name: Name(parts[0])}
		indexedWires[Name(parts[2])] = &Wire{Name: Name(parts[2])}
		indexedWires[Name(parts[4])] = &Wire{Name: Name(parts[4])}
	}

	return indexedWires
}

func solvePart1(inputFile string) {
	blocks := utils.ReadFileAsBlocks(inputFile)
	universe := ParseInput(blocks)
	registers := ResolveZRegisters(universe)

	fmt.Printf("Result of day-%s / part-1: %d\n", Day, ConvertToDecimal(registers))
}

func solvePart2(inputFile string) {
	_ = inputFile

	var count int = 0
	fmt.Printf("Result of day-%s / part-2: %d\n", Day, count)
}

func init() {
	register.Day(Day+"a", solvePart1)
	// register.Day(Day+"b", solvePart2)
}
