# autoupd

<p align="center">
  <img src="./assets/logo.png" alt="autoupd Logo" width="350">
</p>

<p align="center">
  <strong>A simple, zero-config tool to automatically update your system packages.</strong>
</p>

`autoupd` is a "set it and forget it" utility for keeping your system up-to-date. It automatically detects your system's package manager, performs an update, and sets up a systemd timer to run daily for rolling-release distros or weekly for others.

## How It Works

1.  **Detects Package Manager:** Automatically identifies the package manager on your system (e.g., `apt`, `pacman`, `dnf`).
2.  **Updates Packages:** Runs the appropriate command to update all system packages.
3.  **Automates with Systemd:** On the first run, it installs and enables a `systemd` timer to automate future updates.
    - **Rolling-Release:** Runs daily.
    - **Other Systems:** Runs weekly.

## Installation

### Prerequisites

- [Go](https://golang.org/doc/install) (for building from source)
- `git`

### Build from Source

```bash
# Clone the repository
git clone https://github.com/2SSK/autoupd.git

# Navigate to the project directory
cd autoupd

# Build the binary
go build .

# Move the binary to your PATH
sudo cp autoupd /usr/local/bin/

# Run autoupd for the first time to set up automation
sudo autoupd
```

## Usage

After installation, `autoupd` will run automatically. You can also run it manually.

### First Run

To perform the initial update and activate the systemd timer, run:

```bash
sudo autoupd
```

This command will:

1.  Ask for your password to gain `sudo` privileges.
2.  Update all system packages.
3.  Install and enable a `systemd` timer for automatic updates.

### Manual Updates

To force an update at any time, use the `--force` or `-f` flag:

```bash
sudo autoupd --force
```

### View Status

To view the status of `autoupd` without performing an update, use the `--status` or `-s` flag:

```bash
autoupd --status
```

This will display a dashboard with information about the last and next update times.

## Automatic Updates

`autoupd` uses a `systemd` timer to run automatically.

- **Service:** `/etc/systemd/system/autoupd.service`
- **Timer:** `/etc/systemd/system/autoupd.timer`

You can check the status of the timer with:

```bash
systemctl status autoupd.timer
```

## Supported Package Managers

- `apt`
- `apk`
- `brew`
- `dnf`
- `flatpak`
- `nix`
- `pacman`
- `snap`
- `yay`
- `yum`
- `zypper`

## Logs

Logs are stored in `/var/log/autoupd`. You can view the latest log with:

```bash
cat /var/log/autoupd/<today's-date>.log
```

## Uninstallation

To remove `autoupd` and its related files from your system:

```bash
# Stop and disable the systemd timer
sudo systemctl stop autoupd.timer
sudo systemctl disable autoupd.timer

# Remove the systemd files
sudo rm /etc/systemd/system/autoupd.service
sudo rm /etc/systemd/system/autoupd.timer

# Remove the binary
sudo rm /usr/local/bin/autoupd

# Remove the log directory
sudo rm -rf /var/log/autoupd
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
