package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func run() almanac {
	flag.Parse()
	if len(flag.Args()) != 1 {
		log.Fatal("I need a filename, len args is", flag.Args())
	}

	input, err := getData(flag.Args()[0])
	if err != nil {
		log.Fatal(err)
	}

	almanac, err := unmarshal(input)
	if err != nil {
		log.Fatalf("Error unmarshalling the data: %v\n", err)
	}
	return almanac
}

type almanac struct {
	guides guide
	seeds  []int
}

type (
	guide map[string][]entry
	entry struct {
		dstMin int
		dstMax int
		srcMin int
		srcMax int
	}
)

var (
	guides = []string{
		"seed-to-soil",
		"soil-to-fertilizer",
		"fertilizer-to-water",
		"water-to-light",
		"light-to-temperature",
		"temperature-to-humidity",
		"humidity-to-location",
	}
	seeds = []int{}
)

func seedsPart1(l string) []int {
	res := []int{}
	t := strings.Split(l, ": ")
	t2 := strings.Split(t[1], " ")

	for _, n := range t2 {
		x, err := strconv.Atoi(strings.TrimSpace(n))
		if err != nil {
			log.Panic("converting: ", n, " from: ", t2, " got: ", err)
		}
		res = append(res, x)
	}

	return res
}

func seedsPart2(seeds []int) []int {
	res := []int{}
	fmt.Println("received:", seeds)

	i := 0
	for i < len(seeds) {
		from := seeds[i]
		to := seeds[i+1]

		for s := 0; s < to; s++ {
			res = append(res, s+from)
		}
		i += 2
	}

	return res
}

func unmarshal(input []string) (almanac, error) {
	a := almanac{
		seeds:  []int{},
		guides: map[string][]entry{},
	}
	inSection := false
	section := ""

	for _, l := range input {
		if l == "" {
			inSection = false
			continue
		}

		if strings.HasPrefix(l, "seeds:") {
			a.seeds = seedsPart1(l)
		}

		if strings.HasSuffix(l, "map:") {
			inSection = true
			t := strings.Split(l, " ")
			section = t[0]
			continue
		}

		if inSection {
			ns := []int{}
			t := strings.Split(l, " ")
			for _, n := range t {
				x, err := strconv.Atoi(strings.TrimSpace(n))
				if err != nil {
					log.Panic("converting", n, "got", err)
				}
				ns = append(ns, x)
			}
			e := entry{
				dstMin: ns[0],
				dstMax: ns[0] + ns[2] - 1,
				srcMin: ns[1],
				srcMax: ns[1] + ns[2] - 1,
			}
			a.guides[section] = append(a.guides[section], e)
		}
	}

	return a, nil
}

func loadInt(s string) []int {
	res := []int{}
	numbers := strings.Split(s, " ")

	for _, w := range numbers {
		if w == "" {
			continue
		}
		n, err := strconv.Atoi(w)
		if err != nil {
			log.Fatalf("%q is not a number, from %s | %v\n", w, s, err)
		}
		res = append(res, n)
	}
	return res
}

func getData(fn string) ([]string, error) {
	res := []string{}
	fd, err := os.Open(fn)
	if err != nil {
		return []string{}, fmt.Errorf("can't access %s: %v", fn, err)
	}

	defer fd.Close()

	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		if l := scanner.Text(); l != "" {
			res = append(res, l)
		}
	}
	if err := scanner.Err(); err != nil {
		return []string{}, fmt.Errorf("error scanning %s: %v", fn, err)
	}

	return res, nil
}

func contains(n int, played []int) bool {
	for _, i := range played {
		if i == n {
			return true
		}
	}
	return false
}
