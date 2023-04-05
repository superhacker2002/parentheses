package controller

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGenerateHandler(t *testing.T) {
	testCases := []struct {
		name     string
		url      string
		status   int
		expected string
	}{
		{
			name:     "Correct length",
			url:      "/generate?n=10",
			status:   http.StatusOK,
			expected: ")])){({[}(",
		},
		{
			name:     "Zero length",
			url:      "/generate?n=0",
			status:   http.StatusBadRequest,
			expected: "length must be greater than 0\n",
		},
		{
			name:     "Negative length",
			url:      "/generate?n=-1",
			status:   http.StatusBadRequest,
			expected: "length must be greater than 0\n",
		},
		{
			name:     "Not a number length",
			url:      "/generate?n=invalid",
			status:   http.StatusBadRequest,
			expected: "invalid length parameter: invalid\n",
		},
		{
			name:     "No length",
			url:      "/generate?n=",
			status:   http.StatusBadRequest,
			expected: "invalid length parameter: \n",
		},
		{
			name:     "No length parameter",
			url:      "/generate?",
			status:   http.StatusBadRequest,
			expected: "invalid length parameter: \n",
		},
	}
	rand.Seed(1)

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			request, _ := http.NewRequest(http.MethodGet, testCase.url, nil)
			response := httptest.NewRecorder()
			generateHandler(response, request)
			assert.Equal(t, testCase.status, response.Code, "generateHandler returned wrong status code")

			assert.Equal(t, "text/plain; charset=utf-8", response.Header().Get("Content-Type"),
				"generateHandler returned wrong content type")

			assert.Equal(t, testCase.expected, response.Body.String())
		})
	}
}
