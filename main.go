package main

import (
	"fmt"

	cmd "github.com/linuxer77/aictl/cmd"
)

func main() {
	getTree := cmd.Tree()
	fmt.Println(len(getTree))
}
