package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/eotsn/advent_of_code/2023/file"
	"github.com/eotsn/advent_of_code/2023/parser"
)

type Map []Range

type Range struct {
	start int
	end   int
	delta int
}

func main() {
	lines, err := file.ReadLines("input/05.txt")
	if err != nil {
		panic(err)
	}

	var maps []Map
	for _, line := range lines[1:] {
		if line == "" {
			continue
		}
		if line[len(line)-1] == ':' {
			maps = append(maps, Map{})
			continue
		}
		n := parser.ParseInts(line)
		r := Range{start: n[1], end: n[1] + n[2], delta: n[1] - n[0]}
		maps[len(maps)-1] = append(maps[len(maps)-1], r)
	}

	seeds := parser.ParseInts(strings.Split(lines[0], ":")[1])
	for i := range seeds {
		for _, ranges := range maps {
			for _, r := range ranges {
				if seeds[i] >= r.start && seeds[i] < r.end {
					seeds[i] -= r.delta
					break
				}
			}
		}
	}
	slices.Sort(seeds)
	fmt.Println(seeds[0])
}
