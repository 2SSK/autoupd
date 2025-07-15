package utils

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func SetupSystemdService() error {
	// Write service file if not exists
	if _, err := os.Stat(SystemdServicePath); os.IsNotExist(err) {
		if err := os.WriteFile(SystemdServicePath, []byte(serviceUnit), 0644); err != nil {
			return fmt.Errorf("failed to write service file: %w", err)
		}
	}

	// Detect package manager to update the timer unit
	DetectRollingRelease()

	// Write timer file if not exists
	if _, err := os.Stat(SystemdTimerPath); os.IsNotExist(err) {
		if err := os.WriteFile(SystemdTimerPath, []byte(GetTimerUnit()), 0644); err != nil {
			return fmt.Errorf("failed to write timer file: %w", err)
		}
	}

	// Reload systemd
	if err := exec.Command("systemctl", "daemon-reexec").Run(); err != nil {
		return fmt.Errorf("failed to reload systemd: %w", err)
	}

	// Enable and start timer
	if err := exec.Command("systemctl", "enable", "--now", "autoupd.timer").Run(); err != nil {
		return fmt.Errorf("failed to enable/start timer: %w", err)
	}

	Logger.Println("Systemd service and timer are now installed and active.")
	return nil
}

func IsTimerActive() bool {
	out, err := exec.Command("systemctl", "is-active", "autoupd.timer").Output()
	if err != nil {
		return false
	}
	return strings.TrimSpace(string(out)) == "active"
}

func GetTimerUnit() string {
	if IsRollingRelease {
		return `[Unit]
Description=Daily run of autoupd

[Timer]
OnCalendar=*-*-* 00:00:00
Persistent=true

[Install]
WantedBy=timers.target
`
	} else {
		return `[Unit]
Description=Weekly run of autoupd

[Timer]
OnCalendar=Mon *-*-* 00:00:00
Persistent=true

[Install]
WantedBy=timers.target
`
	}
}
