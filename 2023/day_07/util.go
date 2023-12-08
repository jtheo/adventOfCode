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

func run() listHands {
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

type listHands []*handBid

type handBid struct {
	hand     string
	bid      int
	value    int
	typeHand int
}

func unmarshal(data []string) listHands {
	res := listHands{}
	for _, d := range data {
		l := strings.Split(d, " ")
		hand := l[0]
		bid, err := strconv.Atoi(l[1])
		if err != nil {
			log.Fatalf("Error converting %q: %v\n", bid, err)
		}
		res = append(res, &handBid{
			hand: hand,
			bid:  bid,
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
