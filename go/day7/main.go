package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Hand struct {
	Cards []int
	Bid   int
	Rank  int
	Type  Type
}

func (h Hand) String() string {
	return fmt.Sprintf("Cards: %v, Bid: %d, Rank: %d, Type: %d", h.Cards, h.Bid, h.Rank, h.Type)
}

var CardValues = map[string]int{
	"2": 0,
	"3": 1,
	"4": 2,
	"5": 3,
	"6": 4,
	"7": 5,
	"8": 6,
	"9": 7,
	"T": 8,
	"J": 9,
	"Q": 10,
	"K": 11,
	"A": 12,
}

var CardValuesWithJoker = map[string]int{
	"J": 0,
	"2": 1,
	"3": 2,
	"4": 3,
	"5": 4,
	"6": 5,
	"7": 6,
	"8": 7,
	"9": 8,
	"T": 9,
	"Q": 10,
	"K": 11,
	"A": 12,
}

type Type int

const (
	HighCard Type = iota
	OnePair
	TwoPair
	ThreeOAK
	FullHouse
	FourOAK
	FiveOAK
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <input file>\n", os.Args[0])
		return
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("Open file error: %s\n", err)
		return
	}
	defer f.Close()

	fmt.Println(getSolutions(f))
}

func getSolutions(f *os.File) int {
	scanner := bufio.NewScanner(f)
	hands := []Hand{}
	typeMapper := map[string]int{}
	for scanner.Scan() {
		line := scanner.Text()
		hands = append(hands, parseHand(line, typeMapper))
		clear(typeMapper)
	}
	slices.SortFunc(hands, func(h1, h2 Hand) int {
		return compareHands(h1, h2)
	})
	return getTotalWinnings(hands)
}

func parseHand(line string, typeMapper map[string]int) Hand {
	cards := []int{}
	input := strings.Split(line, " ")
	bid, _ := strconv.Atoi(input[1])
	hasJoker := false
	for _, r := range input[0] {
		card := string(r)
		if card == "J" {
			hasJoker = true
		}
		// cards = append(cards, CardValues[card])
		cards = append(cards, CardValuesWithJoker[card])
		typeMapper[card]++
	}
	if hasJoker {
		return Hand{Cards: cards, Bid: bid, Type: getTypeWithJoker(typeMapper)}
	}
	return Hand{Cards: cards, Bid: bid, Type: getType(typeMapper)}
}

func getTotalWinnings(hands []Hand) int {
	total := 0
	for idx, h := range hands {
		total += h.Bid * (idx + 1)
	}
	return total
}

func getTypeWithJoker(typeMapper map[string]int) Type {
	jokerCount := typeMapper["J"]
	switch len(typeMapper) {
	case 5:
		return OnePair
	case 4:
		return ThreeOAK
	case 3:
		for _, v := range typeMapper {
			if v == 3 {
				return FourOAK
			}
		}
		if jokerCount == 2 {
			return FourOAK
		}
		return FullHouse
	default:
		return FiveOAK
	}
}

func getType(typeMapper map[string]int) Type {
	switch len(typeMapper) {
	case 5:
		return HighCard
	case 4:
		return OnePair
	case 3:
		for _, v := range typeMapper {
			if v == 3 {
				return ThreeOAK
			}
		}
		return TwoPair
	case 2:
		for _, v := range typeMapper {
			if v == 3 {
				return FullHouse
			}
		}
		return FourOAK
	default:
		return FiveOAK
	}
}

func compareHands(h1, h2 Hand) int {
	if h1.Type > h2.Type {
		return 1
	} else if h1.Type < h2.Type {
		return -1
	}
	for i := 0; i < len(h1.Cards); i++ {
		if h1.Cards[i] > h2.Cards[i] {
			return 1
		} else if h1.Cards[i] < h2.Cards[i] {
			return -1
		}
	}
	return 0
}
