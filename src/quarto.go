package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/nextmove", DoNextMove)

	port := GetListeningPort()
	fmt.Println("Server started on port : " + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

// Index return the main page
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to Quarto-py!")
}

// DoNextMove return the next move for given grid
func DoNextMove(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Play the next move!")
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
