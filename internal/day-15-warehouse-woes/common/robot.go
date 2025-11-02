package day15warehousewoes_common

import "image"

func (r Robot) CanMoveInDirection(d Direction, u Universe) bool {
	newLocation := r.Point.Add(toVector(d))

	if object, exists := u.Objects[newLocation]; exists {
		return object.CanMoveInDirection(d, u)
	}

	return true
}

func (r *Robot) Move(d Direction, u Universe) {
	newLocation := r.Point.Add(toVector(d))
	if object, ok := u.Objects[newLocation]; ok {
		object.Move(d, u)
	}
	r.Point = r.Point.Add(toVector(d))
}

func toVector(d Direction) image.Point {
	return map[Direction]image.Point{
		DirectionUp:    image.Pt(0, -1),
		DirectionRight: image.Pt(1, 0),
		DirectionDown:  image.Pt(0, 1),
		DirectionLeft:  image.Pt(-1, 0),
	}[d]
}
