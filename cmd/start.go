package cmd

import (
	"fmt"
	"strconv"

	"github.com/kasslima/pomodoro-cli/internal/pomodoro"
	"github.com/kasslima/pomodoro-cli/internal/timer"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use: "start [minutes]",
	Run: func(cmd *cobra.Command, args []string) {

		minutesStr := args[0]

		minutes, err := strconv.Atoi(minutesStr)
		if err != nil {
			fmt.Println("minutes must be a number")
			return
		}

		timer := &timer.RealTimer{}
		pomodoro := pomodoro.NewPomodoro(timer)

		pomodoro.StartPomodoro(minutes)
	},
}

func init() {

	rootCmd.AddCommand(startCmd)
}
