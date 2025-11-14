package day21keypadconundrum_common

import "image"

type (
	KeypadLayout uint
	Move         string
	Moves        map[Move]image.Point
	Path         struct {
		CurrentLocation Location
		NrSteps         int
		Moves           Move
	}
	Location = image.Point
	Keypad   struct {
		Start Location
		Keys  map[Location]string
	}
)

const (
	KeypadLayoutNumeric KeypadLayout = iota + 1
	KeypadLayoutDirectional
)
