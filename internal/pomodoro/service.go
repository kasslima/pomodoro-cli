package pomodoro

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"

	"github.com/kasslima/pomodoro-cli/internal/blocks"
)

type Timer interface {
	Wait(duration time.Duration)
}

type Pomodoro struct {
	Duration      int
	Timer         Timer
	BlocksService blocks.BlocksService
}

func NewService(blocksService blocks.BlocksService) *Pomodoro {
	return &Pomodoro{
		BlocksService: blocksService,
	}
}

func pidFilePath() string {
	dir, err := os.UserCacheDir() 
	if err != nil {
		dir = os.TempDir() 
	}
	return filepath.Join(dir, "pomodoro.pid")
}

func (s *Pomodoro) Start(minutes int) {
	if s.BlocksService != nil {
		if err := s.BlocksService.ApplyBlocks(); err != nil {
			fmt.Println("Warning: Could not apply blocks:", err)
		}
	}

	exe, _ := os.Executable()
	cmd := exec.Command(exe, "run-timer", strconv.Itoa(minutes))
	cmd.Start()
	os.WriteFile(pidFilePath(), []byte(strconv.Itoa(cmd.Process.Pid)), 0644)
}

func (s *Pomodoro) Stop() {
	if s.BlocksService != nil {
		if err := s.BlocksService.RemoveAppliedBlocks(); err != nil {
			fmt.Println("Warning: Could not remove applied blocks:", err)
		}
	}

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
	if err := process.Kill(); err != nil {
		fmt.Println("error stopping timer:", err)
	}
	os.Remove(pidFilePath())
}
