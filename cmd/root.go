package cmd

import (
	"os"
	"github.com/kasslima/pomodoro-cli/internal/blocks"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pomodoro",
	Short: "A simple pomodoro CLI",
}

func getBlocksService() blocks.BlocksService {
	repo := blocks.NewBlocksRepository(blocks.BlocksFilePath())
	return blocks.NewBlocksService(repo)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}