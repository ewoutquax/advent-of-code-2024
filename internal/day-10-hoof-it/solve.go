package day10hoofit

import (
	"container/heap"
	"fmt"
	"strings"

	"github.com/ewoutquax/advent-of-code-2024/pkg/register"
	"github.com/ewoutquax/advent-of-code-2024/pkg/utils"
)

const (
	Day string = "10"
)

type PathHeap []Location

func (p PathHeap) Len() int           { return len(p) }
func (p PathHeap) Less(i, j int) bool { return p[i].Height < p[j].Height }
func (p PathHeap) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p *PathHeap) Push(x any)        { *p = append(*p, x.(Location)) }

func (p *PathHeap) Pop() any {
	old := *p
	n := len(old)
	x := old[n-1]
	*p = old[0 : n-1]
	return x
}

type Vector [2]int

type Coordinate struct {
	X int
	Y int
}

type Location struct {
	Coordinate
	Height int
}

type Trailhead struct {
	Coordinate
	Ends   map[Coordinate]bool
	Rating int
}

func (t Trailhead) score() int {
	return len(t.Ends)
}

type TopographicMap struct {
	Locations  map[Coordinate]Location
	Trailheads map[Coordinate]*Trailhead
}

func SumTrailheadRatings(topographicMap TopographicMap) int {
	var sum int = 0
	for _, trail := range topographicMap.Trailheads {
		sum += trail.Rating
	}

	return sum
}

func SumTrailheadScores(topographicMap TopographicMap) int {
	var sum int = 0
	for _, trail := range topographicMap.Trailheads {
		sum += trail.score()
	}

	return sum
}

func FindTrailheadsEnds(topographicMap TopographicMap) {
	for _, trailhead := range topographicMap.Trailheads {
		findTrailheadEnds(trailhead, topographicMap)
	}
}

func ParseInput(lines []string) TopographicMap {
	var topographicMap TopographicMap = TopographicMap{
		Locations:  make(map[Coordinate]Location, len(lines)*len(lines[0])),
		Trailheads: make(map[Coordinate]*Trailhead),
	}

	for y, line := range lines {
		for x, char := range strings.Split(line, "") {
			coor := Coordinate{x, y}
			topographicMap.Locations[coor] = Location{
				Coordinate: coor,
				Height:     utils.ConvStrToI(char),
			}
			if char == "0" {
				topographicMap.Trailheads[coor] = &Trailhead{
					Coordinate: coor,
					Ends:       make(map[Coordinate]bool),
					Rating:     0,
				}
			}
		}
	}

	return topographicMap
}

func findTrailheadEnds(trailhead *Trailhead, topographicMap TopographicMap) {
	paths := make(PathHeap, 0)
	paths = append(paths, topographicMap.Locations[trailhead.Coordinate])
	heap.Init(&paths)

	for len(paths) > 0 {
		currentLoc := heap.Pop(&paths).(Location)
		// fmt.Printf("currentLoc: %v\n", currentLoc.toS())
		for _, vector := range allVectors() {
			newCoor := Coordinate{
				X: currentLoc.Coordinate.X + vector[0],
				Y: currentLoc.Coordinate.Y + vector[1],
			}
			// fmt.Printf("Evaluate newCoor: %v; ", newCoor.ToS())

			newLoc, exists := topographicMap.Locations[newCoor]
			if exists && newLoc.Height-currentLoc.Height == 1 {
				if newLoc.Height == 9 {
					// fmt.Print("This is an end-location; add it to the ends; ")
					trailhead.Ends[newLoc.Coordinate] = true
					trailhead.Rating++
				} else {
					// fmt.Print("New loc is one heigher; add it to heap")
					heap.Push(&paths, newLoc)
				}
				// } else {
				// fmt.Print("rejecting loc;")
			}
			// fmt.Println()
		}
	}
}

func allVectors() [4]Vector {
	return [4]Vector{
		[2]int{0, -1},
		[2]int{1, 0},
		[2]int{0, 1},
		[2]int{-1, 0},
	}
}

func solvePart1(inputFile string) {
	lines := utils.ReadFileAsLines(inputFile)
	topographicMap := ParseInput(lines)
	FindTrailheadsEnds(topographicMap)

	fmt.Printf("Result of day-%s / part-1: %d\n", Day, SumTrailheadScores(topographicMap))
}

func solvePart2(inputFile string) {
	lines := utils.ReadFileAsLines(inputFile)
	topographicMap := ParseInput(lines)
	FindTrailheadsEnds(topographicMap)

	fmt.Printf("Result of day-%s / part-2: %d\n", Day, SumTrailheadRatings(topographicMap))
}

func init() {
	register.Day(Day+"b", solvePart2)
	register.Day(Day+"a", solvePart1)
}
