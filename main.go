package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"shortener-url/api"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var port string

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	} else {
		port = os.Getenv("PORT")
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/short", api.CreateShort).Methods("POST")
	router.HandleFunc("/{id}", api.Redirect).Methods("GET")
	fmt.Println("API Running on Port", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}
