package day15warehousewoes_test

import (
	. "image"
	"testing"

	. "github.com/ewoutquax/advent-of-code-2024/internal/day-15-warehouse-woes"
	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	assert := assert.New(t)

	universe := ParseInput(testInput())

	assert.IsType(Universe{}, universe)
	assert.Len(universe.Boxes, 21)
	assert.Len(universe.Walls, 37)
	assert.Equal(universe.Robot, Pt(4, 4))
}

func TestMoveRobot(t *testing.T) {
	universe := ParseInput(testInputSmall())

	TakeSteps(&universe)

	// assert.Equal(t, universe.Robot, Pt(1, 1))
	assert.Equal(t, 2028, universe.Hash())
}

func testInput() [][]string {
	return [][]string{
		{
			"##########",
			"#..O..O.O#",
			"#......O.#",
			"#.OO..O.O#",
			"#..O@..O.#",
			"#O#..O...#",
			"#O..O..O.#",
			"#.OO.O.OO#",
			"#....O...#",
			"##########",
		}, {
			"<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^>^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^",
		},
	}
}

func testInputSmall() [][]string {
	return [][]string{
		{
			"########",
			"#..O.O.#",
			"##@.O..#",
			"#...O..#",
			"#.#.O..#",
			"#...O..#",
			"#......#",
			"########",
		}, {
			"<^^>>>vv<v>>v<<",
		},
	}
}
