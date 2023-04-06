package main

import (
	"github.com/superhacker2002/parentheses/internal/client"
	reporter "github.com/superhacker2002/parentheses/internal/controller/reporter"
	"log"
)

func main() {
	requester := client.NewClient("http://localhost:8080")
	reporter := reporter.Reporter{Client: requester}
	if err := reporter.Do(); err != nil {
		log.Fatal(err)
	}
}