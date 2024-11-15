package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type DiceRequest struct {
	Dice []int `json:"dice"`
}

var diceBag []int

func deleteDiceHandler(w http.ResponseWriter, r *http.Request) {
	if len(diceBag) == 0 {
		http.Error(w, "Dice not found", http.StatusNotFound)
		return
	}

	n := rand.Intn(len(diceBag))
	removedDice := diceBag[n]
	if n == len(diceBag)-1 {
		diceBag = diceBag[:n]
	} else {
		diceBag = append(diceBag[:n], diceBag[n+1:]...)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(strconv.Itoa(removedDice)))
}

func addDiceHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var diceRequest DiceRequest
	if err := json.NewDecoder(r.Body).Decode(&diceRequest); err != nil {
		http.Error(w, "Unprocessable Entity", http.StatusUnprocessableEntity)
		return
	}

	if len(diceRequest.Dice) == 0 {
		http.Error(w, "Unprocessable Entity", http.StatusUnprocessableEntity)
		return
	}

	diceBag = append(diceBag, diceRequest.Dice...)
	w.WriteHeader(http.StatusOK)
}
func main() {
	// Assign the helloWorldHandler function to the root URL path
	http.HandleFunc("/delete", deleteDiceHandler)
	http.HandleFunc("/dice/new", addDiceHandler)
	diceBag = append(diceBag, 3, 5, 7)
	// Set the server to listen on port 8020 and log any errors
	fmt.Println("Server starting on port 8020...")
	if err := http.ListenAndServe(":8020", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
