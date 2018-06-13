package game

import (
	"github.com/ahl5esoft/golang-underscore"
	"math/rand"
	"quarto/grid"
	"time"
)

// State define data for a game state
type State struct {
	Grid  [][]int
	Piece int
	Move  [2]int
}

// GetNewState return a blank state of defined size
func GetNewState(size int) State {
	newState := State{}
	newState.Grid = grid.GetNewGrid(size)
	return newState
}

// GetGridSize return the size of a grid in a state
func GetGridSize(state State) int {
	return len(state.Grid)
}

// CopyState create a new state copy of the parameter
func CopyState(state State) State {
	newState := GetNewState(GetGridSize(state))
	newState.Grid = grid.CopyGrid(state.Grid)
	newState.Piece = state.Piece
	newState.Move = state.Move
	return newState
}

// PlayTurn return the next move for given grid
func PlayTurn(state State) State {
	newState := PlacePieceOnGrid(state)
	return DefineNewPiece(newState)
}

// PlacePieceOnGrid add the "Piece" id in an empty place of the Grid array
func PlacePieceOnGrid(state State) State {
	newState := CopyState(state)
	if newState.Piece > 0 {
		coord := ChoosePositionForPiece(state)
		newState.Grid[coord.Y][coord.X] = newState.Piece
		newState.Move = [2]int{coord.Y, coord.X}
		newState.Piece = 0
	}
	return newState
}

// ChoosePositionForPiece return coordinates to place the next piece
func ChoosePositionForPiece(state State) *grid.Point {
	coord := ChooseWinningPositionForPiece(state)
	if coord == nil {
		coord = ChooseRandomPositionForPiece(state)
	}
	return coord
}

// ChooseWinningPositionForPiece return first winning coordinates to place the next piece if exists
func ChooseWinningPositionForPiece(state State) *grid.Point {
	pointList := grid.GetEmptyBoxes(state.Grid)
	for i := 0; i < len(pointList); i++ {
		if grid.IsWinningPosition(pointList[i].X, pointList[i].Y, state.Grid, state.Piece) {
			return &pointList[i]
		}
	}
	return nil
}

// ChooseRandomPositionForPiece return random available coordinates to place the next piece
func ChooseRandomPositionForPiece(state State) *grid.Point {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	pointList := grid.GetEmptyBoxes(state.Grid)
	if len(pointList) == 0 {
		return nil
	}
	return &pointList[r.Intn(len(pointList))]
}

// DefineNewPiece select a new piece for opponent
func DefineNewPiece(state State) State {
	newState := CopyState(state)
	newState.Piece = ChooseNonWinningPiece(newState)
	if newState.Piece == 0 {
		newState.Piece = ChooseRandomPiece(state)
	}
	return newState
}

// ChooseRandomPiece choose a random piece for next opponent turn
func ChooseRandomPiece(state State) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	piecesList := GetRemainingPiecesListFromState(state)
	if len(piecesList) == 0 {
		return 0
	}
	return piecesList[r.Intn(len(piecesList))]
}

// ChooseNonWinningPiece choose a non winning piece for next opponent turn
func ChooseNonWinningPiece(state State) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	piecesList := GetNonWinningPiecesListFromState(state)
	if len(piecesList) == 0 {
		return 0
	}
	return piecesList[r.Intn(len(piecesList))]
}

// GetNonWinningPiecesListFromState generate a list of pieces to play wich can't win on next turn
func GetNonWinningPiecesListFromState(state State) []int {
	var piecesList = GetRemainingPiecesListFromState(state)
	pointList := grid.GetEmptyBoxes(state.Grid)
	for i := 0; i < len(piecesList); i++ {
		for j := 0; j < len(pointList); j++ {
			if grid.IsWinningPosition(pointList[j].X, pointList[j].Y, state.Grid, piecesList[i]) {
				piecesList = append(piecesList[:i], piecesList[i+1:]...)
			}
		}
	}
	return piecesList
}

// GetRemainingPiecesListFromState generate a list of pieces not already in the grid
func GetRemainingPiecesListFromState(state State) []int {
	var piecesList = GetAllPiecesList(state)
	size := GetGridSize(state)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			var index = underscore.FindIndex(piecesList, func(n, _ int) bool {
				return n == state.Grid[i][j]
			})
			if index >= 0 {
				piecesList = append(piecesList[:index], piecesList[index+1:]...)
			}
		}
	}

	return piecesList
}

// GetAllPiecesList generate a list of all pieces
func GetAllPiecesList(state State) []int {
	var piecesList []int
	size := GetGridSize(state)
	for i := 0; i < size*size; i++ {
		piecesList = append(piecesList, i+1)
	}
	return piecesList
}

// IsValid return false if the state is not acceptable
func IsValid(state State) bool {
	size := GetGridSize(state)
	if !IsValidPiece(state.Piece, size) {
		return false
	}
	if !IsValidGrid(state, size) {
		return false
	}
	if !IsValidMove(state.Move, size) {
		return false
	}
	return true
}

// IsValidGrid return false if the piace number is not acceptable
func IsValidGrid(state State, size int) bool {
	var piecesList = GetAllPiecesList(state)
	if len(state.Grid) != size {
		return false
	}
	for i := 0; i < size; i++ {
		if len(state.Grid[i]) != size {
			return false
		}
		for j := 0; j < size; j++ {
			if !IsValidBox(state.Grid[i][j], size, state.Piece) {
				return false
			}
			if state.Grid[i][j] > 0 {
				var pieceIndex = underscore.FindIndex(piecesList, func(n, _ int) bool {
					return n == state.Grid[i][j]
				})
				if pieceIndex < 0 {
					return false
				}
				piecesList = append(piecesList[:pieceIndex], piecesList[pieceIndex+1:]...)
			}
		}
	}
	return true
}

// IsValidPiece return false if the piece number is not acceptable
func IsValidPiece(piece int, size int) bool {
	return piece >= 0 && piece <= (size*size)
}

// IsValidBox return false if the box number is not acceptable
func IsValidBox(box int, size int, piece int) bool {
	if box < 0 || box > (size*size) {
		return false
	}
	if piece > 0 && box == piece {
		return false
	}
	return true
}

// IsValidMove return false if the move is not acceptable
func IsValidMove(move [2]int, size int) bool {
	if move[0] < 0 || move[0] >= size {
		return false
	}
	if move[1] < 0 || move[1] >= size {
		return false
	}
	return true
}
