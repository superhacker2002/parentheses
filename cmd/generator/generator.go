package main

import (
	"github.com/superhacker2002/parentheses/internal/controller/generator"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	generator.SetRoutes()
	rand.Seed(time.Now().UnixNano())
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
