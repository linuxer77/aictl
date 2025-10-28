// Package cmd
package tree

import (
	"log"
	"os"
	"os/exec"
)

func Tree() string {
	cwd := GetCurrDir()
	cmd := exec.Command("tree", cwd)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	return string(output)
}

func GetCurrDir() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

func NoTree() string {
	cwd := GetCurrDir()
	return cwd
}
