package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello world!")

	godotenv.Load()

	portString := os.Getenv("PORT")

	if portString == "" {
		// Log and terminate our program
		log.Fatal("PORT env variable missing")
	}

	fmt.Println("Port:", portString)
}
