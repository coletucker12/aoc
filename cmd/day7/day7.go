package main

import (
	"fmt"
	"github.com/coletucker12/aoc/utils"
	"slices"
	"strings"
)

var cardStrengthP1 = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	"J": 11,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
}

var cardStrengthP2 = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	"T": 11,
	"9": 10,
	"8": 9,
	"7": 8,
	"6": 7,
	"5": 6,
	"4": 5,
	"3": 4,
	"2": 3,
	"J": 2,
}

func calculateHandValue(hand string) int {
	var uniqCards []string
	cardCounts := make(map[string]int)
	for _, ch := range hand {
		if !slices.Contains(uniqCards, string(ch)) {
			uniqCards = append(uniqCards, string(ch))
		}
		cardCounts[string(ch)]++
	}

	if len(uniqCards) == 1 {
		// 5 of a kind
		return 6
	} else if len(uniqCards) == 2 {
		if cardCounts[uniqCards[0]] == 4 || cardCounts[uniqCards[1]] == 4 {
			// 4 of a kind
			return 5
		} else {
			// full house
			return 4
		}
	} else if len(uniqCards) == 3 {
		for _, val := range cardCounts {
			if val == 3 {
				// three of a kind
				return 3
			} else if val == 2 {
				// two pair
				return 2
			} else {
				continue
			}
		}
	} else if len(uniqCards) == 4 {
		return 1
	} else {
		return 0
	}
	return 0
}

func sortHands(hands []string) []string {
	slices.SortFunc(hands, func(handA, handB string) int {
		valA := calculateHandValue(handA)
		valB := calculateHandValue(handB)
		if valA > valB {
			return 1
		} else if valA == valB {
			for i := 0; i < len(handA); i++ {
				if handA[i] == handB[i] {
					continue
				}

				if cardStrengthP1[string(handA[i])] > cardStrengthP1[string(handB[i])] {
					return 1
				} else {
					return -1
				}
			}
		}

		return -1
	})
	return hands
}

func sortHandsPart2(hands []string) []string {
	var newHands []string
	for _, h := range hands {
		newHands = append(newHands, h)
	}

	slices.SortFunc(newHands, func(handA, handB string) int {
		valA := calculateHandValuePart2(handA)
		valB := calculateHandValuePart2(handB)
		if valA > valB {
			return 1
		} else if valA == valB {
			for i := 0; i < len(handA); i++ {
				if handA[i] == handB[i] {
					continue
				}

				if cardStrengthP2[string(handA[i])] > cardStrengthP2[string(handB[i])] {
					return 1
				} else {
					return -1
				}
			}
		}

		return -1
	})
	return newHands
}

func calculateHandValuePart2(hand string) int {
	var uniqCards []string
	cardCounts := make(map[string]int)
	for _, ch := range hand {
		if !slices.Contains(uniqCards, string(ch)) {
			uniqCards = append(uniqCards, string(ch))
		}
		cardCounts[string(ch)]++
	}

	if len(uniqCards) == 1 {
		// 5 of a kind
		return 6
	} else if len(uniqCards) == 2 {
		if cardCounts["J"] > 0 {
			// KKKKJ -> 5 of a kind
			// actually a 5 of a kind now because 4 + 1J
			return 6
		}

		if cardCounts[uniqCards[0]] == 4 || cardCounts[uniqCards[1]] == 4 {
			// 4 of a kind
			return 5
		} else {
			// full house
			return 4
		}
	} else if len(uniqCards) == 3 {
		// AAJ2A
		// KKJJA == KKKJA, always take 4 of a kind over full house
		if cardCounts["J"] == 2 {
			// actually a 4 of a kind now
			return 5
		} else if cardCounts["J"] == 1 {
			// Jxxxy
			// Jxxzy
			// Jxxyy
			// jjxyz
			// could be 4 of a kind or full house
			for k, val := range cardCounts {
				if val == 3 {
					// four of a kind now
					return 5
				} else if val == 2 {
					delete(cardCounts, k)
					for _, v2 := range cardCounts {
						if v2 == 2 {
							return 4
						} else {
							return 3
						}
					}

					// full house
					return 4
				} else {
					continue
				}
			}
		}

		for _, val := range cardCounts {
			if val == 3 {
				// three of a kind
				return 3
			} else if val == 2 {
				// two pair
				return 2
			} else {
				continue
			}
		}
	} else if len(uniqCards) == 4 {
		if cardCounts["J"] == 1 || cardCounts["J"] == 2 {
			// KJJ34
			// three of a kind now
			return 3
		}

		return 1
	} else {
		if cardCounts["J"] == 1 {
			// KJ234
			// actually a one pair now
			return 1
		}
		return 0
	}
	return 0
}

func solvePart1(lines []string) int {
	var hands []string
	bids := make(map[string]int)

	for _, line := range lines {
		s := strings.Split(line, " ")
		hands = append(hands, s[0])
		bids[s[0]] = utils.ParseNumber(s[1])
	}

	sortedHands := sortHands(hands)
	sum := 0
	for i, hand := range sortedHands {
		sum += bids[hand] * (i + 1)
	}
	return sum
}

func solvePart2(lines []string) int {
	var hands []string
	bids := make(map[string]int)

	for _, line := range lines {
		s := strings.Split(line, " ")
		hands = append(hands, s[0])
		bids[s[0]] = utils.ParseNumber(s[1])
	}

	//fmt.Println(hands)
	sortedHands := sortHandsPart2(hands)
	//fmt.Println(sortedHands)
	sum := 0
	for i, hand := range sortedHands {
		sum += bids[hand] * (i + 1)
		fmt.Printf("%s (%d)\n", sortedHands[i], calculateHandValuePart2(sortedHands[i]))
	}
	return sum
}

func main() {
	inputFile := "inputs/day7"
	//inputFile := "input"
	lines := utils.ReadFileLines(inputFile)
	fmt.Printf("Part 1 Answer: %d\n", solvePart1(lines))
	// 250666192
	fmt.Printf("Part 2 Answer: %d\n", solvePart2(lines))
}
