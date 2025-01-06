package day22monkeymarket

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ewoutquax/advent-of-code-2024/pkg/register"
	"github.com/ewoutquax/advent-of-code-2024/pkg/utils"
)

const (
	Day string = "22"
)

type (
	Number              int
	CollectionKey       string
	MasterCollectionKey struct {
		CollectionKey
		Value int
	}
	Collection struct {
		lastDigit int
	}
	MasterCollection map[MasterCollectionKey]int
	Deltas           map[CollectionKey]Collection
)

func (n Number) Mix(secret Number) Number {
	return n ^ secret
}

func CalculateMostBananas(numbers []Number) int {
	var maxBananas int = 0
	var currBananas int
	var rangeBananas = make(map[CollectionKey]int)

	var masterCollection = make(MasterCollection, 0)
	for _, secret := range numbers {
		mergeCollections(masterCollection, CollectDeltas(secret))
	}

	for masterCollectionKey, count := range masterCollection {
		collectionKey := masterCollectionKey.CollectionKey

		if nrBananas, exists := rangeBananas[collectionKey]; exists {
			currBananas = nrBananas + count*masterCollectionKey.Value
		} else {
			currBananas = count * masterCollectionKey.Value
		}
		rangeBananas[collectionKey] = currBananas

		if maxBananas < currBananas {
			maxBananas = currBananas
		}
	}
	return maxBananas
}

func CollectDeltas(secret Number) Deltas {
	var deltas = make([]int, 0, 4)
	var collections = make(Deltas)

	prevSecret := secret
	prevLastDigit := prevSecret % 10
	for idx := 0; idx < 1999; idx++ {
		currSecret := CalculateNextPrice(prevSecret)
		currLastDigit := currSecret % 10
		deltas = append(deltas, int(currLastDigit-prevLastDigit))

		if len(deltas) == 4 {
			key := buildCollectionKey(deltas)
			if _, exists := collections[key]; !exists {
				collections[key] = Collection{
					lastDigit: int(currLastDigit),
				}
			}

			deltas = deltas[1:]
		}

		prevSecret = currSecret
		prevLastDigit = currLastDigit
	}

	return collections
}

func mergeCollections(masterCollection MasterCollection, newCollection Deltas) MasterCollection {
	for collectionKey, collection := range newCollection {
		masterCollectionKey := MasterCollectionKey{
			CollectionKey: collectionKey,
			Value:         collection.lastDigit,
		}

		if count, exists := masterCollection[masterCollectionKey]; exists {
			masterCollection[masterCollectionKey] = count + 1
		} else {
			masterCollection[masterCollectionKey] = 1
		}
	}

	return masterCollection
}

func buildCollectionKey(deltas []int) CollectionKey {
	var ints = make([]string, 0, 4)
	for _, nr := range deltas {
		ints = append(ints, strconv.Itoa(nr))
	}

	return CollectionKey(strings.Join(ints, ","))
}

func Sum2000thSecret(numbers []Number) int {
	var sum int = 0

	for _, number := range numbers {
		sum += int(Calculate2000thSecret(number))
	}

	return sum
}

func Calculate2000thSecret(inputSecret Number) Number {
	currSecret := inputSecret
	for ctr := 0; ctr < 2000; ctr++ {
		currSecret = CalculateNextPrice(currSecret)
	}

	return currSecret
}

func CalculateNextPrice(secret Number) Number {
	secret1 := PruneSecret((secret * 64).Mix(secret))
	secret2 := PruneSecret((secret1 / 32).Mix(secret1))
	secret3 := PruneSecret((secret2 * 2048).Mix(secret2))

	return secret3
}

func ParseInput(lines []string) []Number {
	var numbers = make([]Number, 0, len(lines))

	for _, line := range lines {
		numbers = append(numbers, Number(utils.ConvStrToI(line)))
	}

	return numbers
}

func PruneSecret(secret Number) Number {
	return secret % 16777216
}

func solvePart1(inputFile string) {
	lines := utils.ReadFileAsLines(inputFile)
	numbers := ParseInput(lines)

	fmt.Printf("Result of day-%s / part-1: %d\n", Day, Sum2000thSecret(numbers))
}

func solvePart2(inputFile string) {
	lines := utils.ReadFileAsLines(inputFile)
	numbers := ParseInput(lines)

	fmt.Printf("Result of day-%s / part-2: %d\n", Day, CalculateMostBananas(numbers))
}

func init() {
	register.Day(Day+"a", solvePart1)
	register.Day(Day+"b", solvePart2)
}
