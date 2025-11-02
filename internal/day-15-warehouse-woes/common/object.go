package day15warehousewoes_common

import (
	"image"
	"slices"
)

func (b ObjectBox) CanMoveInDirection(d Direction, u Universe) bool {
	for _, boxLocation := range b.Points {
		newLocation := boxLocation.Add(toVector(d))

		if object, ok := u.Objects[newLocation]; ok && !slices.Contains(b.Points, newLocation) {
			if !object.CanMoveInDirection(d, u) {
				return false
			}
		}
	}

	return true
}

func (b ObjectBox) Move(d Direction, u Universe) {
	// Remove existing location, before creating new
	// To prevent the newly created location being removed directly
	for _, oldLocation := range b.Points {
		delete(u.Objects, oldLocation)
	}

	var newObjectBox = ObjectBox{
		Points: make([]image.Point, 0, 2),
	}

	for _, boxLocation := range b.Points {
		newLocation := boxLocation.Add(toVector(d))
		if object, ok := u.Objects[newLocation]; ok {
			object.Move(d, u)
		}
		newObjectBox.Points = append(newObjectBox.Points, newLocation)
		u.Objects[newLocation] = &newObjectBox
	}
}

func (b ObjectBox) Score() int {
	var minX int = 1<<(32-1) - 1 // MaxInt32 or MaxInt64 depending on intSize.
	for _, boxPoint := range b.Points {
		minX = min(minX, boxPoint.X)
	}

	return b.Points[0].Y*100 + minX
}

func (ObjectWall) CanMoveInDirection(Direction, Universe) bool { return false }
func (ObjectWall) Move(Direction, Universe)                    { panic("Wall can't move") }
func (ObjectWall) Score() int                                  { return 0 }
