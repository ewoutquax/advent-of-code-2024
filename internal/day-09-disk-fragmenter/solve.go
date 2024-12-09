package day09diskfragmenter

import (
	"container/heap"
	"fmt"
	"strconv"
	"strings"

	"github.com/ewoutquax/advent-of-code-2024/pkg/register"
	"github.com/ewoutquax/advent-of-code-2024/pkg/utils"
)

const (
	Day string = "09"
)

type BlockId int

type Block struct {
	Id     BlockId
	Start  int
	Length int
	End    int
}

type Disk struct {
	MaxFileBlockId BlockId
	MaxFreeBlockId BlockId
	Files          map[BlockId]*Block
}

type BlockHeap []Block
type ReverseBlockHeap []Block

func (b BlockHeap) Len() int           { return len(b) }
func (b BlockHeap) Less(i, j int) bool { return b[i].Start < b[j].Start }
func (b BlockHeap) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b *BlockHeap) Push(x any)        { *b = append(*b, x.(Block)) }
func (h *BlockHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (b ReverseBlockHeap) Len() int           { return len(b) }
func (b ReverseBlockHeap) Less(i, j int) bool { return b[i].Start > b[j].Start }
func (b ReverseBlockHeap) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b *ReverseBlockHeap) Push(x any)        { *b = append(*b, x.(Block)) }
func (h *ReverseBlockHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func ChecksumAfterRearranging(disk Disk) int {
	reverseHeap := make(ReverseBlockHeap, 0, len(disk.Files))
	for _, file := range disk.Files {
		reverseHeap = append(reverseHeap, *file)
	}

	heap.Init(&reverseHeap)

	for len(reverseHeap) > 0 {
		copyFile := heap.Pop(&reverseHeap).(Block)

		forwardHeap := make(BlockHeap, 0, len(disk.Files))
		for _, file := range disk.Files {
			forwardHeap = append(forwardHeap, *file)
		}
		heap.Init(&forwardHeap)

		var freeSpacePointer int = -1
		var freeSpaceLength int = -1

		currentFile := heap.Pop(&forwardHeap).(Block)
		for len(forwardHeap) > 0 && freeSpaceLength < copyFile.Length && freeSpacePointer < copyFile.Start {
			freeSpacePointer = currentFile.End + 1
			currentFile = heap.Pop(&forwardHeap).(Block)
			freeSpaceLength = currentFile.Start - freeSpacePointer
		}

		if freeSpacePointer < copyFile.Start && freeSpaceLength >= copyFile.Length {
			disk.Files[copyFile.Id].Start = freeSpacePointer
			disk.Files[copyFile.Id].End = freeSpacePointer + copyFile.Length - 1
		}
	}

	var checksum int = 0
	for _, file := range disk.Files {
		for idx := 0; idx < file.Length; idx++ {
			checksum += int(file.Id) * (file.Start + idx)
		}

	}

	return checksum
}

func ChecksumAfterFragmentation(disk Disk) int {
	forwardHeap := make(BlockHeap, 0, len(disk.Files))
	reverseHeap := make(ReverseBlockHeap, 0, len(disk.Files))
	for _, file := range disk.Files {
		reverseHeap = append(reverseHeap, *file)
		forwardHeap = append(forwardHeap, *file)
	}

	heap.Init(&forwardHeap)
	heap.Init(&reverseHeap)

	currentFile := heap.Pop(&forwardHeap).(Block)
	copyFromFile := heap.Pop(&reverseHeap).(Block)

	var checksum int = 0
	var copyFromPosition = copyFromFile.End
	for diskPosition := 0; diskPosition <= copyFromPosition; diskPosition++ {
		for diskPosition > currentFile.End {
			currentFile = heap.Pop(&forwardHeap).(Block)
		}
		for copyFromPosition < copyFromFile.Start {
			copyFromFile = heap.Pop(&reverseHeap).(Block)
		}
		if copyFromPosition > copyFromFile.End {
			copyFromPosition = copyFromFile.End
		}

		if diskPosition <= copyFromPosition {
			if diskPosition >= currentFile.Start && diskPosition <= currentFile.End {
				// Use currentFile for checksum
				checksum += int(currentFile.Id) * diskPosition
			} else {
				// Use file from the end for checksum
				checksum += int(copyFromFile.Id) * diskPosition
				copyFromPosition--
			}
		}
	}

	return checksum
}

func ParseInput(line string) Disk {
	var disk = Disk{
		Files: make(map[BlockId]*Block, len(line)/2),
	}

	var diskPosition int = 0
	for idx, char := range strings.Split(line, "") {
		length := convAtoi(char)
		if length > 0 {
			if idx%2 == 0 {
				disk.Files[BlockId(idx/2)] = &Block{
					Id:     BlockId(idx / 2),
					Start:  diskPosition,
					Length: length,
					End:    diskPosition + length - 1,
				}
			}
		}
		diskPosition += length
	}

	return disk
}

func convAtoi(s string) int {
	nr, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return nr
}

func solvePart1(inputFile string) {
	line := utils.ReadFileAsLine(inputFile)
	disk := ParseInput(line)

	fmt.Printf("Result of day-%s / part-1: %d\n", Day, ChecksumAfterFragmentation(disk))
}

func solvePart2(inputFile string) {
	line := utils.ReadFileAsLine(inputFile)
	disk := ParseInput(line)

	fmt.Printf("Result of day-%s / part-2: %d\n", Day, ChecksumAfterRearranging(disk))
}

func init() {
	register.Day(Day+"a", solvePart1)
	register.Day(Day+"b", solvePart2)
}
