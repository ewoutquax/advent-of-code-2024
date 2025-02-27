package day23lanparty

import (
	"fmt"
	"sort"
	"strings"

	"github.com/ewoutquax/advent-of-code-2024/pkg/register"
	"github.com/ewoutquax/advent-of-code-2024/pkg/utils"
)

const (
	Day string = "23"
)

type (
	ComputerName string
	Triangle     [3]ComputerName
	Computer     struct {
		Name        ComputerName
		Connections []*Computer
	}
	LAN map[ComputerName]*Computer
)

func (t Triangle) ContainPossibleComputer() bool {
	for _, name := range t {
		if string(name[0]) == "t" {
			return true
		}
	}
	return false
}

func FilterPossibleTriangles(triangles []Triangle) []Triangle {
	var possibleTriangles []Triangle = make([]Triangle, 0, len(triangles))

	for _, triangle := range triangles {
		if triangle.ContainPossibleComputer() {
			possibleTriangles = append(possibleTriangles, triangle)
		}
	}

	return possibleTriangles
}

func FindSetsOf3Computers(lan LAN) []Triangle {
	allNames := make([]ComputerName, 0, len(lan))
	for name := range lan {
		allNames = append(allNames, name)
	}
	sort.Slice(allNames, func(i, j int) bool { return allNames[i] < allNames[j] })
	fmt.Printf("allNames (after sort): %v\n", allNames)

	triangles := make([]Triangle, 0)
	for _, name := range allNames {
		for _, node := range lan[name].Connections {
			if name < node.Name {
				common := intersect(lan[name].Connections, lan[node.Name].Connections)
				for _, thirdName := range common {
					if node.Name < thirdName {
						newTriangle := Triangle{name, node.Name, thirdName}
						triangles = append(triangles, newTriangle)
					}
				}
			}
		}
	}

	return triangles
}

func intersect(leftConnections, rightConnections []*Computer) []ComputerName {
	mappedNames := make(map[ComputerName]int, 0)

	for _, computer := range append(leftConnections, rightConnections...) {
		if count, exists := mappedNames[computer.Name]; exists {
			mappedNames[computer.Name] = count + 1
		} else {
			mappedNames[computer.Name] = 1
		}
	}

	out := make([]ComputerName, 0, len(mappedNames))
	for name, count := range mappedNames {
		if count > 1 {
			out = append(out, name)
		}
	}

	return out
}

func ParseInput(lines []string) LAN {
	var computers = make(LAN, len(lines))

	// First, build a list with all computers by name
	for _, line := range lines {
		parts := strings.Split(line, "-")

		computerName1 := ComputerName(parts[0])
		computerName2 := ComputerName(parts[1])

		computers[computerName1] = &Computer{
			Name:        computerName1,
			Connections: make([]*Computer, 0, 2),
		}
		computers[computerName2] = &Computer{
			Name:        computerName2,
			Connections: make([]*Computer, 0, 2),
		}
	}

	// Link the computers, based on the input
	for _, line := range lines {
		parts := strings.Split(line, "-")

		computerName1 := ComputerName(parts[0])
		computerName2 := ComputerName(parts[1])

		computers[computerName1].Connections = append(computers[computerName1].Connections, computers[computerName2])
		computers[computerName2].Connections = append(computers[computerName2].Connections, computers[computerName1])
	}

	// Testing: print the parsed data
	for _, c := range computers {
		conns := make([]string, 0, 10)
		for _, conn := range c.Connections {
			conns = append(
				conns,
				string(conn.Name),
			)
		}
		fmt.Printf("Computer '%s' is connected to: [%s]\n", c.Name, strings.Join(conns, ", "))
	}

	return computers
}

func solvePart1(inputFile string) {
	lines := utils.ReadFileAsLines(inputFile)
	lan := ParseInput(lines)
	triangles := FindSetsOf3Computers(lan)

	fmt.Printf("Result of day-%s / part-1: %d\n", Day, len(FilterPossibleTriangles(triangles)))
}

func solvePart2(inputFile string) {
	_ = inputFile

	var count int = 0
	fmt.Printf("Result of day-%s / part-2: %d\n", Day, count)
}

func init() {
	register.Day(Day+"a", solvePart1)
	// register.Day(Day+"b", solvePart2)
}
