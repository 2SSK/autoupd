package utils

import "os"

var (
	LogDir             string = "/var/log/autoupd"
	SystemdServiceFile string = "/etc/systemd/system/autoupd.service"
	SystemdTimerFile   string = "/etc/systemd/system/autoupd.timer"
	AutoupdConfigFile  string = os.Getenv("HOME") + "/.config/autoupd/config.yaml"
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
