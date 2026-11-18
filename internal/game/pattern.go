package game

func Blinker(grid *Grid) {
	center := grid.Size / 2

	if grid.Size < 3 {
		return
	}

	grid.Cells[center][center-1].Alive = true
	grid.Cells[center][center].Alive = true
	grid.Cells[center][center+1].Alive = true
}
