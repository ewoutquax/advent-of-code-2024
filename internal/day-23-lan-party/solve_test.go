package day23lanparty_test

import (
	"fmt"
	"testing"

	. "github.com/ewoutquax/advent-of-code-2024/internal/day-23-lan-party"
	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	assert := assert.New(t)

	lan := ParseInput(testInput())

	assert.IsType(LAN{}, lan)
	assert.Len(lan, 16)

	var computerKh *Computer = lan["kh"]

	fmt.Printf("computerKc: %v\n", computerKh)
	assert.Len(computerKh.Connections, 4)
	assert.Equal(ComputerName("tc"), computerKh.Connections[0].Name)
}

func TestFindSetsOf3Computers(t *testing.T) {
	lan := ParseInput(testInput())
	triangles := FindSetsOf3Computers(lan)

	assert.Len(t, triangles, 12)
}

func TestFilterPossibleTriangles(t *testing.T) {
	lan := ParseInput(testInput())
	triangles := FindSetsOf3Computers(lan)

	possibleTriangles := FilterPossibleTriangles(triangles)
	assert.Len(t, possibleTriangles, 7)

}

func testInput() []string {
	return []string{
		"kh-tc",
		"qp-kh",
		"de-cg",
		"ka-co",
		"yn-aq",
		"qp-ub",
		"cg-tb",
		"vc-aq",
		"tb-ka",
		"wh-tc",
		"yn-cg",
		"kh-ub",
		"ta-co",
		"de-co",
		"tc-td",
		"tb-wq",
		"wh-td",
		"ta-ka",
		"td-qp",
		"aq-cg",
		"wq-ub",
		"ub-vc",
		"de-ta",
		"wq-aq",
		"wq-vc",
		"wh-yn",
		"ka-de",
		"kh-ta",
		"co-tc",
		"wh-qp",
		"tb-vc",
		"td-yn",
	}
}
