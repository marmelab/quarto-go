package ai

import (
	"github.com/ahl5esoft/golang-underscore"
	"math/rand"
	"time"
	"quarto/grid"
	"quarto/state"
)

// ChooseWinningPositionForPiece return first winning coordinates to place the next piece if exists
func ChooseWinningPositionForPiece(currentState state.State) *grid.Point {
	boxList := grid.GetEmptyBoxes(currentState.Grid)
	for i := 0; i < len(boxList); i++ {
		if grid.IsWinningPosition(boxList[i].X, boxList[i].Y, currentState.Grid, currentState.Piece) {
			return &boxList[i]
		}
	}
	return nil
}

// ChooseRandomPositionForPiece return random available coordinates to place the next piece
func ChooseRandomPositionForPiece(currentState state.State) *grid.Point {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	boxList := grid.GetEmptyBoxes(currentState.Grid)
	if len(boxList) == 0 {
		return nil
	}
	return &boxList[r.Intn(len(boxList))]
}

// ChooseRandomPiece choose a random piece for next opponent turn
func ChooseRandomPiece(currentState state.State) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	piecesList := state.GetRemainingPiecesListFromState(currentState)
	if len(piecesList) == 0 {
		return 0
	}
	return piecesList[r.Intn(len(piecesList))]
}

// ChooseNonWinningPiece choose a non winning piece for next opponent turn
func ChooseNonWinningPiece(currentState state.State) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	piecesList := GetNonWinningPiecesListFromState(currentState)
	if len(piecesList) == 0 {
		return 0
	}
	return piecesList[r.Intn(len(piecesList))]
}

// GetNonWinningPiecesListFromState generate a list of pieces to play which can't win on next turn
func GetNonWinningPiecesListFromState(currentState state.State) []int {
	var piecesListInitial = state.GetRemainingPiecesListFromState(currentState)
	var piecesListWinning = []int{}
	boxList := grid.GetEmptyBoxes(currentState.Grid)
	for i := 0; i < len(piecesListInitial); i++ {
		for j := 0; j < len(boxList); j++ {
			if grid.IsWinningPosition(boxList[j].X, boxList[j].Y, currentState.Grid, piecesListInitial[i]) {
				piecesListWinning = append(piecesListWinning, piecesListInitial[i])
				break;
			}
		}
	}
	var piecesListNonWinning = underscore.Select(piecesListInitial, func(ni int, _ int) bool {
		var indexWinningPiece = underscore.FindIndex(piecesListWinning, func(nw int, _ int) bool {
			return ni == nw
		})
		return indexWinningPiece < 0
	})
	if piecesListNonWinning == nil {
		return []int{}
	}
	return piecesListNonWinning.([]int)
}
