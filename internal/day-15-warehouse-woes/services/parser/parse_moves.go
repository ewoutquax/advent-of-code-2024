package day15warehousewoes_services_parser

import (
	"strings"

	common "github.com/ewoutquax/advent-of-code-2024/internal/day-15-warehouse-woes/common"
)

func (p Parser) parseMoves(lines []string) []common.Move {
	var line string = strings.Join(lines, "")
	var moves = make([]common.Move, 0, len(line))

	for idx := range len(line) {
		moves = append(moves, common.Move(line[idx]))
	}

	return moves
}
