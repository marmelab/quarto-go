package game

import (
	"quarto/ai"
	"quarto/grid"
	"quarto/state"
	"fmt"
)

// PlayTurn return the next move for given grid
func PlayTurn(currentState state.State) state.State {
	newState, done := ai.StartMiniMax(currentState, 30)
	if !done {
		fmt.Println("minmax killed")
		newState = PlacePieceOnGrid(currentState)
		return DefineNewPiece(newState)
	}
	fmt.Println("minmax worked")
	return newState
}

// PlacePieceOnGrid add the "Piece" id in an empty place of the Grid array
func PlacePieceOnGrid(currentState state.State) state.State {
	newState := state.CopyState(currentState)

	if newState.Piece > 0 {
		coord := ChoosePositionForPiece(currentState)
		newState.Grid[coord.Y][coord.X] = newState.Piece
		newState.Move = [2]int{coord.Y, coord.X}
		newState.Piece = 0
	}
	return newState
}

// ChoosePositionForPiece return coordinates to place the next piece
func ChoosePositionForPiece(currentState state.State) *grid.Point {
	coord := ai.ChooseWinningPositionForPiece(currentState)
	if coord != nil {
		return coord
	}
	loosingBoxList := ai.GetLoosingBoxList(currentState)
	coord = ai.ChooseDefensivePositionForPiece(currentState, loosingBoxList)
	if coord != nil {
		return coord
	}
	return ai.ChooseRandomPositionForPiece(currentState, loosingBoxList)
}

// DefineNewPiece select a new piece for opponent
func DefineNewPiece(currentState state.State) state.State {
	newState := state.CopyState(currentState)
	newState.Piece = ai.ChooseNonWinningPiece(newState)
	if newState.Piece == 0 {
		newState.Piece = ai.ChooseRandomPiece(currentState)
	}
	return newState
}
