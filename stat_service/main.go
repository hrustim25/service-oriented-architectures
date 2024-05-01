package main

import (
	"log"
	"net/http"
)

func main() {
	SetupDB()
	SetupHandlers()

	SetupAndStartStatMessageBrokerConsumer()

	log.Default().Println("Starting stat server...")

	err := http.ListenAndServe("0.0.0.0:12345", nil)
	if err != nil {
		panic("Server falled")
	}
}
