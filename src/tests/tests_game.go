package tests

import "quarto/game"
import "testing"

// TestDoAMove return correct state
func TestDoAMove(t *testing.T) {
	state = game.DoAMove(state)
}

// TestPlacePieceOnGrid return correct state
func TestPlacePieceOnGrid(t *testing.T) {
	state = game.PlacePieceOnGrid(state)
}

// TestChooseNewPiece return correct state
func TestChooseNewPiece(t *testing.T) {
	state = game.ChooseNewPiece(state)
}

// TestInitListOfRemainingPieces return correct list
func TestInitListOfRemainingPieces(t *testing.T) {
	var list = game.InitListOfRemainingPieces(state)
}

// TestInitListOfAllPieces return correct list
func TestInitListOfAllPieces(t *testing.T) {
	var list = game.InitListOfAllPieces(state)
}

func TestContains(t *testing.T) {
	//game.contains([1,4], 1)
}

func TestIndexOf(s []int, e int) int {
	//game.indexOf([1,4], 1)
}
