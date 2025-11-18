package day24crossedwires_test

import (
	"fmt"
	"testing"

	. "github.com/ewoutquax/advent-of-code-2024/internal/day-24-crossed-wires"
	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {

	universe := ParseInput(testInput())

	assert := assert.New(t)
	assert.Len(universe.Wires, 9)
	assert.Contains(universe.Wires, Name("x00"))
	assert.Contains(universe.Wires, Name("y01"))
	assert.Contains(universe.Wires, Name("z02"))

	assert.Equal(1, universe.Wires[Name("x00")].Value())
	assert.Equal(1, universe.Wires[Name("x02")].Value())
	assert.Equal(0, universe.Wires[Name("y02")].Value())

	assert.IsType(Universe{}, universe)
}

func TestResolveZRegisters(t *testing.T) {
	universe := ParseInput(testInput())

	registers := ResolveZRegisters(universe)

	assert := assert.New(t)
	assert.Equal("[]int", fmt.Sprintf("%T", registers))
	assert.Len(registers, 3)
	assert.Equal(0, registers[0])
	assert.Equal(0, registers[1])
	assert.Equal(1, registers[2])
}

func TestConvertToDecimal(t *testing.T) {
	input := []int{0, 0, 1}

	assert.Equal(t, 4, ConvertToDecimal(input))
}

func TestFindSolutionLargeInput(t *testing.T) {
	universe := ParseInput(testInputLarge())

	registers := ResolveZRegisters(universe)

	assert.Equal(t, 2024, ConvertToDecimal(registers))
	assert.True(t, false)
}

func testInput() [][]string {
	return [][]string{
		{
			"x00: 1",
			"x01: 1",
			"x02: 1",
			"y00: 0",
			"y01: 1",
			"y02: 0",
		}, {
			"x00 AND y00 -> z00",
			"x01 XOR y01 -> z01",
			"x02 OR y02 -> z02",
		},
	}
}

func testInputLarge() [][]string {
	return [][]string{
		{
			"x00: 1",
			"x01: 0",
			"x02: 1",
			"x03: 1",
			"x04: 0",
			"y00: 1",
			"y01: 1",
			"y02: 1",
			"y03: 1",
			"y04: 1",
		}, {
			"ntg XOR fgs -> mjb",
			"y02 OR x01 -> tnw",
			"kwq OR kpj -> z05",
			"x00 OR x03 -> fst",
			"tgd XOR rvg -> z01",
			"vdt OR tnw -> bfw",
			"bfw AND frj -> z10",
			"ffh OR nrd -> bqk",
			"y00 AND y03 -> djm",
			"y03 OR y00 -> psh",
			"bqk OR frj -> z08",
			"tnw OR fst -> frj",
			"gnj AND tgd -> z11",
			"bfw XOR mjb -> z00",
			"x03 OR x00 -> vdt",
			"gnj AND wpb -> z02",
			"x04 AND y00 -> kjc",
			"djm OR pbm -> qhw",
			"nrd AND vdt -> hwm",
			"kjc AND fst -> rvg",
			"y04 OR y02 -> fgs",
			"y01 AND x02 -> pbm",
			"ntg OR kjc -> kwq",
			"psh XOR fgs -> tgd",
			"qhw XOR tgd -> z09",
			"pbm OR djm -> kpj",
			"x03 XOR y03 -> ffh",
			"x00 XOR y04 -> ntg",
			"bfw OR bqk -> z06",
			"nrd XOR fgs -> wpb",
			"frj XOR qhw -> z04",
			"bqk OR frj -> z07",
			"y03 OR x01 -> nrd",
			"hwm AND bqk -> z03",
			"tgd XOR rvg -> z12",
			"tnw OR pbm -> gnj",
		},
	}
}
