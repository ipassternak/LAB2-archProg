package lab2

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompute(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
		errMsg   string
	}{
		{
			name:     "The output should match the input expression",
			input:    "4 2 - 3 * 5 +",
			expected: "+ * - 4 2 3 5",
			errMsg:   "",
		},
		{
			name:     "Returns an error for the invalid expression",
			input:    "",
			expected: "",
			errMsg:   "invalid postfix expression: empty input",
		},
	}

	for _, tc := range testCases {
		b := strings.Builder{}
		h := &ComputeHandler{
			R: strings.NewReader(tc.input),
			W: &b,
		}
		t.Run(tc.name, func(t *testing.T) {
			err := h.Compute()
			result := b.String()
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

func ExampleComputeHandler_Compute() {
	h := &ComputeHandler{
		R: strings.NewReader("4 2 - 3 * 5 +"),
		W: os.Stdout,
	}

	h.Compute()
	// Output:
	// + * - 4 2 3 5
}
