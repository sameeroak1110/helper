package helper

import (
	"fmt"
	"os"
	"strconv"
)

func PidFile() (int, error) {
	pid := os.Getpid() // Fetches pid of the server process.

	// Creates a pid file. This file contains pid of the server process.
	pPIDFile, err := os.Create(PID_FILE)
	if (pPIDFile == nil) || (err != nil) {
		return -1, fmt.Errorf("ERROR: PID file create error: %s", err.Error())
	}
	pPIDFile.Close()

	_, err = pPIDFile.WriteString(strconv.Itoa(pid) + "\n")
	if err != nil {
		pPIDFile.Close()
		return -1, fmt.Errorf("ERROR: PID file write error: %s", err.Error())
	}

	return pid, nil
}
