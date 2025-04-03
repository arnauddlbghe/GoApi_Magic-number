package main

import (
	"fmt"
	"net/http"
)

func startServer() {
	fmt.Println("Serveur démarré sur le port 8080...")
	http.ListenAndServe(":8080", nil)
}

func main() {
	http.HandleFunc("/join", joinHandler)
	http.HandleFunc("/leave", leaveHandler)
	http.HandleFunc("/start", startHandler)
	http.HandleFunc("/play", playHandler)
	startServer()
}
