package timer

import (
	"time"
)

type RealTimer struct{}

func (t *RealTimer) Wait(d time.Duration) {
    time.Sleep(d)
}