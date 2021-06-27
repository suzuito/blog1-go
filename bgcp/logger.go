package bgcp

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/suzuito/blog1-go/xlogging"
)

type logEntry struct {
	Message     string         `json:"message,omitempty"`
	Severity    xlogging.Level `json:"severity,omitempty"`
	Trace       string         `json:"logging.googleapis.com/trace,omitempty"`
	JSONPayload interface{}    `json:"jsonPayload,omitempty"`
}

type Logger struct {
	trace string
}

func NewLogger(trace string) *Logger {
	return &Logger{
		trace: trace,
	}
}

func (l *Logger) Payloadf(severity xlogging.Level, format string, a ...interface{}) {
	entry := logEntry{
		Message:  fmt.Sprintf(format, a...),
		Severity: severity,
		Trace:    l.trace,
	}
	body, _ := json.Marshal(&entry)
	fmt.Println(string(body))
}

func ParseTrace(projectID, traceHeader string) string {
	traceParts := strings.Split(traceHeader, "/")
	if len(traceParts) > 0 && len(traceParts[0]) > 0 {
		return fmt.Sprintf("projects/%s/traces/%s", projectID, traceParts[0])
	}
	return ""
}
