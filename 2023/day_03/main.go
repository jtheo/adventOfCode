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
	fmt.Printf("Day 3, part 1. Result: %d\n\n", res)

	res = part2(matrix)
	fmt.Printf("Day 3, part 2. Result: %d\n", res)
}

func part2(input [][]string) int {
	res := 0
	queue := []asterisk{}

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
				isValid, ast := lookAroundForSymbols(i, s, e, input)

				if isValid {
					valid := strings.Join(l[s:e], "")
					n, err := strconv.Atoi(valid)
					if err != nil {
						log.Fatalf("converting %q to int, got %v\ns: %d, e: %d, l: %d, rightStop: %d\n", valid, err, s, e, i, rightStop)
					}

					if ast.present {
						found := false
						delete := 0
						for e, a := range queue {
							if ast.id == a.id {
								n *= a.value
								found = true
								delete = e
							}
						}

						if !found {
							ast.value = n
							queue = append(queue, ast)
							continue
						}
						if found {
							queue = append(queue[:delete], queue[delete+1:]...)
						}
						res += n
					}
				}
			}
		}
	}

	return res
}

func lookAroundForSymbols(i, s, e int, matrix [][]string) (bool, asterisk) {
	hLimit, vLimit := len(matrix[0])-1, len(matrix)-1
	up, down := i-1, i+1
	left, right := s-1, e
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
	isSymbPresent := false
	ast := asterisk{}
	for v := up; v <= down; v++ {
		for h := left; h <= right; h++ {
			if isSymbol(matrix[v][h]) {
				isSymbPresent = true
				if matrix[v][h] == "*" {
					ast.present = true
					ast.id = fmt.Sprintf("%d-%d", v, h)
				}
			}
		}
	}

	return isSymbPresent, ast
}

func isSymbol(c string) bool {
	return !((c >= "0" && c <= "9") || c == ".")
}

type asterisk struct {
	id      string
	present bool
	value   int
}
