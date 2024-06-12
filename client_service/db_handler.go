package main

import (
	"context"
	"log"
	"os"
	"time"

	_ "embed"

	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	Login        string
	PasswordHash string
	Name         string `json:"name"`
	Surname      string `json:"surname"`
	Birthdate    string `json:"birthdate"`
	Email        string `json:"email"`
	PhoneNumber  string `json:"phone_number"`
}

type DBHandler struct {
	db *pgxpool.Pool
}

var clientDB DBHandler

//go:embed sql/create_table.sql
var createTableQuery string

//go:embed sql/lookup_login.sql
var lookupLoginQuery string

//go:embed sql/get_user_creds.sql
var getUserCredsQuery string

//go:embed sql/add_user.sql
var addUserQuery string

//go:embed sql/load_user.sql
var loadUserQuery string

//go:embed sql/update_user.sql
var updateUserQuery string

//go:embed sql/get_login_by_id.sql
var getLoginById string

func (h *DBHandler) LookupLogin(login string) *uint64 {
	rows, err := h.db.Query(context.Background(), lookupLoginQuery, login)
	if err != nil {
		return nil
	}
	defer rows.Close()

	if !rows.Next() {
		return nil
	}

	var userID uint64
	err = rows.Scan(&userID)
	if err != nil {
		return nil
	}
	return &userID
}

func (h *DBHandler) AuthUser(user User) bool {
	rows, err := h.db.Query(context.Background(), getUserCredsQuery, user.Login, user.PasswordHash)

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
	_, err := h.db.Exec(context.Background(), addUserQuery, user.Login, user.PasswordHash)
	return err
}

func (h *DBHandler) LoadUserData(login string) (User, error) {
	var name, surname, birthdate, email, phoneNumber string
	err := h.db.QueryRow(context.Background(), loadUserQuery, login).Scan(&name, &surname, &birthdate, &email, &phoneNumber)
	return User{Name: name, Surname: surname, Birthdate: birthdate, Email: email, PhoneNumber: phoneNumber}, err
}

func (h *DBHandler) UpdateUserData(user User) error {
	_, err := h.db.Exec(context.Background(), updateUserQuery,
		user.Name, user.Surname, user.Birthdate, user.Email, user.PhoneNumber, user.Login)
	return err
}

func (h *DBHandler) GetLoginsByIds(userIDs []uint64) ([]string, error) {
	logins := make([]string, 0, len(userIDs))
	for _, userID := range userIDs {
		var login string
		err := h.db.QueryRow(context.Background(), getLoginById, userID).Scan(&login)
		if err != nil {
			return nil, err
		}
		logins = append(logins, login)
	}
	return logins, nil
}

func SetupDB() {
	if _, ok := os.LookupEnv(DB_URL_ENV); !ok {
		log.Fatalf("'%v' env var not found", DB_URL_ENV)
	}

	time.Sleep(5 * time.Second)

	poolConfig, err := pgxpool.ParseConfig(os.Getenv(DB_URL_ENV))
	if err != nil {
		log.Fatalf("Failed to parse DB config, err: %v", err)
	}

	db, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		log.Fatalf("Failed to connect to DB, err: %v", err)
	}

	_, err = db.Exec(context.Background(), createTableQuery)
	if err != nil {
		log.Fatalf("Failed to create table in DB, err: %v", err)
	}

	clientDB = DBHandler{db}
}
