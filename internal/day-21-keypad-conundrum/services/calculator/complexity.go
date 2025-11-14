package day21keypadconundrum_services_calculator

import (
	"strconv"
	"strings"

	counter "github.com/ewoutquax/advent-of-code-2024/internal/day-21-keypad-conundrum/services/counter"
)

func SumComplexities(codes []string, nrDirectionalKeypads int) (sum int) {
	for _, code := range codes {
		sum += CalculateComplexity(
			code,
			counter.StepsForLayeredCode(code, nrDirectionalKeypads),
		)
	}

	return
}

func CalculateComplexity(code string, lengthSolution int) int {
	var numerics string = ""
	for _, char := range strings.Split(code, "") {
		if strings.Contains("0123456789", char) {
			numerics += char
		}
	}

	number, _ := strconv.Atoi(numerics)

	return number * lengthSolution
}
