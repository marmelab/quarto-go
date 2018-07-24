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
func InitAllTree(currentState state.State, piecesList []int, boxList []grid.Point, quit chan struct{}) StateNode {
	nextState := state.CopyState(currentState)
	tree := InitNode(nextState, false)

	tree = AppendChildNodes(tree, 15, piecesList, boxList, quit)
	return tree
}

// InitNode creates new node
func InitNode(currentState state.State, myNode bool) StateNode {
	return StateNode{State: currentState, Value: UndefinedValue, MyNode: myNode}
}

// AppendChildNodes creates child nodes of current node
func AppendChildNodes(node StateNode, depth int, piecesList []int, boxList []grid.Point, quit chan struct{}) StateNode {
	select {
		case <-quit:
			return node
		default:
			if depth >= 0 {
				node.Childs = []StateNode{}
				if len(piecesList) > 1 && len(boxList) > 1 {
					for j := 0; j < len(piecesList) ; j++ {
						if node.State.Piece != piecesList[j] {
							if node.State.Piece == 0 {
								node.Childs = BuildFirstTurnNode(node, piecesList, j, boxList, depth, quit)
							} else {
								for i := 0; i < len(boxList); i++ {
									node.Childs = BuildNormalTurnNode(node, piecesList, j, i, boxList, depth, quit)
								}
							}
						}
					}
				} else if len(piecesList) == 1 && len(boxList) == 1 {
					node.Childs = BuildLastTurnNode(node, boxList, quit)
				}
			}
			if (node.Value == UndefinedValue) {
				node.Value = CalculateNodeValue(node)
			}
			return node
	}
}

func BuildFirstTurnNode(currentNode StateNode, piecesList []int, pieceIndex int, boxList []grid.Point, depth int, quit chan struct{}) []StateNode {
	nextState := state.CopyState(currentNode.State)
	nextState.Piece = piecesList[pieceIndex]
	childNode := InitNode(nextState, !currentNode.MyNode)
	childPiecesList := grid.GetListPieceMinusPiece(piecesList, nextState.Piece)
	return append(currentNode.Childs, AppendChildNodes(childNode, depth-1, childPiecesList, boxList, quit))
}

func BuildNormalTurnNode(currentNode StateNode, piecesList []int, pieceIndex int, boxIndex int, boxList []grid.Point, depth int, quit chan struct{}) []StateNode {
	nextState := state.CopyState(currentNode.State)
	nextState.Grid[boxList[boxIndex].Y][boxList[boxIndex].X] = currentNode.State.Piece
	nextState.Move = [2]int{boxList[boxIndex].Y, boxList[boxIndex].X}
	nextState.Piece = piecesList[pieceIndex]
	childNode := InitNode(nextState, !currentNode.MyNode)
	if grid.IsWinningPosition(boxList[boxIndex].X, boxList[boxIndex].Y, currentNode.State.Grid, currentNode.State.Piece) {
		if (childNode.MyNode) {
			childNode.Value = WinningLeafValue
		} else {
			childNode.Value = LoosingLeafValue
		}
		return append(currentNode.Childs, childNode)
	} else {
		childPiecesList := grid.GetListPieceMinusPiece(piecesList, nextState.Piece)
		childBoxList := grid.GetListBoxMinusPoint(boxList, boxList[boxIndex])
		return append(currentNode.Childs, AppendChildNodes(childNode, depth-1, childPiecesList, childBoxList, quit ))
	}
}

func BuildLastTurnNode(currentNode StateNode, boxList []grid.Point, quit chan struct{}) []StateNode {
	nextState := state.CopyState(currentNode.State)
	nextState.Grid[boxList[0].Y][boxList[0].X] = currentNode.State.Piece
	nextState.Move = [2]int{boxList[0].Y, boxList[0].X}
	nextState.Piece = 0
	childNode := InitNode(nextState, !currentNode.MyNode)

	if grid.IsWinningPosition(boxList[0].X, boxList[0].Y, currentNode.State.Grid, currentNode.State.Piece) {
		if (childNode.MyNode) {
			childNode.Value = WinningLeafValue
		} else {
			childNode.Value = LoosingLeafValue
		}
	} else {
		childNode.Value = WinningLeafValue
	}
	return append(currentNode.Childs, childNode)
}

func ChooseBestChildState(node StateNode) state.State {
	value := LoosingLeafValue + 1
	bestState := state.State{}
	for i := 0; i < len(node.Childs); i++ {
		if value > node.Childs[i].Value && node.Childs[i].Value != UndefinedValue {
			value = node.Childs[i].Value
			bestState = node.Childs[i].State
		}
	}
	return bestState
}

func CalculateNodeValue(node StateNode) int {
	value := 0
	if node.MyNode {
		value = WinningLeafValue
		for i := 0; i < len(node.Childs); i++ {
			if value < node.Childs[i].Value && node.Childs[i].Value != UndefinedValue {
				value = node.Childs[i].Value
			}
		}
	} else {
		value = LoosingLeafValue
		for i := 0; i < len(node.Childs); i++ {
			if value > node.Childs[i].Value && node.Childs[i].Value != UndefinedValue {
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
	piecesList := state.GetRemainingPiecesListFromState(currentState)
	if len(piecesList) > 7 {
		fmt.Println("minmax ignored")
		return currentState, false
	}
	boxList := grid.GetEmptyBoxes(currentState.Grid)

	stoppedchan := make(chan bool)
	statechan := make(chan state.State)
	quit := make(chan struct{})

	go func() {
		time.Sleep(time.Second * time.Duration(secondNumber))
		stoppedchan <- true
		statechan <- currentState
	}()

	go func() {

		tree := InitAllTree(currentState, piecesList, boxList, quit)
		bestState := ChooseBestChildState(tree)

		stoppedchan <- false
		statechan <- bestState
	}()

	minMaxStopped := <-stoppedchan
	newState := <-statechan
	close(quit)
	return newState, !minMaxStopped
}
