package game

import (
	"fmt"
)

// State define data for a game state
type State struct {
	Grid  [4][4]int
	Piece int
}

// DoAMove return the next move for given grid
func DoAMove(state State) State {

	state.Piece = 3
	fmt.Println("playing")
	return state
}
