package tests

import "quarto/game"
import "testing"
import "strconv"

// TestDoAMove return correct state
func TestDoAMove(t *testing.T) {
	var state = game.DoAMove(game.State{})
	if state.Grid[0][0] != 0 {
		t.Errorf("Grid should be empty at first move")
	}
	if state.Piece == 0 {
		t.Errorf("Piece should'nt be empty at first move")
	}
}

// TestPlacePieceOnGrid return correct state
func TestPlacePieceOnGrid(t *testing.T) {
	var state = game.State{}
	state.Piece = 3
	state = game.PlacePieceOnGrid(state)
	if state.Grid[0][0] != 3 {
		t.Errorf("Piece 3 should be placed in init of the Grid")
	}
	if state.Piece != 0 {
		t.Errorf("Piece should be empty after placed")
	}
}

// TestChooseNewPiece return correct state
func TestChooseNewPiece(t *testing.T) {
	var state = game.State{}
	state.Grid[0][0] = 1
	state = game.ChooseNewPiece(state)
	if state.Piece == 0 {
		t.Errorf("Piece should'nt be empty after choosed")
	}
	if state.Piece == 1 {
		t.Errorf("Piece should'nt be 1 after choosed if 1 is already on the grid")
	}
}

// TestInitListOfRemainingPieces return correct list
func TestInitListOfRemainingPieces(t *testing.T) {
	var list = game.InitListOfRemainingPieces(game.State{})
}

// TestInitListOfAllPieces return correct list
func TestInitListOfAllPieces(t *testing.T) {
	var list = game.InitListOfAllPieces(game.State{})
	if len(list) != game.GridSize * game.GridSize {
		t.Errorf("Pieces list should have " + strconv.Itoa(game.GridSize * game.GridSize) + " elements")
	}
}

func TestContains(t *testing.T) {
	var list = [4]int
	list[3] = 1
	list[2] = 7
	if !game.Contains(list, 7) {
		t.Errorf("Contains should return True")
	}
}

func TestIndexOf(t *testing.T) {
	var list = [4]int
	list[3] = 1
	list[2] = 7
	if game.IndexOf(list, 7) {
		t.Errorf("Contains should return 2")
	}
}
