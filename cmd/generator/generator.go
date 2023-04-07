package main

import (
	"github.com/superhacker2002/parentheses/internal/controller/httphandler"
	"github.com/superhacker2002/parentheses/internal/parentheses"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	handler := httphandler.New(parentheses.New())
	handler.SetRoutes()
	rand.Seed(time.Now().UnixNano())
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
