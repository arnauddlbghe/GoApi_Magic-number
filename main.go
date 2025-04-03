package main

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"net/http"
)

type RequestBody struct {
	Guess uint `json:"guess"`
}

type ResponseBody struct {
	Response string `json:"response"`
}

var toGuessNumber uint

func playHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var reqBody RequestBody
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			http.Error(w, "Requête invalide", http.StatusBadRequest)
			return
		}
		respBody := ResponseBody{Response: validateResponse(reqBody.Guess)}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(respBody)
	default:
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
	}
}

func startServer() {
	fmt.Println("Serveur démarré sur le port 8080...")
	http.ListenAndServe(":8080", nil)
}

func initGame() {
	toGuessNumber = rand.UintN(100)
	fmt.Println("toGuessNumber =", toGuessNumber)
}

func validateResponse(guess uint) string {
	if guess > toGuessNumber {
		return "No.. The number to guess is not that big. Try smaller"
	}
	if guess < toGuessNumber {
		return "No.. Try bigger"
	}
	initGame()
	return "You guessed it ! Play again, a new magic number have been generated"
}

func main() {
	http.HandleFunc("/play", playHandler)
	initGame()
	startServer()
}
