package day11plutonianpebbles

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ewoutquax/advent-of-code-2024/pkg/register"
	"github.com/ewoutquax/advent-of-code-2024/pkg/utils"
)

const (
	Day string = "11"
)

type CacheKey struct {
	Stone
	depth int
}

var cache map[CacheKey]int

type Stone int

func (s Stone) Length(depth int) int {
	if depth == 0 {
		return 1
	}

	cacheKey := CacheKey{
		Stone: s,
		depth: depth,
	}

	if length, exists := cache[cacheKey]; exists {
		return length
	}

	length := 0
	for _, subStone := range s.Resolve() {
		length += subStone.Length(depth - 1)
	}
	cache[cacheKey] = length

	return length
}

func (s Stone) Resolve() []Stone {
	strValue := strconv.Itoa(int(s))
	switch true {
	case s == 0:
		return []Stone{1}
	case len(strValue)%2 == 0:
		middle := len(strValue) / 2

		return []Stone{
			Stone(utils.ConvStrToI(strValue[:middle])),
			Stone(utils.ConvStrToI(strValue[middle:])),
		}
	default:
		return []Stone{s * 2024}
	}
}

func NrStonesAfterBlinks(stones []Stone, nrBlinks int) int {
	cache = make(map[CacheKey]int)

	var count int = 0
	for _, stone := range stones {
		count += stone.Length(nrBlinks)
	}

	return count
}

func ParseInput(line string) []Stone {
	numbers := strings.Split(line, " ")

	var stones []Stone = make([]Stone, 0, len(numbers))

	for _, nr := range numbers {
		stones = append(stones, Stone(utils.ConvStrToI(nr)))
	}

	return stones
}

func solvePart1(inputFile string) {
	line := utils.ReadFileAsLine(inputFile)
	stones := ParseInput(line)

	fmt.Printf("Result of day-%s / part-1: %d\n", Day, NrStonesAfterBlinks(stones, 25))
}

func solvePart2(inputFile string) {
	line := utils.ReadFileAsLine(inputFile)
	stones := ParseInput(line)

	fmt.Printf("Result of day-%s / part-2: %d\n", Day, NrStonesAfterBlinks(stones, 75))
}

func init() {
	register.Day(Day+"a", solvePart1)
	register.Day(Day+"b", solvePart2)
}
