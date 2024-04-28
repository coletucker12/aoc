package main

import (
	"fmt"
	"github.com/coletucker12/aoc/utils"
	"strings"
)

type Race struct {
	Duration       int
	RecordDistance int
}

func calculateMarginOfError(race Race) int {
	var options []int
	for t := 1; t <= race.Duration; t++ {
		speed := t
		distance := speed * (race.Duration - t)
		if distance > race.RecordDistance {
			options = append(options, t)
		}
	}
	return len(options)
}

func solveDay6Part1(lines []string) int {
	var races []Race
	timesStr := strings.TrimSpace(strings.Split(lines[0], ":")[1])
	distancesStr := strings.TrimSpace(strings.Split(lines[1], ":")[1])
	times := utils.DeleteEmptyLines(strings.Split(timesStr, " "))
	distances := utils.DeleteEmptyLines(strings.Split(distancesStr, " "))

	for idx, time := range times {
		races = append(races, Race{
			Duration:       utils.ParseNumber(time),
			RecordDistance: utils.ParseNumber(distances[idx]),
		})
	}

	sum := 1
	for _, race := range races {
		sum *= calculateMarginOfError(race)
	}
	return sum
}

func solveDay4Part2(lines []string) int {
	timesStr := strings.TrimSpace(strings.Split(lines[0], ":")[1])
	distancesStr := strings.TrimSpace(strings.Split(lines[1], ":")[1])
	times := utils.DeleteEmptyLines(strings.Split(timesStr, " "))
	distances := utils.DeleteEmptyLines(strings.Split(distancesStr, " "))

	time := ""
	distance := ""
	for idx, t := range times {
		time = time + t
		distance = distance + distances[idx]
	}

	val := calculateMarginOfError(Race{
		Duration:       utils.ParseNumber(time),
		RecordDistance: utils.ParseNumber(distance),
	})
	return val
}

func main() {
	inputFile := "inputs/day6"
	//inputFile := "input"
	lines := utils.ReadFileLines(inputFile)
	fmt.Printf("Day 6 Part 1 Answer: %d\n", solveDay6Part1(lines))
	fmt.Printf("Day 4 Part 2 Answer: %d\n", solveDay4Part2(lines))
}
