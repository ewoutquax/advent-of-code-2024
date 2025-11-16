package day23lanparty

import (
	"fmt"
	"slices"
	"strings"

	"github.com/ewoutquax/advent-of-code-2024/pkg/register"
	"github.com/ewoutquax/advent-of-code-2024/pkg/utils"
)

type (
	Node       string
	Clique     []Node
	Edge       [2]Node
	Neighbours map[Node][]Node
)

const (
	Day string = "23"
)

func (c Clique) containsComputerT() bool {
	for _, node := range c {
		if string(node[0]) == "t" {
			return true
		}
	}

	return false
}

func (c Clique) deepClone() Clique {
	var out Clique = make(Clique, len(c))

	copy(out, c)

	return out
}

func (c Clique) toS() string {
	var names = make([]string, len(c))

	for idx, name := range c {
		names[idx] = string(name)
	}

	return strings.Join(names, ",")
}

func CountApplicableTriangles(nodes []Node, edges []Edge) int {
	var count int = 0
	var uniqueApplicableCliques = make(map[string]struct{})

	cliques := FindCliques(nodes, edges)
	for _, clique := range cliques {
		sizedCliques := SplitLargeClique(clique)
		for _, sizedClique := range sizedCliques {
			if len(sizedClique) == 3 && sizedClique.containsComputerT() {
				if _, ok := uniqueApplicableCliques[sizedClique.toS()]; !ok {
					uniqueApplicableCliques[sizedClique.toS()] = struct{}{}
					count++
				}
			}
		}
	}

	return count
}

func FindPassword(nodes []Node, edges []Edge) string {
	cliques := FindCliques(nodes, edges)

	var maxLength int = 0
	for _, c := range cliques {
		maxLength = max(maxLength, len(c))
	}

	var largestClique Clique
	for _, c := range cliques {
		if len(c) == maxLength {
			largestClique = c
		}
	}

	slices.Sort(largestClique)

	return largestClique.toS()
}

func FindCliques(nodes []Node, edges []Edge) []Clique {
	neighbours := BuildIndexedNeighbours(nodes, edges)

	return neighbours.BornKerbosch([]Node{}, nodes, []Node{})
}

func SplitLargeClique(largeClique Clique) []Clique {
	var combinations []Clique = make([]Clique, 0)
	n := len(largeClique)
	total := 1<<n - 1

	for mask := 1; mask <= total; mask++ {
		var subClique Clique = Clique{}
		for i := range n {
			if (mask>>i)&1 == 1 {
				subClique = append(subClique, largeClique[i])
			}
		}
		combinations = append(combinations, subClique)
	}

	return combinations
}

/**
* R => Current clique
* P => All vertices, optionally the neighbours of the current node
* X => Excluded vertices, to prevent duplicates
 */
func (n Neighbours) BornKerbosch(R Clique, P, X []Node) []Clique {
	// fmt.Println("BornKerbosch")
	// fmt.Printf("R: %v\nP: %v\nX: %v\n", R, P, X)

	if len(P) == 0 && len(X) == 0 {
		// fmt.Printf("BornKerbosch: returning clique: %v\n", R)
		return []Clique{R}
	}

	foundCliques := make([]Clique, 0)
	localP := deepClone(P)
	for len(localP) > 0 {
		currentNode := localP[0]

		subR := append(R, currentNode)
		subP := Intersection(localP, n[currentNode])
		subX := Intersection(X, n[currentNode])

		subCliques := n.BornKerbosch(subR, subP, subX)
		foundCliques = append(foundCliques, deepCloneCliques(subCliques)...)

		localP = localP[1:]
		X = append(X, currentNode)
	}

	return foundCliques
}

func Intersection(lefts, right []Node) []Node {
	var intersected []Node = make([]Node, 0, min(len(lefts), len(right)))

	for _, left := range lefts {
		if slices.Contains(right, left) {
			intersected = append(intersected, left)
		}
	}

	return intersected
}

func BuildIndexedNeighbours(nodes []Node, edges []Edge) Neighbours {
	var indexedNodes Neighbours = make(Neighbours, len(nodes))

	for _, currentNode := range nodes {
		neighbours := make([]Node, 0, len(nodes))

		for _, nodes := range edges {
			if currentNode == nodes[0] {
				neighbours = append(neighbours, nodes[1])
			}
			if currentNode == nodes[1] {
				neighbours = append(neighbours, nodes[0])
			}
		}

		indexedNodes[currentNode] = neighbours
	}

	return indexedNodes
}

func ParseInput(lines []string) ([]Node, []Edge) {
	var indexedNodes map[Node]struct{} = make(map[Node]struct{}, len(lines)*2)
	var edges []Edge = make([]Edge, 0, len(lines))

	for _, line := range lines {
		parts := strings.Split(line, "-")

		fromNode := Node(parts[0])
		toNode := Node(parts[1])

		indexedNodes[fromNode] = struct{}{}
		indexedNodes[toNode] = struct{}{}

		edges = append(edges, Edge{fromNode, toNode})
	}

	var nodes []Node = make([]Node, 0, len(indexedNodes))
	for node := range indexedNodes {
		nodes = append(nodes, node)
	}

	return nodes, edges
}

func deepClone(P []Node) []Node {
	var out []Node = make([]Node, len(P))

	copy(out, P)

	return out
}

func deepCloneCliques(cliques []Clique) []Clique {
	out := make([]Clique, len(cliques))

	for idx, c := range cliques {
		out[idx] = c.deepClone()
	}

	return out
}

func solvePart1(inputFile string) {
	links := utils.ReadFileAsLines(inputFile)
	nodes, edges := ParseInput(links)

	fmt.Printf("Result of day-%s / part-1: %d\n", Day, CountApplicableTriangles(nodes, edges))
}

func solvePart2(inputFile string) {
	links := utils.ReadFileAsLines(inputFile)
	nodes, edges := ParseInput(links)

	fmt.Printf("Result of day-%s / part-2: %s\n", Day, FindPassword(nodes, edges))
}

func init() {
	register.Day(Day+"a", solvePart1)
	register.Day(Day+"b", solvePart2)
}
