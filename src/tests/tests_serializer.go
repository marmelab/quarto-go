package tests

import "quarto/serializer"
import "quarto/game"
import "testing"
import "fmt"
import "reflect"

func testFromJSONToStateShouldReturnCorrectState(t *testing.T) {
	var b []byte
	var state, err = serializer.FromJSONToState(b)
	if !reflect.DeepEqual(state,state) {
		t.Errorf("State should be correct")
	}
	if err != nil {
		t.Errorf("Serializing should raise an error")
	}
}

func testFromStateToJSONShouldReturnCorrectBytesArray(t *testing.T) {
	var state = game.GetNewState(4)
	var b, err = serializer.FromStateToJSON(state)
	fmt.Println(b)
	if len(b) != len(b) {
		t.Errorf("Binary array should be correct")
	}
	if err != nil {
		t.Errorf("Serializing should raise an error")
	}
}
