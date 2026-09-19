package game

import (
	"tic-tac-toe/internal/board"
	"tic-tac-toe/internal/player"
)

type Game struct {
    Board         board.Board
    Players       [2]player.Player    // two players!
    CurrentPlayer int                 // index: 0 or 1
}