package logs

import "os"

// OpenLogs open log file if not exist then create, it only appends log to log file
func OpenLogs(logFilePath string) (*os.File, error) {
	return os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
}

// ReadLogs will read all log from log file, The problem of too large log files has not been solved yet
func ReadLogs(logFilePath string) ([]byte, error) {
	return os.ReadFile(logFilePath)
}
