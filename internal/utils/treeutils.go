// Package utils
package utils

import (
	"log"
	"os"
)

func GetCurrDir() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return dir
}
