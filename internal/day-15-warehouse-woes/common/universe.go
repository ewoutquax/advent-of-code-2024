package day15warehousewoes_common

func (u Universe) Score() int {
	var total int = 0
	for _, object := range u.Objects {
		total += object.Score()
	}

	return total / u.ScoreDivider
}
