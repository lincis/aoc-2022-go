package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func ReadInputFile(path string) []string {
	fileBytes, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return strings.Split(string(fileBytes), "\n")
}
