package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("I need a filename, only one...")
	}
	input, err := getData(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("The sum of the numbers is", calc(input))
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
		res = append(res, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return []string{}, fmt.Errorf("error scanning %s: %v", fn, err)
	}
	return res, nil
}

func calc(input []string) int {
	ret := 0

	for _, l := range input {
		i := 0
		n := []int{}

		for i < len(l) {
			if l[i] >= '0' && l[i] <= '9' {
				d, err := strconv.Atoi(string(l[i]))
				if err != nil {
					log.Fatal(l[i], "is not a number")
				}
				n = append(n, d)
				i++
				continue
			}

			for k, v := range convMap {
				if strings.HasPrefix(l[i:], k) {
					n = append(n, v)
					i += len(k) - 2
					break
				}
			}

			i++
		}
		ret += (n[0] * 10) + n[len(n)-1]
	}

	return ret
}

var convMap = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}
