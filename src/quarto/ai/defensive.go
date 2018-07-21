package ai

import (
	"math/rand"
	"time"
	"quarto/grid"
	"quarto/state"
)

// ChooseDefensivePositionForPiece return available coordinates to place the next piece where grid is the less filled
func ChooseDefensivePositionForPiece(currentState state.State, loosingBoxList []grid.Point) *grid.Point {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	boxList := grid.GetSafestBoxesIncludingPieceChoice(currentState.Grid, currentState.Piece)
	if len(boxList) == 0 {
		return nil
	}

	nonLoosingBoxList := grid.GetListBoxAMinusListBoxB(boxList, loosingBoxList)

	if len(nonLoosingBoxList) == 0 {
		return &boxList[r.Intn(len(boxList))]
	}

	return &nonLoosingBoxList[r.Intn(len(nonLoosingBoxList))]
}
