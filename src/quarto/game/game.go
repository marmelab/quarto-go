package game

import (
	"github.com/ahl5esoft/golang-underscore"
)

var emptyCoord = [2]int{-1, -1}

// State define data for a game state
type State struct {
	Grid  [][]int
	Piece int
	Move [2]int
}

// GetNewState return a blanck state of defined size
func GetNewState(size int) State {
	newState := State{}
	for i := 0; i < size; i++ {
		newState.Grid = append(newState.Grid, []int{})
		for j := 0; j < size; j++ {
			newState.Grid[i] = append(newState.Grid[i], 0)
		}
	}
	return newState
}

// GetGridSize return the size of a grid in a state
func GetGridSize(state State) int {
	return len(state.Grid)
}

// CopyState create a new state copy of the parameter
func CopyState(state State) State {
	size := GetGridSize(state)
	newState := GetNewState(size)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			newState.Grid[i][j] = state.Grid[i][j]
		}
	}
	newState.Piece = state.Piece
	newState.Move = state.Move
	return newState
}

// PlayTurn return the next move for given grid
func PlayTurn(state State) State {
	newState := PlacePieceOnGrid(state)
	return ChooseNewPiece(newState)
}

// PlacePieceOnGrid add the "Piece" id in an empty place of the Grid array
func PlacePieceOnGrid(state State) State {
	newState := CopyState(state)
	if newState.Piece > 0 {
		coord := ChoosePositionForPiece(state)
		newState.Grid[coord[0]][coord[1]] = newState.Piece
		newState.Move = coord
		newState.Piece = 0
	}
	return newState
}

// ChoosePositionForPiece return coordinates to place the next piece
func ChoosePositionForPiece(state State) [2]int{
	coord := ChooseWinningtPositionForPiece(state)
	if (coord == emptyCoord) {
		coord = ChooseFirstPositionForPiece(state)
	}
	return coord
}

// ChooseWinningtPositionForPiece return first winning coordinates to place the next piece if exists
func ChooseWinningtPositionForPiece(state State) [2]int{
	coord := emptyCoord
	size := GetGridSize(state)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if state.Grid[i][j] == 0 {
				if (IsWinningPosition(j,i,state, size)) {
					coord[0] = i
					coord[1] = j
					return coord
				}
			}
		}
	}
	return coord
}

// IsWinningPosition return true if the piece placed at these coordinates make a winning situation
func IsWinningPosition(x int, y int, state State, size int) bool {
	testState := CopyState(state)
	testState.Grid[x][y] = testState.Piece
	if (IsWinningLine(GetPiecesRaw(x,y,testState.Grid, size))) {
		return true
	}
	if (IsWinningLine(GetPiecesColumn(x,y,testState.Grid, size))) {
		return true
	}
	if (IsWinningLine(GetPiecesSlashDiag(x,y,testState.Grid, size))) {
		return true
	}
	if (IsWinningLine(GetPiecesBackSlashDiag(x,y,testState.Grid, size))) {
		return true
	}
	return false
}

// GetPiecesRaw return an array of the raw of pieces aligned in [x,y] coordinates
func GetPiecesRaw(x int, y int, grid [][]int, size int) []int {
	return grid[y]
}

// GetPiecesColumn return an array of the column of pieces aligned in [x,y] coordinates
func GetPiecesColumn(x int, y int, grid [][]int, size int) []int {
	piecesLine := []int{}
	for i := 0; i < size; i++ {
		piecesLine = append(piecesLine, grid[i][x])
	}
	return piecesLine
}

// GetPiecesSlashDiag return an array of the diag (slash oriented) of pieces aligned in [x,y] coordinates
func GetPiecesSlashDiag(x int, y int, grid [][]int, size int) []int {
	piecesLine := []int{}
	if (x == y) {
		for i := 0; i < size; i++ {
			piecesLine = append(piecesLine, grid[i][i])
		}
	}
	return piecesLine
}

// GetPiecesBackSlashDiag return an array of the diag (backslash oriented) of pieces aligned in [x,y] coordinates
func GetPiecesBackSlashDiag(x int, y int, grid [][]int, size int) []int {
	piecesLine := []int{}
	if (x == size - y - 1) {
		for i := 0; i < size; i++ {
			piecesLine = append(piecesLine, grid[i][size - i - 1])
		}
	}
	return piecesLine
}

// IsWinningLine return true if all piece in array makes a winning situation when aligned
func IsWinningLine(piecesLine []int) bool {
	if (len(piecesLine) == 0) {
		return false
	}
	return false
}


// ChooseFirstPositionForPiece return first available coordinates to place the next piece
func ChooseFirstPositionForPiece(state State) [2]int{
	coord := emptyCoord
	size := GetGridSize(state)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if state.Grid[i][j] == 0 {
				coord[0] = i
				coord[1] = j
				return coord
			}
		}
	}
	return coord
}

// ChooseNewPiece select a new piece for opponent
func ChooseNewPiece(state State) State {
	newState := CopyState(state)
	newState.Piece = GetRemainingPiecesListFromState(newState)[0]
	return newState
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
	if (!IsValidPiece(state.Piece,size)) {
		return false
	}
	if (!IsValidGrid(state, size)) {
		return false
	}
	if (!IsValidMove(state.Move, size)) {
		return false
	}
	return true
}

// IsValidGrid return false if the piace number is not acceptable
func IsValidGrid(state State, size int) bool {
	var piecesList = GetAllPiecesList(state)
	if (len(state.Grid) != size) {
		return false
	}
	for i := 0; i < size; i++ {
		if (len(state.Grid[i]) != size) {
			return false
		}
		for j := 0; j < size; j++ {
			if (!IsValidBox(state.Grid[i][j], size, state.Piece)) {
				return false
			}
			if (state.Grid[i][j] > 0) {
				var pieceIndex = underscore.FindIndex(piecesList, func(n, _ int) bool {
					return n == state.Grid[i][j]
				})
				if (pieceIndex < 0) {
					return false
				}
				piecesList = append(piecesList[:pieceIndex], piecesList[pieceIndex+1:]...)
			}
		}
	}
	return true
}

// IsValidPiece return false if the piace number is not acceptable
func IsValidPiece(piece int, size int) bool {
	return piece >= 0 && piece <= (size * size)
}

// IsValidBox return false if the box number is not acceptable
func IsValidBox(box int, size int, piece int) bool {
	if (box < 0 || box > (size * size)) {
		return false
	}
	if (piece > 0 && box == piece) {
		return false
	}
	return true
}

// IsValidMove return false if the move is not acceptable
func IsValidMove(move [2]int, size int) bool {
	if (move[0] < 0 || move[0] >= size) {
		return false
	}
	if (move[1] < 0 || move[1] >= size) {
		return false
	}
	return true
}
