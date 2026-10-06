package main

import (
	"fmt"
	"tic-tac-toe/internal/game"
	"tic-tac-toe/internal/player"
)

func main() {
	var name string
	fmt.Println("Welcome to the Tic-Tac-Toe game")
	fmt.Println("Enter your name: ")
	fmt.Scan(&name)
	p1 := player.NewHumanPlayer(name, "X")
	p2 := player.NewBotPlayer("Wall-E", "O")
	g := game.NewGame(p1, p2)
	g.Start()
}
