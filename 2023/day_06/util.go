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

func run() races {
	flag.Parse()
	if len(flag.Args()) != 1 {
		log.Fatal("I need a filename, len args is", flag.Args())
	}

	input, err := getData(flag.Args()[0])
	if err != nil {
		log.Fatal(err)
	}

	return unmarshal(input)
}

type race struct {
	time     int
	distance int
}

type races []race

func unmarshal(data []string) races {
	res := races{}
	times := []int{}
	distances := []int{}

	for _, d := range data {
		tmp := []int{}
		t := strings.Split(d, ": ")
		nums := strings.Split(t[1], " ")
		for _, d := range nums {
			if d == "" {
				continue
			}
			n, err := strconv.Atoi(strings.TrimSpace(d))
			if err != nil {
				log.Fatalf("%q is not a number, from %+v | %v\n", d, nums, err)
			}
			tmp = append(tmp, n)
			if t[0] == "Distance" {
				distances = tmp
			} else {
				times = tmp
			}
		}
	}

	for i := range distances {
		res = append(res, race{
			time:     times[i],
			distance: distances[i],
		})
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
