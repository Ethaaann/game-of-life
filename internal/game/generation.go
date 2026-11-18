package game

const (
	MaxGeneration = 10
)

var CurrGeneration = 0

func NextGeneration(grid *Grid) {
	nextGrid := NewGrid(grid.Size)

	for x := 0; x < grid.Size; x++ {
		for y := 0; y < grid.Size; y++ {
			cell := grid.Cells[x][y]
			aliveNeighbours := grid.CountCellNeighbours(x, y)

			nextGrid.Cells[x][y] = Cell{
				Alive: NextCellState(cell, aliveNeighbours),
				Position: Position{
					X: x,
					Y: y,
				},
			}
		}
	}

	grid.Cells = nextGrid.Cells
	CurrGeneration++
}
