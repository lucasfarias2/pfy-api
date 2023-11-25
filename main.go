package main

import (
	"log"
	"net/http"
	"pfy-api/db"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	r := chi.NewRouter()

	db.ConnectDatabase()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	log.Println("Running on localhost:8080")
	http.ListenAndServe(":8080", r)
}
