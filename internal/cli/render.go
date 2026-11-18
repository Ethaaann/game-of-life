package cli

import (
	"fmt"
	"game-of-life/internal/game"
)

func RenderGrid(grid *game.Grid) {
	for x := 0; x < grid.Size; x++ {
		for y := 0; y < grid.Size; y++ {
			fmt.Print(grid.Cells[x][y].Display(), " ")
		}
		fmt.Println()
	}
	fmt.Println()
}
