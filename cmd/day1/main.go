package main

import (
	"fmt"
	"github.com/coletucker12/aoc/utils"
	"strconv"
	"strings"
)

var validStrNums = []string{
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}
var numStrMap = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func isNumber(s string) (bool, int) {
	if mappedVal := numStrMap[s]; mappedVal > 0 {
		return true, mappedVal
	}

	val, err := strconv.Atoi(s)
	if err != nil {
		return false, -1
	}
	return true, val
}

func containsStrNumber(s string) (bool, int) {
	for _, substr := range validStrNums {
		if strings.Contains(s, substr) {
			return true, numStrMap[substr]
		}
	}
	return false, 0
}

func findFirstNum(str string) (int, error) {
	buffer := ""
	for _, c := range str {
		valid, num := isNumber(buffer)
		if valid {
			return num, nil
		}

		valid, num = containsStrNumber(buffer)
		if valid {
			return num, nil
		}

		valid, num = isNumber(string(c))
		if valid {
			return num, nil
		}
		buffer = buffer + string(c)
	}
	return 0, nil
}

func findLastNum(str string) (int, error) {
	buffer := ""
	for i := len(str) - 1; i >= 0; i-- {
		valid, num := isNumber(buffer)
		if valid {
			return num, nil
		}

		valid, num = containsStrNumber(buffer)
		if valid {
			return num, nil
		}

		valid, num = isNumber(string(str[i]))
		if valid {
			return num, nil
		}
		buffer = string(str[i]) + buffer
	}
	return 0, nil
}

func solveDay1Part1(lines []string) int {
	sum := 0

	for _, line := range lines {
		if line == "" {
			continue
		}

		first, err := findFirstNum(line)
		if err != nil {
			utils.Check(err)
		}

		last, err := findLastNum(line)
		if err != nil {
			utils.Check(err)
		}

		combined, err := strconv.Atoi(fmt.Sprintf("%d%d", first, last))
		if err != nil {
			utils.Check(err)
		}
		sum += combined
	}
	return sum
}

func solveDay1Part2(lines []string) int {
	sum := 0
	for _, line := range lines {
		if line == "" {
			continue
		}

		first, err := findFirstNum(line)
		if err != nil {
			utils.Check(err)
		}

		last, err := findLastNum(line)
		if err != nil {
			utils.Check(err)
		}

		combined, err := strconv.Atoi(fmt.Sprintf("%d%d", first, last))
		if err != nil {
			utils.Check(err)
		}
		//fmt.Printf("line: %s \n\tfirst: %d -- last: %d -- combined: %d\n", line, first, last, combined)
		sum += combined
	}
	return sum
}

func main() {
	inputFile := "inputs/day1"
	//inputFile := "input"
	lines := utils.ReadFileLines(inputFile)
	//fmt.Printf("Day 1 Part 1 Answer: %d\n", solveDay1Part1(lines))
	fmt.Printf("Day 1 Part 2 Answer: %d\n", solveDay1Part2(lines))
}
