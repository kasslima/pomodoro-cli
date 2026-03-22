package timer

import (
    "fmt"
    "os"
    "os/signal"
    "syscall"
    "time"

    "github.com/gen2brain/beeep"
    "github.com/kasslima/pomodoro-cli/internal/blocks"
)

func Start(minutes int, blocksService blocks.BlocksService) {
    sig := make(chan os.Signal, 1)
    signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
    defer signal.Stop(sig)

    defer func() {
        if blocksService != nil {
            blocksService.RemoveAppliedBlocks()
        }
    }()

    select {
    case <-time.After(time.Duration(minutes) * time.Minute):
        beeep.Notify("Pomodoro", "Time's up!", "") 
        fmt.Println("🍅 Pomodoro finished")
    case <-sig:
        fmt.Println("⛔ Pomodoro cancelled")
    }
}