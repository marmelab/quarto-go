package tests

import (
	"quarto/ai"
	"quarto/state"
	"quarto/grid"
	"testing"
)

func TestChooseDefensivePositionForPieceShouldPlaceOnFreeCoordinatesWhenCalledFirstWithoutBadBoxList(t *testing.T) {
	var currentState = state.GetNewState(4)
	currentState.Piece = 9
	coord := ai.ChooseDefensivePositionForPiece(currentState, []grid.Point{})
	emptyPoint := grid.Point{-1, -1}
	if *coord == emptyPoint {
		t.Errorf("Piece 9 coordinates should be in the Grid")
	}
}

func TestChooseDefensivePositionForPieceShouldPlaceOnY3CoordinatesWhenCalledFirstWithBadBoxList(t *testing.T) {
	var currentState = state.GetNewState(4)
	currentState.Piece = 9

	var badBoxList = []grid.Point{}
	badBoxList = append(badBoxList, grid.Point{0, 0})
	badBoxList = append(badBoxList, grid.Point{1, 0})
	badBoxList = append(badBoxList, grid.Point{2, 0})
	badBoxList = append(badBoxList, grid.Point{3, 0})
	badBoxList = append(badBoxList, grid.Point{0, 1})
	badBoxList = append(badBoxList, grid.Point{1, 1})
	badBoxList = append(badBoxList, grid.Point{2, 1})
	badBoxList = append(badBoxList, grid.Point{3, 1})
	badBoxList = append(badBoxList, grid.Point{0, 2})
	badBoxList = append(badBoxList, grid.Point{1, 2})
	badBoxList = append(badBoxList, grid.Point{2, 2})
	badBoxList = append(badBoxList, grid.Point{3, 2})

	coord := ai.ChooseDefensivePositionForPiece(currentState, badBoxList)
	if coord.Y != 3 {
		t.Errorf("Piece 9 coordinates should be in Y=3 coordinates")
	}
}

func TestChooseBlockingPositionForPieceShouldBeNil(t *testing.T) {
	var currentState = state.GetNewState(4)
	currentState.Piece = 9
	coord := ai.ChooseBlockingPositionForPiece(currentState, []grid.Point{})
	emptyPoint := grid.Point{-1, -1}
	if *coord != emptyPoint {
		t.Errorf("Piece 9 coordinates should not be in the Grid")
	}
}

func TestChooseBlockingPositionForPieceShouldPlaceIn(t *testing.T) {
	var currentState = state.GetNewState(4)
	currentState.Grid[0] = []int{8, 12, 0, 15}
	currentState.Grid[1] = []int{11, 1, 0, 0}
	currentState.Grid[2] = []int{0, 0, 0, 3}
	currentState.Grid[3] = []int{0, 0, 2, 0}
	currentState.Piece = 9
	coord := ai.ChooseBlockingPositionForPiece(currentState, []grid.Point{})
	if coord.X != 3 && coord.Y != 0 {
		t.Errorf("Piece 12 coordinates should be in XY")
	}
}
