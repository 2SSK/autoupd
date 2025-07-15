package utils

import (
	"os"
	"os/exec"
	"strings"
	"time"
)

func DetectPackageManager() string {
	managers := []string{"yay", "apt", "yum", "dnf", "pacman", "zypper", "apk", "brew", "nix", "flatpak", "snap"}
	for _, m := range managers {
		if _, err := exec.LookPath(m); err == nil {
			return m
		}
	}
	return "unknown package manager"
}

func PerformPackageUpdate() error {
	pkgManager := DetectPackageManager()
	updateCmd := UpdateCmd[pkgManager]

	// Write headers to the log file
	Logger.Printf("\n\n===== [%s] Starting update with %s =====\n", time.Now().Format("15:04:05"), pkgManager)
	Logger.Printf("Running command: %s\n\n", updateCmd)

	// Prepare command to execute
	cmd := exec.Command("sh", "-c", updateCmd)
	cmd.Stdout = Logger.Writer()
	cmd.Stderr = Logger.Writer()

	if err := cmd.Run(); err != nil {
		Logger.Printf("Error executing update command: %v\n", err)
		return err
	}

	Logger.Println("Update completed successfully.")
	return nil
}

func DetectRollingRelease() {
	data, err := os.ReadFile("/etc/os-release")
	if err != nil {
		IsRollingRelease = false
		return
	}

	content := strings.ToLower(string(data))
	rollingIDs := []string{
		"arch", "artix", "manjaro", "void", "gentoo", "nixos", "solus", "opensuse-tumbleweed",
	}

	for _, id := range rollingIDs {
		if strings.Contains(content, id) {
			IsRollingRelease = true
			return
		}
	}

	IsRollingRelease = false
}
