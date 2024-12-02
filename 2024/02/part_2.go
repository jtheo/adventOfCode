package main

import (
	"fmt"
)

func part2(o opts) {
	data := split(o)
	nSafe := 0
	for _, layer := range data {
		for i := range layer {
			tmp := remove(layer, i)
			if safe := checkLayer(tmp); safe {
				if o.verbose {
					fmt.Printf(" Safe\n")
				}
				nSafe++
				break
			}
		}
	}
	fmt.Printf("Safe: %d\n", nSafe)
}

func remove(slice []int, s int) []int {
	tmp := make([]int, 0, len(slice)-1)
	tmp = append(tmp, slice[:s]...)
	tmp = append(tmp, slice[s+1:]...)
	return tmp
}

func checkLayer(layer []int) bool {
	var inc, dec bool
	for i := 1; i < len(layer); i++ {
		diff := layer[i] - layer[i-1]
		switch {
		case diff > 0:
			inc = true
		case diff < 0:
			dec = true
		default:
			inc, dec = true, true
		}

		if abs(diff) > 3 || (dec && inc) {
			return false
		}
	}
	return true
}
