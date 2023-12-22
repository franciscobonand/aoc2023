package main

import (
	"bufio"
	"fmt"
	"os"
)

// mirror '/'
var mirrorReflection1 = map[string]string{
	"R": "U",
	"U": "R",
	"L": "D",
	"D": "L",
}

// mirror '\'
var mirrorReflection2 = map[string]string{
	"R": "D",
	"D": "R",
	"L": "U",
	"U": "L",
}

type point struct {
	x, y      int
	direction string
}

var visitedTiles [][]int

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
	grid := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, line)
	}
	max := 0
	// Top row
	if val := valueStartingFromXEdges(grid, "D", 0, max); val > max {
		max = val
	}
	// Bottom row
	if val := valueStartingFromXEdges(grid, "U", len(grid)-1, max); val > max {
		max = val
	}
	// Left column
	if val := valueStartingFromYEdges(grid, "R", 0, max); val > max {
		max = val
	}
	// Right column
	if val := valueStartingFromYEdges(grid, "L", len(grid[0])-1, max); val > max {
		max = val
	}
	// Corners
	if val := valueStartingFromCorners(grid, max); val > max {
		max = val
	}
	return max
}

func totalEnergizedTiles(grid []string, start point) int {
	visitedTiles = make([][]int, len(grid))
	for i := range visitedTiles {
		visitedTiles[i] = make([]int, len(grid[i]))
	}
	visited := map[string]bool{}
	nextSteps := []point{start}
	for {
		if len(nextSteps) == 0 {
			break
		}
		currStep := nextSteps[0]
		nextSteps = nextSteps[1:]
		if isVisited(visited, currStep.x, currStep.y, currStep.direction) {
			continue
		}
		visited[fmt.Sprintf("x%d,y%d,dir%s", currStep.x, currStep.y, currStep.direction)] = true
		nextSteps = append(nextSteps, getNextSteps(grid, currStep, visited)...)
	}
	total := 0
	for _, row := range visitedTiles {
		for _, val := range row {
			if val > 0 {
				total++
			}
		}
	}
	return total
}

func getNextSteps(grid []string, currStep point, visited map[string]bool) []point {
	visitedTiles[currStep.y][currStep.x]++
	nextSteps := []point{}
	x := currStep.x
	y := currStep.y

	switch grid[y][x] {
	case '.':
		nX, nY := getCoords(x, y, currStep.direction)
		if isInBounds(grid, nX, nY) && !isVisited(visited, nX, nY, currStep.direction) {
			nextSteps = append(nextSteps, point{x: nX, y: nY, direction: currStep.direction})
		}
	case '|':
		if currStep.direction == "U" || currStep.direction == "D" {
			nX, nY := getCoords(x, y, currStep.direction)
			if isInBounds(grid, nX, nY) && !isVisited(visited, nX, nY, currStep.direction) {
				nextSteps = append(nextSteps, point{x: nX, y: nY, direction: currStep.direction})
			}
		} else {
			uX, uY := getCoords(x, y, "U")
			if isInBounds(grid, uX, uY) && !isVisited(visited, uX, uY, "U") {
				nextSteps = append(nextSteps, point{x: uX, y: uY, direction: "U"})
			}
			dX, dY := getCoords(x, y, "D")
			if isInBounds(grid, dX, dY) && !isVisited(visited, dX, dY, "D") {
				nextSteps = append(nextSteps, point{x: dX, y: dY, direction: "D"})
			}
		}
	case '-':
		if currStep.direction == "R" || currStep.direction == "L" {
			nX, nY := getCoords(x, y, currStep.direction)
			if isInBounds(grid, nX, nY) && !isVisited(visited, nX, nY, currStep.direction) {
				nextSteps = append(nextSteps, point{x: nX, y: nY, direction: currStep.direction})
			}
		} else {
			lX, lY := getCoords(x, y, "L")
			if isInBounds(grid, lX, lY) && !isVisited(visited, lX, lY, "L") {
				nextSteps = append(nextSteps, point{x: lX, y: lY, direction: "L"})
			}
			rX, rY := getCoords(x, y, "R")
			if isInBounds(grid, rX, rY) && !isVisited(visited, rX, rY, "R") {
				nextSteps = append(nextSteps, point{x: rX, y: rY, direction: "R"})
			}
		}
	case '\\':
		dir := mirrorReflection2[currStep.direction]
		nX, nY := getCoords(x, y, dir)
		if isInBounds(grid, nX, nY) && !isVisited(visited, nX, nY, dir) {
			nextSteps = append(nextSteps, point{x: nX, y: nY, direction: dir})
		}
	case '/':
		dir := mirrorReflection1[currStep.direction]
		nX, nY := getCoords(x, y, dir)
		if isInBounds(grid, nX, nY) && !isVisited(visited, nX, nY, dir) {
			nextSteps = append(nextSteps, point{x: nX, y: nY, direction: dir})
		}
	}
	return nextSteps
}

