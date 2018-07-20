package tests

import "quarto/serializer"
import "quarto/state"
import "testing"
import "reflect"

func TestFromJSONToStateShouldReturnCorrectStateWithFourthSizedGrid(t *testing.T) {
	referenceState := state.GetNewState(4)
	referenceState.Grid[0] = []int{4, 5, 1, 3}
	referenceState.Grid[2] = []int{9, 7, 10, 0}
	referenceState.Piece = 2
	b := []byte("{\"Grid\":[[4,5,1,3],[0,0,0,0],[9,7,10,0],[0,0,0,0]],\"Piece\":2,\"Move\":[0,0]}")
	var currentState, err = serializer.FromJSONToState(b)
	if !reflect.DeepEqual(currentState, referenceState) {
		t.Error("State should be equal to reference")
	}
	if err != nil {
		t.Error("Unserializing shouldn't raise an error")
	}
}

func TestFromJSONToStateShouldReturnCorrectStateWithThirthSizedGrid(t *testing.T) {
	referenceState := state.GetNewState(3)
	referenceState.Grid[0] = []int{4, 5, 1}
	referenceState.Grid[1] = []int{9, 7, 0}
	referenceState.Piece = 6
	b := []byte("{\"Grid\":[[4,5,1],[9,7,0],[0,0,0]],\"Piece\":6,\"Move\":[0,0]}")
	var currentState, err = serializer.FromJSONToState(b)
	if !reflect.DeepEqual(currentState, referenceState) {
		t.Error("State should be equal to reference")
	}
	if err != nil {
		t.Error("Unserializing shouldn't raise an error")
	}
}

func TestFromJSONToStateShouldReturnCorrectStateWithFithSizedGrid(t *testing.T) {
	referenceState := state.GetNewState(5)
	referenceState.Grid[0] = []int{4, 5, 1, 3, 0}
	referenceState.Grid[2] = []int{9, 7, 10, 0, 12}
	referenceState.Piece = 2
	b := []byte("{\"Grid\":[[4,5,1,3,0],[0,0,0,0,0],[9,7,10,0,12],[0,0,0,0,0],[0,0,0,0,0]],\"Piece\":2,\"Move\":[0,0]}")
	var currentState, err = serializer.FromJSONToState(b)
	if !reflect.DeepEqual(currentState, referenceState) {
		t.Error("State should be equal to reference")
	}
	if err != nil {
		t.Error("Unserializing shouldn't raise an error")
	}
}

func TestFromStateToJSONShouldReturnCorrectBytesArrayWithFourthSizedGrid(t *testing.T) {
	var currentState = state.GetNewState(4)
	currentState.Grid[0] = []int{4, 5, 1, 3}
	currentState.Grid[2] = []int{9, 7, 10, 0}
	currentState.Piece = 2
	referenceBytesArray := []byte("{\"Grid\":[[4,5,1,3],[0,0,0,0],[9,7,10,0],[0,0,0,0]],\"Piece\":2,\"Move\":[0,0]}")
	var b, err = serializer.FromStateToJSON(currentState)
	if !reflect.DeepEqual(referenceBytesArray, b) {
		t.Error("Binary array should be equal to reference")
		t.Error("b : " + string(b))
		t.Error("reference : " + string(referenceBytesArray))
	}
	if err != nil {
		t.Error("Serializing shouldn't raise an error")
	}
}

func TestFromStateToJSONShouldReturnCorrectBytesArrayWithThirthSizedGrid(t *testing.T) {
	var currentState = state.GetNewState(3)
	currentState.Grid[0] = []int{4, 5, 1}
	currentState.Grid[1] = []int{9, 7, 0}
	currentState.Piece = 6
	referenceBytesArray := []byte("{\"Grid\":[[4,5,1],[9,7,0],[0,0,0]],\"Piece\":6,\"Move\":[0,0]}")
	var b, err = serializer.FromStateToJSON(currentState)
	if !reflect.DeepEqual(referenceBytesArray, b) {
		t.Error("Binary array should be equal to reference")
		t.Error("b : " + string(b))
		t.Error("reference : " + string(referenceBytesArray))
	}
	if err != nil {
		t.Error("Serializing shouldn't raise an error")
	}
}

func TestFromStateToJSONShouldReturnCorrectBytesArrayWithFithSizedGrid(t *testing.T) {
	var currentState = state.GetNewState(5)
	currentState.Grid[0] = []int{4, 5, 1, 3, 0}
	currentState.Grid[2] = []int{9, 7, 10, 0, 12}
	currentState.Piece = 2
	referenceBytesArray := []byte("{\"Grid\":[[4,5,1,3,0],[0,0,0,0,0],[9,7,10,0,12],[0,0,0,0,0],[0,0,0,0,0]],\"Piece\":2,\"Move\":[0,0]}")
	var b, err = serializer.FromStateToJSON(currentState)
	if !reflect.DeepEqual(referenceBytesArray, b) {
		t.Error("Binary array should be equal to reference")
		t.Error("b : " + string(b))
		t.Error("reference : " + string(referenceBytesArray))
	}
	if err != nil {
		t.Error("Serializing shouldn't raise an error")
	}
}
