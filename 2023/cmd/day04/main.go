package main

import (
	"fmt"
	"strings"

	"github.com/eotsn/advent_of_code/2023/file"
	"github.com/eotsn/advent_of_code/2023/parser"
)

func main() {
	lines, err := file.ReadLines("input/04.txt")
	if err != nil {
		panic(err)
	}

	var cards []*Card

	var s []string
	for i, line := range lines {
		w := make(map[int]bool) // assemble winning numbers

		s = strings.Split(line, ":")
		s = strings.Split(s[1], "|")

		for _, n := range parser.ParseInts(s[0]) {
			w[n] = true
		}

		var count int
		for _, n := range parser.ParseInts(s[1]) {
			if w[n] {
				count++
			}
		}
		cards = append(cards, &Card{number: i, winners: count})
	}

	for _, card := range cards {
		countCards(card, cards)
	}

	var sum int
	for _, card := range cards {
		sum += card.visits
	}
	fmt.Println(sum)
}

type Card struct {
	number  int
	winners int
	visits  int
}

func countCards(card *Card, winners []*Card) {
	card.visits++
	if card.winners == 0 || card.number >= len(winners) {
		return
	}
	for i := 1; i <= card.winners; i++ {
		countCards(winners[card.number+i], winners)
	}
	return
}
