package logging

import (
	"bytes"
	"encoding/json"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

// TestMarshalZerologObject tests the MarshalZerologObject function of Logging struct
func TestMarshalZerologObject(t *testing.T) {
	index := "test_index"
	labelApp := "test_app"
	httpMethod := "GET"
	path := "/test/path"
	statusCode := 200
	requestID := "test_request_id"
	remoteIP := "127.0.0.1"
	contentType := "application/json"
	latency := "100ms"
	containerName := "test_container"
	pod := "test_pod"
	clusterName := "test_cluster"
	clusterRegion := "us-west-1"
	message := "test message"
	startTime := time.Now()

	logging := &Logging{
		Index:         &index,
		LabelApp:      &labelApp,
		HttpMethod:    &httpMethod,
		Path:          &path,
		StatusCode:    &statusCode,
		RequestID:     &requestID,
		RemoteIP:      &remoteIP,
		ContentType:   &contentType,
		Latency:       &latency,
		ContainerName: &containerName,
		Pod:           &pod,
		ClusterName:   &clusterName,
		ClusterRegion: &clusterRegion,
		Message:       message,
		StartTime:     startTime,
	}

	var buf bytes.Buffer
	logger := zerolog.New(&buf).With().Timestamp().Logger()

	// Log the message
	logger.Log().Object("logging", logging).Send()

	// Log output should be valid JSON
	logOutput := buf.String()

	t.Logf("Logged output: %s", logOutput)

	// Parse the logged output
	var result map[string]interface{}
	err := json.Unmarshal([]byte(logOutput), &result)

	assert.NoError(t, err, "Error parsing logged output")

	// Assert the fields
	loggingFields := result["logging"].(map[string]interface{})
	assert.Equal(t, index, loggingFields["index"].(string))
	assert.Equal(t, labelApp, loggingFields["label_app"].(string))
	assert.Equal(t, httpMethod, loggingFields["http_method"].(string))
	assert.Equal(t, path, loggingFields["path"].(string))
	assert.Equal(t, float64(statusCode), loggingFields["status_code"].(float64)) // JSON numbers are always float64
	assert.Equal(t, requestID, loggingFields["request_id"].(string))
	assert.Equal(t, remoteIP, loggingFields["remote_ip"].(string))
	assert.Equal(t, contentType, loggingFields["content_type"].(string))
	assert.Equal(t, latency, loggingFields["latency"].(string))
	assert.Equal(t, containerName, loggingFields["container_name"].(string))
	assert.Equal(t, pod, loggingFields["pod"].(string))
	assert.Equal(t, clusterName, loggingFields["cluster_name"].(string))
	assert.Equal(t, clusterRegion, loggingFields["cluster_region"].(string))
	assert.Equal(t, startTime.Format(time.RFC3339), loggingFields["start_time"].(string))
}

func TestMarshalZerologObject_Empty(t *testing.T) {
	message := "test message"
	startTime := time.Now()

	logging := &Logging{
		Message:   message,
		StartTime: startTime,
	}

	// Create a zerolog logger with a buffer writer
	var buf bytes.Buffer
	logger := zerolog.New(&buf).With().Timestamp().Logger()

	// Log the message
	logger.Log().Object("logging", logging).Send()

	// Parse the logged output
	logOutput := buf.Bytes()
	var result map[string]interface{}
	err := json.Unmarshal(logOutput, &result)
	assert.NoError(t, err, "Error parsing logged output")

	// Assert the fields
	loggingFields := result["logging"].(map[string]interface{})
	assert.Nil(t, loggingFields["index"])
	assert.Nil(t, loggingFields["label_app"])
	assert.Nil(t, loggingFields["http_method"])
	assert.Nil(t, loggingFields["path"])
	assert.Nil(t, loggingFields["status_code"])
	assert.Nil(t, loggingFields["request_id"])
	assert.Nil(t, loggingFields["remote_ip"])
	assert.Nil(t, loggingFields["content_type"])
	assert.Nil(t, loggingFields["latency"])
	assert.Nil(t, loggingFields["container_name"])
	assert.Nil(t, loggingFields["pod"])
	assert.Nil(t, loggingFields["cluster_name"])
	assert.Nil(t, loggingFields["cluster_region"])
	assert.Equal(t, message, loggingFields["message"])
	assert.Equal(t, startTime.Format(time.RFC3339), loggingFields["start_time"])
}

func TestNewLogging(t *testing.T) {
	message := "test message"
	logging := NewLogging(message)

	assert.Equal(t, message, logging.Message)
	assert.WithinDuration(t, time.Now(), logging.StartTime, time.Second)
}
