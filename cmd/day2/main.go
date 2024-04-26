package main

import (
	"fmt"
	"github.com/coletucker12/aoc/utils"
	"strconv"
	"strings"
)

const (
	MaxRed   = 12
	MaxGreen = 13
	MaxBlue  = 14
)

func getGameInfo(line string) (int, []string) {
	game := strings.Split(line, ":")
	gameId, err := strconv.Atoi(strings.Split(game[0], " ")[1])
	if err != nil {
		utils.Check(err)
		return 0, nil
	}
	return gameId, strings.Split(strings.TrimSpace(game[1]), ";")
}

func getCountAndColor(s string) (int, string) {
	v := strings.Split(s, " ")
	count, err := strconv.Atoi(v[0])
	if err != nil {
		utils.Check(err)
		return 0, ""
	}
	return count, v[1]
}

func checkCounts(counts map[string]int) bool {
	return counts["red"] <= MaxRed &&
		counts["blue"] <= MaxBlue &&
		counts["green"] <= MaxGreen
}

func validGame(gameSets []string) bool {
	for _, set := range gameSets {
		counts := map[string]int{
			"red":   0,
			"blue":  0,
			"green": 0,
		}

		for _, pull := range strings.Split(strings.TrimSpace(set), ",") {
			count, color := getCountAndColor(strings.TrimSpace(pull))
			counts[color] += count
		}

		if !checkCounts(counts) {
			return false
		}
	}
	return true
}

func getMinCounts(gameSets []string) map[string]int {
	counts := map[string]int{
		"red":   0,
		"blue":  0,
		"green": 0,
	}

	for _, set := range gameSets {
		for _, pull := range strings.Split(strings.TrimSpace(set), ",") {
			count, color := getCountAndColor(strings.TrimSpace(pull))
			if counts[color] < count {
				counts[color] = count
			}
		}
	}
	return counts
}

func solveDay2Part1(lines []string) int {
	sum := 0
	for _, line := range lines {
		gameId, gameSets := getGameInfo(line)
		isValid := validGame(gameSets)
		if isValid {
			sum += gameId
		}
	}
	return sum
}

func solveDay2Part2(lines []string) int {
	sum := 0
	for _, line := range lines {
		_, gameSets := getGameInfo(line)
		counts := getMinCounts(gameSets)
		sum += counts["red"] * counts["blue"] * counts["green"]
	}
	return sum
}

func main() {
	inputFile := "inputs/day2"
	//inputFile := "input"
	lines := utils.ReadFileLines(inputFile)
	fmt.Printf("Day 2 Part 1 Answer: %d\n", solveDay2Part1(lines))
	fmt.Printf("Day 2 Part 2 Answer: %d\n", solveDay2Part2(lines))
}
