package day08resonantcollinearity

import (
	"fmt"
	"strings"

	"github.com/ewoutquax/advent-of-code-2024/pkg/register"
	"github.com/ewoutquax/advent-of-code-2024/pkg/utils"
)

const (
	Day string = "08"
)

type Signal string
type Locations []Location
type IndexedLocations map[Location]bool

type Vector struct {
	x int
	y int
}

type Location struct {
	X int
	Y int
}

func (l Location) ToS() string {
	return fmt.Sprintf("[%d, %d]", l.X, l.Y)
}

type Universe struct {
	MaxX      int
	MaxY      int
	Signals   map[Signal]Locations
	Antinodes IndexedLocations
}

func CountAntinodes(u Universe) int {
	return len(u.Antinodes)
}

func AddAntinodes(u *Universe, withRepeater bool) {
	for _, locs := range u.Signals {
		for _, sourceLoc := range locs {
			// for each loc, go by all the other locs
			for _, targetLoc := range locs {
				if sourceLoc != targetLoc {
					vector := Vector{targetLoc.X - sourceLoc.X, targetLoc.Y - sourceLoc.Y}
					repeater := 2
					if withRepeater {
						repeater = 1
					}
					doContinue := true
					for doContinue {
						newLoc := Location{
							X: sourceLoc.X + repeater*vector.x,
							Y: sourceLoc.Y + repeater*vector.y,
						}
						if newLoc.X >= 0 && newLoc.X <= u.MaxX &&
							newLoc.Y >= 0 && newLoc.Y <= u.MaxY {
							repeater++
							doContinue = true
							u.Antinodes[newLoc] = true
						} else {
							doContinue = false
						}
						if !withRepeater {
							doContinue = false
						}
					}
				}
			}
		}
	}
}

func ParseInput(lines []string) Universe {
	var universe Universe = Universe{
		MaxX:      len(lines[0]) - 1,
		MaxY:      len(lines) - 1,
		Signals:   make(map[Signal]Locations),
		Antinodes: make(IndexedLocations),
	}

	for y, line := range lines {
		for x, char := range strings.Split(line, "") {
			if char != "." {
				signal := Signal(char)
				if _, exists := universe.Signals[signal]; !exists {
					universe.Signals[signal] = make(Locations, 0)
				}
				universe.Signals[signal] = append(universe.Signals[signal], Location{x, y})
			}
		}
	}

	return universe
}

func solvePart1(inputFile string) {
	lines := utils.ReadFileAsLines(inputFile)
	universe := ParseInput(lines)

	AddAntinodes(&universe, false)

	fmt.Printf("Result of day-%s / part-1: %d\n", Day, CountAntinodes(universe))
}

func solvePart2(inputFile string) {
	lines := utils.ReadFileAsLines(inputFile)
	universe := ParseInput(lines)
	AddAntinodes(&universe, true)

	fmt.Printf("Result of day-%s / part-2: %d\n", Day, CountAntinodes(universe))
}

func init() {
	register.Day(Day+"a", solvePart1)
	register.Day(Day+"b", solvePart2)
}
