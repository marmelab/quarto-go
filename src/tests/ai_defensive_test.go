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
	if coord == nil {
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