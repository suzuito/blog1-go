package xlogging

type Level string

const (
	SeverityDebug Level = "DEBUG"
	SeverityInfo  Level = "INFO"
	SeverityError Level = "ERROR"
)

type Logger interface {
	Payloadf(severity Level, format string, a ...interface{})
	PayloadJSON(severity Level, v interface{})
}
