package tests

import "quarto/serializer"
import "quarto/game"
import "testing"
import "reflect"

func TestFromJSONToStateShouldReturnCorrectStateWithFourthSizedGrid(t *testing.T) {
	referenceState := game.GetNewState(4)
	referenceState.Grid[0] = []int{4,5,1,3}
	referenceState.Grid[2] = []int{9,7,10,0}
	referenceState.Piece = 2
	b :=[]byte("{\"Grid\":[[4,5,1,3],[0,0,0,0],[9,7,10,0],[0,0,0,0]],\"Piece\":2}")
	var state, err = serializer.FromJSONToState(b)
	if !reflect.DeepEqual(state, referenceState) {
		t.Error("State should be equal to reference")
	}
	if err != nil {
		t.Error("Unserializing shouldn't raise an error")
	}
}

func TestFromJSONToStateShouldReturnCorrectStateWithThirthSizedGrid(t *testing.T) {
	referenceState := game.GetNewState(3)
	referenceState.Grid[0] = []int{4,5,1}
	referenceState.Grid[1] = []int{9,7,0}
	referenceState.Piece = 6
	b :=[]byte("{\"Grid\":[[4,5,1],[9,7,0],[0,0,0]],\"Piece\":6}")
	var state, err = serializer.FromJSONToState(b)
	if !reflect.DeepEqual(state, referenceState) {
		t.Error("State should be equal to reference")
	}
	if err != nil {
		t.Error("Unserializing shouldn't raise an error")
	}
}

func TestFromJSONToStateShouldReturnCorrectStateWithFithSizedGrid(t *testing.T) {
	referenceState := game.GetNewState(5)
	referenceState.Grid[0] = []int{4,5,1,3,0}
	referenceState.Grid[2] = []int{9,7,10,0,12}
	referenceState.Piece = 2
	b :=[]byte("{\"Grid\":[[4,5,1,3,0],[0,0,0,0,0],[9,7,10,0,12],[0,0,0,0,0],[0,0,0,0,0]],\"Piece\":2}")
	var state, err = serializer.FromJSONToState(b)
	if !reflect.DeepEqual(state, referenceState) {
		t.Error("State should be equal to reference")
	}
	if err != nil {
		t.Error("Unserializing shouldn't raise an error")
	}
}

func TestFromStateToJSONShouldReturnCorrectBytesArrayWithFourthSizedGrid(t *testing.T) {
	var state = game.GetNewState(4)
	state.Grid[0] = []int{4,5,1,3}
	state.Grid[2] = []int{9,7,10,0}
	state.Piece = 2
	referenceBytesArray :=[]byte("{\"Grid\":[[4,5,1,3],[0,0,0,0],[9,7,10,0],[0,0,0,0]],\"Piece\":2}")
	var b, err = serializer.FromStateToJSON(state)
	if !reflect.DeepEqual(referenceBytesArray,b) {
		t.Error("Binary array should be equal to reference")
	}
	if err != nil {
		t.Error("Serializing shouldn't raise an error")
	}
}

func TestFromStateToJSONShouldReturnCorrectBytesArrayWithThirthSizedGrid(t *testing.T) {
	var state = game.GetNewState(3)
	state.Grid[0] = []int{4,5,1}
	state.Grid[1] = []int{9,7,0}
	state.Piece = 6
	referenceBytesArray :=[]byte("{\"Grid\":[[4,5,1],[9,7,0],[0,0,0]],\"Piece\":6}")
	var b, err = serializer.FromStateToJSON(state)
	if !reflect.DeepEqual(referenceBytesArray,b) {
		t.Error("Binary array should be equal to reference")
	}
	if err != nil {
		t.Error("Serializing shouldn't raise an error")
	}
}

func TestFromStateToJSONShouldReturnCorrectBytesArrayWithFithSizedGrid(t *testing.T) {
	var state = game.GetNewState(5)
	state.Grid[0] = []int{4,5,1,3,0}
	state.Grid[2] = []int{9,7,10,0,12}
	state.Piece = 2
	referenceBytesArray :=[]byte("{\"Grid\":[[4,5,1,3,0],[0,0,0,0,0],[9,7,10,0,12],[0,0,0,0,0],[0,0,0,0,0]],\"Piece\":2}")
	var b, err = serializer.FromStateToJSON(state)
	if !reflect.DeepEqual(referenceBytesArray,b) {
		t.Error("Binary array should be equal to reference")
	}
	if err != nil {
		t.Error("Serializing shouldn't raise an error")
	}
}
