// This package contains the go struct definition of each entity that is part of the domain problem and can be used across the application.
package logging

import (
	"time"

	"github.com/rs/zerolog"
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
	StatusCode    *int      `json:"status_code,omitempty"`
	RequestID     *string   `json:"request_id,omitempty"`
	Message       string    `json:"message"`
	RemoteIP      *string   `json:"remote_ip,omitempty"`
	ContentType   *string   `json:"content_type,omitempty"`
	Latency       *string   `json:"latency,omitempty"`
	ContainerName *string   `json:"container_name,omitempty"`
	Pod           *string   `json:"pod,omitempty"`
	ClusterName   *string   `json:"cluster_name,omitempty"`
	ClusterRegion *string   `json:"cluster_region,omitempty"`
	StartTime     time.Time `json:"start_time"`
}

// MarshalZerologObject implements the zerolog.LogObjectMarshaler interface
func (l *Logging) MarshalZerologObject(e *zerolog.Event) {
	if l.Index != nil {
		e.Str("index", *l.Index)
	}
	if l.LabelApp != nil {
		e.Str("label_app", *l.LabelApp)
	}
	if l.HttpMethod != nil {
		e.Str("http_method", *l.HttpMethod)
	}
	if l.Path != nil {
		e.Str("path", *l.Path)
	}
	if l.StatusCode != nil {
		e.Int("status_code", *l.StatusCode)
	}
	if l.RequestID != nil {
		e.Str("request_id", *l.RequestID)
	}
	if l.RemoteIP != nil {
		e.Str("remote_ip", *l.RemoteIP)
	}
	if l.ContentType != nil {
		e.Str("content_type", *l.ContentType)
	}
	if l.Latency != nil {
		e.Str("latency", *l.Latency)
	}
	if l.ContainerName != nil {
		e.Str("container_name", *l.ContainerName)
	}
	if l.Pod != nil {
		e.Str("pod", *l.Pod)
	}
	if l.ClusterName != nil {
		e.Str("cluster_name", *l.ClusterName)
	}
	if l.ClusterRegion != nil {
		e.Str("cluster_region", *l.ClusterRegion)
	}
	e.Str("message", l.Message)
	e.Time("start_time", l.StartTime)
}

// NewLogging creates a new Logging instance
func NewLogging(message string) *Logging {
	return &Logging{
		Message:   message,
		StartTime: time.Now(),
	}
}
