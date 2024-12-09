package day09diskfragmenter_test

import (
	"testing"

	. "github.com/ewoutquax/advent-of-code-2024/internal/day-09-disk-fragmenter"
	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	assert := assert.New(t)

	disk := ParseInput(testInput())

	assert.IsType(Disk{}, disk)
	assert.Len(disk.Files, 10)
	assert.Equal(0, disk.Files[0].Start)
	assert.Equal(2, disk.Files[0].Length)
	assert.Equal(1, disk.Files[0].End)
	assert.Equal(40, disk.Files[9].Start)
	assert.Equal(2, disk.Files[9].Length)
	assert.Equal(41, disk.Files[9].End)
}

func TestChecksumAfterFragmentations(t *testing.T) {
	disk := ParseInput(testInput())
	checksum := ChecksumAfterFragmentation(disk)
	assert.Equal(t, 1928, checksum)
}

func TestChecksumAfterRearranging(t *testing.T) {
	disk := ParseInput(testInput())
	checksum := ChecksumAfterRearranging(disk)
	assert.Equal(t, 2858, checksum)
}

func testInput() string {
	return "2333133121414131402"
}
