package cloudlogging

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/pkg/errors"
)

func NewMessageInPayloadFromError(err error) string {
	type stackTracer interface {
		StackTrace() errors.StackTrace
	}
	werr, ok := errors.Cause(err).(stackTracer)
	if ok {
		// Mimic github.com/pkg/errors stack into runtime.Stack
		// https://github.com/googleapis/google-cloud-go/issues/1084#issuecomment-474565019
		stackLines := []string{}
		for _, st := range werr.StackTrace() {
			pc := uintptr(st) - 1
			f := runtime.FuncForPC(pc)
			funcName := f.Name()
			fileName, line := f.FileLine(pc)
			stackLine := fmt.Sprintf("%s()\n\t%s:%d +%#x", funcName, fileName, line, f.Entry())
			stackLines = append(stackLines, stackLine)
		}
		return fmt.Sprintf("%s\ngoroutine 1 [running]:\n%s", err.Error(), strings.Join(stackLines, "\n"))
	}
	buf := make([]byte, 1<<16)
	size := runtime.Stack(buf, false)
	return string(buf[:size])
}
