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
	referenceState.Grid[0] = []int{0, 0, 0, 0}
	referenceState.Grid[1] = []int{0, 0, 0, 0}
	referenceState.Grid[2] = []int{0, 0, 0, 0}
	referenceState.Grid[3] = []int{0, 0, 0, 0}
	referenceState.Piece = 0
	if !reflect.DeepEqual(state, referenceState) {
		t.Errorf("State should be empty at first move")
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

func TestPlacePieceOnGridShouldPlaceOnFreeCaseWhenCalledLast(t *testing.T) {
	var state = game.GetNewState(4)
	state.Grid[0] = []int{1, 2, 3, 4}
	state.Grid[1] = []int{9, 10, 11, 0}
	state.Grid[2] = []int{8, 7, 6, 5}
	state.Grid[3] = []int{16, 15, 14, 13}
	state.Piece = 12
	state = game.PlacePieceOnGrid(state)
	if state.Grid[1][3] != 12 {
		t.Errorf("Piece 12 should be placed in the empty place of the Grid")
	}
	if state.Piece != 0 {
		t.Errorf("Piece should be empty after placed")
	}
}

func TestChoosePositionForPieceShouldPlaceOnFreeCoordinatesWhenCalledFirst(t *testing.T) {
	var state = game.GetNewState(4)
	state.Piece = 9
	coord := game.ChoosePositionForPiece(state)
	if coord.X == -1 || coord.Y == -1 {
		t.Errorf("Piece 9 coordinates should be in the Grid")
	}
}

func TestChooseRandomPositionForPieceShouldReturnAnyCoordinatesInTheGridWhenCalledFirst(t *testing.T) {
	var state = game.GetNewState(4)
	state.Piece = 3
	coord := game.ChooseRandomPositionForPiece(state)
	if coord.X == -1 || coord.Y == -1 {
		t.Errorf("Piece 3 coordinates should be in the Grid")
	}
}

func TestChooseRandomPositionForPieceShouldReturnFreeCoordinatesInTheGrid(t *testing.T) {
	var state = game.GetNewState(4)
	state.Grid[0][0] = 7
	state.Piece = 5
	coord := game.ChooseRandomPositionForPiece(state)
	if coord.X == -1 || coord.Y == -1 {
		t.Errorf("Piece 5 coordinates should be in the Grid")
	}
	if coord.X == 0 && coord.Y == 0 {
		t.Errorf("Piece 5 coordinates should be in a free coordinates in the Grid")
	}
}

func TestChooseRandomPositionForPieceShouldReturnTheOnlyFreeCoordinatesInTheGrid(t *testing.T) {
	var state = game.GetNewState(4)
	state.Grid[0] = []int{1, 2, 3, 4}
	state.Grid[1] = []int{9, 10, 11, 0}
	state.Grid[2] = []int{8, 7, 6, 5}
	state.Grid[3] = []int{16, 15, 14, 13}
	state.Piece = 12
	coord := game.ChooseRandomPositionForPiece(state)
	if coord.Y != 1 || coord.X != 3 {
		t.Errorf("Piece 12 coordinates should be in the only free coordinates in the Grid")
	}
}

func TestDefineNewPieceShouldSelectAnAvailablePiece(t *testing.T) {
	var state = game.GetNewState(4)
	state.Grid[0][0] = 1
	state = game.DefineNewPiece(state)
	if state.Piece == 0 {
		t.Errorf("Piece should'nt be empty after choosed")
	}
	if state.Piece == 1 {
		t.Errorf("Piece should'nt be 1 after choosed if 1 is already on the grid")
	}
}

func TestChooseRandomPieceShouldReturnAnAvailablePiece(t *testing.T) {
	var state = game.GetNewState(4)
	state.Grid[0][0] = 1
	piece := game.ChooseRandomPiece(state)
	if piece == 0 {
		t.Errorf("piece should'nt be empty after choosed")
	}
	if piece == 1 {
		t.Errorf("piece should'nt be 1 after choosed if 1 is already on the grid")
	}
}

func TestChooseRandomPieceShouldSelectAnAvailablePiece(t *testing.T) {
	var state = game.GetNewState(4)
	state.Grid[0] = []int{1, 2, 3, 4}
	state.Grid[1] = []int{9, 10, 11, 0}
	state.Grid[2] = []int{8, 7, 6, 5}
	state.Grid[3] = []int{16, 15, 14, 13}
	piece := game.ChooseRandomPiece(state)
	if piece != 12 {
		t.Errorf("piece should'nt be 12, the only remaining piece")
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
	state.Grid = append(state.Grid, []int{0, 0, 0, 0})
	if game.IsValid(state) {
		t.Errorf("State shouldn't be valid")
	}
}

func TestIsValidShouldReturnFalseWithBadStateWithDuplicateValue(t *testing.T) {
	var state = game.GetNewState(6)
	state.Grid[3][0] = 21
	state.Grid[2][4] = 21
	if game.IsValid(state) {
		t.Errorf("State shouldn't be valid")
	}
}

func TestIsValidPieceShouldReturnTrueWhenPieceNumberIsInGridSize(t *testing.T) {
	if !game.IsValidPiece(2, 6) {
		t.Errorf("Piece id should be valid (2 is in [0, 6*6])")
	}
}

func TestIsValidPieceShouldReturnFalseWhenPieceNumberIsNotInGridSize(t *testing.T) {
	if game.IsValidPiece(41, 6) {
		t.Errorf("Piece id shouldn't be valid (41 is not in [0, 6*6])")
	}
}

func TestIsValidPieceShouldReturnFalseWhenPieceNumberIsUnderZero(t *testing.T) {
	if game.IsValidPiece(-1, 6) {
		t.Errorf("Piece id shouldn't be valid (-1 is not in [0, 6*6])")
	}
}

func TestIsValidPieceShouldReturnTrueWhenPieceNumberIsZero(t *testing.T) {
	if !game.IsValidPiece(0, 5) {
		t.Errorf("Piece id should be valid (0 is not in [0, 5*5])")
	}
}

func TestIsValidMoveShouldReturnTrueWhenMoveCoordinatesAreInGridSize(t *testing.T) {
	if !game.IsValidMove([2]int{3, 2}, 4) {
		t.Errorf("Move should be valid ([2,3] is a good coordinate in 4*4 grid)")
	}
}

func TestIsValidMoveShouldReturnFalseWhenMoveCoordinatesAreNotInGridSize(t *testing.T) {
	if game.IsValidMove([2]int{3, 2}, 3) {
		t.Errorf("Move shouldn't be valid ([2,3] is a bad coordinate in 3*3 grid)")
	}
}

func TestIsValidBoxShouldReturnFalseWhenPieceNumberAreAlreadyUsed(t *testing.T) {
	if game.IsValidBox(3, 4, 3) {
		t.Errorf("Move shouldn't be valid (piece 3 is already used)")
	}
}

func TestIsValidBoxShouldReturnFalseWhenPieceNumberAreNotInGridSize(t *testing.T) {
	if game.IsValidBox(21, 4, 3) {
		t.Errorf("Move shouldn't be valid (piece 21 doesn't exists one 4*4 grid)")
	}
}

func TestIsValidBoxShouldReturnTrueWhenPieceNumberAreInGridSizeAndFree(t *testing.T) {
	if !game.IsValidBox(21, 5, 3) {
		t.Errorf("Move should be valid (piece 21 exists one 5*5 grid)")
	}
}
