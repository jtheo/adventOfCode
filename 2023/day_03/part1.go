package main

import (
	"log"
	"strconv"
	"strings"
)

func part1(input [][]string) int {
	res := 0
	for i, l := range input {
		s, e := 0, 0
		isNum := false
		rightStop := len(l) - 1

		for x, c := range l {
			if c >= "0" && c <= "9" {
				if !isNum {
					s = x
					e = x
					isNum = true
				}
				e++
				if x < rightStop {
					continue
				}
			}

			if isNum {
				isNum = false
				if symb, _ := lookAroundForSymbols(i, s, e, input); symb {

					valid := strings.Join(l[s:e], "")
					n, err := strconv.Atoi(valid)
					if err != nil {
						log.Fatalf("converting %q to int, got %v\ns: %d, e: %d, l: %d, rightStop: %d\n", valid, err, s, e, i, rightStop)
					}
					res += n
				}
			}
		}
	}
	return res
}
