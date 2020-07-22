package logger

// LogLevel - represents log level
type LogLevel string

const (
	// Trace - log level
	Trace LogLevel = "0"
	// Info - log level
	Info LogLevel = "1"
	// Warning - log level
	Warning LogLevel = "2"
	// Error - log level
	Error LogLevel = "3"
	// Fatal - log level
	Fatal LogLevel = "4"
	// Log - used to log events without a level
	Default LogLevel = "5"
)

// LevelMap - maps loglevel to string attribute
var LevelMap = map[LogLevel]string{
	Trace:   "TRACE",
	Info:    "INFO",
	Warning: "WARNING",
	Error:   "ERROR",
	Fatal:   "FATAL",
	Default: "DEFAULT",
}
