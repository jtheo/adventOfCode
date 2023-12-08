package main

import (
	"fmt"
	"sort"
)

func main() {
	r := run()

	res := r.part1()
	fmt.Printf("Day 7, part 1. Result: %d\n\n", res)

	// res = part2(r)
	// fmt.Printf("Day 7, part 2. Result: %d\n\n", res)
}

func (h listHands) part1() int {
	res := 0
	for _, hands := range h {
		for _, r := range hands.hand {
			hands.value += cardValues[string(r)]
		}
	}
	sort.Slice(h, func(i, j int) bool {
		return h[i].value < h[j].value
	})
	for i, hand := range h {
		ht := evaluateHand(hand.hand)
		v := (i + 1) * hand.bid
		fmt.Printf("%+v => %d * %d = %d | hand: %d\n", hand, i+1, hand.bid, v, ht)
		res += v
	}
	return res
}

func evaluateHand(hand string) int {
	max := 0

	return max
}

type combo struct {
	name  string
	value int
}

type typeValue map[int]combo

var cardValues = map[string]int{
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
	"J": 11,
	"Q": 12,
	"K": 13,
	"A": 14,
}
