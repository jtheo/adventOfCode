package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("I need a filename")
	}
	input, err := getData(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	sum := 0
	for _, l := range input {
		numbers := []int{}
		for _, r := range l {
			if unicode.IsDigit(r) {
				v, _ := strconv.Atoi(string(r))
				numbers = append(numbers, v)
			}
		}
		sum += (numbers[0] * 10) + numbers[len(numbers)-1]
	}

	fmt.Println("The sum of the numbers is", sum)
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
