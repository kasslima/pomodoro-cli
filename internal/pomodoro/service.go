package pomodoro

import "time"

type Timer interface {
	Wait(duration time.Duration)
}

type Pomodoro struct {
	Duration int
	Timer    Timer
}

func NewPomodoro(timer Timer) *Pomodoro {
	return &Pomodoro{
		Timer: timer,
	}
}

func (p *Pomodoro) StartPomodoro(duration int) {
	p.Duration = duration
	p.Timer.Wait(time.Duration(duration) * time.Minute)
}