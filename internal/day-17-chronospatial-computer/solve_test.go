package day17chronospatialcomputer_test

import (
	"testing"

	. "github.com/ewoutquax/advent-of-code-2024/internal/day-17-chronospatial-computer"
	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	program := ParseInput(testInput())

	assert := assert.New(t)
	assert.IsType(Program{}, program)
	assert.Equal(0, program.IdxInstruction)

	assert.Len(program.Registers, 3)
	assert.Equal(729, program.Registers["A"])
	assert.Equal(0, program.Registers["B"])
	assert.Equal(0, program.Registers["C"])

	assert.Len(program.Instructions, 6)
	assert.Equal(Instruction("0"), program.Instructions[0])
	assert.Equal(Instruction("0"), program.Instructions[len(program.Instructions)-1])
}

func TestRunProgram(t *testing.T) {
	program := ParseInput(testInput())

	output := RunProgram(program)

	assert.Equal(t, "4,6,3,5,6,3,5,2,1,0", output)
}

func TestRunProgramWithUpdatedRegisterA(t *testing.T) {
	program := ParseInput(testInputClone())

	program.SetRegisterA(117440)
	output := RunProgram(program)

	assert.Equal(t, program.RawInstructions, output)
}

func TestFindCloneValue(t *testing.T) {
	program := ParseInput(testInputClone())

	cloneValue := FindCloneValue(program)

	assert.Equal(t, 117440, cloneValue)
}

func testInput() [][]string {
	return [][]string{
		{
			"Register A: 729",
			"Register B: 0",
			"Register C: 0",
		}, {
			"Program: 0,1,5,4,3,0",
		},
	}
}

func testInputClone() [][]string {
	return [][]string{
		{
			"Register A: 2024",
			"Register B: 0",
			"Register C: 0",
		}, {
			"Program: 0,3,5,4,3,0",
		},
	}
}
