package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	ASH  = '.'
	ROCK = '#'
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
	patterns := []string{}
	reflectionSummary := 0
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			rows, cols := getRowsAndColsValues(patterns)
			if v := getReflectionValue(rows); v > 0 {
				reflectionSummary += (v * 100)
			} else {
				reflectionSummary += (getReflectionValue(cols))
			}
			patterns = nil
			continue
		}
		patterns = append(patterns, line)
	}
	rows, cols := getRowsAndColsValues(patterns)
	if v := getReflectionValue(rows); v > 0 {
		reflectionSummary += (v * 100)
	} else {
		reflectionSummary += (getReflectionValue(cols))
	}
	return reflectionSummary
}

func getReflectionValue(vals []int) int {
	for i := 0; i < len(vals)-1; i++ {
		if isReflection(vals, i, i+1) {
			return (i + 1)
		}
	}
	return -1
}

func isReflection(vals []int, l, r int) bool {
	for l >= 0 && r < len(vals) && vals[l] == vals[r] {
		l--
		r++
	}
	return l < 0 || r == len(vals)
}

func getRowsAndColsValues(pattern []string) ([]int, []int) {
	rows := make([]int, len(pattern))
	cols := make([]int, len(pattern[0]))
	for i, row := range pattern {
		for j, r := range row {
			if r == '#' {
				rows[i] |= 1 << j
				cols[j] |= 1 << i
			}
		}
	}
	return rows, cols
}
