package day02rednosedreports_report

import "strconv"

type Level int

func ConvAtoLevel(s string) Level {
	nr, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return Level(nr)
}

func isLevelsValid(firstLevel, secondLevel Level, d Direction) bool {
	gap := abs(firstLevel - secondLevel)
	return d == DirectionAscending && firstLevel > secondLevel ||
		d == DirectionDescending && firstLevel < secondLevel ||
		gap < 1 || gap > 3
}

func abs(level Level) int {
	if level < 0 {
		return -int(level)
	}
	return int(level)
}

func removeLevelsFromSliceByIdx(levels []Level, idxRemove int) []Level {
	var out []Level = make([]Level, 0, len(levels))

	for idx, nr := range levels {
		if idx != idxRemove {
			out = append(out, nr)
		}
	}

	return out
}
