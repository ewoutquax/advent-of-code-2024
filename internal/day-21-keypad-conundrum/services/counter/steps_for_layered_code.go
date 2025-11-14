package day21keypadconundrum_services_counter

import (
	"math"

	common "github.com/ewoutquax/advent-of-code-2024/internal/day-21-keypad-conundrum/common"
	builder "github.com/ewoutquax/advent-of-code-2024/internal/day-21-keypad-conundrum/services/builder"
	finder "github.com/ewoutquax/advent-of-code-2024/internal/day-21-keypad-conundrum/services/finder"
)

func StepsForLayeredCode(numericalCode string, nrDirectionalKeypads int) int {
	keypadNumeric := builder.BuildKeypad(builder.WithLayout(common.KeypadLayoutNumeric))
	keypadDirectional := builder.BuildKeypad(builder.WithLayout(common.KeypadLayoutDirectional))
	var minNrSteps int = math.MaxInt

	directionalCodes := finder.FindShortestPathsForCode(keypadNumeric, numericalCode)
	for _, directionalCode := range directionalCodes {
		minNrSteps = min(
			minNrSteps,
			StepsForDirectionalCode(directionalCode, nrDirectionalKeypads, keypadDirectional),
		)
	}

	return minNrSteps
}
