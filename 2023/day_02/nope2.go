package main

// func summarize(input []string, red, green, blue int) bag {
// 	b := bag{}
// 	re := regexp.MustCompile(`([0-9]+) ([a-z]+)`)
// 	for _, l := range input {
// 		t := strings.Split(l, ":")
// 		t2 := strings.FieldsFunc(t[1], func(r rune) bool {
// 			return r == ';' || r == ','
// 		})
//
// 		game := cubes{}
// 		for _, g := range t2 {
//
// 			match := re.FindStringSubmatch(strings.TrimSpace(g))
// 			if len(match) != 3 {
// 				for i, x := range match {
// 					log.Println(i, x)
// 				}
// 				log.Fatalf("Something wrong len(i) is %d!\nGame: %s\nCube: %+v\nre: %+v\n%q\n", len(match), l, g, match, match[len(match)-1])
// 			}
//
// 			color := match[2]
// 			n, err := strconv.Atoi(match[1])
// 			if err != nil {
// 				log.Fatalf("number of cubes failed conversion: Color: %s, number: %s, err: %v\n", color, match[0], err)
// 			}
//
// 			switch color {
// 			case "red":
// 				game.red += n
// 			case "blue":
// 				game.blue += n
// 			case "green":
// 				game.green += n
// 			}
// 		}
// 		b = append(b, game)
// 	}
//
// 	return b
// }
//
//
//
// sumCubes := cubes{}
// for i, g := range b {
// 	if g.red > red || g.green > green || g.blue > blue {
// 		continue
// 	}
// 	if sumCubes.red+g.red > red || sumCubes.blue+g.blue > blue || sumCubes.green+g.green > green {
// 		continue
// 	}
// 	sumCubes.red += g.red
// 	sumCubes.green += g.green
// 	sumCubes.blue += g.blue
// 	games = append(games, i+1)
// 	fmt.Printf("%d: %+v\n", i+1, g)
// }
//
//
//
//
