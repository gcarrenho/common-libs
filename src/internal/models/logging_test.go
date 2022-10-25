package models

import (
	"strconv"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
)

func TestInitOurLogging(t *testing.T) {
	monkey.Patch(time.Now, func() time.Time {
		return time.Date(2022, 12, 06, 20, 34, 58, 651387237, time.UTC)
	})
	test := "test"

	tests := []struct {
		name        string
		project     string
		method      string
		uri         string
		clientIP    string
		contentType string
		labelApp    string
		expected    *Logging
	}{
		{
			name:        "Conversion succesfully",
			project:     "test",
			method:      "test",
			uri:         "test",
			clientIP:    "test",
			contentType: "test",
			labelApp:    "test",
			expected: &Logging{
				Index:       &test,
				LabelApp:    &test,
				Path:        &test,
				RemoteIP:    &test,
				ContentType: &test,
				HttpMethod:  &test,
				StartTime:   time.Date(2022, 12, 06, 20, 34, 58, 651387237, time.UTC),
			},
		},
	}

	for _, tt := range tests {

		result := InitOurLogging(&tt.project, &tt.method, &tt.uri, &tt.clientIP, &tt.contentType, &tt.labelApp)
		assert.Equal(t, tt.expected, result)

	}
}

func TestSetLogging(t *testing.T) {
	test := "test"
	monkey.Patch(time.Now, func() time.Time {
		return time.Date(2022, 12, 06, 20, 34, 58, 651387237, time.UTC)
	})
	var status int64 = 200
	latency := strconv.FormatInt(time.Since(time.Now()).Milliseconds(), 10)
	tests := []struct {
		name        string
		project     string
		method      string
		uri         string
		clientIP    string
		contentType string
		labelApp    string
		status      int64
		expected    *Logging
	}{
		{
			name:        "Conversion succesfully",
			project:     "test",
			method:      "test",
			uri:         "test",
			clientIP:    "test",
			contentType: "test",
			labelApp:    "test",
			expected: &Logging{
				Index:       &test,
				LabelApp:    &test,
				Path:        &test,
				RemoteIP:    &test,
				ContentType: &test,
				HttpMethod:  &test,
				Message:     "test",
				Latency:     &latency,
				StatusCode:  &status,
				StartTime:   time.Date(2022, 12, 06, 20, 34, 58, 651387237, time.UTC),
			},
		},
	}

	for _, tt := range tests {
		result := InitOurLogging(&tt.project, &tt.method, &tt.uri, &tt.clientIP, &tt.contentType, &tt.labelApp)
		result.SetLogging("test", &status)
		assert.Equal(t, tt.expected, result)

	}
}
