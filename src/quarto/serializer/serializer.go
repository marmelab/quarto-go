package serializer

import (
	"encoding/json"
	"quarto/state"
)

// FromJSONToState convert a json data into a game state
func FromJSONToState(b []byte) (state.State, error) {
	var currentState state.State
	err := json.Unmarshal(b, &currentState)
	if err != nil {
		return currentState, err
	}
	return currentState, nil
}

// FromStateToJSON convert a game state into a json data
func FromStateToJSON(currentState state.State) ([]byte, error) {
	output, err := json.Marshal(currentState)
	if err != nil {
		return output, err
	}
	return output, nil
}
