package day18ramrun

import (
	"container/heap"
	"fmt"
	"image"
	"math"
	"regexp"

	"github.com/ewoutquax/advent-of-code-2024/pkg/register"
	"github.com/ewoutquax/advent-of-code-2024/pkg/utils"
)

const (
	Day string = "18"
)

type Path struct {
	nrSteps int
	image.Point
}
type PathHeap []Path

func (h PathHeap) Len() int           { return len(h) }
func (h PathHeap) Less(i, j int) bool { return h[i].nrSteps < h[j].nrSteps }
func (h PathHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *PathHeap) Push(x any)        { *h = append(*h, x.(Path)) }

func (h *PathHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type UniverseOpt func(*Universe)

type Universe struct {
	MaxX          int
	MaxY          int
	CorruptSpaces map[image.Point]bool
}

var (
	visitedPoints map[image.Point]int
)

func FindFirstBlockingByte(universe Universe, lines []string) string {
	FindMinSteps(universe)

	for _, line := range lines {
		pnt := parseLine(line)
		universe.CorruptSpaces[pnt] = true

		if _, visited := visitedPoints[pnt]; visited {
			fmt.Printf("Punt %s ligt op het pad: zoek een nieuw pad\n", pnt.String())
			if FindMinSteps(universe) == math.MaxInt {
				return fmt.Sprintf("%d,%d", pnt.X, pnt.Y)
			}
		}
	}

	panic("I keep finding exits...")
}

func FindMinSteps(universe Universe) int {
	var minSteps int = math.MaxInt
	visitedPoints = make(map[image.Point]int)

	var paths PathHeap = make(PathHeap, 0)
	paths = append(paths, Path{
		nrSteps: 0,
		Point:   image.Pt(0, 0),
	})

	heap.Init(&paths)
	for len(paths) > 0 {
		currPath := heap.Pop(&paths).(Path)

		if currPath.X == universe.MaxX &&
			currPath.Y == universe.MaxY &&
			minSteps > currPath.nrSteps {
			minSteps = currPath.nrSteps
		}

		for _, newPoint := range findNewPoints(currPath, universe) {
			if prevNrSteps, exists := visitedPoints[newPoint]; !exists && prevNrSteps > currPath.nrSteps+1 {
				visitedPoints[newPoint] = currPath.nrSteps + 1
				heap.Push(&paths, Path{
					nrSteps: currPath.nrSteps + 1,
					Point:   newPoint,
				})
			}
		}
	}

	return minSteps
}

func findNewPoints(path Path, universe Universe) []image.Point {
	var newPoints = make([]image.Point, 0, 4)
	for _, vector := range allVectors() {
		newPoint := path.Point.Add(vector)
		if _, exists := universe.CorruptSpaces[newPoint]; !exists &&
			newPoint.X >= 0 &&
			newPoint.Y >= 0 &&
			newPoint.X <= universe.MaxX &&
			newPoint.Y <= universe.MaxY {

			newPoints = append(newPoints, newPoint)
		}
	}

	return newPoints
}

func BuildUniverse(funcOpts ...UniverseOpt) Universe {
	universe := Universe{
		MaxX:          0,
		MaxY:          0,
		CorruptSpaces: make(map[image.Point]bool, 0),
	}
	for _, funcOpt := range funcOpts {
		funcOpt(&universe)
	}

	return universe
}

func WithMaxValues(maxX, maxY int) UniverseOpt {
	return func(universe *Universe) {
		universe.MaxX = maxX
		universe.MaxY = maxY
	}
}

func WithCorruptSpaces(lines []string) UniverseOpt {
	return func(universe *Universe) {
		universe.CorruptSpaces = ParseInput(lines)
	}
}

func ParseInput(lines []string) map[image.Point]bool {
	var corruptSpaces = make(map[image.Point]bool)

	for _, line := range lines {
		pnt := parseLine(line)
		corruptSpaces[pnt] = true
	}

	return corruptSpaces
}

func parseLine(line string) image.Point {
	ex := regexp.MustCompile(`(\d+),(\d+)`)
	matches := ex.FindAllStringSubmatch(line, -1)
	return image.Pt(
		utils.ConvStrToI(matches[0][1]),
		utils.ConvStrToI(matches[0][2]),
	)
}

func allVectors() [4]image.Point {
	return [4]image.Point{
		image.Pt(0, -1),
		image.Pt(1, 0),
		image.Pt(0, 1),
		image.Pt(-1, 0),
	}
}

func solvePart1(inputFile string) {
	lines := utils.ReadFileAsLines(inputFile)

	universe := BuildUniverse(
		WithMaxValues(70, 70),
		WithCorruptSpaces(lines[:1024]),
	)

	fmt.Printf("Result of day-%s / part-1: %d\n", Day, FindMinSteps(universe))
}

func solvePart2(inputFile string) {
	lines := utils.ReadFileAsLines(inputFile)
	universe := BuildUniverse(
		WithMaxValues(70, 70),
	)

	fmt.Printf("Result of day-%s / part-2: %s\n", Day, FindFirstBlockingByte(universe, lines))
}

func init() {
	register.Day(Day+"a", solvePart1)
	register.Day(Day+"b", solvePart2)
}
