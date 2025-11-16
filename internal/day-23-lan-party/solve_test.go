package day23lanparty_test

import (
	"testing"

	. "github.com/ewoutquax/advent-of-code-2024/internal/day-23-lan-party"
	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	nodes, edges := ParseInput(testInputSmall())

	assert := assert.New(t)
	assert.Len(nodes, 4)
	assert.Len(edges, 4)

	assert.Contains(nodes, Node("a"))
	assert.Contains(nodes, Node("b"))
	assert.Contains(nodes, Node("c"))
	assert.Contains(nodes, Node("d"))
	assert.Contains(edges, Edge{"a", "b"})
	assert.Contains(edges, Edge{"a", "c"})
	assert.Contains(edges, Edge{"b", "c"})
	assert.Contains(edges, Edge{"b", "d"})
}

func TestBuildIndexedNeighbours(t *testing.T) {
	nodes, edges := ParseInput(testInputSmall())

	indexedNeighours := BuildIndexedNeighbours(nodes, edges)

	assert := assert.New(t)
	assert.Len(indexedNeighours, 4)
	assert.Equal(indexedNeighours["a"], []Node{"b", "c"})
	assert.Equal(indexedNeighours["b"], []Node{"a", "c", "d"})
	assert.Equal(indexedNeighours["c"], []Node{"a", "b"})
	assert.Equal(indexedNeighours["d"], []Node{"b"})
}

func TestFindCliquesSmall(t *testing.T) {
	_, edges := ParseInput(testInputSmall())
	nodes := []Node{"a", "b", "c", "d"} // Declare nodes again, to determine order

	cliques := FindCliques(nodes, edges)

	assert := assert.New(t)
	assert.Len(cliques, 2)
	assert.Contains(cliques, Clique{"a", "b", "c"})
	assert.Contains(cliques, Clique{"b", "d"})
}

func TestCountApplicableTriangles(t *testing.T) {
	nodes, edges := ParseInput(testInput())

	count := CountApplicableTriangles(nodes, edges)

	assert.Equal(t, 7, count)
	assert.True(t, false)
}

func TestSplitLargeCliques(t *testing.T) {
	largeClique := Clique{"co", "de", "ka", "ta"}

	splittedCliques := SplitLargeClique(largeClique)

	assert := assert.New(t)
	assert.Len(splittedCliques, 1+4+6+4)
	assert.Contains(splittedCliques, Clique{"co", "de", "ta"})
	assert.Contains(splittedCliques, Clique{"co", "ka", "ta"})
	assert.Contains(splittedCliques, Clique{"de", "ka", "ta"})
}

func testInputSmall() []string {
	return []string{
		"a-b",
		"a-c",
		"b-c",
		"b-d",
	}
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
