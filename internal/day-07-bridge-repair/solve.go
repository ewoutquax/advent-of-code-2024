package day07bridgerepair

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/ewoutquax/advent-of-code-2024/pkg/register"
	"github.com/ewoutquax/advent-of-code-2024/pkg/utils"
)

type Operator uint

const (
	OperatorAdd Operator = iota + 1
	OperatorMultiply
	OperatorConcat

	Day         string = "07"
	MAX_THREADS int    = 8
)

var (
	wg                           = sync.WaitGroup{}
	lockSummedValidEquations     = sync.Mutex{}
	summedValidEquations     int = 0
	MaxOperators             int = 2
)

type Equation struct {
	Answer  int
	Numbers []int
}

func (e Equation) ToS() string {
	var nrs []string = make([]string, 0, len(e.Numbers))
	for _, nr := range e.Numbers {
		nrs = append(nrs, strconv.Itoa(nr))
	}
	return fmt.Sprintf("%d: %s", e.Answer, strings.Join(nrs, " "))
}

func (e Equation) IsValid() bool {
	resolvements := resolveNumbers(e.Numbers, e.Answer)

	for _, resolvement := range resolvements {
		if resolvement == e.Answer {
			return true
		}
	}

	return false
}

func resolveNumbers(nrs []int, answer int) []int {
	if len(nrs) == 1 {
		return nrs
	}

	lhss := resolveNumbers(nrs[0:len(nrs)-1], answer)
	rhs := nrs[len(nrs)-1]

	var resolvements []int = make([]int, 0, len(lhss)*len(allOperators()))

	for _, lhs := range lhss {
		for _, operator := range allOperators() {
			var resolvement int
			switch operator {
			case OperatorAdd:
				resolvement = lhs + rhs
			case OperatorMultiply:
				resolvement = lhs * rhs
			case OperatorConcat:
				tmps := []string{
					strconv.Itoa(lhs),
					strconv.Itoa(rhs),
				}
				tmp := strings.Join(tmps, "")
				resolvement = convAtoi(tmp)
			default:
				panic("No valid operator found")
			}
			if resolvement <= answer {
				resolvements = append(resolvements, resolvement)
			}
		}
	}

	return resolvements
}

type Equations []Equation

func SumValidEquations(equations Equations) int {
	summedValidEquations = 0

	channel := make(chan Equation)
	for idx := MAX_THREADS; idx > 0; idx-- {
		go resolveEquation(channel)
	}
	wg.Add(MAX_THREADS)

	for _, equation := range equations {
		channel <- equation
	}
	close(channel)

	wg.Wait()

	return summedValidEquations
}

func resolveEquation(channel chan Equation) {
	defer wg.Done()

	for equation := range channel {
		if equation.IsValid() {
			lockSummedValidEquations.Lock()
			summedValidEquations += equation.Answer
			lockSummedValidEquations.Unlock()
		}
	}
}

func ParseInput(lines []string) Equations {
	var equations Equations = make(Equations, 0, len(lines))

	for _, line := range lines {
		parts := strings.Split(line, ": ")
		partNrs := strings.Split(parts[1], " ")
		equation := Equation{
			Answer:  convAtoi(parts[0]),
			Numbers: make([]int, 0, len(partNrs)),
		}

		for _, partNr := range partNrs {
			equation.Numbers = append(equation.Numbers, convAtoi(partNr))
		}

		equations = append(equations, equation)
	}

	return equations
}

func convAtoi(s string) int {
	nr, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return nr
}

func allOperators() []Operator {
	return []Operator{
		OperatorAdd,
		OperatorMultiply,
		OperatorConcat,
	}[:MaxOperators]
}

func solvePart1(inputFile string) {
	lines := utils.ReadFileAsLines(inputFile)
	equations := ParseInput(lines)

	fmt.Printf("Result of day-%s / part-1: %d\n", Day, SumValidEquations(equations))
}

func solvePart2(inputFile string) {
	lines := utils.ReadFileAsLines(inputFile)
	equations := ParseInput(lines)
	MaxOperators = 3

	fmt.Printf("Result of day-%s / part-2: %d\n", Day, SumValidEquations(equations))
}

func init() {
	register.Day(Day+"a", solvePart1)
	register.Day(Day+"b", solvePart2)
}
