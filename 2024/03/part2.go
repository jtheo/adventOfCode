package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func part2(o opts) {
	data, err := load(o.fn)
	must(err)

	r, err := regexp.Compile(`(don't\(\))|(do\(\))|mul\((\d+),(\d+)\)`)
	must(err)
	res := r.FindAllStringSubmatch(string(data), -1)

	sum := 0
	enabled := true
	for i, r := range res {
		if o.verbose {
			fmt.Printf("%d := '%+v' len: %d | %s %d \n", i, r, len(r), string(r[1]), sum)
			for n, v := range r {
				fmt.Printf("%d: %q\n", n, v)
			}
			fmt.Println()
		}
		if r[1] == "don't()" {
			enabled = false
			continue
		}
		if r[2] == "do()" {
			enabled = true
			continue
		}
		if enabled {
			mul := 1
			for i := 3; i < len(r); i++ {
				num, err := strconv.Atoi(r[i])
				must(err)
				mul *= num
			}
			sum += mul
		}
	}
	fmt.Printf("Sum is %d\n", sum)
}
