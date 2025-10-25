// Package cmd
package cmd

import (
	"log"
	"os/exec"

	"github.com/linuxer77/aictl/internal/utils"
)

func Tree() []byte {
	cwd := utils.GetCurrDir()
	cmd := exec.Command("tree", cwd)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	return output
}
