# autoupd

A simple, zero-config tool to automatically update your system packages daily.

`autoupd` automatically detects your system's package manager, performs an update, and sets up a systemd timer to run daily. It's a "set it and forget it" utility for keeping your system up-to-date.

## Key Features

- **Zero-Configuration:** No config files or manual setup needed.
- **Automatic Automation:** Automatically installs a systemd service and timer for daily updates on the first run.
- **Broad Support:** Works with most major Linux package managers.
- **Smart & Safe:** Requires `sudo` for system-level changes and skips updates if a successful one has already run that day.
- **Log Management:** Keeps organized logs in `/var/log/autoupd` with automatic rotation to save space.

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

## Installation

Ensure you have a working Go environment.

```bash
go install github.com/2SSK/autoupd@latest
```

## Usage

To perform the initial update and activate daily automation, run the command with `sudo`. You only need to do this once.

```bash
sudo autoupd
```

The first time you run this command, `autoupd` will:
1.  Ask for your password to gain `sudo` privileges.
2.  Detect your package manager and update all system packages.
3.  Install and enable a `systemd` timer (`autoupd.timer`) that will run `autoupd` automatically every day.

After the first run, the systemd timer will handle the daily updates. No further action is required.

## Manual Updates

If you want to force an update at any time, you can run the command again:

```bash
sudo autoupd
```

If an update was already successfully completed today, it will be skipped unless you wish to run it again.

## Viewing Logs

You can monitor the update history and check for errors in the log files located at:
`/var/log/autoupd/`

Logs are rotated automatically every 45 days or if the total size exceeds 15MB.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.