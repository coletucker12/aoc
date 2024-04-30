package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/coletucker12/aoc/utils"
)

type Galaxy struct {
	Id string
	X  int
	Y  int
}

type GalaxyPair struct {
	PairId   string
	GalaxyA  Galaxy
	GalaxyB  Galaxy
	Distance int
}

func main() {
	start := time.Now()
	defer utils.PrintTimeElapsed(start, "This")

	inputFile := "inputs/day11.txt"
	// inputFile := "input.txt"
	lines := utils.ReadFileLines(inputFile)
	fmt.Printf("Part 1 Answer: %d\n\n", solvePart1(lines))
	// fmt.Printf("Part 2 Answer: %d\n\n", solvePart2(lines))
}

func solvePart1(lines []string) int {
	universe := make([]string, 0)
	galaxies := make([]Galaxy, 0)

	for _, line := range lines {
		universe = append(universe, line)
		// If line contains no galaxies, double it
		if strings.Index(line, "#") == -1 {
			universe = append(universe, line)
		}
	}
	universe = utils.DeleteEmptyLines(universe)

	// check for columns with no galaxy and duplicate them if so
	numCols := len(universe[0])
	for i := 0; i < numCols; i++ {
		hasGalaxy := false
		for _, row := range universe {
			if row[i] == '#' {
				hasGalaxy = true
			}
		}

		if !hasGalaxy {
			for rowIdx, row := range universe {
				newRow := row[:i] + "." + row[i:]
				universe[rowIdx] = newRow
			}
			// Skip the column we are adding
			i++
		}
		hasGalaxy = false
		numCols = len(universe[0])
	}
	// printUniverse(universe)

	// find all the galaxies
	for rowIdx, row := range universe {
		for _, idx := range findAllGalaxyIndexes(row) {
			galaxies = append(galaxies, Galaxy{
				Id: strconv.Itoa(len(galaxies) + 1),
				X:  rowIdx,
				Y:  idx,
			})
		}
	}

	sum := 0
	for i, gA := range galaxies {
		for _, gB := range galaxies[i+1:] {
			sum += calculateDistance(gA, gB)
		}
	}
	return sum
}

func printUniverse(lst []string) {
	for _, r := range lst {
		fmt.Println(r)
	}
}

func findAllGalaxyIndexes(line string) []int {
	indexes := make([]int, 0)
	for i, ch := range line {
		if string(ch) == "#" {
			indexes = append(indexes, i)
		}
	}
	return indexes
}

func calculateDistance(galaxyA, galaxyB Galaxy) int {
	diffX := math.Abs(float64(galaxyA.X) - float64(galaxyB.X))
	diffY := math.Abs(float64(galaxyA.Y) - float64(galaxyB.Y))
	return int(diffX) + int(diffY)
}
