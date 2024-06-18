package errordetails

import (
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

type ErrorDetails struct {
	ErrorMessage string    `json:"error_message"`
	ErrorType    string    `json:"error_type"`
	File         string    `json:"file"`
	Line         int       `json:"line"`
	Function     string    `json:"function"`
	Timestamp    time.Time `json:"timestamp"`
	Context      []Field   `json:"context"`
	StackTrace   string    `json:"stack_trace"`
	err          error
}

// field struct for contextual information
type Field struct {
	Key string `json:"key"`
	Val string `json:"value"`
}

// NewErrorDetails creates a new ErrorDetails instance and captures current context
func NewErrorDetails(err error) *ErrorDetails {
	pc, file, line, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(pc)
	stackBuf := make([]byte, 1024)
	stackLen := runtime.Stack(stackBuf, false)
	stackTrace := strings.TrimSpace(string(stackBuf[:stackLen]))

	return &ErrorDetails{
		ErrorMessage: err.Error(),
		ErrorType:    fmt.Sprintf("%T", err),
		File:         file,
		Line:         line,
		Function:     fn.Name(),
		Timestamp:    time.Now(),
		StackTrace:   stackTrace,
		err:          err,
	}
}

// Error implements the error interface
func (e *ErrorDetails) Error() string {
	return e.err.Error()
}

// Unwrap returns the underlying error
func (e *ErrorDetails) Unwrap() error {
	return e.err
}

// Str adds a string key-value pair to the context
func (e *ErrorDetails) Str(key, val string) *ErrorDetails {
	e.Context = append(e.Context, Field{Key: key, Val: val})
	return e
}

// Int adds an integer key-value pair to the context
func (e *ErrorDetails) Int(key string, val int) *ErrorDetails {
	e.Context = append(e.Context, Field{Key: key, Val: fmt.Sprintf("%d", val)})
	return e
}

// Msg sets the error message
func (e *ErrorDetails) Msg(message string) *ErrorDetails {
	e.ErrorMessage = message
	return e
}

// ToClientError transforms ErrorDetails to a ClientError
func (e *ErrorDetails) ToClientError() *ClientError {
	return &ClientError{
		Message: e.ErrorMessage,
	}
}

// ClientError struct to return error message to the client
type ClientError struct {
	Message string `json:"message"`
}

// MarshalZerologObject implements the zerolog.LogObjectMarshaler interface
func (e *ErrorDetails) MarshalZerologObject(event *zerolog.Event) {
	event.Str("message", e.ErrorMessage)
	event.Str("file", e.File)
	event.Str("method", e.Function)
	event.Int("line", e.Line)
	event.Str("trace", e.StackTrace)
	event.Str("error_type", e.ErrorType)

	for _, f := range e.Context {
		event.Str(f.Key, f.Val)
	}

	if e.err != nil {
		event.Str("error", e.Error())
	}
}
