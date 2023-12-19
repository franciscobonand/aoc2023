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
		rowWeights[lineCount] = strings.Count(line, "O")
		platform = append(platform, strings.Split(line, ""))
		lineCount++
	}
	// After 180 cycles, the total weight starts repeating every 28 cycles
	for i := 0; i < 180; i++ {
		tiltNorth()
		tiltWest()
		tiltSouth()
		tiltEast()
	}
	// 1000000000 - 180 = 999999820, and 999999820 % 28 = 8
	// So after 1000000000 cycles, the total weight is the same as after 8 cycles after the 180th cycle
	for i := 0; i < 8; i++ {
		tiltNorth()
		tiltWest()
		tiltSouth()
		tiltEast()
	}
	fmt.Println("Final platform:")
	return totalWeight()
}

func tiltNorth() {
	for i := 1; i < len(platform); i++ {
		for j := 0; j < len(platform[i]); j++ {
			if platform[i][j] == "O" {
				up := i - 1
				curr := i
				for up >= 0 && platform[up][j] == "." {
					platform[curr][j] = "."
					platform[up][j] = "O"
					rowWeights[curr]--
					rowWeights[up]++
					curr--
					up--
				}
			}
		}
	}
}

func tiltSouth() {
	for i := len(platform) - 1; i >= 0; i-- {
		for j := 0; j < len(platform[i]); j++ {
			if platform[i][j] == "O" {
				down := i + 1
				curr := i
				for down < len(platform) && platform[down][j] == "." {
					platform[curr][j] = "."
					platform[down][j] = "O"
					rowWeights[curr]--
					rowWeights[down]++
					curr++
					down++
				}
			}
		}
	}
}

func tiltWest() {
	for i := 0; i < len(platform); i++ {
		for j := 1; j < len(platform[i]); j++ {
			if platform[i][j] == "O" {
				left := j - 1
				curr := j
				for left >= 0 && platform[i][left] == "." {
					platform[i][curr] = "."
					platform[i][left] = "O"
					curr--
					left--
				}
			}
		}
	}
}

func tiltEast() {
	for i := 0; i < len(platform); i++ {
		for j := len(platform[i]) - 1; j >= 0; j-- {
			if platform[i][j] == "O" {
				right := j + 1
				curr := j
				for right < len(platform[i]) && platform[i][right] == "." {
					platform[i][curr] = "."
					platform[i][right] = "O"
					curr++
					right++
				}
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

func printPlatform() {
	for _, row := range platform {
		fmt.Println(row)
	}
	fmt.Println()
}
