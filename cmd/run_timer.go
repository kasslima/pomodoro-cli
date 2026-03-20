package cmd

import (
    "strconv"
    "github.com/kasslima/pomodoro-cli/internal/timer"
    "github.com/spf13/cobra"
)

var runTimerCmd = &cobra.Command{
    Use:    "run-timer [minutes]",
    Hidden: true, // não aparece no help
    Args:   cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        minutes, _ := strconv.Atoi(args[0])
        timer.Start(minutes, getBlocksService())
    },
}

func init() {
    rootCmd.AddCommand(runTimerCmd)
}