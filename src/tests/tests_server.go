package tests

import "quarto/server"
import "quarto/game"
import "testing"
import "fmt"

// TestFromJSONToState return correct state
func TestFromJSONToState(t *testing.T) {
	var b []byte
	//b = [10]byte
	var state = server.FromJSONToState(b)
	//if b != 10 {
	//	t.Errorf("Sum was incorrect, got: %d, want: %d.", b, 10)
	//}
}

//TestFromStateToJSON return correct bytes array
func TestFromStateToJSON(t *testing.T) {
	var state = game.State{}
	var b = server.FromStateToJSON(state)
	fmt.Println(b)
	//if b != 10 {
	//	t.Errorf("Sum was incorrect, got: %d, want: %d.", b, 10)
	//}
}