<div align="center">

# Pomodoro CLI

**A lightweight command-line Pomodoro timer built in Go.**  
Stay focused, block distractions, and get things done.

[![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go&logoColor=white)](https://go.dev/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
[![Platform](https://img.shields.io/badge/Platform-Windows-0078D6?style=flat&logo=windows&logoColor=white)](https://github.com)

</div>

---

## Overview

Pomodoro CLI is a terminal-based productivity tool that helps you manage focus sessions using the [Pomodoro Technique](https://en.wikipedia.org/wiki/Pomodoro_Technique). During your session, it automatically blocks distracting websites and sends a desktop notification with a sound alert when time is up.

---

## ⚠️ Important: Admin Privileges for Website Blocking

Because Pomodoro CLI modifies your system's `hosts` file to block websites, you **must run the application with elevated privileges** when using the timer with active blocks.

### How to run with elevated privileges:

**Windows**:  
1. Search for your terminal (Command Prompt, PowerShell, or Windows Terminal) in the Start Menu.
2. Right-click and select **Run as administrator**.
3. Run your commands normally:
   ```bash
   pomodoro-cli.exe start 25
   ```

**Linux / macOS**:  
Use `sudo` before your commands:
```bash
sudo ./pomodoro-cli start 25
```

---

## Quick Start

```bash
# Add distracting sites to your block list
pomodoro-cli.exe blocks add facebook.com
pomodoro-cli.exe blocks add youtube.com

# Start a 25-minute focus session
pomodoro-cli.exe start 25

# Stop early at any time (lifts all blocks immediately)
pomodoro-cli.exe stop
```

---

## Commands

### `start [minutes]`

Starts a new Pomodoro session for the specified duration. Website blocks are enforced for the entire session.

```bash
pomodoro-cli.exe start 25
```

---

### `stop`

Manually stops the running timer, triggers a completion notification, and lifts all active blocks.

```bash
pomodoro-cli.exe stop
```

---

### `blocks`

Manages the list of websites to block during focus sessions.

| Subcommand | Description | Example |
|---|---|---|
| `blocks add [link]` | Add a site to the block list | `pomodoro-cli.exe blocks add twitter.com` |
| `blocks remove [link]` | Remove a site from the block list | `pomodoro-cli.exe blocks remove twitter.com` |
| `blocks list` | Show all blocked sites | `pomodoro-cli.exe blocks list` |

---

## Building from Source

Requires [Go](https://go.dev/) to be installed.

```bash
# Clone the repository
git clone https://github.com/kasslima/pomodoro-cli.git
cd pomodoro-cli

# Build the executable
go build -o pomodoro-cli.exe main.go

# Run
./pomodoro-cli.exe start 25
```

---

<div align="center">
  <sub>Built with Go · <a href="https://en.wikipedia.org/wiki/Pomodoro_Technique">Learn about the Pomodoro Technique</a></sub>
</div>
