package day15warehousewoes_common

import "image"

type (
	Direction uint

	Location = image.Point
	Robot    struct {
		Point image.Point
	}
	Move string

	Object interface {
		CanMoveInDirection(Direction, Universe) bool
		Move(Direction, Universe)
		Score() int
	}
	ObjectBox struct {
		Points []image.Point
	}
	ObjectWall struct {
		Points []image.Point
	}

	Objects  map[Location]Object
	Universe struct {
		Robot        Robot
		Objects      Objects
		Moves        []Move
		ScoreDivider int // For universes with doubling 'on', we count blocks double. So we need to divide the score by 2
	}
)

const (
	DirectionUp Direction = iota + 1
	DirectionRight
	DirectionDown
	DirectionLeft
)
