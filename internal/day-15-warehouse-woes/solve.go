package day15warehousewoes

import (
	"fmt"
	. "image"
	"strings"

	"github.com/ewoutquax/advent-of-code-2024/pkg/register"
	"github.com/ewoutquax/advent-of-code-2024/pkg/utils"
)

type Moveable interface {
	CanMove(Direction, *Universe) bool
	Move(Direction, *Universe)
	toS() string
	Hash() int
}

type (
	Direction string
)

type Wall struct {
	Point
}
type Box struct {
	Point
}
type Robot struct {
	Point
}

func (w *Wall) Hash() int { return 0 }
func (b *Box) Hash() int  { return b.Point.X + 100*b.Point.Y }

func (w *Wall) toS() string { return "#" }
func (b *Box) toS() string  { return "O" }

func (w *Wall) CanMove(_ Direction, _ *Universe) bool { return false }
func (w Wall) Move(_ Direction, _ *Universe)          { panic("Walls can't move") }

func (b *Box) CanMove(d Direction, u *Universe) bool {
	nextPoint := b.Point.Add(toVector(d))
	return pointEmptyOrMoveable(nextPoint, d, u)
}

func (r *Robot) CanMove(d Direction, u *Universe) bool {
	nextPoint := r.Point.Add(toVector(d))

	// tmp := pointEmptyOrMoveable(nextPoint, d, u)
	// if !tmp {
	// 	fmt.Printf("Robot can't move %s from %s\n", d.toS(), r.Point.String())
	// }

	return pointEmptyOrMoveable(nextPoint, d, u)
}

func (b *Box) Move(d Direction, u *Universe) {
	nextPoint := b.Point.Add(toVector(d))

	if moveable, exists := u.Moveables[nextPoint]; exists {
		moveable.Move(d, u)
	}

	delete(u.Moveables, b.Point)
	u.Moveables[nextPoint] = b
	b.Point = nextPoint
}

func (r *Robot) Move(d Direction, u *Universe) {
	nextPoint := r.Point.Add(toVector(d))

	if moveable, exists := u.Moveables[nextPoint]; exists {
		moveable.Move(d, u)
	}

	// fmt.Printf("Bot moves %s from %s to %s\n", d.toS(), r.Point.String(), nextPoint.String())
	r.Point = nextPoint
}

func pointEmptyOrMoveable(nextPoint Point, d Direction, u *Universe) bool {
	moveable, exists := u.Moveables[nextPoint]
	if !exists {
		// fmt.Printf("Location %s is empty\n", nextPoint.String())
		return true
	}
	return moveable.CanMove(d, u)
}

const (
	DirectionUp    Direction = "^"
	DirectionRight Direction = ">"
	DirectionDown  Direction = "v"
	DirectionLeft  Direction = "<"

	Day string = "15"

	ObjectBox   string = "O"
	ObjectRobot string = "@"
	ObjectWall  string = "#"
)

func (d Direction) toS() string {
	return map[Direction]string{
		DirectionUp:    "up",
		DirectionRight: "right",
		DirectionDown:  "down",
		DirectionLeft:  "left",
	}[d]
}

type Warehouse struct {
	maxX      int
	maxY      int
	Walls     map[Point]Moveable
	Boxes     map[Point]Moveable
	Moveables map[Point]Moveable
	Robot
}

type Universe struct {
	Warehouse
	Steps []Direction
}

func (u Universe) Hash() int {
	var hash = 0

	for _, moveable := range u.Moveables {
		hash += moveable.Hash()
	}

	return hash
}

func TakeSteps(u *Universe) {
	for _, direction := range u.Steps {
		if u.Robot.CanMove(direction, u) {
			u.Robot.Move(direction, u)
			// draw(*u)
		}
	}
}

func ParseInput(blocks [][]string) Universe {
	var u = Universe{
		Warehouse: parseWarehouse(blocks[0]),
		Steps:     parseDirections(blocks[1]),
	}

	return u
}

func parseWarehouse(lines []string) Warehouse {
	var w = Warehouse{
		maxX:      len(lines[0]),
		maxY:      len(lines),
		Moveables: make(map[Point]Moveable, 0),
	}

	for y, line := range lines {
		for x, char := range strings.Split(line, "") {
			switch char {
			case string(ObjectBox):
				w.Moveables[Pt(x, y)] = &Box{
					Point: Pt(x, y),
				}
			case string(ObjectWall):
				w.Moveables[Pt(x, y)] = &Wall{
					Point: Pt(x, y),
				}
			case string(ObjectRobot):
				w.Robot = Robot{
					Point: Pt(x, y),
				}
			case ".":
				continue
			default:
				panic("No valid case found")
			}
		}
	}

	return w
}

func parseDirections(lines []string) []Direction {
	var directions = make([]Direction, 0, len(lines[0]))

	singleLine := strings.Join(lines, "")
	for _, char := range strings.Split(singleLine, "") {
		directions = append(directions, convToDirection(char))
	}

	return directions
}

func draw(u Universe) {
	fmt.Println("Drawing warehouse")
	fmt.Println("-----------------")
	for y := 0; y < u.maxY; y++ {
		for x := 0; x < u.maxX; x++ {
			point := Pt(x, y)
			moveable, exists := u.Moveables[point]
			switch true {
			case u.Robot.Point.Eq(point):
				fmt.Print("@")
			case !exists:
				fmt.Print(".")
			case exists:
				fmt.Printf("%s", moveable.toS())
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func convToDirection(char string) Direction {
	return map[string]Direction{
		string(DirectionUp):    DirectionUp,
		string(DirectionRight): DirectionRight,
		string(DirectionDown):  DirectionDown,
		string(DirectionLeft):  DirectionLeft,
	}[char]
}

func toVector(d Direction) Point {
	return map[Direction]Point{
		DirectionUp:    Pt(0, -1),
		DirectionRight: Pt(1, 0),
		DirectionDown:  Pt(0, 1),
		DirectionLeft:  Pt(-1, 0),
	}[d]
}

func solvePart1(inputFile string) {
	blocks := utils.ReadFileAsBlocks(inputFile)
	universe := ParseInput(blocks)

	TakeSteps(&universe)

	fmt.Printf("Result of day-%s / part-1: %d\n", Day, universe.Hash())
}

func solvePart2(inputFile string) {
	_ = inputFile

	var count int = 0
	fmt.Printf("Result of day-%s / part-2: %d\n", Day, count)
}

func init() {
	register.Day(Day+"a", solvePart1)
	// register.Day(Day+"b", solvePart2)
}
