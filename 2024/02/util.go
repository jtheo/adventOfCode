package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type opts struct {
	fn      string
	verbose bool
	part    int
}

func run() opts {
	var fn string
	var verb bool
	var part int

	flag.StringVar(&fn, "fn", "", "Filename of the list")
	flag.BoolVar(&verb, "v", false, "Verbose output")
	flag.IntVar(&part, "part", 1, "Part of the day")
	flag.Parse()

	if fn == "" {
		fmt.Println("I need a filename")
		flag.Usage()
	}
	o := opts{}
	o.fn = fn
	o.verbose = verb
	o.part = part
	return o
}

func load(n string) ([]byte, error) {
	c, err := os.ReadFile(n)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func split(o opts) [][]int {
	c, err := load(o.fn)
	must(err)
	lines := strings.Split(string(c), "\n")
	res := [][]int{}

	for _, l := range lines {
		if len(l) == 0 {
			continue
		}
		layer := []int{}
		vals := strings.Fields(l)
		for _, v := range vals {
			n, err := strconv.Atoi(v)
			must(err)
			layer = append(layer, n)
		}
		res = append(res, layer)
	}

	return res
}

func abs(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}
