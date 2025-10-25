package cmd

import (
	"fmt"

	"github.com/linuxer77/aictl/internal/utils"
	"github.com/spf13/cobra"
)

func NewStr() *cobra.Command {
	var str string

	cmd := &cobra.Command{
		Use:   "string",
		Short: "Handle strinr input",
		Run: func(cmd *cobra.Command, args []string) {
			respFromAi := utils.GenText(str)
			fmt.Println(respFromAi)
		},
	}

	cmd.Flags().StringVarP(&str, "str", "s", "", "some string input")
	cmd.MarkFlagRequired("str")
	return cmd
}

func init() {
	rootCmd.AddCommand(NewStr())
}
