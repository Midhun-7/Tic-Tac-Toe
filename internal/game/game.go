package game

import (
	"fmt"
	"tic-tac-toe/internal/board"
	"tic-tac-toe/internal/player"
)

type Game struct {
	Board         board.Board
	Players       [2]player.Player // two players!
	CurrentPlayer int              // index: 0 or 1
}

func NewGame(player1Name, player2Name string) *Game {
	return &Game{
		Board:         board.Board{},
		Players:       [2]player.Player{*player.NewPlayer(player1Name, "X"), *player.NewPlayer(player2Name, "O")},
		CurrentPlayer: 0,
	}
}

func (g *Game) Start() {
	var row, col int
	for {
		g.Board.Display()
		if g.Board.CheckWin() == true {
			g.CurrentPlayer = (g.CurrentPlayer + 1) % 2
			fmt.Printf("%s wins!\n", g.Players[g.CurrentPlayer].Name)
			break
		} else if g.Board.CheckDraw() == true {
			fmt.Println("It's a draw!")
			break
		}
		currentPlayer := g.Players[g.CurrentPlayer]
		fmt.Printf("%s's turn (%s). Enter row and column (0-2): ", currentPlayer.Name, currentPlayer.Mark)
		fmt.Scanf("%d %d", &row, &col)
		if row < 0 || row > 2 || col < 0 || col > 2 {
			fmt.Println("Invalid input. Please enter row and column between 0 and 2.")
			continue
		}
		if err := g.Board.PlaceMark(row, col, currentPlayer.Mark); err != nil {
			fmt.Println(err)
			continue
		}
		g.CurrentPlayer = (g.CurrentPlayer + 1) % 2
	}
}
