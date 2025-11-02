package day15warehousewoes_services

import (
	service "github.com/ewoutquax/advent-of-code-2024/internal/day-15-warehouse-woes/services/parser"
)

type FuncParser func(*service.Parser)

func BuildParser(opts ...FuncParser) service.Parser {
	parser := service.DefaultParser()

	for _, opt := range opts {
		opt(&parser)
	}

	return parser
}

func WithDoubling() FuncParser {
	return func(p *service.Parser) {
		p.Doubling = true
	}
}
