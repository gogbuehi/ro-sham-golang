package main

import (
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
	gameString := "You (" + PlayValue(userPlayValue) + ") vs RoSham (" + PlayValue(playValue) + ") / Outcome: " + WinnerValue(gameOutcome)
	_, err := fmt.Fprint(w, gameString)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
