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

type card struct {
	winning []int
	played  []int
}
type data []card

func run() data {
	flag.Parse()
	if len(flag.Args()) != 1 {
		log.Fatal("I need a filename, len args is", flag.Args())
	}

	input, err := getData(flag.Args()[0])
	if err != nil {
		log.Fatal(err)
	}

	matrix := data{}

	for _, l := range input {
		c := card{}
		l = strings.Join(strings.Fields(l), " ")
		cols := strings.FieldsFunc(l, func(r rune) bool {
			return r == ':' || r == '|'
		})
		if len(cols) != 3 {
			log.Fatalf("cols != 3: %d, %v\n", len(cols), cols)
		}
		c.winning = loadInt(cols[1])
		c.played = loadInt(cols[2])
		matrix = append(matrix, c)
	}
	return matrix
}

func loadInt(s string) []int {
	res := []int{}
	numbers := strings.Split(s, " ")

	for _, w := range numbers {
		if w == "" {
			continue
		}
		// fmt.Printf("processing: %q of %+v\n", w, numbers)
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
