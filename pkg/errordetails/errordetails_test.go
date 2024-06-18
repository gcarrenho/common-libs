package errordetails

import (
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestNewErrorDetails(t *testing.T) {
	err := errors.New("test error")
	details := NewErrorDetails(err)

	// Check basic fields
	assert.Equal(t, "test error", details.ErrorMessage)
	assert.Equal(t, "*errors.errorString", details.ErrorType)
	assert.NotEmpty(t, details.File)
	assert.NotEmpty(t, details.Function)
	assert.True(t, details.Line > 0)
	assert.NotEmpty(t, details.StackTrace)
	assert.Equal(t, err, details.err)

	// Check timestamp
	assert.WithinDuration(t, time.Now(), details.Timestamp, time.Second)
}

func TestErrorDetails_Str(t *testing.T) {
	err := errors.New("test error")
	details := NewErrorDetails(err)

	details.Str("key1", "value1").Str("key2", "value2")

	expectedContext := []Field{
		{Key: "key1", Val: "value1"},
		{Key: "key2", Val: "value2"},
	}

	assert.Equal(t, expectedContext, details.Context)
}

func TestErrorDetails_Int(t *testing.T) {
	err := errors.New("test error")
	details := NewErrorDetails(err)

	details.Int("key1", 1).Int("key2", 2)

	expectedContext := []Field{
		{Key: "key1", Val: "1"},
		{Key: "key2", Val: "2"},
	}

	assert.Equal(t, expectedContext, details.Context)
}

func TestErrorDetails_Msg(t *testing.T) {
	err := errors.New("test error")
	details := NewErrorDetails(err)

	details.Msg("new message")

	assert.Equal(t, "new message", details.ErrorMessage)
}

func TestErrorDetails_ToClientError(t *testing.T) {
	err := errors.New("test error")
	details := NewErrorDetails(err)

	clientError := details.ToClientError()

	assert.Equal(t, details.ErrorMessage, clientError.Message)
}

func TestErrorDetails_Error(t *testing.T) {
	err := errors.New("test error")
	details := NewErrorDetails(err)

	assert.Equal(t, err.Error(), details.Error())
}

func TestErrorDetails_Unwrap(t *testing.T) {
	err := errors.New("test error")
	details := NewErrorDetails(err)

	assert.Equal(t, err, details.Unwrap())
}

func TestErrorDetails_MarshalZerologObject(t *testing.T) {
	err := errors.New("test error")
	details := NewErrorDetails(err)
	details.Str("context_key", "context_value")

	output := strings.Builder{}
	logger := zerolog.New(&output)
	logger.Error().Object("error_details", details).Msg("")

	logOutput := output.String()
	assert.Contains(t, logOutput, `"message":"test error"`)
	assert.Contains(t, logOutput, `"file"`)
	assert.Contains(t, logOutput, `"method"`)
	assert.Contains(t, logOutput, `"line"`)
	assert.Contains(t, logOutput, `"trace"`)
	assert.Contains(t, logOutput, `"error_type":"*errors.errorString"`)
	assert.Contains(t, logOutput, `"context_key":"context_value"`)
}
