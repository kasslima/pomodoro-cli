package pomodoro

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/kasslima/pomodoro-cli/internal/timer"
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


func (s *Pomodoro) Start(minutes int) {

	pid := os.Getpid()
	os.WriteFile("/tmp/pomodoro.pid", []byte(strconv.Itoa(pid)), 0644)

	go timer.Start(minutes)
}

func (s *Pomodoro) Stop() {

	data, err := os.ReadFile("/tmp/pomodoro.pid")
	if err != nil {
		fmt.Println("error reading pid file")
		return
	}

	pid, _ := strconv.Atoi(string(data))

	process, _ := os.FindProcess(pid)

	process.Signal(os.Interrupt)
}