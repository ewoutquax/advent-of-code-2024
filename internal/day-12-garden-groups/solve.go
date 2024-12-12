package day12gardengroups

import (
	"container/heap"
	"fmt"
	. "image"
	"strings"

	"github.com/ewoutquax/advent-of-code-2024/pkg/register"
	"github.com/ewoutquax/advent-of-code-2024/pkg/utils"
)

type Orientation uint

type Side uint

const (
	SideUp Side = iota + 1
	SideRight
	SideDown
	SideLeft
)

const (
	OrientationHorizontal Orientation = iota + 1
	OrientationVertical

	Day string = "12"
)

type PerimiterHeap []Perimiter

func (h PerimiterHeap) Len() int      { return len(h) }
func (h PerimiterHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *PerimiterHeap) Push(x any)   { *h = append(*h, x.(Perimiter)) }
func (h PerimiterHeap) Less(i, j int) bool {
	if h[i].Orientation == OrientationVertical && h[j].Orientation == OrientationHorizontal {
		return true
	}

	if h[i].Orientation == OrientationHorizontal && h[j].Orientation == OrientationVertical {
		return false
	}

	if h[i].Orientation == OrientationVertical {
		// `j` MUST also be vertical
		if h[i].Side == SideLeft && h[j].Side == SideRight {
			return true
		}
		if h[i].Side == SideRight && h[j].Side == SideLeft {
			return false
		}

		return h[i].Point.X < h[j].Point.X ||
			(h[i].Point.X == h[j].Point.X &&
				h[i].Point.Y < h[j].Point.Y)
	}

	// Both MUST be horizontal
	if h[i].Side == SideUp && h[j].Side == SideDown {
		return true
	}
	if h[i].Side == SideDown && h[j].Side == SideUp {
		return false
	}

	return h[i].Point.Y < h[j].Point.Y ||
		(h[i].Point.Y == h[j].Point.Y &&
			h[i].Point.X < h[j].Point.X)
}

func (h *PerimiterHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type Perimiter struct {
	Point
	Orientation
	Side
}

type Plant string
type Region int

type Location struct {
	Point
	Region
	Plant
}

type Garden struct {
	Locations map[Point]*Location
	Regions   map[Region][]Point
}

func SumPricesByPerimiter(garden Garden) int {
	var sum = 0

	for region := range garden.Regions {
		sum += CalculatePriceByPerimiter(region, garden)
	}

	return sum
}

func SumPricesBySides(garden Garden) int {
	var sum = 0

	for region := range garden.Regions {
		sum += CalculatePriceBySide(region, garden)
	}

	return sum
}

func CalculatePriceByPerimiter(region Region, garden Garden) int {
	return len(garden.Regions[region]) * CalculatePerimiter(region, garden)
}

func CalculatePriceBySide(region Region, garden Garden) int {
	return len(garden.Regions[region]) * CalculateSides(region, garden)
}

func CalculateSides(region Region, garden Garden) int {
	var perimiters = make(PerimiterHeap, 0)

	for _, point := range garden.Regions[region] {
		currLoc := garden.Locations[point]
		for _, vector := range allVectors() {
			if otherLocation, exists := garden.Locations[point.Add(vector)]; !exists || exists && currLoc.Plant != otherLocation.Plant {
				p := Perimiter{
					Point:       point.Add(vector),
					Orientation: OrientationVertical,
					Side:        toSide(vector),
				}

				if vector.X == 0 {
					p.Orientation = OrientationHorizontal
				}

				perimiters = append(perimiters, p)
			}
		}
	}

	heap.Init(&perimiters)

	var nrSides = 1
	prev := heap.Pop(&perimiters).(Perimiter)
	for len(perimiters) > 0 {
		curr := heap.Pop(&perimiters).(Perimiter)

		nrSides++
		if prev.Orientation == curr.Orientation {
			if curr.Orientation == OrientationVertical &&
				curr.Point.X == prev.Point.X &&
				curr.Point.Y-prev.Point.Y == 1 {
				nrSides--
			}
			if curr.Orientation == OrientationHorizontal &&
				curr.Point.Y == prev.Point.Y &&
				curr.Point.X-prev.Point.X == 1 {
				nrSides--
			}
		}
		prev = curr
	}

	return nrSides
}

func toSide(vector Point) Side {
	return map[Point]Side{
		{0, -1}: SideUp,
		{1, 0}:  SideRight,
		{0, 1}:  SideDown,
		{-1, 0}: SideLeft,
	}[vector]
}

func CalculatePerimiter(region Region, garden Garden) int {
	var perimiter = 0

	for _, point := range garden.Regions[region] {
		currLoc := garden.Locations[point]
		for _, vector := range allVectors() {
			if otherLocation, exists := garden.Locations[point.Add(vector)]; !exists || exists && currLoc.Plant != otherLocation.Plant {
				perimiter++
			}
		}
	}

	return perimiter
}

func ParseInput(lines []string) Garden {
	var maxRegion = Region(0)
	_ = maxRegion

	var garden = Garden{
		Locations: make(map[Point]*Location, len(lines)*len(lines[0])),
		Regions:   make(map[Region][]Point),
	}

	for y, line := range lines {
		for x, char := range strings.Split(line, "") {
			currPoint := Pt(x, y)
			currLoc := Location{
				Point:  currPoint,
				Region: 0,
				Plant:  Plant(char),
			}
			garden.Locations[currPoint] = &currLoc

			// Connect to an existing region
			if upLoc, exists := garden.Locations[currPoint.Add(Pt(0, -1))]; exists && upLoc.Plant == Plant(char) {
				currLoc.Region = upLoc.Region
			}
			if leftLoc, exists := garden.Locations[currPoint.Add(Pt(-1, 0))]; exists && leftLoc.Plant == Plant(char) {
				if currLoc.Region == 0 {
					currLoc.Region = leftLoc.Region
				} else {
					// Region of upperloc is leading; change all location with the same region as the one on the left
					changeLocationRegions(garden, leftLoc.Region, currLoc.Region)
				}
			}
			if currLoc.Region == 0 {
				maxRegion++
				currLoc.Region = maxRegion
			}
		}
	}

	// Build the map with all the found regions
	for _, loc := range garden.Locations {
		if _, exists := garden.Regions[loc.Region]; !exists {
			garden.Regions[loc.Region] = make([]Point, 0)
		}
		garden.Regions[loc.Region] = append(garden.Regions[loc.Region], loc.Point)
	}

	return garden
}

func changeLocationRegions(g Garden, regionFrom, regionTo Region) {
	for _, loc := range g.Locations {
		if loc.Region == regionFrom {
			loc.Region = regionTo
		}
	}
}

func allVectors() [4]Point {
	return [4]Point{
		{0, -1},
		{1, 0},
		{0, 1},
		{-1, 0},
	}
}

// Calculate the solution of part 1
func solvePart1(inputFile string) {
	lines := utils.ReadFileAsLines(inputFile)
	garden := ParseInput(lines)

	fmt.Printf("Result of day-%s / part-1: %d\n", Day, SumPricesByPerimiter(garden))
}

func solvePart2(inputFile string) {
	lines := utils.ReadFileAsLines(inputFile)
	garden := ParseInput(lines)

	fmt.Printf("Result of day-%s / part-2: %d\n", Day, SumPricesBySides(garden))
}

func init() {
	register.Day(Day+"a", solvePart1)
	register.Day(Day+"b", solvePart2)
}
