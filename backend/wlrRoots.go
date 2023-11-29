package backend

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/DebuggerAndrzej/displayer/backend/entities"
)

func getWlrRandrOutput() string {
	cmd := exec.Command("wlr-randr")
	stdout, err := cmd.Output()

	if err != nil {
		return "Failed to run command"
	}
	return string(stdout)
}

func GetMonitorData() []entities.MonitorData {
	wrlRandrOutput := getWlrRandrOutput()
	for _, line := range strings.Split(wrlRandrOutput, "\n") {
		if isNewMonitorData(line) {
			fmt.Println("NEW MONITOR")
		}
	}

	return nil
}

func isNewMonitorData(line string) bool {
	return !strings.HasPrefix(line, " ") && line != ""
}
