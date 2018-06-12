package game

import (
	"github.com/ahl5esoft/golang-underscore"
)

// GridSize is the Size of a grid
const GridSize = 4

// State define data for a game state
type State struct {
	Grid  [GridSize][GridSize]int
	Piece int
}

// CopyState create a new state copy of the parameter
func CopyState(state State) State {
	newState := State{}
	for i := 0; i < GridSize; i++ {
		for j := 0; j < GridSize; j++ {
			newState.Grid[i][j] = state.Grid[i][j]
		}
	}
	newState.Piece = state.Piece
	return newState
}

// DoAMove return the next move for given grid
func DoAMove(state State) State {

	state = PlacePieceOnGrid(state)
	state = ChooseNewPiece(state)
	return state
}

// PlacePieceOnGrid add the "Piece" id in an empty place of the Grid array
func PlacePieceOnGrid(state State) State {
	newState := CopyState(state)
	if (newState.Piece > 0) {
		for i := 0; i < GridSize; i++ {
			for j := 0; j < GridSize; j++ {
				if (newState.Grid[i][j] == 0) {
					newState.Grid[i][j] = newState.Piece
					newState.Piece = 0
					return newState
				}
			}
		}
	}
	return newState
}

// ChooseNewPiece select a new piece for opponent
func ChooseNewPiece(state State) State {
	newState := CopyState(state)
	newState.Piece = InitListOfRemainingPieces(newState)[0]
	return newState
}

// InitListOfRemainingPieces generate a list of pieces not already in the grid
func InitListOfRemainingPieces(state State) []int {
	var piecesList = InitListOfAllPieces(state)

	for i := 0; i < GridSize; i++ {
		for j := 0; j < GridSize; j++ {
			var index = underscore.FindIndex(piecesList, func(n, _ int) bool {
				return n == state.Grid[i][j]
			})
			if (index >= 0) {
				piecesList = append(piecesList[:index], piecesList[index+1:]...)
			}
		}
	}

	return piecesList
}

// InitListOfAllPieces generate a list of all pieces
func InitListOfAllPieces(state State) []int {
	var piecesList []int
	for i := 0; i < GridSize*GridSize; i++ {
		piecesList = append(piecesList, i+1)
	}
	return piecesList
}
