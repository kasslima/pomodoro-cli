package cmd

import (
	"fmt"
	"strconv"

	"github.com/kasslima/pomodoro-cli/internal/pomodoro"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:  "start [minutes]",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		minutes, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("minutes must be a number")
			return
		}

		service := pomodoro.NewService()
		service.Start(minutes)
	},
}

func init() {

	rootCmd.AddCommand(startCmd)
}
