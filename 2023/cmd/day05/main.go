package main

import (
	"fmt"
	"math"
	"strings"
	"sync"

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

	n := parser.ParseInts(strings.Split(lines[0], ":")[1])

	var seedRanges []Range
	for i := 0; i < len(n); i += 2 {
		seedRanges = append(seedRanges, Range{start: n[i], end: n[i] + n[i+1]})
	}

	var ch = make(chan int)
	var wg sync.WaitGroup

	for _, seed := range seedRanges {
		wg.Add(1)
		go func(seed Range) {
			defer wg.Done()
			for i := seed.start; i < seed.end; i++ {
				ch <- findLocation(i, maps)
			}
		}(seed)
	}

	go func() {
		wg.Wait()
		close(ch) // make sure to close the channel once we're done
	}()

	location := math.MaxInt
	for l := range ch {
		if l < location {
			location = l
		}
	}

	fmt.Println(location)
}

func findLocation(seed int, maps []Map) int {
	for _, ranges := range maps {
		for _, r := range ranges {
			if seed >= r.start && seed < r.end {
				seed -= r.delta
				break
			}
		}
	}
	return seed
}
