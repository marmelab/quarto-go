package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"quarto/game"
	"quarto/serializer"
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
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")

	b, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Error invalid request", 500)
		return
	}

	state, err := serializer.FromJSONToState(b)
	if (!game.IsValid(state)) {
		http.Error(w, "Error invalid state in request", 500)
		return
	}

	if err != nil {
		http.Error(w, "Error invalid json request", 500)
		return
	}

	serializedState, err := serializer.FromStateToJSON(game.PlayTurn(state))

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write(serializedState)
}
