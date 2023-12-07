package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/eotsn/advent_of_code/2023/file"
)

var Values = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	"J": 11,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
	"1": 1,
}

type Hand struct {
	Cards  []string
	Counts []int
	Bid    int
}

func main() {
	lines, err := file.ReadLines("input/07.txt")
	if err != nil {
		panic(err)
	}

	var hands []Hand
	for _, line := range lines {
		s := strings.Split(line, " ")
		s2 := strings.Split(s[0], "")

		bid, _ := strconv.Atoi(s[1])
		hands = append(hands, Hand{Cards: s2, Counts: getCounts(s2), Bid: bid})
	}

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].Counts[0] == hands[j].Counts[0] {
			if len(hands[i].Counts) == len(hands[j].Counts) {
				for k := 0; k < 5; k++ {
					if Values[hands[i].Cards[k]] != Values[hands[j].Cards[k]] {
						return Values[hands[i].Cards[k]] < Values[hands[j].Cards[k]]
					}
				}
			}
			return len(hands[i].Counts) > len(hands[j].Counts)
		}
		return hands[i].Counts[0] < hands[j].Counts[0]
	})

	var sum int
	for i, hand := range hands {
		sum += hand.Bid * (i + 1)
	}
	fmt.Println(sum)
}

func getCounts(hand []string) []int {
	cards := make(map[string]int)
	for _, card := range hand {
		cards[card]++
	}
	var counts []int
	for _, v := range cards {
		counts = append(counts, v)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(counts)))
	return counts
}
