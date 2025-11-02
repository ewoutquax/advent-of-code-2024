package day15warehousewoes_services_parser

import (
	"image"
	"strings"

	common "github.com/ewoutquax/advent-of-code-2024/internal/day-15-warehouse-woes/common"
)

func (p Parser) parseObjects(lines []string) (common.Objects, common.Robot) {
	var objects = make(common.Objects)
	var robot common.Robot

	for y, line := range lines {
		for x, char := range strings.Split(line, "") {
			var newObject common.Object

			switch char {
			case "O":
				if p.Doubling {
					newObject = common.ObjectBox{
						Points: []image.Point{
							image.Pt(x*2, y),
							image.Pt(x*2+1, y),
						},
					}
				} else {
					newObject = common.ObjectBox{
						Points: []image.Point{
							image.Pt(x, y),
						},
					}
				}
			case "#":
				if p.Doubling {
					newObject = common.ObjectWall{
						Points: []image.Point{
							image.Pt(x*2, y),
							image.Pt(x*2+1, y),
						},
					}
				} else {
					newObject = common.ObjectWall{
						Points: []image.Point{
							image.Pt(x, y),
						},
					}
				}
			case "@":
				if p.Doubling {
					robot = common.Robot{Point: image.Pt(x*2, y)}
				} else {
					robot = common.Robot{Point: image.Pt(x, y)}
				}
			}

			switch o := newObject.(type) {
			case common.ObjectBox:
				for _, point := range o.Points {
					objects[point] = &o
				}
			case common.ObjectWall:
				for _, point := range o.Points {
					objects[point] = &o
				}
			}
		}
	}

	return objects, robot
}
