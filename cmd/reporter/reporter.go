package main

import (
	"github.com/joho/godotenv"
	"github.com/superhacker2002/parentheses/internal/client"
	"github.com/superhacker2002/parentheses/internal/controller/reporter"
	"github.com/superhacker2002/parentheses/internal/parentheses"
	"log"
	"os"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("no .env file found")
	}
}

func main() {
	serverURL, exists := os.LookupEnv("SERVER_URL")
	if !exists {
		log.Fatal("no server address provided in .env file")
	}
	client := client.New(serverURL)
	reporter := reporter.New(parentheses.New(), client)
	if err := reporter.Do(); err != nil {
		log.Fatal(err)
	}
}
