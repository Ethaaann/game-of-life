package game

import "math/rand"

type Grid struct {
	Cells [][]Cell
	Size  int
}

func NewGrid(size int) *Grid {
	cells := make([][]Cell, size)
	for i := range cells {
		cells[i] = make([]Cell, size)
	}
	return &Grid{
		Cells: cells,
		Size:  size,
	}
}

func (grid *Grid) FillGridWithDiedCells() {
	for x := 0; x < grid.Size; x++ {
		for y := 0; y < grid.Size; y++ {
			grid.Cells[x][y] = Cell{
				Alive: false,
				Position: Position{
					X: x,
					Y: y,
				},
			}
		}
	}
}

func (grid *Grid) AddRandomAliveCells(aliveCells int) {
	added := 0
	for added < aliveCells {
		x := rand.Intn(grid.Size)
		y := rand.Intn(grid.Size)
		if !grid.Cells[x][y].Alive {
			grid.Cells[x][y].Alive = true
			added++
		}
	}
}

func (grid *Grid) CountCellNeighbours(x, y int) int {
	count := 0
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx == 0 && dy == 0 {
				continue
			}

			nx := x + dx
			ny := y + dy
			if nx < 0 || nx >= grid.Size || ny < 0 || ny >= grid.Size {
				continue
			}

			if grid.Cells[nx][ny].Alive {
				count++
			}
		}
	}

	return count
}

func (grid *Grid) CountAliveCells() int {
	aliveCells := 0
	for x := 0; x < grid.Size; x++ {
		for y := 0; y < grid.Size; y++ {
			if grid.Cells[x][y].Alive {
				aliveCells++
			}
		}
	}
	return aliveCells
}

type Cell struct {
	Alive    bool
	Position Position
}

func (cell Cell) Display() string {
	if cell.Alive {
		return "A"
	}
	return "D"
}

type Position struct {
	X int
	Y int
}
