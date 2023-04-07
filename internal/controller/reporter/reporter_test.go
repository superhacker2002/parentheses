package controller

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockClient struct {
	generateFunc func(length uint) (string, error)
}

func (m mockClient) Generate(length uint) (string, error) {
	return m.generateFunc(length)
}

func TestDo(t *testing.T) {
	testCases := []struct {
		name         string
		generateFunc func(length uint) (string, error)
		expectedErr  error
	}{
		{
			name: "Successful report",
			generateFunc: func(length uint) (string, error) {
				return "()()()", nil
			},
			expectedErr: nil,
		},
		{
			name: "Error generating string",
			generateFunc: func(length uint) (string, error) {
				return "", errors.New("error generating string")
			},
			expectedErr: errors.New("error generating string"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c := Reporter{
				Client: mockClient{generateFunc: tc.generateFunc},
			}
			err := c.Do()
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}
