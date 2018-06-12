package tests

import (
	"quarto/game"
	"testing"
	"strconv"
	"reflect"
)

func testDoAMoveShouldReturnAnEmptyGridWithSelectedPieceWhenCalledFirst(t *testing.T) {
	var state = game.DoAMove(game.State{})
	var referenceState = game.State{}
	if (state.Grid != referenceState.Grid) {
		t.Errorf("Grid should be empty at first move")
	}
	if (state.Piece == 0) {
		t.Errorf("Piece should'nt be empty at first move")
	}
}

func testPlacePieceOnGridShouldPlaceOnFirstCaseWhenCalledFirst(t *testing.T) {
	var state = game.State{}
	state.Piece = 3
	state = game.PlacePieceOnGrid(state)
	if (state.Grid[0][0] != 3) {
		t.Errorf("Piece 3 should be placed in init of the Grid")
	}
	if (state.Piece != 0) {
		t.Errorf("Piece should be empty after placed")
	}
}

func testChooseNewPieceShouldSelectAnAvailablePiece(t *testing.T) {
	var state = game.State{}
	state.Grid[0][0] = 1
	state = game.ChooseNewPiece(state)
	if (state.Piece == 0) {
		t.Errorf("Piece should'nt be empty after choosed")
	}
	if (state.Piece == 1) {
		t.Errorf("Piece should'nt be 1 after choosed if 1 is already on the grid")
	}
}

func testInitListOfRemainingPiecesShouldReturnAGridSizedZeroFilledListWhenCalledFirst(t *testing.T) {
	var list = game.InitListOfRemainingPieces(game.State{})
	var referenceList []int
	//game.GridSize * game.GridSize
	for i := 0; i < game.GridSize*game.GridSize; i++ {
		referenceList = append(referenceList, 0)
	}
	if (reflect.DeepEqual(list,referenceList)) {
		t.Errorf("Pieces list should have " + strconv.Itoa(game.GridSize * game.GridSize) + " 0 elements at the beginning")
	}
}

func testInitListOfAllPiecesShouldReturnAGridSizedZeroFilledList(t *testing.T) {
	var list = game.InitListOfAllPieces(game.State{})
	var referenceList [game.GridSize * game.GridSize]int
	if (reflect.DeepEqual(list,referenceList)) {
		t.Errorf("Pieces list should have " + strconv.Itoa(game.GridSize * game.GridSize) + " 0 elements")
	}
}
