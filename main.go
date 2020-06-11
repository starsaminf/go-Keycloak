package main

import (
	"fmt"
	"net/http"
	"time"
	"log"

	"github.com/Nerzal/gocloak/v5"
	"github.com/gorilla/mux"
)

func main() {
	m := mux.NewRouter()
	m.HandleFunc("/register", register)

	server := &http.Server{
		Addr:           ":8080",
		Handler:        m,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}

func register(w http.ResponseWriter, r *http.Request) {
	basePath     := "https://keycloak.instance"
	clientId     := "your-client-id-name"
	clientSecret := "your-client-secret"
	realm        := "your-realm"

	client := gocloak.NewClient(basePath)
	token, err := client.LoginClient(clientId, clientSecret, realm)
	if err != nil {
		panic(err)
	}

	user := gocloak.User{
		Username:      gocloak.StringP("sam"),
		EmailVerified: gocloak.BoolP(true),
		FirstName:     gocloak.StringP("sam"),
		LastName:      gocloak.StringP("sam"),
		Email:         gocloak.StringP("sam@sam.com"),
		Enabled:       gocloak.BoolP(true),
	}

	userID, err := client.CreateUser(token.AccessToken, realm, user)
	if err != nil {
		panic(err)
	}
	fmt.Println(userID)
}
