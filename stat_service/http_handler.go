package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type AuthResponseBody struct {
	Token string `json:"token"`
}

func SetupHandlers() {
	router := chi.NewRouter()

	router.Get("/stats", GetStatsHandler)

	http.Handle("/", router)
}

func GetStatsHandler(w http.ResponseWriter, req *http.Request) {
	events, err := statDB.GetEvents()
	if err != nil {
		log.Printf("Get stats from db failed, err: %v", err)
		return
	}

	respBody, err := json.Marshal(events)
	if err != nil {
		log.Default().Printf("Response json marshal error: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Write(respBody)
}
