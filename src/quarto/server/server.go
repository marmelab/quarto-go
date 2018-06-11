package server

import (
	"fmt"
    "github.com/gorilla/mux"
    "io/ioutil"
	"log"
	"net/http"
	"os"
    "strconv"
	//"encoding/json"
	"quarto/game"
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

    vars := mux.Vars(r)

    fmt.Fprintln(w, vars)

    //var m Member
    b, _ := ioutil.ReadAll(r.Body)
    fmt.Fprintln(w, b)
    //json.Unmarshal(b, &m)
    //members = append(members, m)
    //j, _ := json.Marshal(m)
    //w.Write(j)
    //j, _ := json.Marshal(members)
    //login := r.FormValue("login")

	fmt.Fprintln(w, FromStateToJSON(game.DoAMove(FromJSONToState("toto"))))
}

// FromJSONToState convert a json data into a game state
func FromJSONToState(j string) game.State {
	var grid [4][4]int
	return game.State{grid, 20}
}

// FromStateToJSON convert a game state into a json data 
func FromStateToJSON(state game.State) string {
	return "toto"
}