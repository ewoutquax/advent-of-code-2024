package day17chronospatialcomputer

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/ewoutquax/advent-of-code-2024/pkg/register"
	"github.com/ewoutquax/advent-of-code-2024/pkg/utils"
)

const (
	Day string = "17"
)

type (
	Instruction string
	Program     struct {
		IdxInstruction  int
		Instructions    []Instruction
		Registers       map[string]int
		Output          []string
		RawInstructions string
	}
)

func FindCloneValue(program Program) int {
	var result int = 0
	for currentIndex := len(program.RawInstructions) - 1; currentIndex >= 0; currentIndex -= 2 {
		doContinue := true
		var outcome string = ""

		offset := 0
		result *= 8
		for doContinue {
			program.SetRegisterA(result + offset)
			outcome = RunProgram(program)

			if program.RawInstructions[currentIndex:] == outcome {
				doContinue = false
			} else {
				offset++
			}
		}
		result += offset
	}

	return result
}

func RunProgram(program Program) string {
	program.Output = make([]string, 0)
	program.IdxInstruction = 0

	for program.IdxInstruction < len(program.Instructions) {
		value := program.Instructions[program.IdxInstruction+1]

		switch program.Instructions[program.IdxInstruction] {
		case "0":
			program.runInstructionAdv(value)
		case "1":
			program.runInstructionBxl(value)
		case "2":
			program.runInstructionBst(value)
		case "3":
			program.runInstructionJnz(value)
		case "4":
			program.runInstructionBxc(value)
		case "5":
			program.runInstructionOut(value)
		case "6":
			program.runInstructionBdv(value)
		case "7":
			program.runInstructionCdv(value)
		default:
			panic("No valid case found")
		}
	}

	return strings.Join(program.Output, ",")
}

func (p *Program) runInstructionAdv(value Instruction) {
	var divider int = int(math.Pow(2, float64(comboValue(value, p.Registers))))

	p.Registers["A"] = p.Registers["A"] / divider
	p.IdxInstruction += 2
}

func (p *Program) runInstructionBxl(value Instruction) {
	p.Registers["B"] = p.Registers["B"] ^ convInstructiontoI(value)
	p.IdxInstruction += 2
}

func (p *Program) runInstructionBst(value Instruction) {
	p.Registers["B"] = comboValue(value, p.Registers) % 8
	p.IdxInstruction += 2
}

func (p *Program) runInstructionJnz(value Instruction) {
	if p.Registers["A"] == 0 {
		p.IdxInstruction += 2
	} else {
		p.IdxInstruction = convInstructiontoI(value)
	}
}

func (p *Program) runInstructionBxc(Instruction) {
	p.Registers["B"] = p.Registers["B"] ^ p.Registers["C"]
	p.IdxInstruction += 2
}

func (p *Program) runInstructionOut(value Instruction) {
	newOut := comboValue(value, p.Registers) % 8
	p.Output = append(p.Output, fmt.Sprintf("%d", newOut))

	p.IdxInstruction += 2
}

func (p *Program) runInstructionBdv(value Instruction) {
	var divider int = int(math.Pow(2, float64(comboValue(value, p.Registers))))

	p.Registers["B"] = p.Registers["A"] / divider
	p.IdxInstruction += 2
}

func (p *Program) runInstructionCdv(value Instruction) {
	var divider int = int(math.Pow(2, float64(comboValue(value, p.Registers))))

	p.Registers["C"] = p.Registers["A"] / divider
	p.IdxInstruction += 2
}

func comboValue(value Instruction, registers map[string]int) int {
	return map[Instruction]int{
		"0": 0,
		"1": 1,
		"2": 2,
		"3": 3,
		"4": registers["A"],
		"5": registers["B"],
		"6": registers["C"],
	}[value]
}

func (p *Program) SetRegisterA(newValue int) {
	p.Registers["A"] = newValue
}

func ParseInput(blocks [][]string) Program {
	rawInstructionsParts := strings.Split(blocks[1][0], ": ")
	return Program{
		IdxInstruction:  0,
		Instructions:    parseInstructions(blocks[1]),
		Registers:       parseRegisters(blocks[0]),
		RawInstructions: rawInstructionsParts[1],
	}
}

func parseRegisters(lines []string) map[string]int {
	var registers = make(map[string]int, 3)
	var parts []string

	parts = strings.Split(lines[0], ": ")
	registers["A"] = convAtoI(parts[1])
	parts = strings.Split(lines[1], ": ")
	registers["B"] = convAtoI(parts[1])
	parts = strings.Split(lines[2], ": ")
	registers["C"] = convAtoI(parts[1])

	return registers
}

func parseInstructions(lines []string) []Instruction {
	programParts := strings.Split(lines[0], ": ")
	chars := strings.Split(programParts[1], ",")

	var instructions = make([]Instruction, 0, len(chars))
	for _, char := range chars {
		instructions = append(instructions, Instruction(char))
	}

	return instructions
}

func convInstructiontoI(i Instruction) int {
	return convAtoI(string(i))
}

func convAtoI(s string) int {
	nr, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return nr
}

func solvePart1(inputFile string) {
	blocks := utils.ReadFileAsBlocks(inputFile)
	program := ParseInput(blocks)
	output := RunProgram(program)

	fmt.Printf("Result of day-%s / part-1: %s\n", Day, output)
}

func solvePart2(inputFile string) {
	blocks := utils.ReadFileAsBlocks(inputFile)
	program := ParseInput(blocks)
	value := FindCloneValue(program)

	fmt.Printf("Result of day-%s / part-2: %d\n", Day, value)
}

func init() {
	register.Day(Day+"a", solvePart1)
	register.Day(Day+"b", solvePart2)
}
