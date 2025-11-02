package day15warehousewoes_services

import (
	service "github.com/ewoutquax/advent-of-code-2024/internal/day-15-warehouse-woes/services/solver"
)

type FuncSolver func(*service.Solver)

func BuildSolver(opts ...FuncSolver) service.Solver {
	solver := service.DefaultSolver()

	for _, fnOpt := range opts {
		fnOpt(&solver)
	}

	return solver
}
