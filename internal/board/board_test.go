package board

import (
	"testing"
)

func TestPlaceMark_Occupied(t *testing.T) {
	// 1. Create a blank board
	b := Board{}

	// 2. Claim the cell for player X
	b.PlaceMark(0, 0, "X")

	// 3. Try to steal the cell for player O, and save the returned error
	err := b.PlaceMark(0, 0, "X")

	// 4. Assert: We EXPECTED an error. If err is nil, something is broken!
	if err == nil {
		t.Errorf("Expected an error because the cell was taken, but got nil!")
	}

	// 5. Assert: Check the grid to make sure "O" didn't overwrite "X"
	if b.grid[0][0] == "O" {
		t.Errorf("The cell was overwritten! It should still be X.")
	}
}

func TestCheckWin(t *testing.T) {
	b := Board{}

	// 1. Assert that a brand new board does NOT have a win.
	// If b.CheckWin() returns true here, throw an error!
	if b.CheckWin() == true {
		t.Errorf("Expected a brand new board to NOT have a win, but CheckWin returned true")
	}

	// 2. Place three "X"s in a row (e.g. at [0,0], [0,1], [0,2])
	for i := 0; i < 3; i++ {
		b.PlaceMark(0, i, "X")
	}

	// 3. Assert that the board NOW has a win.
	// If b.CheckWin() returns false here, throw an error!
	if b.CheckWin() == false {
		t.Errorf("Expected a board with 3 Xs in a row to have a win, but CheckWin returned false")
	}
}
