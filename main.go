package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"task-session-1/router"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system env")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.RegisterRoutes()

	fmt.Println("Server running at http://localhost:" + port)
	http.ListenAndServe(":"+port, nil)
}