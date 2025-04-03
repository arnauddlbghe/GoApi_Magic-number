package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func playHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	var reqBody PlayRequestBody
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "Requête invalide", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	for i, player := range players {
		if player.Name == reqBody.Name {
			result := validateResponse(reqBody.Guess)
			if result.Success {
				respBody := ResponseBody{
					Response: result.Message + ". " + strconv.Itoa(int(player.GuessLeft)) + " tries were left.",
				}
				json.NewEncoder(w).Encode(respBody)
				removePlayer(player.Name)
				return
			}
			players[i].GuessLeft--
			respBody := ResponseBody{Response: result.Message + ". " + strconv.Itoa(int(players[i].GuessLeft)) + " tries left"}
			json.NewEncoder(w).Encode(respBody)
			return
		}
	}
	json.NewEncoder(w).Encode(ResponseBody{Response: "You are not playing. Join the game via /join"})
}

func startHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}
	initGame()
	json.NewEncoder(w).Encode(ResponseBody{Response: "Game started! Guess the number"})
}

func joinHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	var reqBody NameRequestBody
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "Requête invalide", http.StatusBadRequest)
		return
	}

	if !gameStarted {
		json.NewEncoder(w).Encode(ResponseBody{Response: "Game hasn't started yet."})
		return
	}

	if playerIsPlaying(reqBody.Name) {
		json.NewEncoder(w).Encode(ResponseBody{Response: "You are already in the game."})
	} else {
		players = append(players, Player{Name: reqBody.Name, GuessLeft: startGuessNumbers})
		json.NewEncoder(w).Encode(ResponseBody{Response: "Player added. Total: " + strconv.Itoa(len(players))})
	}
}

func leaveHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	var reqBody NameRequestBody
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "Requête invalide", http.StatusBadRequest)
		return
	}

	if playerIsPlaying(reqBody.Name) {
		removePlayer(reqBody.Name)
		json.NewEncoder(w).Encode(ResponseBody{Response: "You left the game."})
	} else {
		json.NewEncoder(w).Encode(ResponseBody{Response: "You are not in the game."})
	}
}
