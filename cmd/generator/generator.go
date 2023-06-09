package main

import (
	"github.com/joho/godotenv"
	"github.com/superhacker2002/parentheses/internal/controller/httphandler"
	"github.com/superhacker2002/parentheses/internal/parentheses"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("no .env file found")
	}
}

func main() {
	port, exists := os.LookupEnv("PORT")
	if !exists {
		log.Fatal("no server port provided in .env file")
	}

	httphandler.New(parentheses.New())
	rand.Seed(time.Now().UnixNano())
	log.Fatal(http.ListenAndServe("localhost:"+port, nil))
}
