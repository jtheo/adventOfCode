package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func run() [][]string {
	flag.Parse()
	if len(flag.Args()) != 1 {
		log.Fatal("I need a filename, len args is", flag.Args())
	}

	input, err := getData(flag.Args()[0])
	if err != nil {
		log.Fatal(err)
	}

	matrix := [][]string{}
	for _, l := range input {
		t := strings.Split(l, "")
		matrix = append(matrix, t)
	}
	return matrix
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
