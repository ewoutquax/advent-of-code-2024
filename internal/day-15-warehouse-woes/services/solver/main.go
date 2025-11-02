package day15warehousewoes_services_solver

import common "github.com/ewoutquax/advent-of-code-2024/internal/day-15-warehouse-woes/common"

type Solver struct{}

func DefaultSolver() Solver {
	return Solver{}
}

func (Solver) FollowMoves(u common.Universe) {
	for _, currentMove := range u.Moves {
		currentDirection := toDirection(currentMove)
		if u.Robot.CanMoveInDirection(currentDirection, u) {
			u.Robot.Move(currentDirection, u)
		}
	}
}

func toDirection(m common.Move) common.Direction {
	return map[common.Move]common.Direction{
		"^": common.DirectionUp,
		">": common.DirectionRight,
		"v": common.DirectionDown,
		"<": common.DirectionLeft,
	}[m]
}
