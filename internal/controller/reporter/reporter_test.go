package reporter

import (
	"errors"
	"github.com/superhacker2002/parentheses/internal/parentheses"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockClient struct {
	res string
	err error
}

func (m mockClient) Generate(length uint) (string, error) {
	return m.res, m.err
}

func TestDo(t *testing.T) {
	testCases := []struct {
		name        string
		res         string
		err         error
		expectedErr error
	}{
		{
			name:        "Successful report",
			res:         "()()()",
			err:         nil,
			expectedErr: nil,
		},
		{
			name:        "Error generating string",
			res:         "",
			err:         errors.New("error generating string"),
			expectedErr: errors.New("error generating string"),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c := New(parentheses.New(), mockClient{res: tc.res, err: tc.err})
			err := c.Do()
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}
