package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	_ "github.com/ewoutquax/advent-of-code-2024/internal/day-01-historian-hysteria"
	_ "github.com/ewoutquax/advent-of-code-2024/internal/day-02-red-nosed-reports"
	_ "github.com/ewoutquax/advent-of-code-2024/internal/day-03-mull-it-over"
	_ "github.com/ewoutquax/advent-of-code-2024/internal/day-04-ceres-search"
	_ "github.com/ewoutquax/advent-of-code-2024/internal/day-05-print-queue"
	_ "github.com/ewoutquax/advent-of-code-2024/internal/day-06-guard-gallivant"
	_ "github.com/ewoutquax/advent-of-code-2024/internal/day-07-bridge-repair"
	_ "github.com/ewoutquax/advent-of-code-2024/internal/day-08-resonant-collinearity"
	_ "github.com/ewoutquax/advent-of-code-2024/internal/day-09-disk-fragmenter"
	_ "github.com/ewoutquax/advent-of-code-2024/internal/day-10-hoof-it"
	"github.com/ewoutquax/advent-of-code-2024/pkg/register"
)

func main() {
	for _, puzzle := range getPuzzles() {
		register.ExecDay(puzzle)
	}
}

func getPuzzles() (puzzles []string) {
	var allPuzzles []string = register.GetAllDays()

	selection := readUserInput(fmt.Sprintf("Which puzzle to run %s:\n", allPuzzles))
	switch selection {
	case "":
		latestPuzzle := allPuzzles[len(allPuzzles)-1]
		fmt.Printf("Running latest puzzle: %s\n\n", latestPuzzle)
		puzzles = []string{latestPuzzle}
	case "all":
		fmt.Printf("Running all puzzles\n\n")
		puzzles = allPuzzles
	default:
		fmt.Printf("Running selected puzzle: '%s'\n\n", selection)
		puzzles = []string{selection}
	}

	return
}

func readUserInput(question string) string {
	if len(os.Args) == 2 {
		return os.Args[1]
	}

	fmt.Printf("%s", question)
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	return strings.Trim(text, "\n")
}
