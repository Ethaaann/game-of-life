package main

import (
	"fmt"
	"game-of-life/internal/cli"
	"game-of-life/internal/game"
)

func main() {
	fmt.Println("Preparing game...")

	grid := game.NewGrid(5)
	grid.FillGridWithDiedCells()

	// Test pattern
	game.Blinker(grid)

	const maxGenerations = 10
	for generation := 0; generation < maxGenerations; generation++ {
		if grid.CountAliveCells() == 0 {
			break
		}

		game.NextGeneration(grid)
		cli.RenderGrid(grid)
	}
}
