package main

import (
	"fmt"
	"io"
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
	bContent, err := io.ReadAll(f)
	if err != nil {
		panic(err.Error())
	}
	strContent := strings.Split(string(bContent), "\n")
	// times := getValues(strContent[0])
	// distances := getValues(strContent[1])
	times := getUnifiedValue(strContent[0])
	distances := getUnifiedValue(strContent[1])
	worstTimeCount := 0
	betterTimes := []int{}
	for idx, t := range times {
		for ms := 0; ms <= t/2; ms++ {
			if ms*(t-ms) > distances[idx] {
				betterTimes = append(betterTimes, (t+1)-(worstTimeCount*2))
				worstTimeCount = 0
				break
			}
			worstTimeCount++
		}
	}
	waysToBeat := 1
	for _, bt := range betterTimes {
		waysToBeat *= bt
	}
	return waysToBeat
}

// Use for part 1
func getValues(s string) []int {
	vals := []int{}
	nums := strings.Split(s, ":")[1]
	strVals := strings.Split(nums, " ")
	for _, sv := range strVals {
		if sv == "" {
			continue
		}
		v, _ := strconv.Atoi(sv)
		vals = append(vals, v)
	}
	return vals
}

// Use for part 2
func getUnifiedValue(s string) []int {
	vals := []int{}
	nums := strings.Split(s, ":")[1]
	strVal := ""
	for _, sv := range nums {
		if isDigit(sv) {
			strVal += string(sv)
		}
	}
	v, _ := strconv.Atoi(strVal)
	vals = append(vals, v)
	return vals
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}
