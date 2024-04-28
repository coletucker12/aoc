package main

import (
	"fmt"
	"github.com/coletucker12/aoc/utils"
	"slices"
	"strconv"
)

type Part struct {
	Value                 int
	Valid                 bool
	LineIndex             int
	Line                  string
	PreviousLine          string
	NextLine              string
	CharIndexes           []int
	SurroundingCharacters []SurroundingCharacter
}

type SurroundingCharacter struct {
	Value     string
	LineIndex int
	CharIndex int
}

type Gear struct {
	PartOne Part
	PartTwo Part
	Value   int
}

func isNumber(s string) bool {
	_, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	return true
}

func parseNumber(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		utils.Check(err)
		return -1
	}
	return val
}

func parseLine(line string, lineIndex int, prevLine string, nextLine string) []Part {
	var parts []Part
	buffer := ""
	var indexes []int
	for i, ch := range line {
		if isNumber(string(ch)) {
			buffer = buffer + string(ch)
			indexes = append(indexes, i)
			if i < len(line)-1 {
				continue
			}
		}
		if isNumber(buffer) {
			num := parseNumber(buffer)
			parts = append(parts, Part{
				Value:                 num,
				Valid:                 false,
				Line:                  line,
				LineIndex:             lineIndex,
				PreviousLine:          prevLine,
				NextLine:              nextLine,
				CharIndexes:           indexes,
				SurroundingCharacters: nil,
			})
			buffer = ""
			indexes = []int{}
		}
	}
	return parts
}

func getPartSurroundingCharacters(part Part) []SurroundingCharacter {
	var surroundingChars []SurroundingCharacter
	for _, index := range part.CharIndexes {
		if index-1 >= 0 {
			surroundingChars = append(surroundingChars, SurroundingCharacter{
				Value:     string(part.Line[index-1]),
				LineIndex: part.LineIndex,
				CharIndex: index - 1,
			})

			if part.PreviousLine != "" {
				surroundingChars = append(surroundingChars, SurroundingCharacter{
					Value:     string(part.PreviousLine[index-1]),
					LineIndex: part.LineIndex - 1,
					CharIndex: index - 1,
				})
			}

			if part.NextLine != "" {
				surroundingChars = append(surroundingChars, SurroundingCharacter{
					Value:     string(part.NextLine[index-1]),
					LineIndex: part.LineIndex + 1,
					CharIndex: index - 1,
				})
			}
		}

		if index+1 < len(part.Line) {
			surroundingChars = append(surroundingChars, SurroundingCharacter{
				Value:     string(part.Line[index+1]),
				LineIndex: part.LineIndex,
				CharIndex: index + 1,
			})
			if part.PreviousLine != "" {
				surroundingChars = append(surroundingChars, SurroundingCharacter{
					Value:     string(part.PreviousLine[index+1]),
					LineIndex: part.LineIndex - 1,
					CharIndex: index + 1,
				})
			}
			if part.NextLine != "" {
				surroundingChars = append(surroundingChars, SurroundingCharacter{
					Value:     string(part.NextLine[index+1]),
					LineIndex: part.LineIndex + 1,
					CharIndex: index + 1,
				})
			}
		}

		if part.PreviousLine != "" {
			surroundingChars = append(surroundingChars, SurroundingCharacter{
				Value:     string(part.PreviousLine[index]),
				LineIndex: part.LineIndex - 1,
				CharIndex: index,
			})
		}

		if part.NextLine != "" {
			surroundingChars = append(surroundingChars, SurroundingCharacter{
				Value:     string(part.NextLine[index]),
				LineIndex: part.LineIndex + 1,
				CharIndex: index,
			})
		}
	}
	// check if any surrounding char is not a '.'
	return surroundingChars
}

func findContiguousDigits(lines []string, initialVal string, lnIdx int, chIdx int) int {
	buffer := initialVal
	i := chIdx - 1
	ch := string(lines[lnIdx][i])
	isNum := isNumber(ch)
	for isNum {
		buffer = string(lines[lnIdx][i]) + buffer
		i--
		if i < 0 {
			break
		}

		ch = string(lines[lnIdx][i])
		isNum = isNumber(ch)
	}

	i = chIdx + 1
	ch = string(lines[lnIdx][i])
	isNum = isNumber(ch)
	for isNum {
		buffer = buffer + string(lines[lnIdx][i])
		i++
		if i >= len(lines[lnIdx]) {
			break
		}
		ch = string(lines[lnIdx][i])
		isNum = isNumber(ch)
	}
	return parseNumber(buffer)
}

