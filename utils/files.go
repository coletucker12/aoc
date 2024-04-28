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

func ReadFileLinesUnsplit(fileName string) string {
	file, err := os.ReadFile(fileName)
	if err != nil {
		Check(err)
	}
	return string(file)
}

func DeleteEmptyLines(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}
