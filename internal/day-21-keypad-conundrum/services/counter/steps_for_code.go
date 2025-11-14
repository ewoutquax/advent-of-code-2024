package day21keypadconundrum_services_counter

import (
	"math"
	"strings"

	common "github.com/ewoutquax/advent-of-code-2024/internal/day-21-keypad-conundrum/common"
	finder "github.com/ewoutquax/advent-of-code-2024/internal/day-21-keypad-conundrum/services/finder"
)

type (
	CacheKey struct {
		segment string
		nrLayer int
	}
)

var (
	Cache map[CacheKey]int
)

func StepsForDirectionalCode(code string, nrLayer int, keypad common.Keypad) int {
	var nrSteps int = 0
	for _, segment := range CodeToSegments(code) {
		nrSteps += leastStepsForSegment(segment, nrLayer, keypad)
	}

	return nrSteps
}

func leastStepsForSegment(segment string, nrLayer int, keypad common.Keypad) int {
	if Cache == nil {
		Cache = make(map[CacheKey]int)
	}

	cacheKey := CacheKey{
		segment: segment,
		nrLayer: nrLayer,
	}

	if cachedNrSteps, ok := Cache[cacheKey]; ok {
		return cachedNrSteps
	}

	shortestPaths := finder.FindShortestPathsForCode(keypad, segment)

	var nrSteps int = math.MaxInt
	if nrLayer == 1 {
		nrSteps = len(shortestPaths[0])
	} else {
		for _, shortestPath := range shortestPaths {
			var sum int = 0
			for _, subSegment := range CodeToSegments(shortestPath) {
				sum += leastStepsForSegment(subSegment, nrLayer-1, keypad)
			}
			nrSteps = min(nrSteps, sum)
		}
	}

	Cache[cacheKey] = nrSteps
	return nrSteps
}

func CodeToSegments(path string) []string {
	parts := strings.Split(path, "A")

	var segments []string = make([]string, 0, len(parts))
	for idx := 0; idx < len(parts)-1; idx++ {
		segments = append(segments, parts[idx]+"A")
	}

	return segments
}
