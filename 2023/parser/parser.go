package parser

import (
	"regexp"
	"strconv"
)

func ParseInts(s string) (ns []int) {
	r := regexp.MustCompile(`\d+`)
	for _, m := range r.FindAllString(s, -1) {
		n, _ := strconv.Atoi(m)
		ns = append(ns, n)
	}
	return
}
