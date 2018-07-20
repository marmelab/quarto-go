package tests

import (
	"fmt"
	"quarto/ai"
	"quarto/state"
	"testing"
	"strconv"
)

func TestInitAllTreeShouldReturnTree(t *testing.T) {
	var state = state.GetNewState(2)
	quit := make(chan struct{})
	tree := ai.InitAllTree(state, quit)
	close(quit)
	fmt.Println("tree")
	ai.PrintTree(tree, 0)
	if tree.Value != 0 {
		t.Errorf("Bad tree (" + strconv.Itoa(tree.Value) + ")")
	}
}
