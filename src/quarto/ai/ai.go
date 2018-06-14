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

func InitAllTree(currentState state.State) StateNode {
	nextState := state.CopyState(currentState)
	tree := InitNode(nextState)
	tree = AppendChildNodes(tree, 2)
	return tree
}

func InitNode(state state.State) StateNode {
	return StateNode{State: state, Value: 0}
}

func AppendChildNodes(node StateNode, depth int) StateNode {
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
					node.Childs = append(node.Childs, AppendChildNodes(childNode, depth-1))
				} else {
					for i := 0; i < len(boxList); i++ {
						nextState := state.CopyState(node.State)
						nextState.Grid[boxList[i].Y][boxList[i].X] = node.State.Piece
						nextState.Piece = piecesList[j]
						childNode := InitNode(nextState)
						node.Childs = append(node.Childs, AppendChildNodes(childNode, depth-1))
					}
				}
			}
		}
	}
	return node
}

func PrintTree(node StateNode, depth int) {

	fmt.Println(RandStringBytes(depth*2) + " DEPTH[" + strconv.Itoa(depth) + "] / " + strconv.Itoa(len(node.Childs)) + " childs")
	fmt.Println(node.State.Grid)
	fmt.Println(RandStringBytes(depth*2) + " Piece : " + strconv.Itoa(node.State.Piece) + "(value : " + strconv.Itoa(node.Value) + ")")
	for i := 0; i < len(node.Childs); i++ {
		PrintTree(node.Childs[i], depth+1)
	}
}

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = " "[0]
	}
	return string(b)
}

func CountTimeElapsed(second_number int) {
	time.Sleep(time.Second * time.Duration(second_number))
	fmt.Println("Ended time")
}
