package player

import (
	"fmt"
	"tic-tac-toe/internal/board"
)

type HumanPlayer struct {
	Name string
	Mark string
}

type BotPlayer struct {
	Name string
	Mark string
}

type Player interface {
	GetMove(b *board.Board) (int, int)
	GetName() string
	GetMark() string
}

func NewHumanPlayer(name, mark string) *HumanPlayer {
	return &HumanPlayer{
		Name: name,
		Mark: mark,
	}
}

func (h *HumanPlayer) GetName() string {
	return h.Name
}

func (h *HumanPlayer) GetMark() string {
	return h.Mark
}

func (h *HumanPlayer) GetMove(b *board.Board) (int, int) {
	var row, col int
	for {
		// Prompt the player for input
		fmt.Printf("%s's turn (%s). Enter row and column (0-2): ", h.Name, h.Mark)
		fmt.Scanf("%d %d", &row, &col)
		return row, col
	}
}
