package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
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

func split(o opts) ([]int, []int) {
	c, err := load(o.fn)
	must(err)
	lines := strings.Split(string(c), "\n")
	var left, right []int
	for i, l := range lines {
		if len(l) == 0 {
			continue
		}
		if o.verbose {
			fmt.Printf("%d :=%s\n", i, l)
		}
		lf, rt := 0, 0
		t := strings.NewReader(l)
		n, err := fmt.Fscanf(t, "%d   %d", &lf, &rt)
		if err != nil {
			log.Panic(err)
		}
		if n != 2 {
			log.Fatalf("Number of items not right: %d != 2\n", n)
		}
		left = append(left, lf)
		right = append(right, rt)
	}
	sort.Ints(left)
	sort.Ints(right)
	return left, right
}
