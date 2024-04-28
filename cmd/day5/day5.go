package main

import (
	"fmt"
	"github.com/coletucker12/aoc/utils"
	"strings"
)

func solveDay5Part1(lines string) int {
	maps := strings.Split(lines, "\n")
	seeds := strings.Split(strings.TrimSpace(strings.Split(maps[0], ":")[1]), " ")
	fmt.Println(seeds)

	// seed -> soil
	// soil -> fertilizer
	// fertilizer -> water
	// water -> light
	// light -> temp
	// temp -> humidity
	// humidity -> location
	//var seedToSoil map[string]string
	//var soilToFertilizer map[string]string
	//var fertilizerToWater map[string]string
	//var waterToLight map[string]string
	//var lightToTemp map[string]string
	//var tempToHumidity map[string]string
	//var humidityToLocation map[string]string
	for _, line := range maps[2 : len(maps)-1] {
		fmt.Println(" - " + line)
		if line == "" {
			continue
		}
	}
	return 0
}

func main() {
	//inputFile := "inputs/day5"
	inputFile := "input"
	lines := utils.ReadFileLinesUnsplit(inputFile)
	fmt.Printf("Day 5 Part 1 Answer: %d\n", solveDay5Part1(lines))
	//fmt.Printf("Day 4 Part 2 Answer: %d\n", solveDay4Part2(lines))
}
