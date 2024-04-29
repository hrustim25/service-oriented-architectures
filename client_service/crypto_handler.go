package main

import (
	"crypto/rsa"
	"log"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type AuthData struct {
	pubKey  *rsa.PublicKey
	privKey *rsa.PrivateKey
}

var serviceKeys AuthData

type LoginCustomClaims struct {
	Login string `json:"login"`
	jwt.RegisteredClaims
}

func CreateToken(login string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{"login": login})
	tokenString, err := token.SignedString(serviceKeys.privKey)
	if err != nil {
		log.Default().Printf("Sign token error: %v", err)
		return "", err
	}
	return tokenString, nil
}

func DecryptToken(token string) (string, error) {
	claims := LoginCustomClaims{}
	_, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return serviceKeys.pubKey, nil
	})
	if err != nil {
		return "", err
	}
	return claims.Login, nil
}

func SetupAuthData() {
	if _, ok := os.LookupEnv(PUBLIC_KEY_PATH); !ok {
		log.Fatalf("'%v' env var not found", PUBLIC_KEY_PATH)
	}
	if _, ok := os.LookupEnv(PRIVATE_KEY_PATH); !ok {
		log.Fatalf("'%v' env var not found", PRIVATE_KEY_PATH)
	}
	publicKeyData, err := os.ReadFile(os.Getenv(PUBLIC_KEY_PATH))
	if err != nil {
		log.Fatal("Failed to find public key")
	}
	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyData)
	if err != nil {
		log.Fatal("Failed to parse public key")
	}

	privateKeyData, err := os.ReadFile(os.Getenv(PRIVATE_KEY_PATH))
	if err != nil {
		log.Fatal("Failed to find private key")
	}
	privKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyData)
	if err != nil {
		log.Fatal("Failed to parse private key")
	}

	serviceKeys = AuthData{pubKey: pubKey, privKey: privKey}
}
