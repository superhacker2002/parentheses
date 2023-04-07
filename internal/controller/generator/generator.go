package generator

import (
	"fmt"
	"net/http"
	"strconv"
)

type parenthesesGenerator interface {
	GenerateRandom(uint) string
}

type generator struct {
	parentheses parenthesesGenerator
}

func New(gen parenthesesGenerator) generator {
	return generator{parentheses: gen}
}

func (g generator) SetRoutes() {
	http.HandleFunc("/generate", g.generateHandler)
}

func (g generator) generateHandler(w http.ResponseWriter, r *http.Request) {
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
	fmt.Fprintln(w, g.parentheses.GenerateRandom(uint(length)))
}
