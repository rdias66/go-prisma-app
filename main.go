package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not load environmental variables")
	}

	fmt.Println("Server will be running on Port: " + os.Getenv("PORT"))

	server := &http.Server{
		Addr:           ":" + os.Getenv("PORT"),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	server_err := server.ListenAndServe()

	if server_err != nil {
		panic(server_err)
	}
}
