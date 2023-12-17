package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	port := os.Getenv("PORT")

	if port == "" {
		// Log and terminate our program
		log.Fatal("PORT env variable missing")
	}

	router := chi.NewRouter()

	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	// Start serving an http server
	fmt.Printf("Server starting on port %v\n", port)
	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

}
