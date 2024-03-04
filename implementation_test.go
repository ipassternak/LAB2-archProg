package lab2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostfixToPrefix(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
		errMsg   string
	}{
		{
			name:     "Simple expression",
			input:    "4 2 - 3 * 5 +",
			expected: "+ * - 4 2 3 5",
			errMsg:   "",
		},
		{
			name:     "Complex expression",
			input:    "2 3 4 * + 5 / 6 ^",
			expected: "^ / + 2 * 3 4 5 6",
			errMsg:   "",
		},
		{
			name:     "Invalid expression: insufficient operands",
			input:    "4 2 - * 5 +",
			expected: "",
			errMsg:   "invalid postfix expression: insufficient operands",
		},
		{
			name:     "Invalid expression: too many operands",
			input:    "4 2 - 3 * 5 + 6",
			expected: "",
			errMsg:   "invalid postfix expression: too many operands",
		},
		{
			name:     "Invalid expression: empty input",
			input:    "",
			expected: "",
			errMsg:   "invalid postfix expression: empty input",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := PostfixToPrefix(tc.input)
			if tc.errMsg != "" {
				assert.Error(t, err)
				assert.Equal(t, tc.errMsg, err.Error())
				assert.Equal(t, tc.expected, result)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, result)
			}
		})
	}
}
