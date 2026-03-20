package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var blocksCmd = &cobra.Command{
	Use:   "blocks",
	Short: "Manage website blocks during pomodoro",
}

var blocksAddCmd = &cobra.Command{
	Use:   "add [link]",
	Short: "Add a website to the block list",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		link := args[0]
		service := getBlocksService()
		
		fmt.Printf("Verifying %s...\n", link)
		err := service.AddBlock(link)
		if err != nil {
			fmt.Printf("Error adding block: %v\n", err)
			return
		}
		
		fmt.Printf("Successfully added %s to block list.\n", link)
	},
}

var blocksRemoveCmd = &cobra.Command{
	Use:   "remove [link]",
	Short: "Remove a website from the block list",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		link := args[0]
		service := getBlocksService()
		err := service.RemoveBlock(link)
		if err != nil {
			fmt.Printf("Error removing block: %v\n", err)
			return
		}
		
		fmt.Printf("Successfully removed %s from block list.\n", link)
	},
}

var blocksListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all blocked websites",
	Run: func(cmd *cobra.Command, args []string) {
		service := getBlocksService()
		blocks, err := service.GetBlocks()
		if err != nil {
			fmt.Printf("Error getting blocks: %v\n", err)
			return
		}
		
		fmt.Println("Blocked websites during pomodoro:")
		if len(blocks) == 0 {
			fmt.Println("  (No websites currently blocked)")
			return
		}
		for _, b := range blocks {
			fmt.Printf("  - %s\n", b)
		}
	},
}

func init() {
	blocksCmd.AddCommand(blocksAddCmd)
	blocksCmd.AddCommand(blocksRemoveCmd)
	blocksCmd.AddCommand(blocksListCmd)
	rootCmd.AddCommand(blocksCmd)
}
