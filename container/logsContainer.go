package container

import (
	"fmt"
	"io/ioutil"
	"os"
	"poker/alert"
)

func PrintContainerLogs(ContainerID string) {
	logFilePath := CONTAINER_FOLDER_PATH + ContainerID + "/stdout.log"
	f, err := os.Open(logFilePath)
	if err != nil {
		alert.Fatal(err.Error())
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		alert.Fatal(err.Error())
	}
	fmt.Fprintf(os.Stdout, string(b))
}
