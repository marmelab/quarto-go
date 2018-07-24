package tests

import (
	"quarto/ai"
	"quarto/state"
	"quarto/grid"
	"reflect"
	"strconv"
	"testing"
)

func TestChooseRandomPositionForPieceShouldReturnAnyCoordinatesInTheGridWhenCalledFirst(t *testing.T) {
	var currentState = state.GetNewState(4)
	currentState.Piece = 3
	coord := ai.ChooseRandomPositionForPiece(currentState, []grid.Point{})
	if coord == nil {
		t.Errorf("Piece 3 coordinates should be in the Grid")
	}
}

func TestChooseRandomPositionForPieceShouldReturnFreeCoordinatesInTheGrid(t *testing.T) {
	var currentState = state.GetNewState(4)
	currentState.Grid[0][0] = 7
	currentState.Piece = 5
	coord := ai.ChooseRandomPositionForPiece(currentState, []grid.Point{})
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
	coord := ai.ChooseRandomPositionForPiece(currentState, []grid.Point{})
	if coord.Y != 1 || coord.X != 3 {
		t.Errorf("Piece 12 coordinates should be in the only free coordinates in the Grid")
	}
}

func TestChooseRandomPieceShouldReturnAnAvailablePiece(t *testing.T) {
	var currentState = state.GetNewState(4)
	currentState.Grid[0][0] = 1
	piece := ai.ChooseRandomPiece(currentState)
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
	piece := ai.ChooseRandomPiece(currentState)
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
	piece := ai.ChooseNonWinningPiece(currentState)
	if piece != 0 {
		t.Errorf("piece should be O, because the only remaining piece is winning (" + strconv.Itoa(piece) + ")")
	}
}

func TestGetNonWinningPiecesShouldReturnFullListWhenGameIsEmptyn(t *testing.T) {
	var currentState = state.GetNewState(4)
	var list = ai.GetNonWinningPiecesListFromState(currentState)
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
	var list = ai.GetNonWinningPiecesListFromState(currentState)
	if len(list)!= 0 {
		t.Errorf("Pieces list should have 0 elements when the game is lost next turn (" + strconv.Itoa(len(list)) + ")")
	}
}

func TestGetNonWinningPiecesShouldReturnNoPiece(t *testing.T) {
	var currentState = state.GetNewState(4)
	currentState.Grid[0] = []int{11, 8, 12, 0}
	currentState.Grid[1] = []int{6, 0, 1, 13}
	currentState.Grid[2] = []int{9, 16, 0, 2}
	currentState.Grid[3] = []int{15, 10, 3, 0}
	var list = ai.GetNonWinningPiecesListFromState(currentState)
	if len(list) != 0 {
		t.Errorf("Selection of piece should return no piece , all are winning (" + strconv.Itoa(len(list)) + ")")
	}
}


func TestGetLoosingBoxListShouldReturnX2Y2(t *testing.T) {
	var currentState = state.GetNewState(4)
	currentState.Grid[0] = []int{11, 8, 12, 0}
	currentState.Grid[1] = []int{6, 0, 1, 13}
	currentState.Grid[2] = []int{9, 16, 0, 2}
	currentState.Grid[3] = []int{15, 10, 3, 0}

	var listResult = []grid.Point{}
	listResult = append(listResult, grid.Point{3, 0})
	listResult = append(listResult, grid.Point{1, 1})
	listResult = append(listResult, grid.Point{2, 2})
	listResult = append(listResult, grid.Point{3, 3})

	var testList = ai.GetLoosingBoxList(currentState)

	if !reflect.DeepEqual(testList, listResult) {
		t.Errorf("Loosing box isn't well founded (" + strconv.Itoa(len(testList)) + ")")
	}
}
