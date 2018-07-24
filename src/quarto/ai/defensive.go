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
		var nilPoint = grid.Point{-1, -1}
		return &nilPoint
	}

	nonLoosingBoxList := grid.GetListBoxAMinusListBoxB(boxList, loosingBoxList)
	if len(nonLoosingBoxList) == 0 {
		return &boxList[r.Intn(len(boxList))]
	}

	return &nonLoosingBoxList[r.Intn(len(nonLoosingBoxList))]
}

// ChooseBlockingPositionForPiece return available coordinates to place the next piece to block a line of 4
func ChooseBlockingPositionForPiece(currentState state.State, loosingBoxList []grid.Point) *grid.Point {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	boxList := grid.GetBlockingBoxesIncludingPieceChoice(currentState.Grid, currentState.Piece)
	if len(boxList) == 0 {
		var nilPoint = grid.Point{-1, -1}
		return &nilPoint
	}

	nonLoosingBoxList := grid.GetListBoxAMinusListBoxB(boxList, loosingBoxList)
	if len(nonLoosingBoxList) == 0 {
		return &boxList[r.Intn(len(boxList))]
	}
	return &nonLoosingBoxList[r.Intn(len(nonLoosingBoxList))]
}
