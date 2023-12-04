package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	matrix := run()

	res := part1(matrix)
	fmt.Printf("Day 3, part 1. Result: %d\n", res)
}

func part1(input [][]string) int {
	res := 0
	for i, l := range input {
		// fmt.Println(i, l, len(l))
		s, e := 0, 0
		isNum := false
		rightStop := len(l) - 1
		fmt.Print("   ")
		for i := range l {
			fmt.Printf("%d ", i)
		}
		fmt.Println()
		fmt.Println(i, l)

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
				if lookAroundForSymbols(i, s, e, input) {
					valid := strings.Join(l[s:e], "")
					n, err := strconv.Atoi(valid)
					if err != nil {
						log.Fatalf("converting %q to int, got %v\ns: %d, e: %d, l: %d, rightStop: %d\n", valid, err, s, e, i, rightStop)
					}
					fmt.Println("found", n)
					res += n
				}
			}
		}
	}
	return res
}

func lookAroundForSymbols(i, s, e int, matrix [][]string) bool {
	hLimit, vLimit := len(matrix[0])-1, len(matrix)-1
	up, down := i-1, i+1
	left, right := s-1, e+1
	if up < 0 {
		up = 0
	}

	if down > vLimit {
		down = vLimit
	}

	if left < 0 {
		left = 0
	}

	if right > hLimit {
		right = hLimit
	}
	for v := up; v <= down; v++ {
		for h := left; h <= right; h++ {
			if isSymbol(matrix[v][h]) {
				return true
			}
		}
	}

	return false
}

func isSymbol(c string) bool {
	return !((c >= "0" && c <= "9") || c == ".")
}