func findSurroundingParts(part Part, lines []string, char SurroundingCharacter) Part {
	// find parts around this character
	for i := char.LineIndex - 1; i <= char.LineIndex+1 && i >= 0 && i < len(lines); i++ {
		for j := char.CharIndex - 1; j <= char.CharIndex+1 && j >= 0 && j < len(lines[i]); j++ {
			c := string(lines[i][j])
			if isNumber(c) {
				num := findContiguousDigits(lines, c, i, j)
				if num != part.Value {
					return Part{
						Value:                 num,
						Valid:                 true,
						LineIndex:             i,
						Line:                  lines[i],
						PreviousLine:          "",
						NextLine:              "",
						CharIndexes:           nil,
						SurroundingCharacters: nil,
					}
				}
			}
		}
	}
	return Part{}
}

func solveDay3Part1(lines []string) int {
	sum := 0
	for lineIdx, line := range lines {
		prevLine := ""
		if lineIdx-1 >= 0 {
			prevLine = lines[lineIdx-1]
		}
		nextLine := ""
		if lineIdx+1 < len(lines) {
			nextLine = lines[lineIdx+1]
		}

		parts := parseLine(line, lineIdx, prevLine, nextLine)
		for _, part := range parts {
			part.SurroundingCharacters = getPartSurroundingCharacters(part)
			for _, ch := range part.SurroundingCharacters {
				if ch.Value != "." && !isNumber(ch.Value) {
					part.Valid = true
				}
			}
			if part.Valid {
				sum += part.Value
			}
		}
	}
	return sum
}

func solveDay3Part2(lines []string) int64 {
	var sum int64
	var gears []Gear
	counter := make(map[string]int)
	for lineIdx, line := range lines {
		prevLine := ""
		if lineIdx-1 >= 0 {
			prevLine = lines[lineIdx-1]
		}
		nextLine := ""
		if lineIdx+1 < len(lines) {
			nextLine = lines[lineIdx+1]
		}

		parts := parseLine(line, lineIdx, prevLine, nextLine)
		for _, part := range parts {
			part.SurroundingCharacters = getPartSurroundingCharacters(part)
			for _, ch := range part.SurroundingCharacters {
				if ch.Value == "*" {
					part2 := findSurroundingParts(part, lines, ch)
					gear := Gear{
						PartOne: part,
						PartTwo: part2,
						Value:   part.Value * part2.Value,
					}
					if !slices.ContainsFunc(gears, func(g Gear) bool {
						return (g.PartOne.Value == gear.PartOne.Value && g.PartTwo.Value == gear.PartTwo.Value) || (g.PartTwo.Value == gear.PartOne.Value && g.PartOne.Value == gear.PartTwo.Value)
					}) {
						//fmt.Printf("adding %d * %d = %d\n", gear.PartOne.Value, gear.PartTwo.Value, gear.Value)
						counter[fmt.Sprintf("%d * %d", part.Value, part2.Value)]++
						gears = append(gears, gear)
					} else {
						//fmt.Printf("not adding %d * %d = %d\n", gear.PartOne.Value, gear.PartTwo.Value, gear.Value)
					}
				}
			}
		}
	}

	gearCount := make(map[string]int)
	for _, g := range gears {
		//fmt.Printf("%d * %d = %d\n", g.PartOne.Value, g.PartTwo.Value, g.Value)
		if g.Value > 0 {
			sum += int64(g.Value)
		}
		gearCount[strconv.Itoa(g.Value)]++
	}
	return sum
}

// 81,517,405 -- 81,709,807
func main() {
	inputFile := "inputs/day3"
	//inputFile := "input"
	lines := utils.ReadFileLines(inputFile)
	fmt.Printf("Day 3 Part 1 Answer: %d\n", solveDay3Part1(lines))
	fmt.Printf("Day 3 Part 2 Answer: %d\n", solveDay3Part2(lines))
}
