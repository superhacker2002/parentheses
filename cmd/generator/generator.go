package main

import (
	"github.com/superhacker2002/parentheses/internal/controller"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	controller.SetRoutes()
	rand.Seed(time.Now().UnixNano())
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
