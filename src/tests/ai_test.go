package tests

import (
	"fmt"
	"quarto/ai"
	"quarto/state"
	"testing"
)

func estInitAllTreeShouldReturnTree(t *testing.T) {
	var state = state.GetNewState(2)
	quit := make(chan struct{})
	tree := ai.InitAllTree(state, quit)
	close(quit)
	fmt.Println("tree")
	ai.PrintTree(tree, 0)
	if 1 != 4 {
		t.Errorf("Bad tree")
	}
}
