package utils

import "os"

var (
	LogDir                string = "/var/log/autoupd"
	SystemdServicePath    string = "/etc/systemd/system/autoupd.service"
	SystemdTimerPath      string = "/etc/systemd/system/autoupd.timer"
	AutoupdConfigFilePath string = os.Getenv("HOME") + "/.config/autoupd/config.yaml"
	IsRollingRelease      bool   = false // Set to true if the system is a rolling release
)

var UpdateCmd = map[string]string{
	"yay":     "yay -Syu --noconfirm",
	"apt":     "apt update && apt upgrade -y",
	"yum":     "yum update -y",
	"dnf":     "dnf upgrade --refresh -y",
	"pacman":  "pacman -Syu --noconfirm",
	"zypper":  "zypper --non-interactive refresh && zypper --non-interactive update",
	"apk":     "apk update && apk upgrade",
	"brew":    "HOMEBREW_NO_INTERACTION=1 brew update && brew upgrade",
	"nix":     "nix-channel --update && nix-env -u '*'",
	"flatpak": "flatpak update --noninteractive",
	"snap":    "snap refresh",
}

var serviceUnit = `[Unit]
Description=Auto system update with autoupd
After=network.target

[Service]
Type=oneshot
ExecStart=/usr/local/bin/autoupd
`

var timerUnit = `
[Unit]
Description=Daily run of autoupd

[Timer]
OnBootSec=10min
OnCalendar=*-*-* 00:00:00
Persistent=true

[Install]
WantedBy=timers.target
`
