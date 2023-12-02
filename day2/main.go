package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var MAX_CUBES = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
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

	fmt.Println(sumOfPossibleGames(f))
}

func sumOfPossibleGames(f *os.File) int {
	scanner := bufio.NewScanner(f)
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		game := strings.Split(line, ": ")
		// if gameIsValid(game[1]) {
		// 	g := strings.Split(game[0], " ")
		// 	val, _ := strconv.Atoi(g[1])
		// 	total += val
		// }
		total += minCubeAmount(game[1])
	}
	return total
}

func gameIsValid(s string) bool {
	rounds := strings.Split(s, "; ")
	for _, round := range rounds {
		cubes := strings.Split(round, ", ")
		for _, cube := range cubes {
			c := strings.Split(cube, " ")
			val, _ := strconv.Atoi(c[0])
			if val > MAX_CUBES[c[1]] {
				return false
			}
		}
	}
	return true
}

func minCubeAmount(s string) int {
	minCubes := map[string]int{}
	rounds := strings.Split(s, "; ")
	for _, round := range rounds {
		cubes := strings.Split(round, ", ")
		for _, cube := range cubes {
			c := strings.Split(cube, " ")
			val, _ := strconv.Atoi(c[0])
			if currVal, ok := minCubes[c[1]]; !ok || val > currVal {
				minCubes[c[1]] = val
			}
		}
	}
	return minCubes["red"] * minCubes["green"] * minCubes["blue"]
}
