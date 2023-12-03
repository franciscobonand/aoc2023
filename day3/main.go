package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// line => EnginePart
type EnginePart struct {
	Value    int
	From, To int
}

type Symbol struct {
	Position int
	Value    string
}

func (ep EnginePart) String() string {
	return fmt.Sprintf("%d from index %d to %d", ep.Value, ep.From, ep.To)
}

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
	lineCount := 0
	possibleEngParts := map[int][]EnginePart{}
	symbols := map[int][]Symbol{}
	// Extracts all numbers
	for scanner.Scan() {
		line := scanner.Text()
		n := ""
		readingNum := false
		for idx, r := range line {
			if isDigit(r) {
				readingNum = true
				n += string(r)
				if idx < len(line)-1 {
					continue
				}
			}
			if isSymbol(r) {
				if _, ok := symbols[lineCount]; !ok {
					symbols[lineCount] = []Symbol{}
				}
				symbols[lineCount] = append(symbols[lineCount], Symbol{Position: idx, Value: string(r)})
			}
			if readingNum {
				v, _ := strconv.Atoi(n)
				ep := EnginePart{
					Value: v,
				}
				if isDigit(r) {
					ep.From = idx - (len(n) - 1)
					ep.To = idx
				} else {
					ep.From = idx - len(n)
					ep.To = idx - 1
				}
				if _, ok := possibleEngParts[lineCount]; !ok {
					possibleEngParts[lineCount] = []EnginePart{}
				}
				possibleEngParts[lineCount] = append(possibleEngParts[lineCount], ep)
				readingNum = false
				n = ""
			}
		}
		lineCount++
	}
	// Search for engine parts
	engPartCount, gearRatio := getEnginePartSumAndGearRatio(lineCount, symbols, possibleEngParts)
	return engPartCount, gearRatio
}

func getEnginePartSumAndGearRatio(lineCount int, symbols map[int][]Symbol, possibleEngParts map[int][]EnginePart) (int, int) {
	partSum := 0
	gearRatio := 0
	for line := 0; line <= lineCount; line++ {
		if _, ok := symbols[line]; !ok {
			continue
		}
		v := symbols[line]
		for _, symb := range v {
			isPossibleGear := symb.Value == "*"
			adjPartCount := 0
			partValues := []int{}
			// line above
			if _, ok := possibleEngParts[line-1]; ok {
				val, count := getEnginePartValueAndCount(line-1, symb.Position, possibleEngParts)
				partValues = append(partValues, val...)
				adjPartCount += count
			}
			// same line
			if _, ok := possibleEngParts[line]; ok {
				val, count := getEnginePartValueAndCount(line, symb.Position, possibleEngParts)
				partValues = append(partValues, val...)
				adjPartCount += count
			}
			// line below
			if _, ok := possibleEngParts[line+1]; ok {
				val, count := getEnginePartValueAndCount(line+1, symb.Position, possibleEngParts)
				partValues = append(partValues, val...)
				adjPartCount += count
			}
			if isPossibleGear && adjPartCount == 2 {
				gearRatio += partValues[0] * partValues[1]
			}
			for _, v := range partValues {
				partSum += v
			}
		}
	}
	return partSum, gearRatio
}

func getEnginePartValueAndCount(line, symbolPos int, possibleEngParts map[int][]EnginePart) ([]int, int) {
	val := []int{}
	count := 0
	for _, n := range possibleEngParts[line] {
		if isAdjencent(n.To, n.From, symbolPos) {
			val = append(val, n.Value)
			count++
		}
		if n.From > symbolPos+1 {
			break
		}
	}
	return val, count
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func isSymbol(r rune) bool {
	return !isDigit(r) && r != '.'
}

func isAdjencent(to, from, target int) bool {
	return from <= target && target <= to ||
		to == target-1 ||
		from == target+1
}
