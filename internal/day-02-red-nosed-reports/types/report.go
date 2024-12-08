package day02rednosedreports_types

type Report struct {
	Levels []Level
}

type Direction string

const (
	DirectionAscending  Direction = "ascending"
	DirectionDescending Direction = "descending"
)

func (r Report) probableDirection() Direction {
	if r.Levels[0] < r.Levels[1] {
		return DirectionAscending
	}
	return DirectionDescending
}

func (r Report) IsSafe(dampeningLevel int) bool {
	d := r.probableDirection()

	for idx := 0; idx < len(r.Levels)-1; idx++ {
		if isLevelsValid(r.Levels[idx], r.Levels[idx+1], d) {
			if dampeningLevel > 0 {
				return isAnySubreportSafe(r, dampeningLevel-1)
			}
			return false
		}
	}
	return true
}

func isAnySubreportSafe(r Report, dampeningLevel int) bool {
	for idxRemove := 0; idxRemove < len(r.Levels); idxRemove++ {
		dampenedLevels := removeLevelsFromSliceByIdx(r.Levels, idxRemove)
		dampenedReport := Report{Levels: dampenedLevels}
		if dampenedReport.IsSafe(dampeningLevel) {
			return true
		}
	}

	return false
}
