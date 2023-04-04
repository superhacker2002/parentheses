package parentheses

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsBalanced(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Empty string",
			input:    "",
			expected: true,
		},
		{
			name:     "Balanced string",
			input:    "{[{}]()}",
			expected: true,
		},
		{
			name:     "Unbalanced string",
			input:    "[(])",
			expected: false,
		},
		{
			name:     "String with only opening brackets",
			input:    "{{[[((",
			expected: false,
		},
		{
			name:     "String with only closing brackets",
			input:    "))]}",
			expected: false,
		},
		{
			name:     "Balanced string with other characters",
			input:    "a(b[c{d}e]f)g",
			expected: true,
		},
		{
			name:     "Unbalanced string with other characters",
			input:    "a(b[c{d}e]f)g))){{]]",
			expected: false,
		},
		{
			name:     "String without parenthesis",
			input:    "abcdef",
			expected: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			isBalanced := IsBalanced(testCase.input)
			assert.Equal(t, isBalanced, testCase.expected, "IsBalanced(%q) ", testCase.input)
		})
	}
}
