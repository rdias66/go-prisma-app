package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not load environmental variables")
	}

	fmt.Print("Server will be running on Port: " + os.Getenv("PORT"))
}
