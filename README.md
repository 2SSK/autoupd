# autoupd

A simple, zero-config tool to automatically update your system packages.

`autoupd` detects your system's package manager and runs the appropriate update commands. It's designed to be run as a daily automated task.

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

## How It Works

1.  **Detects:** Automatically finds the available package manager on your system.
2.  **Updates:** Runs the standard update and upgrade commands.
3.  **Logs:** Stores a log of the update process in `/var/log/autoupd/`.
4.  **Skips:** If an update has already run successfully today, it will skip subsequent runs.

## Installation

```bash
go install github.com/2SSK/autoupd@latest
```

## Usage

The command must be run with `sudo` to have the necessary permissions to update system packages.

```bash
sudo autoupd
```

## Automation

You can automate this tool to run daily using `cron` or a `systemd` timer.

**Example Cron Job (runs daily at 2:00 AM):**

```
0 2 * * * /path/to/your/go/bin/autoupd
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
