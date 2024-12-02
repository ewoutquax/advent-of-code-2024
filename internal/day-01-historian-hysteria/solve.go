package day01historianhysteria

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"

	"github.com/ewoutquax/advent-of-code-2024/pkg/utils"
)

type Lists struct {
	Left  []int
	Right []int
}

func SumDistanceBetweenSmallest(lists Lists) int {
	slices.Sort(lists.Left)
	slices.Sort(lists.Right)

	var sum int = 0
	for idx := 0; idx < len(lists.Left); idx++ {
		sum += abs(lists.Left[idx] - lists.Right[idx])
	}
	return sum
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func SimilarityScore(lists Lists) int {
	type Occurences map[int]int
	var occurences Occurences = make(Occurences, len(lists.Right))
	for _, nr := range lists.Right {
		if count, exists := occurences[nr]; exists {
			occurences[nr] = count + 1
		} else {
			occurences[nr] = 1
		}
	}

	var score int = 0
	for _, nr := range lists.Left {
		if count, exists := occurences[nr]; exists {
			score += nr * count
		}
	}

	return score
}

func ParseInput(lines []string) Lists {
	var lists Lists = Lists{
		Left:  make([]int, 0, len(lines)),
		Right: make([]int, 0, len(lines)),
	}

	regex := regexp.MustCompile(`(\d+)\s+(\d+)`)

	for _, line := range lines {
		matches := regex.FindAllStringSubmatch(line, -1)

		lists.Left = append(lists.Left, convAtoi(matches[0][1]))
		lists.Right = append(lists.Right, convAtoi(matches[0][2]))
	}

	return lists
}

func convAtoi(s string) int {
	nr, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return nr
}

func solvePart1(inputFile string) {
	lines := utils.ReadFileAsLines(inputFile)
	lists := ParseInput(lines)

	fmt.Printf("Result of day-%s / part-1: %d\n", Day, SumDistanceBetweenSmallest(lists))
}

func solvePart2(inputFile string) {
	lines := utils.ReadFileAsLines(inputFile)
	lists := ParseInput(lines)

	fmt.Printf("Result of day-%s / part-2: %d\n", Day, SimilarityScore(lists))
}
