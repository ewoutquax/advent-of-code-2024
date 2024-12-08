package day06guardgallivant

import (
	"fmt"
	"strings"

	"github.com/ewoutquax/advent-of-code-2024/pkg/register"
	"github.com/ewoutquax/advent-of-code-2024/pkg/utils"
)

type Direction uint

const (
	DirectionUp Direction = iota + 1
	DirectionRight
	DirectionDown
	DirectionLeft

	Day string = "06"
)

type Location struct {
	X int
	Y int
}

type Guard struct {
	Location
	Direction
}

func (g Guard) NextLocation() Location {
	vector := toVector(g.Direction)

	return Location{
		X: g.Location.X + vector[0],
		Y: g.Location.Y + vector[1],
	}
}

func (g Guard) isOnMap(u *Universe) bool {
	return g.Location.X >= 0 &&
		g.Location.X <= u.MaxX &&
		g.Location.Y >= 0 &&
		g.Location.Y <= u.MaxY
}

type Universe struct {
	MaxX   int
	MaxY   int
	Blocks map[Location]bool
	Guard
	VisitedLocations map[Location]Direction

	origGuardX         int
	origGuardY         int
	origGuardDirection Direction
}

func (u *Universe) Reset() {
	u.Guard.Location.X = u.origGuardX
	u.Guard.Location.Y = u.origGuardY
	u.Guard.Direction = u.origGuardDirection

	u.VisitedLocations = make(map[Location]Direction, 0)
}

func CountLoopingBlocks(universe Universe) int {
	MoveGuardOffMap(&universe)

	// Copy all visited locations to a list
	var visitedLocations []Location = make([]Location, 0, len(universe.VisitedLocations))
	for loc := range universe.VisitedLocations {
		visitedLocations = append(visitedLocations, loc)
	}

	// Place a block on each of the visited locations
	var count int = 0
	for _, loc := range visitedLocations {
		universe.Blocks[loc] = true
		if GuardIsLooping(&universe) {
			count++
		}
		universe.Reset()
		delete(universe.Blocks, loc)
	}

	return count
}

func MoveGuardOffMap(universe *Universe) {
	for universe.Guard.isOnMap(universe) {
		MoveGuard(universe)
	}
}

func MoveGuard(universe *Universe) {
	nextLoc := universe.Guard.NextLocation()
	if _, exists := universe.Blocks[nextLoc]; exists {
		universe.Guard.Direction = (universe.Guard.Direction % 4) + 1
	} else {
		universe.Guard.Location = nextLoc
		universe.VisitedLocations[nextLoc] = universe.Guard.Direction
	}
}

func GuardIsLooping(universe *Universe) bool {
	for universe.Guard.isOnMap(universe) {
		MoveGuard(universe)

		if universe.VisitedLocations[universe.Guard.NextLocation()] == universe.Guard.Direction {
			return true
		}
	}

	return false
}

func ParseInput(lines []string) Universe {
	var universe Universe = Universe{
		MaxX:             len(lines[0]) - 1,
		MaxY:             len(lines) - 1,
		Blocks:           make(map[Location]bool),
		Guard:            Guard{Location: Location{}, Direction: DirectionUp},
		VisitedLocations: make(map[Location]Direction),
	}

	for y, line := range lines {
		for x, char := range strings.Split(line, "") {
			if char == "#" {
				universe.Blocks[Location{x, y}] = true
			}
			if char == "^" {
				universe.Guard.Location.X = x
				universe.Guard.Location.Y = y
			}
		}
	}

	universe.VisitedLocations[universe.Guard.Location] = universe.Guard.Direction
	universe.origGuardX = universe.Guard.Location.X
	universe.origGuardY = universe.Guard.Location.Y
	universe.origGuardDirection = universe.Guard.Direction

	return universe
}

func toVector(direction Direction) [2]int {
	return map[Direction][2]int{
		DirectionUp:    {0, -1},
		DirectionRight: {1, 0},
		DirectionDown:  {0, 1},
		DirectionLeft:  {-1, 0},
	}[direction]
}

func solvePart1(inputFile string) {
	lines := utils.ReadFileAsLines(inputFile)
	universe := ParseInput(lines)

	MoveGuardOffMap(&universe)

	fmt.Printf("Result of day-%s / part-1: %d\n", Day, len(universe.VisitedLocations)-1)
}

func solvePart2(inputFile string) {
	lines := utils.ReadFileAsLines(inputFile)
	universe := ParseInput(lines)

	fmt.Printf("Result of day-%s / part-2: %d\n", Day, CountLoopingBlocks(universe))
}

func init() {
	register.Day(Day+"a", solvePart1)
	register.Day(Day+"b", solvePart2)
}
