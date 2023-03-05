package container

import "poker/internal/logs"

func Logs(id string) ([]byte, error) {
	containerId, err := find(id)
	if err != nil {
		return nil, err
	}

	logFilePath := CONTAINER_FOLDER_PATH + containerId
	return logs.ReadLogs(logFilePath)
}
