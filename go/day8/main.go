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

func getSolutions(f *os.File) int {
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	instructions := []string{}
	for _, r := range scanner.Text() {
		instructions = append(instructions, string(r))
	}
	scanner.Scan()
	network := map[string][]string{}
	startingNodes := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		node := strings.Split(line, " = ")
		orig := node[0]
		aux := strings.ReplaceAll(node[1], "(", "")
		aux = strings.ReplaceAll(aux, ")", "")
		dests := strings.Split(aux, ", ")
		network[orig] = dests
		if strings.LastIndex(orig, "A") == 2 {
			startingNodes = append(startingNodes, orig)
		}
	}
	stepsByStartingNode := []int{}
	for _, node := range startingNodes {
		stepsByStartingNode = append(stepsByStartingNode, getSteps(instructions, network, node))
	}
	if len(stepsByStartingNode) == 1 {
		return stepsByStartingNode[0]
	} else if len(stepsByStartingNode) == 2 {
		return LCM(stepsByStartingNode[0], stepsByStartingNode[1])
	}
	return LCM(stepsByStartingNode[0], stepsByStartingNode[1], stepsByStartingNode[2:]...)
}

func getSteps(instructions []string, network map[string][]string, startingNode string) int {
	steps := 0
	currInstruction := 0
	currNode := startingNode
	for {
		nextNodes := network[currNode]
		if instructions[currInstruction] == "L" {
			currNode = nextNodes[0]
			if strings.LastIndex(currNode, "Z") == 2 {
				return steps + 1
			}
		} else if instructions[currInstruction] == "R" {
			currNode = nextNodes[1]
			if strings.LastIndex(currNode, "Z") == 2 {
				return steps + 1
			}
		}
		steps++
		currInstruction = (currInstruction + 1) % len(instructions)
	}
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
