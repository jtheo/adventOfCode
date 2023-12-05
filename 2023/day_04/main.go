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
		win := 0
		for _, n := range c.winning {
			if contains(n, c.played) {
				win++
				wins[i]++
			}
		}

		fmt.Printf("Card %d: %+v | %+v | %-3d\n", i+1, c.winning, c.played, win)
		for x := 0; x < wins[i]; x++ {
			card := (i + x + 1) % nGames
			wins[card]++
		}
	}

	res := 0

	for _, w := range wins {
		// fmt.Println(p+1, w)
		res += w
	}

	return res
}
