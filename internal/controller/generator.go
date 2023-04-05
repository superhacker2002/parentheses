package controller

import (
	"fmt"
	"github.com/superhacker2002/parentheses/internal/parentheses"
	"net/http"
	"strconv"
)

func SetRoutes() {
	http.HandleFunc("/generate", generateHandler)
}

func generateHandler(w http.ResponseWriter, r *http.Request) {
	lengthParameter := r.URL.Query().Get("n")
	length, err := strconv.Atoi(lengthParameter)
	if err != nil {
		http.Error(w, "invalid length parameter: "+lengthParameter, http.StatusBadRequest)
		return
	}
	if length <= 0 {
		http.Error(w, "length must be greater than 0", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, parentheses.GenerateRandomSequence(uint(length)))
}
