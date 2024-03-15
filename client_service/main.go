package main

import (
	"net/http"
)

const (
	DB_URL_ENV       = "POSTGRES_URL"
	PRIVATE_KEY_PATH = "PRIVATE_KEY_PATH"
	PUBLIC_KEY_PATH  = "PUBLIC_KEY_PATH"
)

func main() {
	SetupAuthData()
	SetupDB()
	SetupHandlers()

	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		panic("Server falled")
	}
}
