package main

import (
	"fmt"
	"github.com/coletucker12/aoc/utils"
	"math"
	"slices"
	"strings"
)

type Card struct {
	Id             int
	PointValue     int
	NumMatches     int
	NumCopies      int
	Winning        bool
	WinningNumbers []int
	MyNumbers      []int
}

func parseLine(line string) Card {
	card := Card{
		Id:             0,
		PointValue:     0,
		NumMatches:     0,
		NumCopies:      1,
		WinningNumbers: nil,
		MyNumbers:      nil,
		Winning:        false,
	}
	s1 := strings.Split(line, ":")
	cardInfo := strings.Split(s1[0], " ")
	card.Id = utils.ParseNumber(cardInfo[len(cardInfo)-1])
	cardNumbers := strings.Split(s1[1], "|")
	for _, num := range strings.Split(strings.TrimSpace(cardNumbers[0]), " ") {
		if num != "" {
			card.WinningNumbers = append(card.WinningNumbers, utils.ParseNumber(strings.TrimSpace(num)))
		}
	}

	for _, num := range strings.Split(strings.TrimSpace(cardNumbers[1]), " ") {
		if num != "" {
			card.MyNumbers = append(card.MyNumbers, utils.ParseNumber(strings.TrimSpace(num)))
		}
	}
	return card
}

func calculateCardValue(card Card) Card {
	numMatches := 0
	for _, num := range card.MyNumbers {
		if slices.Contains(card.WinningNumbers, num) {
			numMatches++
		}
	}
	card.NumMatches = numMatches
	card.PointValue = int(math.Pow(2, float64(numMatches-1)))
	if numMatches > 0 {
		card.Winning = true
	}
	return card
}

func solveDay4Part1(lines []string) int {
	sum := 0
	for _, line := range lines {
		card := parseLine(line)
		card = calculateCardValue(card)
		sum += card.PointValue
	}
	return sum
}

func solveDay4Part2(lines []string) int {
	var cards []Card
	for _, line := range lines {
		card := parseLine(line)
		card = calculateCardValue(card)
		cards = append(cards, card)
	}

	for cardIdx, card := range cards {
		for j := 0; j < card.NumCopies; j++ {
			for i := 0; i < card.NumMatches; i++ {
				cards[i+cardIdx+1].NumCopies++
			}
		}
	}
	sum := 0
	for _, card := range cards {
		sum += card.NumCopies
	}
	return sum
}

func main() {
	inputFile := "inputs/day4"
	//inputFile := "input"
	lines := utils.ReadFileLines(inputFile)
	fmt.Printf("Day 4 Part 1 Answer: %d\n", solveDay4Part1(lines))
	fmt.Printf("Day 4 Part 2 Answer: %d\n", solveDay4Part2(lines))
}
