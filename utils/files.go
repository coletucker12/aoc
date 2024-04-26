package utils

import (
	"os"
	"strings"
)

func ReadFileLines(fileName string) []string {
	file, err := os.ReadFile(fileName)
	if err != nil {
		Check(err)
	}
	lines := strings.Split(string(file), "\n")
	return lines
}
