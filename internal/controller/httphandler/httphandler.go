package httphandler

import (
	"fmt"
	"net/http"
	"strconv"
)

type generator interface {
	GenerateRandom(uint) string
}

type httpHandler struct {
	parentheses generator
}

func New(generator generator) httpHandler {
	handler := httpHandler{parentheses: generator}
	handler.setRoutes()
	return handler
}

func (h httpHandler) setRoutes() {
	http.HandleFunc("/generate", h.generateHandler)
}

func (h httpHandler) generateHandler(w http.ResponseWriter, r *http.Request) {
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
	fmt.Fprintln(w, h.parentheses.GenerateRandom(uint(length)))
}
