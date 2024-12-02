package day01historianhysteria

import "github.com/ewoutquax/advent-of-code-2024/pkg/register"

const (
	Day string = "01"
)

func init() {
	register.Day(Day+"a", solvePart1)
	register.Day(Day+"b", solvePart2)
}
