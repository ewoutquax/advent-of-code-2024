package day15warehousewoes_services_parser

import (
	common "github.com/ewoutquax/advent-of-code-2024/internal/day-15-warehouse-woes/common"
)

type Parser struct {
	Doubling bool
}

func DefaultParser() Parser {
	return Parser{
		Doubling: false,
	}
}

func (p Parser) ParseInput(blocks [][]string) common.Universe {
	objects, robot := p.parseObjects(blocks[0])

	var scoreDivider int = 1
	if p.Doubling {
		scoreDivider = 2
	}

	return common.Universe{
		Robot:        robot,
		Objects:      objects,
		Moves:        p.parseMoves(blocks[1]),
		ScoreDivider: scoreDivider,
	}
}
