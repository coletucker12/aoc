package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/coletucker12/aoc/utils"
)

type Pos struct {
	x int
	y int
}

type Pipe struct {
	x    int
	y    int
	pipe byte
	dist int
}

var pipeTypeList = []byte{'-', '|', 'L', 'J', 'F', '7'}

func main() {
	start := time.Now()
	defer utils.PrintTimeElapsed(start, "This")

	// inputFile := "inputs/day10.txt"
	inputFile := "input.txt"
	lines := utils.ReadFileLines(inputFile)
	fmt.Printf("Part 1 Answer: %d\n\n", solvePart1(lines))
	// fmt.Printf("Part 2 Answer: %d\n\n", solvePart2(lines))
}

func solvePart1(lines []string) int {
	maze := make([]string, 0)
	startPos := Pos{
		x: 0,
		y: 0,
	}

	i := 0
	for _, line := range lines {
		maze = append(maze, line)

		idx := strings.Index(line, "S")
		if idx != -1 {
			startPos.x = idx
			startPos.y = i
		}
		i++
	}

	fmt.Println(startPos)
	return 0
}
func identifyStartingPipeType(maze []string, pos Pos) (Pos, Pos) {
	for _, pipe := range pipeTypeList {
		if found, res1, res2 := findConnectedPipes(maze, pipe, pos); found {
			fmt.Println("Found Starting Pipe Type: ", string(pipe))
			return res1, res2
		}
	}
	fmt.Println("Failed to identify the starting pipe")
	return Pos{0, 0}, Pos{0, 0}
}

func findConnectedPipes(maze []string, currentPipe byte, pos Pos) (bool, Pos, Pos) {
	up := Pos{pos.x, pos.y - 1}
	down := Pos{pos.x, pos.y + 1}
	left := Pos{pos.x - 1, pos.y}
	right := Pos{pos.x + 1, pos.y}

	switch currentPipe {
	case '-':
		return false, left, right
	case '|':
		return false, up, down
	case 'L':
		return false, down, right
	}

	return false, up, down
}
