package main

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
)

type AuthResponseBody struct {
	Token string `json:"token"`
}

func SetupHandlers() {
	http.HandleFunc("/register", RegisterHandler)
	http.HandleFunc("/auth", AuthHandler)
	http.HandleFunc("/update", UpdateHandler)
}

func RegisterHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if !req.URL.Query().Has("login") || !req.URL.Query().Has("password") {
		http.Error(w, "Request data is invalid", http.StatusBadRequest)
		return
	}

	login := req.URL.Query().Get("login")
	tempHash := sha256.Sum256([]byte(req.URL.Query().Get("password")))
	passwordHash := base64.URLEncoding.EncodeToString(tempHash[:])

	if clientDB.LookupLogin(login) {
		http.Error(w, "Login already used", http.StatusForbidden)
		return
	}

	err := clientDB.AddUser(User{Login: login, PasswordHash: passwordHash})
	if err != nil {
		log.Default().Printf("Add user error: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	token, err := CreateToken(login)
	if err != nil {
		log.Default().Printf("Create jwt token error: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	respBody, err := json.Marshal(AuthResponseBody{Token: token})
	if err != nil {
		log.Default().Printf("Response json marshal error: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Write(respBody)
}

func AuthHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if !req.URL.Query().Has("login") || !req.URL.Query().Has("password") {
		http.Error(w, "Request data is invalid", http.StatusBadRequest)
		return
	}

	login := req.URL.Query().Get("login")
	tempHash := sha256.Sum256([]byte(req.URL.Query().Get("password")))
	passwordHash := base64.URLEncoding.EncodeToString(tempHash[:])

	if !clientDB.AuthUser(User{Login: login, PasswordHash: passwordHash}) {
		http.Error(w, "Login or password is incorrent", http.StatusNotFound)
		return
	}

	token, err := CreateToken(login)
	if err != nil {
		log.Default().Printf("Create jwt token error: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	respBody, err := json.Marshal(AuthResponseBody{Token: token})
	if err != nil {
		log.Default().Printf("Response json marshal error: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Write(respBody)
}

func UpdateHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != "PUT" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if !req.URL.Query().Has("token") {
		http.Error(w, "Request data is invalid", http.StatusBadRequest)
		return
	}
	token := req.URL.Query().Get("token")
	login, err := DecryptToken(token)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if !clientDB.LookupLogin(login) {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	userData, err := clientDB.LoadUserData(login)
	if err != nil {
		log.Default().Printf("Load user data error: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	decoder := json.NewDecoder(req.Body)
	err = decoder.Decode(&userData)
	if err != nil {
		http.Error(w, "Request data is invalid", http.StatusBadRequest)
		return
	}

	userData.Login = login

	err = clientDB.UpdateUserData(userData)
	if err != nil {
		log.Default().Printf("Update error: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
