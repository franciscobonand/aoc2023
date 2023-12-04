package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// CardMatches represent the relation of Card N (idx + 1) => quantity of matched numbers
type CardMatches []int

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

func getSolutions(f *os.File) (int, int) {
	scanner := bufio.NewScanner(f)
	total := 0
	currCard := 1
	origCards := CardMatches{}
	cardCount := map[int]int{}
	for scanner.Scan() {
		line := scanner.Text()
		cardNums := strings.Split(line, ": ")
		wCards, recCards := getCardNumbers(cardNums[1])
		p, c := getCardPointsAndMatchCount(wCards, recCards)
		origCards = append(origCards, c)
		cardCount[currCard] = 1
		total += p
		currCard++
	}
	return total, countTotalCards(origCards, cardCount)
}

func countTotalCards(oCards CardMatches, cCount map[int]int) int {
	for cardNum, matchCount := range oCards {
		currCardQnt := cCount[cardNum+1]
		for i := 1; i <= matchCount; i++ {
			cCount[cardNum+i+1] += currCardQnt
		}
	}
	count := 0
	for _, val := range cCount {
		count += val
	}
	return count
}

func getCardPointsAndMatchCount(wCards, recCards []string) (int, int) {
	points := 0
	count := 0
	for _, wnum := range wCards {
		if wnum == "" {
			continue
		}
		for _, recnum := range recCards {
			if wnum == recnum {
				count++
				if points != 0 {
					points *= 2
					continue
				}
				points++
			}
		}
	}
	return points, count
}

func getCardNumbers(card string) ([]string, []string) {
	tmp := strings.Split(card, " | ")
	wCards := strings.Split(tmp[0], " ")
	recCards := strings.Split(tmp[1], " ")
	return wCards, recCards
}
