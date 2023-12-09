package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	nextValues := []int{}
	for scanner.Scan() {
		initValues := []int{}
		line := scanner.Text()
		for _, strn := range strings.Split(line, " ") {
			n, _ := strconv.Atoi(strn)
			initValues = append(initValues, (n))
		}
		nextValues = append(nextValues, getNextValue(initValues))
	}
	sum := 0
	for _, val := range nextValues {
		sum += val
	}
	return sum
}

func getNextValue(initValues []int) int {
	matrix := [][]int{}
	currVals := initValues
	matrix = append(matrix, initValues)
	for {
		nextVals := []int{}
		for i := 0; i < len(currVals)-1; i++ {
			nextVals = append(nextVals, currVals[i+1]-currVals[i])
		}
		matrix = append(matrix, nextVals)
		if allZeroes(nextVals) {
			break
		}
		currVals = nextVals
	}
	matrix[len(matrix)-1] = append(matrix[len(matrix)-1], 0)
	for i := len(matrix) - 2; i >= 0; i-- {
		// Part 1
		// nextVal := matrix[i+1][len(matrix[i])-1] + matrix[i][len(matrix[i])-1]
		// Part 2
		nextVal := matrix[i][0] - matrix[i+1][len(matrix[i])-1]
		matrix[i] = append(matrix[i], nextVal)
	}
	return matrix[0][len(matrix[0])-1]
}

func allZeroes(vals []int) bool {
	for _, v := range vals {
		if v != 0 {
			return false
		}
	}
	return true
}
