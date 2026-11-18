package game

func NextCellState(cell Cell, aliveNeighbours int) bool {
	if isUnderpopulated(cell.Alive, aliveNeighbours) {
		return false
	}

	if canSurvive(cell.Alive, aliveNeighbours) {
		return true
	}

	if isOverpopulated(cell.Alive, aliveNeighbours) {
		return false
	}

	if canReproduce(cell.Alive, aliveNeighbours) {
		return true
	}

	return false
}

func isUnderpopulated(alive bool, aliveNeighbours int) bool {
	return alive && aliveNeighbours < 2
}

func canSurvive(alive bool, aliveNeighbours int) bool {
	return alive && (aliveNeighbours == 2 || aliveNeighbours == 3)
}

func isOverpopulated(alive bool, aliveNeighbours int) bool {
	return alive && aliveNeighbours > 3
}

func canReproduce(alive bool, aliveNeighbours int) bool {
	return !alive && aliveNeighbours == 3
}
