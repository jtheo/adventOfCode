package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func convert_not_working(input []string) []string {
	re := regexp.MustCompile("(one|two|three|four|five|six|seven|eight|nine)")
	ret := []string{}
	for _, l := range input {
		l := re.ReplaceAllStringFunc(l, replMap)
		ret = append(ret, l)
	}
	return ret
}

func replMap(s string) string {
	return fmt.Sprintf("%d", convMap[s])
}

func calculate(input []string) int {
	sum := 0
	for _, l := range input {
		numbers := []int{}
		for _, r := range l {
			if unicode.IsDigit(r) {
				v, _ := strconv.Atoi(string(r))
				numbers = append(numbers, v)
			}
		}
		tmp := (numbers[0] * 10) + numbers[len(numbers)-1]

		sum += tmp
	}
	return sum
}

func convert(input []string) []string {
	ret := []string{}

	for _, l := range input {
		i := 0
		n := ""

		for i < len(l) {
			if l[i] >= '0' && l[i] <= '9' {
				n += string(l[i])
				i++
				continue
			}

			for k, v := range convMap {
				if strings.HasPrefix(l[i:], k) {
					n += fmt.Sprintf("%d", v)
					i += len(k) - 2
					break
				}
			}

			i++
		}
		ret = append(ret, n)
	}

	return ret
}
