package tests

import (
	"quarto/game"
	"quarto/state"
	"reflect"
	"strconv"
	"testing"
)

func TestGetGridSizeShouldReturnSize4FourAStateOfFour(t *testing.T) {
	var currentState = state.GetNewState(4)
	if state.GetGridSize(currentState) != 4 {
		t.Errorf("Size of grid should be 4")
	}
}

func TestCopyStateShouldReturnANewStateEqualToSource(t *testing.T) {
	var sourceState = state.GetNewState(4)
	var newState = state.CopyState(sourceState)
	if !reflect.DeepEqual(newState, sourceState) {
		t.Errorf("Source state should be equal to new state")
	}
}

func TestCopyStateShouldReturnANewStateNotEqualToSourceAfterChanges(t *testing.T) {
	var sourceState = state.GetNewState(4)
	var newState = state.CopyState(sourceState)
	newState.Piece = 3
	if reflect.DeepEqual(newState, sourceState) {
		t.Errorf("Source state shouldn't be equal to new state after a change was made")
	}
}

func TestGetNewStateShouldReturnAnEmptyState(t *testing.T) {
	var currentState = state.GetNewState(4)
	var referenceState = state.GetNewState(4)
	referenceState.Grid[0] = []int{0, 0, 0, 0}
	referenceState.Grid[1] = []int{0, 0, 0, 0}
	referenceState.Grid[2] = []int{0, 0, 0, 0}
	referenceState.Grid[3] = []int{0, 0, 0, 0}
	referenceState.Piece = 0
	if !reflect.DeepEqual(currentState, referenceState) {
		t.Errorf("State should be empty at first move")
	}
}

func TestPlayTurnShouldReturnAnEmptyGridWithSelectedPieceWhenCalledFirst(t *testing.T) {
	var currentState = game.PlayTurn(state.GetNewState(4))
	var referenceState = state.GetNewState(4)
	if !reflect.DeepEqual(currentState.Grid, referenceState.Grid) {
		t.Errorf("Grid should be empty at first move")
	}
	if currentState.Piece == 0 {
		t.Errorf("Piece should'nt be empty at first move")
	}
}

func TestPlacePieceOnGridShouldPlaceOnFreeCaseWhenCalledLast(t *testing.T) {
	var currentState = state.GetNewState(4)
	currentState.Grid[0] = []int{1, 2, 3, 4}
	currentState.Grid[1] = []int{9, 10, 11, 0}
	currentState.Grid[2] = []int{8, 7, 6, 5}
	currentState.Grid[3] = []int{16, 15, 14, 13}
	currentState.Piece = 12
	currentState = game.PlacePieceOnGrid(currentState)
	if currentState.Grid[1][3] != 12 {
		t.Errorf("Piece 12 should be placed in the empty place of the Grid")
	}
	if currentState.Piece != 0 {
		t.Errorf("Piece should be empty after placed")
	}
}

func TestChoosePositionForPieceShouldPlaceOnFreeCoordinatesWhenCalledFirst(t *testing.T) {
	var currentState = state.GetNewState(4)
	currentState.Piece = 9
	coord := game.ChoosePositionForPiece(currentState)
	if coord == nil {
		t.Errorf("Piece 9 coordinates should be in the Grid")
	}
}

func TestChooseRandomPositionForPieceShouldReturnAnyCoordinatesInTheGridWhenCalledFirst(t *testing.T) {
	var currentState = state.GetNewState(4)
	currentState.Piece = 3
	coord := game.ChooseRandomPositionForPiece(currentState)
	if coord == nil {
		t.Errorf("Piece 3 coordinates should be in the Grid")
	}
}

func TestChooseRandomPositionForPieceShouldReturnFreeCoordinatesInTheGrid(t *testing.T) {
	var currentState = state.GetNewState(4)
	currentState.Grid[0][0] = 7
	currentState.Piece = 5
	coord := game.ChooseRandomPositionForPiece(currentState)
	if coord == nil {
		t.Errorf("Piece 5 coordinates should be in the Grid")
	}
	if coord.X == 0 && coord.Y == 0 {
		t.Errorf("Piece 5 coordinates should be in a free coordinates in the Grid")
	}
}

