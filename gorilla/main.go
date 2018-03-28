package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type player struct {
	Name   string `json:"name" schema:"name"`
	Number int    `json:"number" schema:"number"`
}

var players map[int]player

func main() {
	players = map[int]player{
		87: player{"Sidney Crosby", 87},
		71: player{"Evgeni Malkin", 71},
		81: player{"Phil Kessel", 81},
	}

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", index)
	router.HandleFunc("/players", getPlayers).Methods("GET")
	router.HandleFunc("/players/{number}", getPlayer).Methods("GET")
	router.HandleFunc("/players", createPlayer).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/players", 302)
}

func getPlayers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	results := make([]player, 0, len(players))
	for _, value := range players {
		results = append(results, value)
	}

	if err := json.NewEncoder(w).Encode(results); err != nil {
		panic(err)
	}
}

func getPlayer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	number, err := strconv.Atoi(vars["number"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid number"))
		return
	}

	fmt.Println("looking for", number)

	result, ok := players[number]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Player not found"))
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		panic(err)
	}
}

func createPlayer(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	result := new(player)
	err := decoder.Decode(result)

	if err != nil || result.Number == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid request"))
		return
	}

	players[result.Number] = *result

	http.Redirect(w, r, "/players", http.StatusOK)
}
