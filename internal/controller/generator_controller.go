package controller

import (
	"fmt"
	"github.com/superhacker2002/parentheses/internal/parentheses"
	"net/http"
	"strconv"
)

func SetRoutes() {
	http.HandleFunc("/generate", GenerateHandler)
}

func GenerateHandler(w http.ResponseWriter, r *http.Request) {
	length, err := strconv.Atoi(r.URL.Query().Get("n"))
	if err != nil {
		http.Error(w, "invalid length parameter", http.StatusBadRequest)
		return
	}
	if length <= 0 {
		http.Error(w, "length must be greater than 0", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, parentheses.Generate(length))
}
