package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

// 36920, 108332 is too low

var boxes = map[int]map[string]int{}
var positions = make([][]string, 256)

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
	for i := 0; i < 256; i++ {
		positions[i] = make([]string, 0)
	}
	scanner := bufio.NewScanner(f)
	totalValue := 0
	scanner.Scan()
	line := scanner.Text()
	seqs := strings.Split(line, ",")
	for _, seq := range seqs {
		sigIdx := strings.Index(seq, "-")
		if sigIdx == -1 {
			sigIdx = strings.Index(seq, "=")
		}
		hash := getSequenceValue(seq, sigIdx)
		handleLens(seq, hash, sigIdx)
	}
	totalValue = getTotalFocusingPower()
	return totalValue
}

func handleLens(seq string, hash, sigIdx int) {
	if boxes[hash] == nil {
		boxes[hash] = map[string]int{}
	}
	if seq[sigIdx] == '-' {
		delete(boxes[hash], seq[:sigIdx])
		idx := slices.Index(positions[hash], seq[:sigIdx])
		if idx != -1 {
			positions[hash] = append(positions[hash][:idx], positions[hash][idx+1:]...)
		}
		return
	}
	val, _ := strconv.Atoi(seq[sigIdx+1:])
	if _, ok := boxes[hash][seq[:sigIdx]]; !ok {
		positions[hash] = append(positions[hash], seq[:sigIdx])
	}
	boxes[hash][seq[:sigIdx]] = val
}

func getSequenceValue(seq string, sigIdx int) int {
	currentValue := 0
	for _, c := range seq[:sigIdx] {
		currentValue += int(c)
		currentValue *= 17
		currentValue %= 256
	}
	return currentValue
}

func getTotalFocusingPower() int {
	totalValue := 0
	for hash, box := range boxes {
		for lens, lenval := range box {
			totalValue += (hash + 1) * (slices.Index(positions[hash], lens) + 1) * lenval
		}
	}
	return totalValue
}
