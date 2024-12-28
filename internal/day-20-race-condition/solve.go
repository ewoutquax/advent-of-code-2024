package day20racecondition

import (
	"container/heap"
	"fmt"
	. "image"
	"math"
	"strings"

	"github.com/ewoutquax/advent-of-code-2024/pkg/register"
	"github.com/ewoutquax/advent-of-code-2024/pkg/utils"
)

const (
	Day string = "20"
)

type Track struct {
	Walls  map[Point]bool
	Start  Point
	Finish Point

	MaxX int
	MaxY int
}

func (t Track) inRange(point Point) bool {
	return point.In(Rectangle{
		Min: Pt(0, 0),
		Max: Pt(t.MaxX, t.MaxY),
	})
}

type Path struct {
	Point
	NrSteps  int
	prevPath *Path
}

func (p Path) ToKey() PathKey {
	return PathKey(p.Point.String())
}

func manhattanDistance(source, target Point) int {
	return utils.Abs(target.X-source.X) + utils.Abs(target.Y-source.Y)
}

type PathHeap []Path
type PathKey string

func (h PathHeap) Len() int           { return len(h) }
func (h PathHeap) Less(i, j int) bool { return h[i].NrSteps < h[j].NrSteps }
func (h PathHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *PathHeap) Push(x any)        { *h = append(*h, x.(Path)) }

func (h *PathHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func FindPaths(track Track, minStepsSkipped, maxStepsCheating int) (map[int]int, int) {
	var visitedPoints map[PathKey]int = make(map[PathKey]int, 0)
	var solutionsGroupedBySteps = make(map[int]int)
	var solutionBase = Path{
		Point:    Point{},
		NrSteps:  math.MaxInt,
		prevPath: &Path{},
	}

	var pathHeap = make(PathHeap, 0)

	pathHeap = append(pathHeap, Path{
		Point:    track.Start,
		NrSteps:  0,
		prevPath: nil,
	})

	heap.Init(&pathHeap)

	for len(pathHeap) > 0 {
		currPath := heap.Pop(&pathHeap).(Path)

		if currPath.Point.Eq(track.Finish) {
			if solutionBase.NrSteps > currPath.NrSteps {
				solutionBase = currPath
			}
		}

		// Check reachable locations
		for _, nextPoint := range findNextPoints(currPath, track) {
			newPath := Path{
				Point:    nextPoint,
				NrSteps:  currPath.NrSteps + 1,
				prevPath: &currPath,
			}

			if prevNrSteps, exists := visitedPoints[newPath.ToKey()]; !exists || exists && prevNrSteps > newPath.NrSteps {
				visitedPoints[newPath.ToKey()] = newPath.NrSteps
				heap.Push(&pathHeap, newPath)
			}
		}
	}

	// Store each point of the solution in trackPoints
	currPath := &solutionBase
	trackPoints := make(map[Point]int, 0)
	for currPath != nil {
		trackPoints[currPath.Point] = currPath.NrSteps
		currPath = currPath.prevPath
	}

	for pointSource, nrStepsSource := range trackPoints {
		for pointTarget, nrStepsTarget := range trackPoints {
			distance := manhattanDistance(pointSource, pointTarget)
			// source and target should not be the same point
			profit := nrStepsTarget - nrStepsSource - distance
			if distance > 0 && distance <= maxStepsCheating && profit >= minStepsSkipped {
				if count, exists := solutionsGroupedBySteps[profit]; exists {
					solutionsGroupedBySteps[profit] = count + 1
				} else {
					solutionsGroupedBySteps[profit] = 1
				}
				// fmt.Printf("From %s to %s yields a profit of: %d\n", pointSource.String(), pointTarget.String(), profit)
			}
		}
	}

	return solutionsGroupedBySteps, solutionBase.NrSteps
}

func findNextPoints(currPath Path, track Track) []Point {
	var nextPoints = make([]Point, 0, 4)
	for _, vector := range allVectors() {
		nextPoint := currPath.Point.Add(vector)
		_, inWallNext := track.Walls[nextPoint]

		if track.inRange(nextPoint) && !inWallNext {
			nextPoints = append(nextPoints, nextPoint)
		}
	}

	return nextPoints
}

func allVectors() [4]Point {
	return [4]Point{
		Pt(0, -1),
		Pt(1, 0),
		Pt(0, 1),
		Pt(-1, 0),
	}
}

func ParseInput(lines []string) Track {
	var track = Track{
		Walls: make(map[Point]bool, 0),
		MaxX:  len(lines[0]) - 1,
		MaxY:  len(lines) - 1,
	}

	for y, line := range lines {
		for x, char := range strings.Split(line, "") {
			switch char {
			case "#":
				track.Walls[Pt(x, y)] = true
			case "S":
				track.Start = Pt(x, y)
			case "E":
				track.Finish = Pt(x, y)
			case ".":
				continue
			default:
				panic("No valid case found")
			}
		}
	}

	return track
}

func solvePart1(inputFile string) {
	lines := utils.ReadFileAsLines(inputFile)
	track := ParseInput(lines)

	solutions, nrStepsBaseSolution := FindPaths(track, 100, 2)
	fmt.Printf("nrStepsBaseSolution: %v\n", nrStepsBaseSolution)

	count := 0
	for _, profitCount := range solutions {
		count += profitCount
	}

	fmt.Printf("Result of day-%s / part-1: %d\n", Day, count)
}

func solvePart2(inputFile string) {
	lines := utils.ReadFileAsLines(inputFile)
	track := ParseInput(lines)

	solutions, nrStepsBaseSolution := FindPaths(track, 100, 20)
	fmt.Printf("nrStepsBaseSolution: %v\n", nrStepsBaseSolution)

	count := 0
	for _, profitCount := range solutions {
		count += profitCount
	}
	fmt.Printf("Result of day-%s / part-2: %d\n", Day, count)

	/*
	  Answers:
	  --------
	  1001786: too low
	*/
}

func init() {
	register.Day(Day+"a", solvePart1)
	register.Day(Day+"b", solvePart2)
}
