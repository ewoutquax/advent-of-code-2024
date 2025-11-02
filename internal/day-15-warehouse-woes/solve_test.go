package day15warehousewoes_test

import (
	"fmt"
	"image"
	"testing"

	common "github.com/ewoutquax/advent-of-code-2024/internal/day-15-warehouse-woes/common"
	services "github.com/ewoutquax/advent-of-code-2024/internal/day-15-warehouse-woes/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestParseInput(t *testing.T) {
	universe := services.BuildParser().ParseInput(testInput())

	assert := assert.New(t)
	assert.IsType(common.Universe{}, universe)
	assert.True(universe.Robot.Point.Eq(image.Pt(4, 4)))
	assert.Equal(1, universe.ScoreDivider)

	// Split objects into boxes and walls; just for testing purposes
	var countBoxes, countWalls = 0, 0
	for _, object := range universe.Objects {
		switch object.(type) {
		case *common.ObjectBox:
			countBoxes++
		case *common.ObjectWall:
			countWalls++
		default:
			panic("No valid case found")
		}
	}
	assert.Equal(21, countBoxes)
	assert.Equal(37, countWalls)

	assert.Contains(universe.Objects, image.Pt(3, 1))
	assert.Contains(universe.Objects, image.Pt(8, 7))
	assert.Contains(universe.Objects, image.Pt(5, 8))

	assert.Len(universe.Moves, 700)
	assert.Equal(common.Move("<"), universe.Moves[0])
	assert.Equal(common.Move("^"), universe.Moves[len(universe.Moves)-1])
}

func TestParseInputDoubled(t *testing.T) {
	universe := services.BuildParser(services.WithDoubling()).ParseInput(testInput())

	assert := assert.New(t)
	assert.IsType(common.Universe{}, universe)
	assert.True(universe.Robot.Point.Eq(image.Pt(8, 4)))
	assert.Equal(2, universe.ScoreDivider)

	// Split objects into boxes and walls; just for testing purposes
	var countBoxes, countWalls = 0, 0
	for _, object := range universe.Objects {
		switch v := object.(type) {
		case *common.ObjectBox:
			countBoxes++
		case *common.ObjectWall:
			countWalls++
		default:
			panic(fmt.Sprintf("No valid case found: %T", v))
		}
	}
	assert.Equal(42, countBoxes)
	assert.Equal(74, countWalls)

	assert.Contains(universe.Objects, image.Pt(6, 1))
	assert.Contains(universe.Objects, image.Pt(7, 1))
	assert.Contains(universe.Objects, image.Pt(16, 7))
	assert.Contains(universe.Objects, image.Pt(17, 7))
	assert.Contains(universe.Objects, image.Pt(10, 8))
	assert.Contains(universe.Objects, image.Pt(11, 8))

	assert.Same(universe.Objects[image.Pt(6, 1)], universe.Objects[image.Pt(7, 1)])
	assert.Same(universe.Objects[image.Pt(16, 7)], universe.Objects[image.Pt(17, 7)])
	assert.Same(universe.Objects[image.Pt(10, 8)], universe.Objects[image.Pt(11, 8)])

	assert.Len(universe.Moves, 700)
	assert.Equal(common.Move("<"), universe.Moves[0])
	assert.Equal(common.Move("^"), universe.Moves[len(universe.Moves)-1])
}

type RobotCanPushTestSuite struct{ suite.Suite }

func TestRunSuiteRobotCanPush(t *testing.T) {
	suite.Run(t, new(RobotCanPushTestSuite))
}

func (s *RobotCanPushTestSuite) TestRobotCanMove_IsFree_ReturnsTrue() {
	universe := services.BuildParser().ParseInput(testInput())

	robot := common.Robot{Point: image.Pt(2, 2)}
	s.True(robot.CanMoveInDirection(common.DirectionUp, universe))
}

func (s *RobotCanPushTestSuite) TestRobotCanMove_PushesBlockedBox_ReturnsFalse() {
	universe := services.BuildParser().ParseInput(testInput())

	robot := common.Robot{Point: image.Pt(3, 2)}
	s.False(robot.CanMoveInDirection(common.DirectionUp, universe))
}

func (s *RobotCanPushTestSuite) TestRobotCanMove_PushesFreeBox_ReturnsTrue() {
	universe := services.BuildParser().ParseInput(testInput())

	robot := common.Robot{Point: image.Pt(2, 1)}
	s.True(robot.CanMoveInDirection(common.DirectionRight, universe))
}

func (s *RobotCanPushTestSuite) TestRobotCanMove_PushesBlockedBoxes_ReturnsFalse() {
	universe := services.BuildParser().ParseInput(testInput())

	robot := common.Robot{Point: image.Pt(6, 7)}
	s.False(robot.CanMoveInDirection(common.DirectionRight, universe))
}

func (s *RobotCanPushTestSuite) TestRobotCanMove_PushesFreeBoxes_ReturnsTrue() {
	universe := services.BuildParser().ParseInput(testInput())

	robot := common.Robot{Point: image.Pt(4, 7)}
	s.True(robot.CanMoveInDirection(common.DirectionLeft, universe))
}

type RobotMoveTestSuite struct{ suite.Suite }

func TestRunSuiteRobotMove(t *testing.T) {
	suite.Run(t, new(RobotMoveTestSuite))
}

func (s *RobotMoveTestSuite) TestRobotMove_IsFree_Succeeds() {
	universe := services.BuildParser().ParseInput(testInput())

	robot := common.Robot{Point: image.Pt(2, 2)}
	robot.Move(common.DirectionUp, universe)

	s.True(robot.Point.Eq(image.Pt(2, 1)))
}

func (s *RobotMoveTestSuite) TestRobotMove_PushesBox_Success() {
	universe := services.BuildParser().ParseInput(testInput())

	robot := common.Robot{Point: image.Pt(5, 4)}
	robot.Move(common.DirectionDown, universe)

	// Robot is moved to new location
	s.True(robot.Point.Eq(image.Pt(5, 5)), robot.Point.String())

	// Box is deleted from starting location
	s.NotContains(universe.Objects, image.Pt(5, 5))

	// Box is created at new location
	newBox, newBoxExists := universe.Objects[image.Pt(5, 6)]
	s.True(newBoxExists)
	s.True(newBox.(*common.ObjectBox).Points[0].Eq(image.Pt(5, 6)))
}

func TestCalculateScoreSmall(t *testing.T) {
	universe := services.BuildParser().ParseInput(testInputSmall())

	services.BuildSolver().FollowMoves(universe)

	assert.Equal(t, 2028, universe.Score())
}

func TestCalculateScore(t *testing.T) {
	universe := services.BuildParser().ParseInput(testInput())

	services.BuildSolver().FollowMoves(universe)

	assert.Equal(t, 10092, universe.Score())
}

func TestCalculateScore_WithDoubling(t *testing.T) {
	universe := services.BuildParser(
		services.WithDoubling(),
	).ParseInput(testInput())

	services.BuildSolver().FollowMoves(universe)

	assert.Equal(t, 9021, universe.Score())
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
			"<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^",
			"vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v",
			"><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<",
			"<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^",
			"^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><",
			"^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^",
			">^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^",
			"<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>",
			"^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>",
			"v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^",
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
