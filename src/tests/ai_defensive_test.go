package tests

import (
	"quarto/ai"
	"quarto/state"
	"testing"
)

func TestChooseDefensivePositionForPieceShouldPlaceOnFreeCoordinatesWhenCalledFirst(t *testing.T) {
	var currentState = state.GetNewState(4)
	currentState.Piece = 9
	coord := ai.ChooseDefensivePositionForPiece(currentState)
	if coord == nil {
		t.Errorf("Piece 9 coordinates should be in the Grid")
	}
}