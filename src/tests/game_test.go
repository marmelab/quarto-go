package tests

import (
	"quarto/game"
	"reflect"
	"testing"
)

func TestGetNewStateShouldReturnAnEmptyState(t *testing.T) {
	var state = game.PlayTurn(game.GetNewState(4))
	var referenceState = game.GetNewState(4)
	if !reflect.DeepEqual(state.Grid, referenceState.Grid) {
		t.Errorf("Grid should be empty at first move")
	}
	if state.Piece == 0 {
		t.Errorf("Piece should'nt be empty at first move")
	}
}

func TestPlayTurnShouldReturnAnEmptyGridWithSelectedPieceWhenCalledFirst(t *testing.T) {
	var state = game.PlayTurn(game.GetNewState(4))
	var referenceState = game.GetNewState(4)
	if !reflect.DeepEqual(state.Grid, referenceState.Grid) {
		t.Errorf("Grid should be empty at first move")
	}
	if state.Piece == 0 {
		t.Errorf("Piece should'nt be empty at first move")
	}
}

func TestPlacePieceOnGridShouldPlaceOnFirstCaseWhenCalledFirst(t *testing.T) {
	var state = game.GetNewState(4)
	state.Piece = 3
	state = game.PlacePieceOnGrid(state)
	if state.Grid[0][0] != 3 {
		t.Errorf("Piece 3 should be placed in init of the Grid")
	}
	if state.Piece != 0 {
		t.Errorf("Piece should be empty after placed")
	}
}

func TestChooseNewPieceShouldSelectAnAvailablePiece(t *testing.T) {
	var state = game.GetNewState(4)
	state.Grid[0][0] = 1
	state = game.ChooseNewPiece(state)
	if state.Piece == 0 {
		t.Errorf("Piece should'nt be empty after choosed")
	}
	if state.Piece == 1 {
		t.Errorf("Piece should'nt be 1 after choosed if 1 is already on the grid")
	}
}

func TestInitListOfRemainingPiecesShouldReturnAGridSizedZeroFilledListWhenCalledFirst(t *testing.T) {
	var list = game.GetRemainingPiecesListFromState(game.State{})
	var referenceList []int
	for i := 0; i < 16; i++ {
		referenceList = append(referenceList, 0)
	}
	if reflect.DeepEqual(list, referenceList) {
		t.Errorf("Pieces list should have 16 elements of 0 at the beginning")
	}
}

func TestInitListOfAllPiecesShouldReturnAGridSizedZeroFilledList(t *testing.T) {
	var list = game.GetAllPiecesList(game.GetNewState(4))
	var referenceList []int
	for i := 0; i < 16; i++ {
		referenceList = append(referenceList, 0)
	}
	if reflect.DeepEqual(list, referenceList) {
		t.Errorf("Pieces list should have 16 elements of 0")
	}
}