func isVisited(visited map[string]bool, x, y int, dir string) bool {
	if _, ok := visited[fmt.Sprintf("x%d,y%d,dir%s", x, y, dir)]; ok {
		return true
	}
	return false
}

func isInBounds(grid []string, x, y int) bool {
	if x < 0 || y < 0 || y >= len(grid) || x >= len(grid[y]) {
		return false
	}
	return true
}

func getCoords(currX, currY int, dir string) (int, int) {
	var x, y int
	switch dir {
	case "R":
		x = currX + 1
		y = currY
	case "U":
		x = currX
		y = currY + -1
	case "L":
		x = currX - 1
		y = currY
	case "D":
		x = currX
		y = currY + 1
	}
	return x, y
}

func valueStartingFromXEdges(grid []string, dir string, y, currMax int) int {
	max := currMax
	for x := 1; x < len(grid[0])-1; x++ {
		tilesVal := totalEnergizedTiles(grid, point{x: x, y: y, direction: dir})
		if tilesVal > max {
			max = tilesVal
		}
	}
	return max
}

func valueStartingFromYEdges(grid []string, dir string, x, currMax int) int {
	max := currMax
	for y := 1; y < len(grid)-1; y++ {
		tilesVal := totalEnergizedTiles(grid, point{x: x, y: y, direction: dir})
		if tilesVal > max {
			max = tilesVal
		}
	}
	return max
}

func valueStartingFromCorners(grid []string, currMax int) int {
	max := currMax
	// Top left corner
	tilesVal := totalEnergizedTiles(grid, point{x: 0, y: 0, direction: "R"})
	if tilesVal > max {
		max = tilesVal
	}
	tilesVal = totalEnergizedTiles(grid, point{x: 0, y: 0, direction: "D"})
	if tilesVal > max {
		max = tilesVal
	}
	// Top right corner
	tilesVal = totalEnergizedTiles(grid, point{x: len(grid[0]) - 1, y: 0, direction: "L"})
	if tilesVal > max {
		max = tilesVal
	}
	tilesVal = totalEnergizedTiles(grid, point{x: len(grid[0]) - 1, y: 0, direction: "D"})
	if tilesVal > max {
		max = tilesVal
	}
	// Bottom left corner
	tilesVal = totalEnergizedTiles(grid, point{x: 0, y: len(grid) - 1, direction: "R"})
	if tilesVal > max {
		max = tilesVal
	}
	tilesVal = totalEnergizedTiles(grid, point{x: 0, y: len(grid) - 1, direction: "U"})
	if tilesVal > max {
		max = tilesVal
	}
	// Bottom right corner
	tilesVal = totalEnergizedTiles(grid, point{x: len(grid[0]) - 1, y: len(grid) - 1, direction: "L"})
	if tilesVal > max {
		max = tilesVal
	}
	tilesVal = totalEnergizedTiles(grid, point{x: len(grid[0]) - 1, y: len(grid) - 1, direction: "U"})
	if tilesVal > max {
		max = tilesVal
	}
	return max
}
