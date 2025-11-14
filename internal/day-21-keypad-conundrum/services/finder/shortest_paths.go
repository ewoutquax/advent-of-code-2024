package day21keypadconundrum_services_finder

import (
	"container/heap"
	"image"
	"math"

	common "github.com/ewoutquax/advent-of-code-2024/internal/day-21-keypad-conundrum/common"
)

type (
	PathHeap []common.Path
)

func (h PathHeap) Len() int           { return len(h) }
func (h PathHeap) Less(i, j int) bool { return h[i].NrSteps < h[j].NrSteps }
func (h PathHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *PathHeap) Push(x any)        { *h = append(*h, x.(common.Path)) }

func (h *PathHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func FindShortestPathsForCode(keypad common.Keypad, code string) []string {
	allCodes := findAllPathsForCode(keypad, code)

	minLength := math.MaxInt
	for _, code := range allCodes {
		minLength = min(minLength, len(code))
	}

	var shortestCodes []string
	for _, foundCode := range allCodes {
		if minLength == len(foundCode) {
			shortestCodes = append(shortestCodes, string(foundCode))
		}
	}

	return shortestCodes
}

func findAllPathsForCode(keypad common.Keypad, code string) []common.Move {
	var parts [][]common.Move = make([][]common.Move, 0, len(code))

	fullCode := "A" + code
	for idx := 0; idx < len(fullCode)-1; idx++ {
		from, to := string(fullCode[idx]), string(fullCode[idx+1])

		moves := FindPaths(keypad, from, to)
		parts = append(parts, moves)
	}

	return combineParts(parts)
}

func FindPaths(keypad common.Keypad, from, to string) []common.Move {
	if from == to {
		return []common.Move{"A"}
	}

	var visitedLocations = make(map[common.Location]int)
	var pathHeap PathHeap
	var foundPaths = make([]common.Move, 0)

	var fromKey, toKey common.Location
	for loc, key := range keypad.Keys {
		if key == from {
			fromKey = loc
		}
		if key == to {
			toKey = loc
		}
	}

	pathHeap = PathHeap{
		{
			CurrentLocation: fromKey,
			NrSteps:         0,
		},
	}

	heap.Init(&pathHeap)

	for len(pathHeap) > 0 {
		currentPath := heap.Pop(&pathHeap).(common.Path)
		for nextMove, nextLocation := range findValidNextLocations(currentPath.CurrentLocation, keypad) {
			if prevNrSteps, ok := visitedLocations[nextLocation]; !ok || ok && currentPath.NrSteps <= prevNrSteps {
				nextPath := common.Path{
					CurrentLocation: nextLocation,
					NrSteps:         currentPath.NrSteps + 1,
					Moves:           currentPath.Moves + nextMove,
				}

				visitedLocations[nextLocation] = nextPath.NrSteps

				if nextLocation == toKey {
					foundPaths = append(foundPaths, nextPath.Moves+"A")
				} else {
					if len(foundPaths) == 0 {
						heap.Push(&pathHeap, nextPath)
					}
				}
			}
		}
	}

	return foundPaths
}

func findValidNextLocations(currentLocation common.Location, keypad common.Keypad) common.Moves {
	var validNextLocations = make(common.Moves, 4)

	for move, vector := range vectors() {
		nextLocation := currentLocation.Add(vector)
		if _, ok := keypad.Keys[nextLocation]; ok {
			validNextLocations[move] = nextLocation
		}
	}

	return validNextLocations
}

func combineParts(parts [][]common.Move) []common.Move {
	if len(parts) == 1 {
		return parts[0]
	}

	var combinedParts []common.Move = make([]common.Move, 0)
	for _, topPart := range parts[0] {
		for _, lowerPart := range combineParts(parts[1:]) {
			combinedParts = append(combinedParts, topPart+lowerPart)
		}
	}

	return combinedParts
}

func vectors() common.Moves {
	return common.Moves{
		"^": image.Pt(0, -1),
		">": image.Pt(1, 0),
		"v": image.Pt(0, 1),
		"<": image.Pt(-1, 0),
	}
}
