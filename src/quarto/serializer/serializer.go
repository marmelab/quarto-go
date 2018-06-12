package serializer

import (
	"encoding/json"
	"quarto/game"
)

// FromJSONToState convert a json data into a game state
func FromJSONToState(b []byte) (game.State, error) {
	var state game.State
	err := json.Unmarshal(b, &state)
	if err != nil {
		return state, err
	}
	return state, nil
}

// FromStateToJSON convert a game state into a json data
func FromStateToJSON(state game.State) ([]byte, error) {
	output, err := json.Marshal(state)
	if err != nil {
		return output, err
	}
	return output, nil
}
