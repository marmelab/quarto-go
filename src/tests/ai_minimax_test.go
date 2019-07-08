package tests

import (
	"fmt"
	"quarto/ai"
	"quarto/state"
	"quarto/grid"
	"testing"
	"strconv"
)

func TestInitAllTreeShouldReturnTree(t *testing.T) {
	var state = state.GetNewState(2)
	quit := make(chan struct{})
	tree := ai.InitAllTree(state, []int{}, []grid.Point{}, quit)
	close(quit)
	fmt.Println("tree")
	ai.PrintTree(tree, 0, 1)
	if tree.MyNode != false {
		t.Errorf("Bad tree (" + strconv.Itoa(tree.Value) + ")")
	}
}
