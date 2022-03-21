package cloudlogging

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"runtime"

	"github.com/pkg/errors"
)

type logEntryLevel string

const (
	levelError logEntryLevel = "error"
	levelInfo  logEntryLevel = "info"
)

type logEntryPayload struct {
	Level          logEntryLevel                  `json:"level,omitempty"`
	ServiceContext *logEntryPayloadServiceContext `json:"serviceContext,omitempty"`
	Message        string                         `json:"message,omitempty"`
	Context        *logEntryPayloadContext        `json:"context,omitempty"`
}

type logEntryPayloadServiceContext struct {
	Service string `json:"service,omitempty"`
	Version string `json:"version,omitempty"`
}

type logEntryPayloadContext struct {
	HTTPRequest *logEntryPayloadContextHTTPRequest
}

type logEntryPayloadContextHTTPRequest struct {
	Method             string `json:"method,omitempty"`
	URL                string `json:"url,omitempty"`
	UserAgent          string `json:"userAgent,omitempty"`
	Referrer           string `json:"referrer,omitempty"`
	RemoteIP           string `json:"remoteIp,omitempty"`
	ResponseStatusCode int    `json:"responseStatusCode,omitempty"`
}

func newErrorLogEntryPayload(
	err error,
) *logEntryPayload {
	type stackTracer interface {
		StackTrace() errors.StackTrace
	}
	message := ""
	errWithStackTrace, ok := errors.Cause(err).(stackTracer)
	if !ok {
		buf := make([]byte, 1<<16)
		size := runtime.Stack(buf, false)
		message = fmt.Sprintf(string(buf[:size]))
	} else {
		message = fmt.Sprintf("%+v", errWithStackTrace.StackTrace())
	}
	payload := logEntryPayload{
		Level: levelError,
		ServiceContext: &logEntryPayloadServiceContext{
			Service: serviceName,
			Version: os.Getenv("COMMIT_SHA"),
		},
		Message: message,
	}
	return &payload
}

var serviceName = "blog"

func Error(
	err error,
) {
	payload := newErrorLogEntryPayload(err)
	payloadBytes, _ := json.Marshal(payload)
	fmt.Println(string(payloadBytes))
}

func ErrorWithReq(
	err error,
	req *http.Request,
) {
	payload := newErrorLogEntryPayload(err)
	payload.Context = &logEntryPayloadContext{
		HTTPRequest: &logEntryPayloadContextHTTPRequest{
			Method:    req.Method,
			URL:       req.URL.String(),
			UserAgent: req.UserAgent(),
			Referrer:  req.Referer(),
			RemoteIP:  req.RemoteAddr,
		},
	}
	payloadBytes, _ := json.Marshal(payload)
	fmt.Println(string(payloadBytes))
}

func ErrorWithReqRes(
	err error,
	req *http.Request,
	resCode int,
) {
	payload := newErrorLogEntryPayload(err)
	payload.Context = &logEntryPayloadContext{
		HTTPRequest: &logEntryPayloadContextHTTPRequest{
			Method:             req.Method,
			URL:                req.URL.String(),
			UserAgent:          req.UserAgent(),
			Referrer:           req.Referer(),
			RemoteIP:           req.RemoteAddr,
			ResponseStatusCode: resCode,
		},
	}
	payloadBytes, _ := json.Marshal(payload)
	fmt.Println(string(payloadBytes))
}
