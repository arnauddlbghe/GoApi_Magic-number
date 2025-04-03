package main

import (
	"fmt"
	"math/rand/v2"
)

var startGuessNumbers uint = 6
var toGuessNumber uint
var players []Player
var gameStarted = false

func initGame() {
	toGuessNumber = rand.UintN(100)
	gameStarted = true
	fmt.Println("Number to guess:", toGuessNumber)
}

func removePlayer(name string) {
	for i, player := range players {
		if player.Name == name {
			players = append(players[:i], players[i+1:]...)
			return
		}
	}
}

func playerIsPlaying(name string) bool {
	for _, player := range players {
		if player.Name == name {
			return true
		}
	}
	return false
}

func validateResponse(guess uint) GuessResult {
	if guess > toGuessNumber {
		return GuessResult{false, "Too high! Try a smaller number."}
	}
	if guess < toGuessNumber {
		return GuessResult{false, "Too low! Try a larger number."}
	}
	return GuessResult{true, "Correct! You won!"}
}
