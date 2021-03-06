package ai

import (
	"fmt"
	"quarto/grid"
	"quarto/state"
	"strconv"
	"time"
)

// StateNode define data for a node a the tree used in minmax algorithm
type StateNode struct {
	State  state.State
	Value  int
	Childs []StateNode
}

// InitAllTree creates tree of possibilities
func InitAllTree(currentState state.State, quit chan struct{}) StateNode {
	nextState := state.CopyState(currentState)
	tree := InitNode(nextState)
	tree = AppendChildNodes(tree, 2, quit)
	return tree
}

// InitNode creates new node
func InitNode(currentState state.State) StateNode {
	return StateNode{State: currentState, Value: 0}
}

// AppendChildNodes creates nodes of current node
func AppendChildNodes(node StateNode, depth int, quit chan struct{}) StateNode {
	select {
		case <-quit:
			return node
		default:
			if depth >= 0 {
				node.Childs = []StateNode{}
				piecesList := state.GetRemainingPiecesListFromState(node.State)
				boxList := grid.GetEmptyBoxes(node.State.Grid)
				if len(piecesList) > 0 && len(boxList) > 0 {
					for j := 0; j < len(piecesList); j++ {
						if node.State.Piece == 0 {
							nextState := state.CopyState(node.State)
							nextState.Piece = piecesList[j]
							childNode := InitNode(nextState)
							node.Childs = append(node.Childs, AppendChildNodes(childNode, depth-1, quit))
						} else {
							for i := 0; i < len(boxList); i++ {
								nextState := state.CopyState(node.State)
								nextState.Grid[boxList[i].Y][boxList[i].X] = node.State.Piece
								nextState.Piece = piecesList[j]
								childNode := InitNode(nextState)
								node.Childs = append(node.Childs, AppendChildNodes(childNode, depth-1, quit ))
							}
						}
					}
				}
			}
			return node
	}
}

func PrintTree(node StateNode, depth int) {

	fmt.Println(FixedStringBytes(depth*2) + " DEPTH[" + strconv.Itoa(depth) + "] / " + strconv.Itoa(len(node.Childs)) + " childs")
	fmt.Println(node.State.Grid)
	fmt.Println(FixedStringBytes(depth*2) + " Piece : " + strconv.Itoa(node.State.Piece) + "(value : " + strconv.Itoa(node.Value) + ")")
	for i := 0; i < len(node.Childs); i++ {
		PrintTree(node.Childs[i], depth+1)
	}
}

func FixedStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = " "[0]
	}
	return string(b)
}

// StartMiniMax tries to perform a very good move with minimax in imparted time
func StartMiniMax(currentState state.State, secondNumber int) (returnState state.State, err bool) {
	newState := state.CopyState(currentState)

	stoppedchan := make(chan bool)
	statechan := make(chan state.State)
	quit := make(chan struct{})

	go func() {
		time.Sleep(time.Second * time.Duration(secondNumber))
		stoppedchan <- true
		statechan <- currentState
	}()

	go func() {
		bestState := state.CopyState(newState)

		InitAllTree(bestState, quit)
		
		stoppedchan <- false
		statechan <- bestState
	}()

	minMaxStopped := <-stoppedchan
	newState = <-statechan
	close(quit)
	//False added to force fail of minmax until algorithm have been finished
	return newState, !minMaxStopped && false
}
