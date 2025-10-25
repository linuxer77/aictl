package cmd

import (
	"fmt"

	"github.com/linuxer77/aictl/internal/prompt"
	"github.com/linuxer77/aictl/internal/utils"
	"github.com/spf13/cobra"
)

func CommandRes() *cobra.Command {
	var query string
	cmd := &cobra.Command{
		Use:   "go",
		Short: "i know its just a wrapper",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			query = args[0]
			p := prompt.GetPrompt(query)
			llmRes := utils.GenText(p)
			fmt.Println(llmRes)
		},
	}
	return cmd
}

func init() {
	rootCmd.AddCommand(CommandRes())
}
