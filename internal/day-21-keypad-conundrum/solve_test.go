package day21keypadconundrum_test

import (
	"image"
	"testing"

	common "github.com/ewoutquax/advent-of-code-2024/internal/day-21-keypad-conundrum/common"
	builder "github.com/ewoutquax/advent-of-code-2024/internal/day-21-keypad-conundrum/services/builder"
	calculator "github.com/ewoutquax/advent-of-code-2024/internal/day-21-keypad-conundrum/services/calculator"
	counter "github.com/ewoutquax/advent-of-code-2024/internal/day-21-keypad-conundrum/services/counter"
	finder "github.com/ewoutquax/advent-of-code-2024/internal/day-21-keypad-conundrum/services/finder"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type (
	BuildKeypadTestSuite            struct{ suite.Suite }
	FindPathsNumericKeypadTestSuite struct{ suite.Suite }
)

func TestRunSuiteBuildKeypad(t *testing.T) {
	suite.Run(t, new(BuildKeypadTestSuite))
}

func (s *BuildKeypadTestSuite) TestBuildKeypad_Numeric() {
	keypad := builder.BuildKeypad(builder.WithLayout(common.KeypadLayoutNumeric))

	s.Len(keypad.Keys, 11)
	s.True(keypad.Start.Eq(image.Pt(2, 3)), keypad.Start.String())
	s.Equal("7", keypad.Keys[image.Pt(0, 0)])
	s.Equal("3", keypad.Keys[image.Pt(2, 2)])
	s.Equal("A", keypad.Keys[image.Pt(2, 3)])
	s.NotContains(keypad.Keys, image.Pt(0, 3))
}

func (s *BuildKeypadTestSuite) TestBuildKeypad_Directional() {
	keypad := builder.BuildKeypad(builder.WithLayout(common.KeypadLayoutDirectional))

	s.Len(keypad.Keys, 5)
	s.True(keypad.Start.Eq(image.Pt(2, 0)), keypad.Start.String())
	s.Equal("A", keypad.Keys[image.Pt(2, 0)])
	s.Equal("^", keypad.Keys[image.Pt(1, 0)])
	s.Equal(">", keypad.Keys[image.Pt(2, 1)])
	s.NotContains(keypad.Keys, image.Pt(0, 0))
}

func TestRunSuiteFindPathsNumericKeypad(t *testing.T) {
	suite.Run(t, new(FindPathsNumericKeypadTestSuite))
}

func (s *FindPathsNumericKeypadTestSuite) TestNumericKeyPad_MoveFrom7To8() {
	keypad := builder.BuildKeypad(builder.WithLayout(common.KeypadLayoutNumeric))

	var paths []common.Move = finder.FindPaths(keypad, "7", "8")
	s.Len(paths, 1)
	s.Equal(common.Move(">A"), paths[0])
}

func (s *FindPathsNumericKeypadTestSuite) TestNumericKeyPad_MoveFrom7To9() {
	keypad := builder.BuildKeypad(builder.WithLayout(common.KeypadLayoutNumeric))

	var paths []common.Move = finder.FindPaths(keypad, "7", "9")
	s.Len(paths, 1)
	s.Equal(common.Move(">>A"), paths[0])
}

func (s *FindPathsNumericKeypadTestSuite) TestNumericKeyPad_MoveFrom7To6() {
	keypad := builder.BuildKeypad(builder.WithLayout(common.KeypadLayoutNumeric))

	var paths []common.Move = finder.FindPaths(keypad, "7", "6")
	s.Len(paths, 3)
	s.Contains(paths, common.Move(">>vA"))
	s.Contains(paths, common.Move(">v>A"))
	s.Contains(paths, common.Move("v>>A"))
}

func (s *FindPathsNumericKeypadTestSuite) TestNumericKeyPad_MoveFrom7To7() {
	keypad := builder.BuildKeypad(builder.WithLayout(common.KeypadLayoutNumeric))

	var paths []common.Move = finder.FindPaths(keypad, "7", "7")
	s.Equal(common.Move("A"), paths[0])
}

func TestNumericKeypad_FindPathsForCodes(t *testing.T) {
	keypad := builder.BuildKeypad(builder.WithLayout(common.KeypadLayoutNumeric))

	var paths []string = finder.FindShortestPathsForCode(keypad, "029A")

	assert := assert.New(t)
	assert.Len(paths, 3)

	assert.Contains(paths, "<A^A>^^AvvvA")
	assert.Contains(paths, "<A^A^>^AvvvA")
	assert.Contains(paths, "<A^A^^>AvvvA")
}

func TestDirectionalKeypad_FindPathsForCodes(t *testing.T) {
	keypad := builder.BuildKeypad(builder.WithLayout(common.KeypadLayoutDirectional))

	directionalCode := "<A^A>^^AvvvA"
	var paths []string = finder.FindShortestPathsForCode(keypad, directionalCode)

	assert := assert.New(t)
	assert.Contains(paths, "v<<A>>^A<A>AvA<^AA>A<vAAA>^A")
}

func TestPathToSegments(t *testing.T) {
	path := "<A^A>^^AvvvA"

	segments := counter.CodeToSegments(path)

	assert := assert.New(t)
	assert.Len(segments, 4)
	assert.Equal("<A", segments[0])
	assert.Equal("^A", segments[1])
	assert.Equal(">^^A", segments[2])
	assert.Equal("vvvA", segments[3])
}

func TestFindShortestTriplePathForCode(t *testing.T) {
	testCases := map[string]int{
		"029A": 68,
		"980A": 60,
		"179A": 68,
		"456A": 64,
		"379A": 64,
	}

	for inputCode, expectedLength := range testCases {
		actualLength := counter.StepsForLayeredCode(inputCode, 2)
		assert.Equal(t, expectedLength, actualLength)
	}
}

func TestCalculateComplexity(t *testing.T) {
	type Input struct {
		code    string
		nrSteps int
	}

	testCases := map[Input]int{
		{"029A", 68}: 29 * 68,
		{"980A", 60}: 980 * 60,
		{"179A", 68}: 179 * 68,
		{"456A", 64}: 456 * 64,
		{"379A", 64}: 379 * 64,
	}

	for inputs, expectedScore := range testCases {
		actualScore := calculator.CalculateComplexity(inputs.code, inputs.nrSteps)
		assert.Equal(t, expectedScore, actualScore)
	}
}

func TestSumComplexities(t *testing.T) {
	assert.Equal(t, 126384, calculator.SumComplexities(testInput(), 2))

}

func testInput() []string {
	return []string{
		"029A",
		"980A",
		"179A",
		"456A",
		"379A",
	}
}
