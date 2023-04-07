package main

import (
	"github.com/superhacker2002/parentheses/internal/client"
	"github.com/superhacker2002/parentheses/internal/controller/reporter"
	"log"
)

func main() {
	client := client.New("http://localhost:8080")
	reporter := reporter.New(client)
	if err := reporter.Do(); err != nil {
		log.Fatal(err)
	}
}
