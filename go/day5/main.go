package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Item struct {
	Init, End int
	Remapped  bool
}

type Mapper struct {
	Items []Item
}

// Use for part1 solution
func NewMapper(s string) *Mapper {
	seeds := []Item{}
	strNums := strings.Split(s, ": ")[1]
	for _, sn := range strings.Split(strNums, " ") {
		n, _ := strconv.Atoi(sn)
		seeds = append(seeds, Item{Init: n, End: n, Remapped: false})
	}
	return &Mapper{Items: seeds}
}

// Use for part2 solution
func NewNewMapper(s string) *Mapper {
	seeds := []Item{}
	strNums := strings.Split(strings.Split(s, ": ")[1], " ")
	for i := 0; i < len(strNums); i += 2 {
		n1, _ := strconv.Atoi(strNums[i])
		n2, _ := strconv.Atoi(strNums[i+1])
		seeds = append(seeds, Item{Init: n1, End: n1 + n2 - 1, Remapped: false})
	}
	return &Mapper{Items: seeds}
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
	scanner.Scan()
	// mp := NewMapper(scanner.Text())
	mp := NewNewMapper(scanner.Text())
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || string(line[len(line)-1]) == ":" {
			mp.startNewMapping()
			continue
		}
		dest, src, length := getMapLineValues(line)
		mp.updateMappedValue(dest, src, length)
	}
	return mp.getLowest()
}

func (m *Mapper) updateMappedValue(dest, src, length int) {
	destEnd := dest + length - 1
	srcEnd := src + length - 1
	for idx, i := range m.Items {
		if i.Remapped {
			continue
		}
		// maps will be updated here. If a Item from Mapper is not
		// completely within a given range (src <= x <= srcEnd), the part that
		// is outside the range will generate a new Item in the array
		if src <= i.Init && i.End <= srcEnd {
			m.Items[idx].Init = destEnd - (srcEnd - i.Init)
			m.Items[idx].End = destEnd - (srcEnd - i.End)
			m.Items[idx].Remapped = true
		} else if src <= i.Init && srcEnd > i.Init && i.End > srcEnd {
			newItem := Item{Init: srcEnd + 1, End: i.End, Remapped: false}
			m.Items = append(m.Items, newItem)
			m.Items[idx].Init = destEnd - (srcEnd - i.Init)
			m.Items[idx].End = destEnd
			m.Items[idx].Remapped = true
		} else if src > i.Init && i.End > src && i.End <= srcEnd {
			newItem := Item{Init: i.Init, End: src - 1, Remapped: false}
			m.Items = append(m.Items, newItem)
			m.Items[idx].Init = dest
			m.Items[idx].End = destEnd - (srcEnd - i.End)
			m.Items[idx].Remapped = true
		} else if src > i.Init && i.End > srcEnd {
			newItemLeft := Item{Init: i.Init, End: src - 1, Remapped: false}
			m.Items = append(m.Items, newItemLeft)
			newItemRight := Item{Init: srcEnd + 1, End: i.End, Remapped: false}
			m.Items = append(m.Items, newItemRight)
			m.Items[idx].Init = dest
			m.Items[idx].End = destEnd
			m.Items[idx].Remapped = true
		}
	}
}

func (m *Mapper) startNewMapping() {
	for idx := range m.Items {
		m.Items[idx].Remapped = false
	}
}

func (m *Mapper) getLowest() int {
	l := m.Items[0].Init
	for i := 1; i < len(m.Items); i++ {
		if m.Items[i].Init < l {
			l = m.Items[i].Init
		}
	}
	return l
}

func getMapLineValues(s string) (int, int, int) {
	values := []int{}
	for _, sn := range strings.Split(s, " ") {
		n, _ := strconv.Atoi(sn)
		values = append(values, n)
	}
	return values[0], values[1], values[2]
}
