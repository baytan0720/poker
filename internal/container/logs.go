package container

import "poker/internal/logs"

func Logs(containerId string) ([]byte, error) {
	containerPath, err := findPath(containerId)
	if err != nil {
		return nil, err
	}

	logFilePath := containerPath + "/stdout.log"
	return logs.ReadLogs(logFilePath)
}
