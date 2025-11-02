package day11plutonianpebbles

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ewoutquax/advent-of-code-2024/pkg/register"
)

const (
	Day string = "11"
)

type (
	Rock int
)

func (r Rock) NextRocks() []Rock {
	stringRock := fmt.Sprintf("%d", r)
	nrDigits := len(strings.Split(stringRock, ""))

	switch true {
	case r == 0:
		return []Rock{1}
	case (nrDigits % 2) == 0:
		leftRock := stringRock[:nrDigits/2]
		rightRock := stringRock[nrDigits/2:]

		return []Rock{
			convAtoRock(leftRock),
			convAtoRock(rightRock),
		}
	default:
		return []Rock{r * 2024}
	}
}

func (r Rock) CountAfterBlinks(blinks int) int {
	if blinks == 0 {
		return 1
	}

	return CountRocksAfterBlinks(r.NextRocks(), blinks-1)
}

func CountRocksAfterBlinks(rocks []Rock, blinks int) int {
	var sum int = 0

	for _, rock := range rocks {
		sum += rock.CountAfterBlinks(blinks)
	}

	return sum
}

func ParseInput(line string) []Rock {
	parts := strings.Split(line, " ")

	var rocks = make([]Rock, 0, len(parts))

	for _, nr := range parts {
		rocks = append(rocks, convAtoRock(nr))
	}

	return rocks
}

func convAtoRock(s string) Rock {
	nr, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return Rock(nr)
}

func solvePart1(inputFile string) {
	_ = inputFile

	rocks := ParseInput("0 7 198844 5687836 58 2478 25475 894")
	fmt.Printf("Result of day-%s / part-1: %d\n", Day, CountRocksAfterBlinks(rocks, 25))
}

func solvePart2(inputFile string) {
	_ = inputFile

	rocks := ParseInput("0 7 198844 5687836 58 2478 25475 894")
	fmt.Printf("Result of day-%s / part-1: %d\n", Day, CountRocksAfterBlinks(rocks, 75))
}

func init() {
	register.Day(Day+"a", solvePart1)
	register.Day(Day+"b", solvePart2)
}
