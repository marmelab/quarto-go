package tests

import (
	"quarto/game"
	"reflect"
	"testing"
)

func TestGetGridSizeShouldReturnSize4FourAStateOfFour(t *testing.T) {
	var state = game.GetNewState(4)
	if game.GetGridSize(state) != 4 {
		t.Errorf("Size of grid should be 4")
	}
}

func TestCopyStateShouldReturnANewStateEqualToSource(t *testing.T) {
	var sourceState = game.GetNewState(4)
	var newState = game.CopyState(sourceState)
	if !reflect.DeepEqual(newState, sourceState) {
		t.Errorf("Source state should be equal to new state")
	}
}

func TestCopyStateShouldReturnANewStateNotEqualToSourceAfterChanges(t *testing.T) {
	var sourceState = game.GetNewState(4)
	var newState = game.CopyState(sourceState)
	newState.Piece = 3
	if reflect.DeepEqual(newState, sourceState) {
		t.Errorf("Source state shouldn't be equal to new state after a change was made")
	}
}

func TestGetNewStateShouldReturnAnEmptyState(t *testing.T) {
	var state = game.GetNewState(4)
	var referenceState = game.GetNewState(4)
	referenceState.Grid[0] = []int{0,0,0,0}
	referenceState.Grid[1] = []int{0,0,0,0}
	referenceState.Grid[2] = []int{0,0,0,0}
	referenceState.Grid[3] = []int{0,0,0,0}
	referenceState.Piece = 0
	if !reflect.DeepEqual(state, referenceState) {
		t.Errorf("Grid should be empty at first move")
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

func TestIsValidShouldReturnTrueWithGoodState(t *testing.T) {
	var state = game.GetNewState(6)
	if !game.IsValid(state) {
		t.Errorf("State should be valid")
	}
}

func TestIsValidShouldReturnFalseWithBadState(t *testing.T) {
	var state = game.GetNewState(6)
	state.Grid = append(state.Grid, []int{0,0,0,0})
	if game.IsValid(state) {
		t.Errorf("State shouldn't be valid")
	}
}