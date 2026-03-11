package timer

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/gen2brain/beeep"
)

func Start(minutes int) {

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	select {

	case <-time.After(time.Duration(minutes) * time.Minute):
		beeep.Notify("Pomodoro", "Tempo acabou!", "")
		fmt.Println("🍅 Pomodoro finalizado")

	case <-sig:
		fmt.Println("Pomodoro cancelado")
	}
}