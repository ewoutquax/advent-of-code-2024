package day15warehousewoes

import (
	"fmt"

	services "github.com/ewoutquax/advent-of-code-2024/internal/day-15-warehouse-woes/services"
	"github.com/ewoutquax/advent-of-code-2024/pkg/register"
	"github.com/ewoutquax/advent-of-code-2024/pkg/utils"
)

const (
	Day string = "15"
)

func solvePart1(inputFile string) {
	blocks := utils.ReadFileAsBlocks(inputFile)
	universe := services.BuildParser().ParseInput(blocks)

	services.BuildSolver().FollowMoves(universe)

	fmt.Printf("Result of day-%s / part-1: %d\n", Day, universe.Score())
}

func solvePart2(inputFile string) {
	blocks := utils.ReadFileAsBlocks(inputFile)
	universe := services.BuildParser(services.WithDoubling()).ParseInput(blocks)

	services.BuildSolver().FollowMoves(universe)

	fmt.Printf("Result of day-%s / part-2: %d\n", Day, universe.Score()/2)
}

func init() {
	register.Day(Day+"a", solvePart1)
	register.Day(Day+"b", solvePart2)
}
