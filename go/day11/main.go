package main

import (
	"bufio"
	"fmt"
	"os"
)

// Part 1
// const ExpansionSize = 1

// Part 2
// This value must be 10^6 - 1 so the first step in the expansion is not counted twice
const ExpansionSize = 1000000 - 1

type Coord struct {
	x, y int
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

func getSolutions(f *os.File) int {
	scanner := bufio.NewScanner(f)
	lineCount := 0
	maze := []string{}
	vExp := map[int]bool{}
	galaxies := []Coord{}
	for scanner.Scan() {
		line := scanner.Text()
		expand := true
		for idx, r := range line {
			if r == '#' {
				expand = false
				galaxies = append(galaxies, Coord{idx, lineCount})
			}
		}
		if expand {
			vExp[lineCount] = true
		}
		maze = append(maze, line)
		lineCount++
	}
	hExp := getHorizontalExpansions(maze)
	return sumGalaxyDistances(galaxies, hExp, vExp)
}

func getHorizontalExpansions(maze []string) map[int]bool {
	verticalExpansions := map[int]bool{}
	for i := 0; i < len(maze[0]); i++ {
		isExpantion := true
		for j := 0; j < len(maze); j++ {
			if maze[j][i] == '#' {
				isExpantion = false
				break
			}
		}
		if isExpantion {
			verticalExpansions[i] = true
		}
	}
	return verticalExpansions
}

func sumGalaxyDistances(galaxies []Coord, hExp map[int]bool, vExp map[int]bool) int {
	sum := 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			acc := abs(galaxies[j].x-galaxies[i].x) + abs(galaxies[j].y-galaxies[i].y)
			if galaxies[j].x != galaxies[i].x {
				acc += getExpansionSize(galaxies[i].x, galaxies[j].x, hExp)
			}
			if galaxies[j].y != galaxies[i].y {
				acc += getExpansionSize(galaxies[i].y, galaxies[j].y, vExp)
			}
			sum += acc
		}
	}
	return sum
}

func getExpansionSize(n1 int, n2 int, exp map[int]bool) int {
	sum := 0
	a := n1
	b := n2
	if n1 > n2 {
		a = n2
		b = n1
	}
	for i := a + 1; i < b; i++ {
		if exp[i] {
			sum += ExpansionSize
		}
	}
	return sum
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
