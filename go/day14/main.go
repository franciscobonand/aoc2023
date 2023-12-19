package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

var platform = [][]string{}
var rowWeights []int

func getSolutions(f *os.File) int {
	scanner := bufio.NewScanner(f)
	lineCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		if lineCount == 0 {
			// Sample and input are square matrices
			rowWeights = make([]int, len(line))
		}
		addToPlatformTiltingNorth(line, lineCount)
		lineCount++
	}
	return totalWeight()
}

func addToPlatformTiltingNorth(line string, idx int) {
	rowWeights[idx] = strings.Count(line, "O")
	platform = append(platform, strings.Split(line, ""))
	if idx == 0 {
		return
	}
	for i := idx; i > 0; i-- {
		for j := 0; j < len(platform[i]); j++ {
			if platform[i][j] == "O" && platform[i-1][j] == "." {
				platform[i][j] = "."
				platform[i-1][j] = "O"
				rowWeights[i]--
				rowWeights[i-1]++
			}
		}
	}
}

func totalWeight() int {
	total := 0
	for idx, w := range rowWeights {
		total += w * (len(rowWeights) - idx)
	}
	return total
}
