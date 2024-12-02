package main

import (
	"fmt"
)

func main() {
	o := run()

	fmt.Printf("Running Part %d\n", o.part)
	switch o.part {
	case 1:
		part1(o)
	case 2:
		part2(o)
	}
}
