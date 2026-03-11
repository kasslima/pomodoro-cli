package cmd

import (
	"github.com/gen2brain/beeep"
	"github.com/kasslima/pomodoro-cli/internal/pomodoro"
	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use: "stop",
	Run: func(cmd *cobra.Command, args []string) {

		
		pomodoro := pomodoro.NewService()

		pomodoro.Stop()

		beeep.AppName = "pomodoro-cli"
		beeep.Notify("Pomodoro", "Pomodoro finalizado!", "")
	},
}

func init() {

	rootCmd.AddCommand(stopCmd)
}
