package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringToPointer(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		expected *string
	}{
		{
			name:     "Conversion succesfully",
			value:    "string to test",
			expected: StringToPointer("string to test"),
		},
	}

	for _, tt := range tests {

		result := StringToPointer(tt.value)
		assert.Equal(t, tt.expected, result)

	}
}
