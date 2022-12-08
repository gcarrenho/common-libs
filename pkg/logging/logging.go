// This package contains the go struct definition of each entity that is part of the domain problem and can be used across the application.
package logging

import (
	"strconv"
	"time"
)

/*
	index: <project>-<environment> (ie stratus-dev, largeformat-dev)
	label_app: service name (ie udc-uploader)
	time_stamp: time at which the event took place
	level: log level (debug, info, warn, error)
	http_method: HTTP request/response method type
	path: HTTP URL/path
	status_code: HTTP request/response status code
	request_id
	message: free text to provide more context to the event
	remote_ip: client IP who initiated the request
	content_type: HTTP content-type
	latency: duration of the event (in nanoseconds)
	container_name
	pod
	cluster_name
	cluster_region
*/
type Logging struct {
	Index         *string   `json:"index,omitempty"`
	LabelApp      *string   `json:"label_app,omitempty"`
	HttpMethod    *string   `json:"http_method,omitempty"`
	Path          *string   `json:"path,omitempty"`
	StatusCode    *int64    `json:"status_code,omitempty"`
	RequestID     *string   `json:"request_id,omitempty"`
	Message       string    `json:"-"`
	RemoteIP      *string   `json:"remote_ip,omitempty"`
	ContentType   *string   `json:"content_type,omitempty"`
	Latency       *string   `json:"latency,omitempty"`
	ContainerName *string   `json:"container_name,omitempty"`
	Pod           *string   `json:"pod,omitempty"`
	ClusterName   *string   `json:"cluster_name,omitempty"`
	ClusterRegion *string   `json:"cluster_region,omitempty"`
	StartTime     time.Time `json:"-"`
}

func InitOurLogging(project *string, method *string, uri *string, clientIP *string, contentType *string, labelApp *string) *Logging {
	return &Logging{
		Index:       project,
		HttpMethod:  method,
		Path:        uri,
		RemoteIP:    clientIP,
		ContentType: contentType,
		LabelApp:    labelApp,
		StartTime:   time.Now(),
	}
}

func (l *Logging) SetLogging(message string, statusCode *int64) {
	latency := strconv.FormatInt(time.Since(l.StartTime).Milliseconds(), 10)
	l.Latency = &latency
	l.Message = message
	l.StatusCode = statusCode
}
