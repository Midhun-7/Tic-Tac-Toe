package main

import "tic-tac-toe/internal/game"

func main() {
	g := game.NewGame("Player 1", "Player 2")
	g.Start()
}
