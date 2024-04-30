package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/coletucker12/aoc/utils"
)

func main() {
	inputFile := "inputs/day9.txt"
	// inputFile := "input.txt"
	lines := utils.ReadFileLines(inputFile)
	fmt.Printf("Part 1 Answer: %d\n\n", solvePart1(lines))
	fmt.Printf("Part 2 Answer: %d\n\n", solvePart2(lines))
}

func solvePart1(lines []string) int {
	sum := 0
	for _, line := range lines {
		val := predictNextValue(line)
		sum += val
	}
	return sum
}

func solvePart2(lines []string) int {
	sum := 0
	for _, line := range lines {
		val := predictNextValueBackwards(line)
		sum += val
	}
	return sum
}

func allZero(numbers []int) bool {
	for _, i := range numbers {
		if i != 0 {
			return false
		}
	}
	return true
}

func predictNextValue(line string) int {
	var startingNums []int
	for _, s := range utils.DeleteEmptyLines(strings.Split(strings.TrimSpace(line), " ")) {
		startingNums = append(startingNums, utils.ParseNumber(s))
	}

	if len(startingNums) == 0 {
		return 0
	}

	var numsToAdd []int
	cur := startingNums
	for !allZero(cur) {
		var nums []int
		for i := 0; i < len(cur)-1; i++ {
			newNum := cur[i+1] - cur[i]
			nums = append(nums, newNum)
		}
		cur = nums
		numsToAdd = append(numsToAdd, cur[len(cur)-1])
	}
	nextVal := startingNums[len(startingNums)-1]
	for _, num := range numsToAdd {
		nextVal += num
	}
	return nextVal
}

func predictNextValueBackwards(line string) int {
	var startingNums []int
	for _, s := range utils.DeleteEmptyLines(strings.Split(strings.TrimSpace(line), " ")) {
		startingNums = append(startingNums, utils.ParseNumber(s))
	}

	if len(startingNums) == 0 {
		return 0
	}

	var numsToSub []int
	numsToSub = append(numsToSub, startingNums[0])
	cur := startingNums
	for !allZero(cur) {
		var nums []int
		for i := 0; i < len(cur)-1; i++ {
			newNum := cur[i+1] - cur[i]
			nums = append(nums, newNum)
		}
		cur = nums
		numsToSub = append(numsToSub, cur[0])
	}
	slices.Reverse(numsToSub)
	sum := 0
	for i := 0; i < len(numsToSub); i++ {
		sum = numsToSub[i] - sum
	}
	return sum
}