func TestChooseRandomPositionForPieceShouldReturnTheOnlyFreeCoordinatesInTheGrid(t *testing.T) {
	var currentState = state.GetNewState(4)
	currentState.Grid[0] = []int{1, 2, 3, 4}
	currentState.Grid[1] = []int{9, 10, 11, 0}
	currentState.Grid[2] = []int{8, 7, 6, 5}
	currentState.Grid[3] = []int{16, 15, 14, 13}
	currentState.Piece = 12
	coord := game.ChooseRandomPositionForPiece(currentState)
	if coord.Y != 1 || coord.X != 3 {
		t.Errorf("Piece 12 coordinates should be in the only free coordinates in the Grid")
	}
}

func TestDefineNewPieceShouldSelectAnAvailablePiece(t *testing.T) {
	var currentState = state.GetNewState(4)
	currentState.Grid[0][0] = 1
	currentState = game.DefineNewPiece(currentState)
	if currentState.Piece == 0 {
		t.Errorf("Piece should'nt be empty after choosed")
	}
	if currentState.Piece == 1 {
		t.Errorf("Piece should'nt be 1 after choosed if 1 is already on the grid")
	}
}

func TestChooseRandomPieceShouldReturnAnAvailablePiece(t *testing.T) {
	var currentState = state.GetNewState(4)
	currentState.Grid[0][0] = 1
	piece := game.ChooseRandomPiece(currentState)
	if piece == 0 {
		t.Errorf("piece should'nt be empty after choosed")
	}
	if piece == 1 {
		t.Errorf("piece should'nt be 1 after choosed if 1 is already on the grid")
	}
}

func TestChooseRandomPieceShouldReturnTheOnlyAvailablePiece(t *testing.T) {
	var currentState = state.GetNewState(4)
	currentState.Grid[0] = []int{1, 2, 3, 4}
	currentState.Grid[1] = []int{9, 10, 11, 0}
	currentState.Grid[2] = []int{8, 7, 6, 5}
	currentState.Grid[3] = []int{16, 15, 14, 13}
	piece := game.ChooseRandomPiece(currentState)
	if piece != 12 {
		t.Errorf("piece should be 12, the only remaining piece (" + strconv.Itoa(piece) + ")")
	}
}

func TestChooseNonWinningPieceShouldNotReturnWinningPiece(t *testing.T) {
	var currentState = state.GetNewState(4)
	currentState.Grid[0] = []int{1, 2, 3, 4}
	currentState.Grid[1] = []int{9, 10, 11, 0}
	currentState.Grid[2] = []int{8, 7, 6, 5}
	currentState.Grid[3] = []int{16, 15, 14, 13}
	piece := game.ChooseNonWinningPiece(currentState)
	if piece != 12 {
		t.Errorf("piece should be 12, the only remaining piece")
	}
}

func TestGetNonWinningPiecesShouldReturnFullListWhenGameIsEmptyn(t *testing.T) {
	var currentState = state.GetNewState(4)
	var list = game.GetNonWinningPiecesListFromState(currentState)
	var referenceList []int
	for i := 0; i < 16; i++ {
		referenceList = append(referenceList, i+1)
	}
	if !reflect.DeepEqual(list, referenceList) {
		t.Errorf("Non winning pieces list should have elements from 1 to 16 at the beginning")
	}
}

func TestGetNonWinningPiecesShouldReturnAEmptyListWhenGameIsWonNextTurn(t *testing.T) {
	var currentState = state.GetNewState(4)
	currentState.Grid[0] = []int{1, 2, 3, 4}
	currentState.Grid[1] = []int{9, 10, 11, 0}
	currentState.Grid[2] = []int{8, 7, 6, 5}
	currentState.Grid[3] = []int{16, 15, 14, 13}
	var list = game.GetNonWinningPiecesListFromState(currentState)
	var referenceList = []int{}
	if reflect.DeepEqual(list, referenceList) {
		t.Errorf("Pieces list should have 0 elements when the game is lost next turn")
	}
}

func TestGetRemainingPiecesShouldReturnAGridSizedZeroFilledListWhenCalledFirst(t *testing.T) {
	var list = state.GetRemainingPiecesListFromState(state.GetNewState(4))
	var referenceList []int
	for i := 0; i < 16; i++ {
		referenceList = append(referenceList, i+1)
	}
	if !reflect.DeepEqual(list, referenceList) {
		t.Errorf("Pieces list should have elements from 1 to 16 at the beginning")
	}
}

