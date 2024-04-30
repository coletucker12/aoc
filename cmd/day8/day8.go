package main

import (
	"fmt"
	"github.com/coletucker12/aoc/utils"
	"strings"
)

type Instruction struct {
	Value string
}

type Node struct {
	Id    string
	Left  Instruction
	Right Instruction
}

func parseLine(line string) (string, string, string) {
	s := strings.Split(line, "=")
	id := s[0]
	values := utils.DeleteEmptyLines(strings.Split(strings.TrimSpace(s[1]), ","))
	leftValue := strings.TrimSpace(strings.Replace(values[0], "(", "", 1))
	rightValue := strings.TrimSpace(strings.Replace(values[1], ")", "", 1))
	return strings.TrimSpace(id), leftValue, rightValue
}

func solvePart1(lines []string) int {
	instructions := strings.TrimSpace(lines[0])
	nodeLefts := make(map[string]string)
	nodeRights := make(map[string]string)
	for _, line := range lines[2 : len(lines)-1] {
		val, left, right := parseLine(line)
		nodeLefts[val] = left
		nodeRights[val] = right
	}

	cur := "AAA"
	finish := "ZZZ"
	i := 0
	numSteps := 0
	for cur != finish {
		if cur == "" {
			break
		}
		if i == len(instructions) {
			i = 0
		}

		if string(instructions[i]) == "L" {
			cur = nodeLefts[cur]
		} else if string(instructions[i]) == "R" {
			cur = nodeRights[cur]
		}
		i++
		numSteps++
	}
	return numSteps
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func solvePart2(lines []string) int {
	instructions := strings.TrimSpace(lines[0])
	var nodes []string
	nodeLefts := make(map[string]string)
	nodeRights := make(map[string]string)
	for _, line := range lines[2 : len(lines)-1] {
		val, left, right := parseLine(line)
		if strings.HasSuffix(val, "A") {
			nodes = append(nodes, val)
		}
		nodeLefts[val] = left
		nodeRights[val] = right
	}
	fmt.Println(len(nodes))

	var distances []int
	for _, node := range nodes {
		cur := node
		i := 0
		numSteps := 0
		for !strings.HasSuffix(cur, "Z") {
			if cur == "" {
				break
			}

			if i == len(instructions) {
				i = 0
			}

			if string(instructions[i]) == "L" {
				cur = nodeLefts[cur]
			} else if string(instructions[i]) == "R" {
				cur = nodeRights[cur]
			}

			i++
			numSteps++
		}
		distances = append(distances, numSteps)
	}
	return LCM(distances[0], distances[1], distances[2], distances[3], distances[4], distances[5])
}

func main() {
	inputFile := "inputs/day8"
	// inputFile := "input"
	lines := utils.ReadFileLines(inputFile)
	fmt.Printf("Part 1 Answer: %d\n", solvePart1(lines))
	fmt.Printf("Part 2 Answer: %d\n", solvePart2(lines))
}
