package board

import "fmt"

type Board struct {
    grid  [3][3]string
}

func (b *Board) PlaceMark(row, col int, mark string) {
	if b.grid[row][col] != "" {
		fmt.Printf("Cell (%d, %d) is already occupied.\n", row, col)
	} else {
		b.grid[row][col] = mark
	}
	fmt.Printf("Board: (%d, %d): %s\n", row, col, mark)
}