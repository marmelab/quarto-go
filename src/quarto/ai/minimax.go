package ai

import (
	"fmt"
	"quarto/grid"
	"quarto/state"
	"strconv"
	"time"
)

const UndefinedValue = -1
const WinningLeafValue = 0
const LoosingLeafValue = 100

// StateNode define data for a node a the tree used in minmax algorithm
type StateNode struct {
	State  state.State
	Value  int
	MyNode bool
	Childs []StateNode
}

// InitAllTree creates tree of possibilities
func InitAllTree(currentState state.State, quit chan struct{}) StateNode {
	nextState := state.CopyState(currentState)
	tree := InitNode(nextState, false)
	tree = AppendChildNodes(tree, 5, quit)
	PrintTree(tree, 0, 1) 
	return tree
}

// InitNode creates new node
func InitNode(currentState state.State, myNode bool) StateNode {
	return StateNode{State: currentState, Value: UndefinedValue, MyNode: myNode}
}

// AppendChildNodes creates child nodes of current node
func AppendChildNodes(node StateNode, depth int, quit chan struct{}) StateNode {
	select {
		case <-quit:
			return node
		default:
			if depth >= 0 {
				node.Childs = []StateNode{}
				piecesList := state.GetRemainingPiecesListFromState(node.State)
				boxList := grid.GetEmptyBoxes(node.State.Grid)
				if len(piecesList) > 1 && len(boxList) > 1 {
					for j := 0; j < len(piecesList) ; j++ {
						if node.State.Piece != piecesList[j] {
							if node.State.Piece == 0 {
								nextState := state.CopyState(node.State)
								nextState.Piece = piecesList[j]
								childNode := InitNode(nextState, !node.MyNode)
								node.Childs = append(node.Childs, AppendChildNodes(childNode, depth-1, quit))
							} else {
								for i := 0; i < len(boxList); i++ {
									nextState := state.CopyState(node.State)
									nextState.Grid[boxList[i].Y][boxList[i].X] = node.State.Piece
									nextState.Piece = piecesList[j]
									childNode := InitNode(nextState, !node.MyNode)
									if grid.IsWinningPosition(boxList[i].X, boxList[i].Y, node.State.Grid, node.State.Piece) {
										if (childNode.MyNode) {
											childNode.Value = WinningLeafValue
										} else {
											childNode.Value = LoosingLeafValue
										}
										node.Childs = append(node.Childs, childNode)
									} else {
										node.Childs = append(node.Childs, AppendChildNodes(childNode, depth-1, quit ))
									}
								}
							}
						}
					}
				} else if len(piecesList) == 1 && len(boxList) == 1 {
					nextState := state.CopyState(node.State)
					nextState.Grid[boxList[0].Y][boxList[0].X] = node.State.Piece
					nextState.Piece = 0
					childNode := InitNode(nextState, !node.MyNode)

					if grid.IsWinningPosition(boxList[0].X, boxList[0].Y, node.State.Grid, node.State.Piece) {
						if (childNode.MyNode) {
							childNode.Value = WinningLeafValue
						} else {
							childNode.Value = LoosingLeafValue
						}
					} else {
						childNode.Value = WinningLeafValue
					}
					node.Childs = append(node.Childs, childNode)
				}
			}
			if (node.Value == UndefinedValue) {
				node.Value = CalculateNodeValue(node)
			}
			return node
	}
}

func CalculateNodeValue(node StateNode) int {
	value := 0
	if node.MyNode {
		value = WinningLeafValue
		for i := 0; i < len(node.Childs); i++ {
			if value < node.Childs[i].Value {
				value = node.Childs[i].Value
			}
		}
	} else {
		value = LoosingLeafValue
		for i := 0; i < len(node.Childs); i++ {
			if value > node.Childs[i].Value {
				value = node.Childs[i].Value
			}
		}
	}
	return value
}

func PrintTree(node StateNode, depth int, maxDisplayDepth int) {

	fmt.Println(FixedStringBytes(depth*2) + " DEPTH[" + strconv.Itoa(depth) + "] / " + strconv.Itoa(len(node.Childs)) + " childs")
	fmt.Println(node.State.Grid)
	if node.MyNode {
		fmt.Println(FixedStringBytes(depth*2) + " Piece : " + strconv.Itoa(node.State.Piece) + "(value : " + strconv.Itoa(node.Value) + " / my node)")
	} else {
		fmt.Println(FixedStringBytes(depth*2) + " Piece : " + strconv.Itoa(node.State.Piece) + "(value : " + strconv.Itoa(node.Value) + " / opponent node)")
	}
	if (maxDisplayDepth > depth) {
		for i := 0; i < len(node.Childs); i++ {
			PrintTree(node.Childs[i], depth+1, maxDisplayDepth)
		}
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
	return newState, !minMaxStopped
}
