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
	totalArrangements := 0
	for scanner.Scan() {
		line := scanner.Text()
		tmp := strings.Split(line, " ")
		springs := tmp[0]
		dmgLocations := []int{}
		for _, s := range strings.Split(tmp[1], ",") {
			v, _ := strconv.Atoi(s)
			dmgLocations = append(dmgLocations, v)
		}
		pt2DmgLoc := []int{}
		pt2Springs := ""
		for i := 0; i < 5; i++ {
			pt2DmgLoc = append(pt2DmgLoc, dmgLocations...)
			pt2Springs += springs + "?"
		}

		totalArrangements += buildArrangements(pt2Springs[:len(pt2Springs)-1], pt2DmgLoc)
	}
	return totalArrangements
}

var arrangements = make(map[string]int)

func buildArrangements(springs string, dmgLocs []int) int {
	key := springs
	for _, group := range dmgLocs {
		key += fmt.Sprintf("%d-", group)
	}
	if v, ok := arrangements[key]; ok {
		return v
	}
	if len(springs) == 0 {
		if len(dmgLocs) == 0 {
			return 1
		}
		return 0
	}
	if springs[0] == '?' {
		return buildArrangements(strings.Replace(springs, "?", ".", 1), dmgLocs) +
			buildArrangements(strings.Replace(springs, "?", "#", 1), dmgLocs)
	}
	if springs[0] == '.' {
		pcount := 1
		for i := 1; i < len(springs); i++ {
			if springs[i] == '.' {
				pcount++
				continue
			}
			break
		}
		res := buildArrangements(springs[pcount:], dmgLocs)
		arrangements[key] = res
		return res
	}
	if springs[0] == '#' {
		if len(dmgLocs) == 0 {
			arrangements[key] = 0
			return 0
		}
		if len(springs) < dmgLocs[0] {
			arrangements[key] = 0
			return 0
		}
		if strings.Contains(springs[:dmgLocs[0]], ".") {
			arrangements[key] = 0
			return 0
		}
		if len(dmgLocs) == 1 {
			res := buildArrangements(springs[dmgLocs[0]:], dmgLocs[1:])
			arrangements[key] = res
			return res
		}
		if len(springs) < dmgLocs[0]+1 || springs[dmgLocs[0]] == '#' {
			arrangements[key] = 0
			return 0
		}
		res := buildArrangements(springs[dmgLocs[0]+1:], dmgLocs[1:])
		arrangements[key] = res
		return res
	}
	return 0
}
