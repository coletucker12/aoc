package main

import (
	"fmt"
	"github.com/coletucker12/aoc/utils"
	"strconv"
)

const (
	nonSymbol = ""
)

type ValidNumber struct {
	Value int
	Valid bool
	Char  string
}

func isNumber(s string) bool {
	_, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	return true
}

func isValidNumber(lines []string, indexes []int, lineIndex int) bool {
	// indexes = [0, 1, 2]
	for _, index := range indexes {
		// check left
		if index-1 >= 0 && string(lines[lineIndex][index-1]) != "." {
			return true
		}
		// check up (starting it -1,-1)
	}
	return false
}

func solveDay3Part1(lines []string) int64 {
	// 467..114..
	// ...*......
	// ..35..633.
	//
	// we loop through each line ->
	// 		we loop through each char ->
	//  		if we hit a number, continue processing until end of number
	// 			at the same time checking if at anypoint it touches something else
	var validNums []ValidNumber
	for lineIdx, line := range lines {
		leftBound := -1
		rightBound := len(line)
		topBound := -1
		bottomBound := len(lines)
		buffer := ""
		isValid := false
		symbol := ""
		for charIdx, ch := range line {
			if isNumber(string(ch)) {
				buffer = buffer + string(ch)
				// check if this character is touching any special characters

				// check left
				if charIdx-1 > leftBound && !isNumber(string(lines[lineIdx][charIdx-1])) && string(lines[lineIdx][charIdx-1]) != "." {
					isValid = true
					symbol = string(lines[lineIdx][charIdx-1])
				}

				// check right
				if charIdx+1 < rightBound && !isNumber(string(lines[lineIdx][charIdx+1])) && string(lines[lineIdx][charIdx+1]) != "." {
					isValid = true
					symbol = string(lines[lineIdx][charIdx+1])
				}

				// check up
				if lineIdx-1 > topBound && !isNumber(string(lines[lineIdx-1][charIdx])) && string(lines[lineIdx-1][charIdx]) != "." {
					isValid = true
					symbol = string(lines[lineIdx-1][charIdx])
				}

				// check bottom
				if lineIdx+1 < bottomBound && !isNumber(string(lines[lineIdx+1][charIdx])) && string(lines[lineIdx+1][charIdx]) != "." {
					isValid = true
					symbol = string(lines[lineIdx+1][charIdx])
				}

				// check top left
				if charIdx-1 > leftBound && lineIdx-1 > topBound && !isNumber(string(lines[lineIdx-1][charIdx-1])) && string(lines[lineIdx-1][charIdx-1]) != "." {
					isValid = true
					symbol = string(lines[lineIdx-1][charIdx-1])
				}

				// check top right
				if charIdx+1 < rightBound && lineIdx-1 > topBound && !isNumber(string(lines[lineIdx-1][charIdx+1])) && string(lines[lineIdx-1][charIdx+1]) != "." {
					isValid = true
					symbol = string(lines[lineIdx-1][charIdx+1])
				}

				// check bottom left
				if charIdx-1 > leftBound && lineIdx+1 < bottomBound && !isNumber(string(lines[lineIdx+1][charIdx-1])) && string(lines[lineIdx+1][charIdx-1]) != "." {
					isValid = true
					symbol = string(lines[lineIdx+1][charIdx-1])
				}

				// check bottom right
				if charIdx+1 < rightBound && lineIdx+1 < bottomBound && !isNumber(string(lines[lineIdx+1][charIdx+1])) && string(lines[lineIdx+1][charIdx+1]) != "." {
					isValid = true
					symbol = string(lines[lineIdx+1][charIdx+1])
				}

				if charIdx < rightBound-1 {
					continue
				}
			}

			if isNumber(buffer) {
				num, err := strconv.Atoi(buffer)
				if err != nil {
					utils.Check(err)
				}
				validNums = append(validNums, ValidNumber{
					Value: num,
					Valid: isValid,
					Char:  symbol,
				})
				buffer = ""
				symbol = ""
				isValid = false
			}
		}
	}
	var sum int64
	for _, num := range validNums {
		if num.Valid {
			sum += int64(num.Value)
		}
	}
	return sum
}

func solveDay3Part2(lines []string) int64 {
	fmt.Println(lines)
	var sum int64
	var validNums []ValidNumber
	for _, num := range validNums {
		if num.Valid {
			sum += int64(num.Value)
		}
	}
	return sum
}

func main() {
	//inputFile := "inputs/day3"
	inputFile := "input"
	lines := utils.ReadFileLines(inputFile)
	fmt.Printf("Day 3 Part 1 Answer: %d\n", solveDay3Part1(lines))
	fmt.Printf("Day 3 Part 2 Answer: %d\n", solveDay3Part2(lines))
}
