package client

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/generate", r.URL.Path)
		assert.Equal(t, "10", r.URL.Query().Get("n"))
		fmt.Fprint(w, "random-string")
	}))
	defer server.Close()

	c := NewClient(server.URL)
	result, err := c.Generate(10)

	assert.NoError(t, err)
	assert.Equal(t, "random-string", result)
}
