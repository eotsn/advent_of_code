package main

import (
	"fmt"
	"strconv"
	"strings"
	"text/scanner"

	"github.com/eotsn/advent_of_code/2023/file"
)

var cubes = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func main() {
	lines, err := file.ReadLines("input/02.txt")
	if err != nil {
		panic(err)
	}

	games := make(map[int]bool)
	for _, line := range lines {
		var sc scanner.Scanner
		sc.Init(strings.NewReader(line))

		var game int
		var prevToken string
	parse:
		for token := sc.Scan(); token != scanner.EOF; token = sc.Scan() {
			text := sc.TokenText()

			switch token {
			case scanner.Ident:
				switch text {
				case "red", "green", "blue":
					count, _ := strconv.Atoi(prevToken)
					if count > cubes[text] {
						games[game] = false
						break parse
					}
				default:
				}
			case scanner.Int:
				if prevToken == "Game" {
					game, _ = strconv.Atoi(text)
					games[game] = true // assume game is valid
				}
			default:
			}
			prevToken = text
		}
	}

	var sum int
	for game, valid := range games {
		if valid {
			sum += game
		}
	}
	fmt.Println(sum)
}
