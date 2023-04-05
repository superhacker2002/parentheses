package main

import (
	"github.com/superhacker2002/parentheses/internal/controller"
	"log"
	"net/http"
)

func main() {
	controller.SetRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
