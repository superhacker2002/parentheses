package main

import (
	"github.com/superhacker2002/parentheses/internal/client"
	"github.com/superhacker2002/parentheses/internal/controller/reporter"
	"github.com/superhacker2002/parentheses/internal/parentheses"
	"log"
)

func main() {
	client := client.New("http://localhost:8080")
	reporter := reporter.New(parentheses.New(), client)
	if err := reporter.Do(); err != nil {
		log.Fatal(err)
	}
}
