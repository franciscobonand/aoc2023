package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/exp/maps"
)

type CalibrationValue struct {
	Value int
	Index int
}

func NewCalibrationValue() *CalibrationValue {
	return &CalibrationValue{
		Value: -1,
		Index: -1,
	}
}

var nums = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
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

	fmt.Println(getCalibrationValue(f))
}

func getCalibrationValue(f *os.File) int {
	scanner := bufio.NewScanner(f)
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		left := NewCalibrationValue()
		right := NewCalibrationValue()
		handleStringifiedValues(line, left, right)
		for idx, rn := range line {
			if isDigit(rn) {
				if left.Value == -1 || idx < left.Index {
					left.Value = int(rn - '0')
					left.Index = idx
				}
				if right.Value == -1 || idx > right.Index {
					right.Value = int(rn - '0')
					right.Index = idx
				}
			}
		}
		total += left.Value*10 + right.Value
	}
	return total
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func handleStringifiedValues(s string, l, r *CalibrationValue) {
	for _, key := range maps.Keys(nums) {
		if idx := strings.LastIndex(s, key); idx != -1 {
			if r.Value == -1 || idx > r.Index {
				r.Value = nums[key]
				r.Index = idx
			}
		}
		if idx := strings.Index(s, key); idx != -1 {
			if l.Value == -1 || idx < l.Index {
				l.Value = nums[key]
				l.Index = idx
			}
		}
	}
}
