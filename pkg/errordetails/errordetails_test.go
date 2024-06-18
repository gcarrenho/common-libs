package errordetails

/*
import (
	"errors"
	"fmt"
	"testing"

	"github.com/huandu/go-assert"
)

func TestNewErrorDetails(t *testing.T) {
	baseErr := errors.New("an error occurred")
	customErr := NewErrorDetails(baseErr)

	if customErr.err != baseErr {
		t.Errorf("expected %v, got %v", baseErr, customErr.err)
	}

	/*if customErr.Message == "" {
		t.Errorf("expected title to be initialized, got empty")
	}
}

func TestErrorDetails_Str(t *testing.T) {
	baseErr := errors.New("an error occurred")
	errorDetails := NewErrorDetails(baseErr)

	// Add fields using Str method
	errorDetails.Str("OrderItemID", "123").
		Str("PresetID", "preset123").
		Str("SettingsJSON", `{"key":"value"}`)

	// Check if the fields were added correctly
	expectedFields := []field{
		{key: "OrderItemID", val: "123"},
		{key: "PresetID", val: "preset123"},
		{key: "SettingsJSON", val: `{"key":"value"}`},
	}

	if len(errorDetails.fields) != len(expectedFields) {
		t.Fatalf("expected %d fields, got %d", len(expectedFields), len(errorDetails.fields))
	}

	for i, field := range errorDetails.fields {
		if field.key != expectedFields[i].key || field.val != expectedFields[i].val {
			t.Errorf("expected field %d to be %v, got %v", i, expectedFields[i], field)
		}
	}
}

func TestErrorDetails_Int(t *testing.T) {
	baseErr := errors.New("an error occurred")
	errorDetails := NewErrorDetails(baseErr)

	// Add integer fields using Int method
	errorDetails.Int("OrderItemID", 123).
		Int("ErrorCode", 404)

	// Check if the fields were added correctly
	expectedFields := []field{
		{key: "OrderItemID", val: "123"},
		{key: "ErrorCode", val: "404"},
	}

	if len(errorDetails.fields) != len(expectedFields) {
		t.Fatalf("expected %d fields, got %d", len(expectedFields), len(errorDetails.fields))
	}

	for i, field := range errorDetails.fields {
		if field.key != expectedFields[i].key || field.val != expectedFields[i].val {
			t.Errorf("expected field %d to be %v, got %v", i, expectedFields[i], field)
		}
	}
}

/*func TestErrorDetails_Error(t *testing.T) {
	baseErr := errors.New("an error occurred")
	customErr := NewErrorDetails(baseErr).
		Str("OrderItemID", "123").
		Str("PresetID", "preset123").
		Str("SettingsJSON", `{"key":"value"}`).Msg("message error example")

	expectedMessage := "message error example | OrderItemID: 123 | PresetID: preset123 | SettingsJSON: {\"key\":\"value\"} --> an error occurred"

	if customErr.Error() != expectedMessage {
		t.Errorf("expected %s, got %s", expectedMessage, customErr.Error())
	}
}

func TestErrorDetails_Error(t *testing.T) {
	baseErr := errors.New("base error")
	file := "file.go"
	line := 42

	tests := []struct {
		name           string
		errorDetails   *ErrorDetails
		expectedOutput string
	}{
		{
			name: "Base Error",
			errorDetails: &ErrorDetails{
				Message: "An error occurred",
				fields:  nil,
				err:     baseErr,
				file:    file,
				line:    line,
			},
			expectedOutput: fmt.Sprintf("%s:%d: An error occurred --> %s", file, line, baseErr.Error()),
		},
		{
			name: "With Fields",
			errorDetails: &ErrorDetails{
				Message: "An error occurred",
				fields: []field{
					{key: "OrderNumber", val: "12345"},
				},
				err:  baseErr,
				file: file,
				line: line,
			},
			expectedOutput: fmt.Sprintf("%s:%d: An error occurred | OrderNumber: 12345 --> %s", file, line, baseErr.Error()),
		},
		{
			name: "Without Base Error",
			errorDetails: &ErrorDetails{
				Message: "An error occurred",
				fields:  nil,
				err:     nil,
				file:    file,
				line:    line,
			},
			expectedOutput: fmt.Sprintf("%s:%d: An error occurred", file, line),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expectedOutput, tt.errorDetails.Error())
		})
	}
}
*/
