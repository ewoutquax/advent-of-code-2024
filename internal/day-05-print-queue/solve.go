package day05printqueue

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ewoutquax/advent-of-code-2024/pkg/register"
	"github.com/ewoutquax/advent-of-code-2024/pkg/utils"
)

const (
	Day string = "05"
)

type PageNr int

type Order struct {
	Before PageNr
	After  PageNr
}

type Update []PageNr

func (u Update) middlePage() PageNr {
	middle := (len(u) - 1) / 2
	return u[middle]
}

func (u Update) IsValid(orders []Order) bool {
	type Index map[PageNr]int

	index := make(Index, len(u))

	for idx, pageNr := range u {
		index[pageNr] = idx
	}

	for _, order := range orders {
		idxBefore, existsBefore := index[order.Before]
		idxAfter, existsAfter := index[order.After]

		if existsBefore && existsAfter && idxBefore > idxAfter {
			return false
		}
	}

	return true
}

type Manual struct {
	Orders  []Order
	Updates []Update
}

func OrderUpdate(update Update, orders []Order) Update {
	type Index map[PageNr]int
	type IndexPage map[int]PageNr

	var orderedUpdate Update = update
	var doContinue bool = true

	for doContinue {
		index := make(Index, len(update))
		for idx, pageNr := range orderedUpdate {
			index[pageNr] = idx
		}

		for _, order := range orders {
			idxBefore, existsBefore := index[order.Before]
			idxAfter, existsAfter := index[order.After]

			if existsBefore && existsAfter && idxBefore > idxAfter {
				index[order.Before], index[order.After] = index[order.After], index[order.Before]
			}
		}

		indexPage := make(IndexPage, len(update))
		for pageNr, idx := range index {
			indexPage[idx] = pageNr
		}

		orderedUpdate = make(Update, 0, len(update))
		for idx := 0; idx < len(update); idx++ {
			orderedUpdate = append(orderedUpdate, indexPage[idx])
		}

		doContinue = !(orderedUpdate.IsValid(orders))
	}

	return orderedUpdate
}

func SumMiddlePagesOfCorrectedInvalidReports(blocks [][]string) int {
	manual := ParseInput(blocks)
	invalidUpdates := make([]Update, 0, len(manual.Updates))

	for _, update := range manual.Updates {
		if !(update.IsValid(manual.Orders)) {
			invalidUpdates = append(invalidUpdates, OrderUpdate(update, manual.Orders))
		}
	}

	return SumMiddlePages(invalidUpdates)
}

func SumMiddlePagesOfValidReports(blocks [][]string) int {
	manual := ParseInput(blocks)

	validUpdates := make([]Update, 0, len(manual.Updates))

	for _, update := range manual.Updates {
		if update.IsValid(manual.Orders) {
			validUpdates = append(validUpdates, update)
		}
	}

	return SumMiddlePages(validUpdates)
}

func SumMiddlePages(updates []Update) int {
	var sum int = 0

	for _, update := range updates {
		sum += int(update.middlePage())
	}

	return sum
}

func ParseInput(blocks [][]string) Manual {
	var manual Manual = Manual{
		Orders:  parseOrders(blocks[0]),
		Updates: parseUpdate(blocks[1]),
	}

	return manual
}

func parseOrders(lines []string) []Order {
	var orders []Order = make([]Order, 0, len(lines))

	for _, line := range lines {
		parts := strings.Split(line, "|")
		orders = append(
			orders,
			Order{
				Before: convAtoPageNr(parts[0]),
				After:  convAtoPageNr(parts[1]),
			},
		)
	}

	return orders
}

func parseUpdate(lines []string) []Update {
	var updates []Update = make([]Update, 0, len(lines))

	for _, line := range lines {
		nrs := strings.Split(line, ",")
		update := make([]PageNr, 0, len(nrs))

		for _, nr := range nrs {
			update = append(update, convAtoPageNr(nr))
		}
		updates = append(updates, update)
	}

	return updates
}

func convAtoPageNr(s string) PageNr {
	nr, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return PageNr(nr)
}

func solvePart1(inputFile string) {
	blocks := utils.ReadFileAsBlocks(inputFile)

	fmt.Printf(
		"Result of day-%s / part-1: %d\n",
		Day,
		SumMiddlePagesOfValidReports(blocks),
	)
}

func solvePart2(inputFile string) {
	blocks := utils.ReadFileAsBlocks(inputFile)

	fmt.Printf(
		"Result of day-%s / part-2: %d\n",
		Day,
		SumMiddlePagesOfCorrectedInvalidReports(blocks),
	)
}

func init() {
	register.Day(Day+"a", solvePart1)
	register.Day(Day+"b", solvePart2)
}
