package main

import (
	"fmt"
)

func main() {
	games := run()

	res := part1(games)
	fmt.Printf("Day 4, part 1. Result: %d\n\n", res)

	res = part2(games)
	fmt.Printf("Day 4, part 2. Result: %d\n", res)
}

func part2(games data) int {
	nGames := len(games)
	wins := make([]int, nGames)
	for i, c := range games {
		wins[i]++
		nextGames := 0
		for _, n := range c.winning {
			if contains(n, c.played) {
				nextGames++
			}
		}

		for o := 0; o < wins[i]; o++ {
			for x := 0; x < nextGames; x++ {
				card := (i + 1 + x)
				wins[card]++
			}
		}
	}

	res := 0

	for _, w := range wins {
		res += w
	}

	return res
}
