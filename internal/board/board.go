package board

import "fmt"

type Board struct {
	grid [3][3]string
}

func (b *Board) PlaceMark(row, col int, mark string) {
	if b.grid[row][col] != "" {
		fmt.Printf("Cell (%d, %d) is already occupied.\n", row, col)
	} else {
		b.grid[row][col] = mark
		fmt.Printf("Board: (%d, %d): %s\n", row, col, mark)
	}
}

func (b *Board) Display() {
	fmt.Println("Current Board:")
	fmt.Println("-------------")
	for i := 0; i < 3; i++ {
		fmt.Printf("| ")
		for j := 0; j < 3; j++ {
			cell := b.grid[i][j]
			if cell == "" {
				cell = " "
			}
			fmt.Printf("%s | ", cell)
		}
		fmt.Println()
		fmt.Println("-------------")
	}
}

func (b *Board) CheckWin() bool {
	for i := 0; i < 3; i++ {
		if b.grid[0][i] != "" && b.grid[0][i] == b.grid[1][i] && b.grid[0][i] == b.grid[2][i] {
			return true
		} else if b.grid[i][0] != "" && b.grid[i][0] == b.grid[i][1] && b.grid[i][0] == b.grid[i][2] {
			return true
		} else if b.grid[0][0] != "" && b.grid[0][0] == b.grid[1][1] && b.grid[0][0] == b.grid[2][2] {
			return true
		} else if b.grid[0][2] != "" && b.grid[0][2] == b.grid[1][1] && b.grid[0][2] == b.grid[2][0] {
			return true
		}
	}
	return false
}