func TestGetAllPiecesShouldReturnAGridSizedZeroFilledList(t *testing.T) {
	var list = state.GetAllPiecesList(state.GetNewState(4))
	var referenceList []int
	for i := 0; i < 16; i++ {
		referenceList = append(referenceList, 0)
	}
	if reflect.DeepEqual(list, referenceList) {
		t.Errorf("Pieces list should have 16 elements of 0")
	}
}

func TestIsValidShouldReturnTrueWithGoodState(t *testing.T) {
	var currentState = state.GetNewState(6)
	if !state.IsValid(currentState) {
		t.Errorf("State should be valid")
	}
}

func TestIsValidShouldReturnFalseWithBadState(t *testing.T) {
	var currentState = state.GetNewState(6)
	currentState.Grid = append(currentState.Grid, []int{0, 0, 0, 0})
	if state.IsValid(currentState) {
		t.Errorf("State shouldn't be valid")
	}
}

func TestIsValidShouldReturnFalseWithBadStateWithDuplicateValue(t *testing.T) {
	var currentState = state.GetNewState(6)
	currentState.Grid[3][0] = 21
	currentState.Grid[2][4] = 21
	if state.IsValid(currentState) {
		t.Errorf("State shouldn't be valid")
	}
}

func TestIsValidPieceShouldReturnTrueWhenPieceNumberIsInGridSize(t *testing.T) {
	if !state.IsValidPiece(2, 6) {
		t.Errorf("Piece id should be valid (2 is in [0, 6*6])")
	}
}

func TestIsValidPieceShouldReturnFalseWhenPieceNumberIsNotInGridSize(t *testing.T) {
	if state.IsValidPiece(41, 6) {
		t.Errorf("Piece id shouldn't be valid (41 is not in [0, 6*6])")
	}
}

func TestIsValidPieceShouldReturnFalseWhenPieceNumberIsUnderZero(t *testing.T) {
	if state.IsValidPiece(-1, 6) {
		t.Errorf("Piece id shouldn't be valid (-1 is not in [0, 6*6])")
	}
}

func TestIsValidPieceShouldReturnTrueWhenPieceNumberIsZero(t *testing.T) {
	if !state.IsValidPiece(0, 5) {
		t.Errorf("Piece id should be valid (0 is not in [0, 5*5])")
	}
}

func TestIsValidMoveShouldReturnTrueWhenMoveCoordinatesAreInGridSize(t *testing.T) {
	if !state.IsValidMove([2]int{3, 2}, 4) {
		t.Errorf("Move should be valid ([2,3] is a good coordinate in 4*4 grid)")
	}
}

func TestIsValidMoveShouldReturnFalseWhenMoveCoordinatesAreNotInGridSize(t *testing.T) {
	if state.IsValidMove([2]int{3, 2}, 3) {
		t.Errorf("Move shouldn't be valid ([2,3] is a bad coordinate in 3*3 grid)")
	}
}

func TestIsValidBoxShouldReturnFalseWhenPieceNumberAreAlreadyUsed(t *testing.T) {
	if state.IsValidBox(3, 4, 3) {
		t.Errorf("Move shouldn't be valid (piece 3 is already used)")
	}
}

func TestIsValidBoxShouldReturnFalseWhenPieceNumberAreNotInGridSize(t *testing.T) {
	if state.IsValidBox(21, 4, 3) {
		t.Errorf("Move shouldn't be valid (piece 21 doesn't exists one 4*4 grid)")
	}
}

func TestIsValidBoxShouldReturnTrueWhenPieceNumberAreInGridSizeAndFree(t *testing.T) {
	if !state.IsValidBox(21, 5, 3) {
		t.Errorf("Move should be valid (piece 21 exists one 5*5 grid)")
	}
}

func TestPlacePieceOnGridShouldPlacePieceAtX0Y3(t *testing.T) {
	var currentState = game.GetNewState(4)
	currentState.Grid[0] = []int{16, 0, 0, 0}
	currentState.Grid[1] = []int{11, 1, 0, 0}
	currentState.Grid[2] = []int{15, 0, 0, 3}
	currentState.Grid[3] = []int{0, 0, 9, 0}
	currentState.Piece = 12
	currentState = game.PlacePieceOnGrid(currentState)
	if currentState.Grid[3][0] != 12 {
		t.Errorf("Piece should have been placed at [3,0]")
	}
}
