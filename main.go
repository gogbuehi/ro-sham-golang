package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Ro Sham Golang")
	http.HandleFunc("/", IndexHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Printf("Open http://localhost:%s in the browser", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

type PlayResult struct {
	ComputerPlayValue int
	UserPlayValue     int
	PlayOutcome       int
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	queryData := r.URL.Query()
	userPlayValue, queryErr := strconv.Atoi(queryData.Get("play"))
	if queryErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if userPlayValue < 0 || userPlayValue > 2 {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	playValue := RockPaperScissor()

	gameOutcome := PlayAgainst(userPlayValue, playValue)
	gameResult := PlayResult{
		ComputerPlayValue: playValue,
		UserPlayValue:     userPlayValue,
		PlayOutcome:       gameOutcome,
	}
	jsonGame, _ := json.Marshal(gameResult)
	w.Header().Set("Content-Type", "application/json")
	_, err := fmt.Fprint(w, string(jsonGame))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
