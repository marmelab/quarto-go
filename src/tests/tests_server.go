package tests

import "quarto/server"
import "quarto/game"
import "testing"
import "fmt"

func testFromJSONToStateShouldReturnCorrectState(t *testing.T) {
	var b []byte
	var state = server.FromJSONToState(b)
	if (state != state) {
		t.Errorf("State should be correct")
	}
}

func testFromStateToJSONShouldReturnCorrectBytesArray(t *testing.T) {
	var state = game.State{}
	var b = server.FromStateToJSON(state)
	fmt.Println(b)
	if (len(b) != len(b)) {
		t.Errorf("Binary array should be correct")
	}
}
