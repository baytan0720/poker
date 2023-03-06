package container

import "poker/internal/logs"

func Logs(containerIdOrName string) ([]byte, error) {
	containerId := checkName(containerIdOrName)
	containerPath, err := findPath(containerId)
	if err != nil {
		return nil, err
	}

	logFilePath := containerPath + "/stdout.log"
	return logs.ReadLogs(logFilePath)
}
