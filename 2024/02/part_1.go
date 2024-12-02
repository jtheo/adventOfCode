package main

import "fmt"

func part1(o opts) {
	data := split(o)
	nSafe := 0
	for _, layer := range data {
		var inc, dec bool
		safe := true
		for i := 1; i < len(layer); i++ {
			diff := layer[i] - layer[i-1]
			if abs(diff) > 3 {
				if o.verbose {
					fmt.Printf("%d - %d = %d - Inc: %t | Dec: %t\n", layer[i], layer[i-1], diff, inc, dec)
				}
				safe = false
				break
			}

			switch {
			case diff > 0:
				inc = true
			case diff < 0:
				dec = true
			default:
				inc, dec = true, true
			}
			if dec && inc {
				safe = false
				break
			}
		}
		if safe {
			safe = true
			nSafe++
		}
		if o.verbose {
			fmt.Printf("layer: %v %t\n\n", layer, safe)
		}
		safe = false
	}
	fmt.Printf("Safe: %d\n", nSafe)
}
