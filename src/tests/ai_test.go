package tests

import (
	"fmt"
	"quarto/ai"
	"quarto/state"
	"testing"
)

func estInitAllTreeShouldReturnTree(t *testing.T) {
	var state = state.GetNewState(4)
	tree := ai.InitAllTree(state)
	fmt.Println("tree")
	ai.PrintTree(tree, 0)
	if 1 != 4 {
		t.Errorf("Bad tree")
	}
}
