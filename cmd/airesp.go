package cmd

import (
	"fmt"

	"github.com/linuxer77/aictl/internal/prompt"
	"github.com/linuxer77/aictl/internal/utils"
	"github.com/spf13/cobra"
)

func CommandRes() *cobra.Command {
	var notree bool
	cmd := &cobra.Command{
		Use:   "go",
		Short: "i know its just a wrapper",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			query := args[0]
			p := prompt.GetPrompt(query, notree)
			llmRes := utils.GenText(p)
			fmt.Println(llmRes)
		},
	}
	cmd.Flags().BoolVarP(&notree, "treecond", "n", false, "tree or not tree?")

	return cmd
}

func init() {
	rootCmd.AddCommand(CommandRes())
}
