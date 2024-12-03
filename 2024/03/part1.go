package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func part1(o opts) {
	data, err := load(o.fn)
	must(err)

	r, err := regexp.Compile(`mul\((\d+),(\d+)\)`)
	must(err)
	res := r.FindAllStringSubmatch(string(data), -1)

	sum := 0
	for i, r := range res {
		if o.verbose {
			fmt.Printf("%d := %v len: %d | %s %d \n", i, r, len(r), string(r[1]), sum)
		}
		mul := 1
		for i := 1; i < len(r); i++ {
			num, err := strconv.Atoi(r[i])
			must(err)
			mul *= num
		}
		sum += mul
	}
	fmt.Printf("Sum is %d\n", sum)
}
