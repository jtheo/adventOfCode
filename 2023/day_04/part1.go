package main

func part1(games data) int {
	res := 0
	for _, c := range games {
		tmp := 0
		for _, n := range c.winning {
			if contains(n, c.played) {
				if tmp == 0 {
					tmp = 1
				} else {
					tmp *= 2
				}
				continue
			}
		}
		res += tmp
	}

	return res
}
