package pomodoro

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"
)

type Timer interface {
	Wait(duration time.Duration)
}

type Pomodoro struct {
	Duration int
	Timer    Timer
}

func NewService() *Pomodoro {
	return &Pomodoro{}
}

func pidFilePath() string {
	dir, err := os.UserCacheDir() // ~AppData/Local on Windows, ~/.cache on Linux/Mac
	if err != nil {
		dir = os.TempDir() // fallback to os temp dir
	}
	return filepath.Join(dir, "pomodoro.pid")
}

func (s *Pomodoro) Start(minutes int) {
	exe, _ := os.Executable()
	cmd := exec.Command(exe, "run-timer", strconv.Itoa(minutes))
	cmd.Start()
	os.WriteFile(pidFilePath(), []byte(strconv.Itoa(cmd.Process.Pid)), 0644)
}

func (s *Pomodoro) Stop() {
	data, err := os.ReadFile(pidFilePath())
	if err != nil {
		fmt.Println("error reading pid file")
		return
	}
	pid, _ := strconv.Atoi(string(data))
	process, err := os.FindProcess(pid)
	if err != nil {
		fmt.Println("error finding process")
		return
	}
	if err := process.Kill(); err != nil { // Kill() is cross-platform unlike Signal(SIGINT)
		fmt.Println("error stopping timer:", err)
	}
	os.Remove(pidFilePath())
}
