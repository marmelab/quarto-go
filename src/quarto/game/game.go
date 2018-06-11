package game

// GridSize is the Size of a grid
const GridSize = 4

// State define data for a game state
type State struct {
	Grid  [GridSize][GridSize]int
	Piece int
}

// DoAMove return the next move for given grid
func DoAMove(state State) State {

	state = PlacePieceOnGrid(state)
	state = ChooseNewPiece(state)
	return state
}

// PlacePieceOnGrid add the "Piece" id in an empty place of the Grid array
func PlacePieceOnGrid(state State) State {
    if state.Piece > 0 {
        for i := 0; i < GridSize; i++ {
            for j := 0; j < GridSize; j++ {
                if state.Grid[i][j] == 0 {
                    state.Grid[i][j] = state.Piece
                    state.Piece = 0
                    return state
                }
            }
        }
    }
	return state
}

// ChooseNewPiece select a new piece for opponent
func ChooseNewPiece(state State) State {
	state.Piece = InitListOfRemainingPieces(state)[0]
	return state
}

// InitListOfRemainingPieces generate a list of pieces not already in the grid
func InitListOfRemainingPieces(state State) []int {
	var piecesList = InitListOfAllPieces(state)

	for i := 0; i < GridSize; i++ {
		for j := 0; j < GridSize; j++ {
			var index = IndexOf(piecesList, state.Grid[i][j])
			if index >= 0 {
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

// Contains returns the presence of an element in a list
func Contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// IndexOf returns index of an element in a list
func IndexOf(s []int, e int) int {
	for i := 0; i < len(s); i++ {
		if s[i] == e {
			return i
		}
	}
	return -1
}
