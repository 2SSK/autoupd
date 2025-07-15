package status

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/2SSK/autoupd/internal/utils"
)

func getTodaysLogContent() string {
	today := time.Now().Format("2006-01-02")
	logFilePath := fmt.Sprintf("%s/%s.log", utils.LogDir, today)

	data, err := os.ReadFile(logFilePath)
	if err != nil {
		return "Error reading log file"
	}

	lines := strings.Split(string(data), "\n")
	if len(lines) > 20 { // show last 20 lines max
		lines = lines[len(lines)-20:]
	}
	return strings.Join(lines, "\n")
}

func getLastRun() string {
	today := time.Now().Format("2006-01-02")
	logFilePath := fmt.Sprintf("%s/%s.log", utils.LogDir, today)

	info, err := os.Stat(logFilePath)
	if err != nil {
		return "Not Found"
	}
	return info.ModTime().Format("Mon 15:04:05")
}

func getNextRun() string {
	// since it's daily via systemd timer
	now := time.Now()
	next := now.Add(24 * time.Hour)
	return next.Format("Mon 15:04:05")
}

func getOSName() string {
	data, err := os.ReadFile("/etc/os-release")
	if err != nil {
		return "Unknown"
	}
	for _, line := range strings.Split(string(data), "\n") {
		if strings.HasPrefix(line, "PRETTY_NAME=") {
			return strings.Trim(line[13:], `"`)
		}
	}
	return "Unknown"
}

func getKernelVersion() string {
	out, err := exec.Command("uname", "-r").Output()
	if err != nil {
		return "Unknown"
	}
	return strings.TrimSpace(string(out))
}
