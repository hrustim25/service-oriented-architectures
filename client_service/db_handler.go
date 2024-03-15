package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	Login        string
	PasswordHash string
	Name         string `json:"name"`
	Surname      string `json:"surname"`
	Birthdate    string `json:"birthdate"`
	Email        string `json:"email"`
	PhoneNumber  string `json:"phone-number"`
}

type DBHandler struct {
	db *pgxpool.Pool
}

var clientDB DBHandler

func LoadQuery(filename string) string {
	result, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("file not found")
	}
	return string(result)
}

func (h *DBHandler) LookupLogin(login string) bool {
	rows, err := h.db.Query(context.Background(), "select login from users where login=$1", login)

	if err != nil {
		return false
	}
	defer rows.Close()

	if !rows.Next() {
		return false
	}
	values, _ := rows.Values()
	return len(values) > 0
}

func (h *DBHandler) AuthUser(user User) bool {
	rows, err := h.db.Query(context.Background(), "SELECT login FROM users WHERE login=$1 AND pwd_hash=$2", user.Login, user.PasswordHash)

	if err != nil {
		return false
	}
	defer rows.Close()

	if !rows.Next() {
		return false
	}
	values, _ := rows.Values()
	return len(values) > 0
}

func (h *DBHandler) AddUser(user User) error {
	_, err := h.db.Exec(context.Background(), "INSERT INTO users(login, pwd_hash) VALUES($1, $2)", user.Login, user.PasswordHash)
	return err
}

func (h *DBHandler) LoadUserData(login string) (User, error) {
	var name, surname, birthdate, email, phoneNumber string
	err := h.db.QueryRow(context.Background(), "SELECT name, surname, birthdate, email, phone_number FROM users WHERE login=$1", login).Scan(&name, &surname, &birthdate, &email, &phoneNumber)
	return User{Name: name, Surname: surname, Birthdate: birthdate, Email: email, PhoneNumber: phoneNumber}, err
}

func (h *DBHandler) UpdateUserData(user User) error {
	_, err := h.db.Exec(context.Background(), "UPDATE users SET name=$1, surname=$2, birthdate=$3, email=$4, phone_number=$5 WHERE login=$6",
		user.Name, user.Surname, user.Birthdate, user.Email, user.PhoneNumber, user.Login)
	return err
}

func SetupDB() {
	if _, ok := os.LookupEnv(DB_URL_ENV); !ok {
		log.Fatalf("'%v' env var not found", DB_URL_ENV)
	}

	time.Sleep(5 * time.Second)

	poolConfig, err := pgxpool.ParseConfig(os.Getenv("POSTGRES_URL"))
	if err != nil {
		log.Fatalf("Failed to parse DB config, err: %v", err)
	}

	db, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		log.Fatalf("Failed to connect to DB, err: %v", err)
	}

	clientDB = DBHandler{db}
}