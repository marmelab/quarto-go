package server

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"quarto/game"
	"strconv"
)

// Start launch the server
func Start() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/suggestMove", SuggestMove).
		Methods("POST")

	port := GetListeningPort()
	fmt.Println("Server started on port : " + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

// GetListeningPort determine the port to start the server (8080 is the default)
func GetListeningPort() string {
	const defaultPort = 8080

	currentPort := defaultPort
	var err error

	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) > 0 {
		stringPort := argsWithoutProg[0]
		currentPort, err = strconv.Atoi(stringPort)
		if err != nil {
			currentPort = defaultPort
		}
	}

	return strconv.Itoa(currentPort)
}

// SuggestMove wait for a new move request and return a move
func SuggestMove(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	b, err := ioutil.ReadAll(r.Body)

	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write(FromStateToJSON(game.DoAMove(FromJSONToState(b))))
}

// FromJSONToState convert a json data into a game state
func FromJSONToState(b []byte) game.State {
	var state game.State
	err := json.Unmarshal(b, &state)
	if err != nil {
		//return
	}
	return state
}

// FromStateToJSON convert a game state into a json data
func FromStateToJSON(state game.State) []byte {
	output, err := json.Marshal(state)
	if err != nil {
		//return
	}
	return output
}
