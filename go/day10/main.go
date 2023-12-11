package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Coords struct {
	X, Y     int
	Distance int
}

var coordToPipe = map[string]string{
	"NS": "|",
	"EW": "-",
	"NW": "J",
	"NE": "L",
	"SW": "7",
	"SE": "F",
}
var directions = map[string][]Coords{
	"|": {{X: 0, Y: 1}, {X: 0, Y: -1}},
	"-": {{X: 1, Y: 0}, {X: -1, Y: 0}},
	"L": {{X: 0, Y: -1}, {X: 1, Y: 0}},
	"J": {{X: 0, Y: -1}, {X: -1, Y: 0}},
	"7": {{X: 0, Y: 1}, {X: -1, Y: 0}},
	"F": {{X: 0, Y: 1}, {X: 1, Y: 0}},
	".": {},
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
	startPoint := Coords{}
	maze := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		if idx := strings.Index(line, "S"); idx != -1 {
			startPoint = Coords{X: idx, Y: lineCount, Distance: 0}
		}
		maze = append(maze, line)
		lineCount++
	}
	return getStepsToFartestPoint(maze, startPoint)
}

func getStepsToFartestPoint(maze []string, startPoint Coords) int {
	visited := map[string]int{}
	sp := getStartingPipe(maze, startPoint)
	maze[startPoint.Y] = strings.Replace(maze[startPoint.Y], "S", sp, 1)
	nextSteps := []Coords{startPoint}
	for {
		if len(nextSteps) == 0 {
			break
		}
		currStep := nextSteps[0]
		nextSteps = nextSteps[1:]
		if _, ok := visited[fmt.Sprintf("x%d,y%d", currStep.X, currStep.Y)]; ok {
			continue
		}
		visited[fmt.Sprintf("x%d,y%d", currStep.X, currStep.Y)] = currStep.Distance
		nextSteps = append(nextSteps, getNextSteps(maze, currStep, visited)...)
	}
	maxDistance := 0
	for _, v := range visited {
		if v > maxDistance {
			maxDistance = v
		}
	}
	fmt.Println(countEnclosedTiles(maze, startPoint, visited))
	return maxDistance
}

func getNextSteps(maze []string, currLocation Coords, visited map[string]int) []Coords {
	nextSteps := []Coords{}
	for _, dir := range directions[string(maze[currLocation.Y][currLocation.X])] {
		nextLocation := Coords{X: currLocation.X + dir.X, Y: currLocation.Y + dir.Y, Distance: currLocation.Distance + 1}
		if nextLocation.X < 0 || nextLocation.X >= len(maze[0]) || nextLocation.Y < 0 || nextLocation.Y >= len(maze) {
			continue
		}
		nextSteps = append(nextSteps, nextLocation)
	}
	return nextSteps
}

func countEnclosedTiles(maze []string, start Coords, visited map[string]int) int {
	count := 0
	for i := 0; i < len(maze); i++ {
		for j := 0; j < len(maze[i]); j++ {
			if _, ok := visited[fmt.Sprintf("x%d,y%d", j, i)]; ok {
				continue
			}
			collisions := 0
			for p := 0; p < j; p++ {
				if _, ok := visited[fmt.Sprintf("x%d,y%d", p, i)]; ok && isPipePointingNorth(string(maze[i][p])) {
					collisions++
				}
			}
			if collisions%2 == 1 {
				count++
			}
		}
	}
	return count
}

func isPipePointingNorth(pipe string) bool {
	return pipe == "|" || pipe == "L" || pipe == "J"
}

func isPipePointingSouth(pipe string) bool {
	return pipe == "|" || pipe == "7" || pipe == "F"
}

func isPipePointingWest(pipe string) bool {
	return pipe == "-" || pipe == "J" || pipe == "7"
}

func isPipePointingEast(pipe string) bool {
	return pipe == "-" || pipe == "F" || pipe == "L"
}

func getStartingPipe(maze []string, startPoint Coords) string {
	dir := ""
	if startPoint.Y > 0 && isPipePointingSouth(string(maze[startPoint.Y-1][startPoint.X])) {
		dir += "N"
	}
	if startPoint.Y < len(maze)-1 && isPipePointingNorth(string(maze[startPoint.Y+1][startPoint.X])) {
		dir += "S"
	}
	if startPoint.X > 0 && isPipePointingEast(string(maze[startPoint.Y][startPoint.X-1])) {
		dir += "W"
	}
	if startPoint.X < len(maze[0])-1 && isPipePointingWest(string(maze[startPoint.Y][startPoint.X+1])) {
		dir += "E"
	}
	return coordToPipe[dir]
}
