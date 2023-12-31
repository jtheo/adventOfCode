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

func main() {
	var red, green, blue int
	flag.IntVar(&red, "red", 0, "how many red cubes")
	flag.IntVar(&blue, "blue", 0, "how many blue cubes")
	flag.IntVar(&green, "green", 0, "how many green cubes")
	flag.Parse()

	if red == 0 || blue == 0 || green == 0 {
		flag.Usage()
		return
	}

	if len(flag.Args()) != 1 {
		log.Fatal("I need a filename, len args is", flag.Args())
	}

	input, err := getData(flag.Args()[0])
	if err != nil {
		log.Fatal(err)
	}

	want := map[string]int{
		"red":   red,
		"blue":  blue,
		"green": green,
	}
	fmt.Printf("Looking for Blue: %d, Green: %d, Red: %d\n", want["blue"], want["green"], want["red"])
	b := summarize(input, want)

	sum := 0
	for _, i := range b {
		sum += i
	}
	fmt.Printf("List of valid Games: %+v\n\nResult is: %d\n", b, sum)

	part2 := summarize2(input, want)
	fmt.Printf("Result part 2: %d\n", part2)
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

type bag []cubes

type cubes struct {
	red, green, blue int
}

func summarize(input []string, want map[string]int) []int {
	b := []int{}
	for i, l := range input {
		game := i + 1
		tmp := strings.Split(l, ": ")
		if checkGame(tmp[1], want) {
			// fmt.Println(l)
			b = append(b, game)
		}
	}

	return b
}

func checkGame(game string, want map[string]int) bool {
	ret := true
	sets := strings.Split(game, "; ")

	for _, s := range sets {
		c := strings.Split(s, ", ")
		for _, nc := range c {
			t := strings.Split(nc, " ")
			colour := t[1]
			num, err := strconv.Atoi(strings.TrimSpace(t[0]))
			if err != nil {
				for i := range t {
					fmt.Printf("%d: %v\n", i, t[i])
				}
				log.Fatalf("%q is not a number\nline: %v\n%v\n", t[0], t, err)
			}
			// fmt.Printf("num: %d > want[colour]: %v, colour: %s\n", num, want[colour], colour)
			if num > want[colour] {
				// fmt.Printf("%s is too big: %d\n", colour, num)
				ret = false
				break
			}
		}
	}

	return ret
}

func getCubes(game string, want map[string]int) int {
	sets := strings.Split(game, "; ")
	maxForColour := map[string]int{
		"red":   0,
		"blue":  0,
		"green": 0,
	}

	for _, s := range sets {
		c := strings.Split(s, ", ")
		for _, nc := range c {
			t := strings.Split(nc, " ")
			colour := t[1]
			num, err := strconv.Atoi(strings.TrimSpace(t[0]))
			if err != nil {
				for i := range t {
					fmt.Printf("%d: %v\n", i, t[i])
				}
				log.Fatalf("%q is not a number\nline: %v\n%v\n", t[0], t, err)
			}
			if maxForColour[colour] < num {
				maxForColour[colour] = num
			}
		}
	}

	sumCubes := 1
	for _, n := range maxForColour {
		sumCubes *= n
	}
	return sumCubes
}

func summarize2(input []string, want map[string]int) int {
	sum := 0
	for _, l := range input {
		tmp := strings.Split(l, ": ")
		g := getCubes(tmp[1], want)
		sum += g
	}

	return sum
}
