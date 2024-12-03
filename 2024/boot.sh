#!/usr/bin/env bash

function err() {
  echo "[ERROR]: ${*}"
  exit 1
}

today=$(date +%d)
newDay=${1:-${today}}

mkdir "${newDay}" || err "Can't create ${newDay}"
cd "${newDay}" || err "Can't cd to ${newDay}"

cat <<EOF >main.go
package main

import (
	"fmt"
)

func main() {
	o := run()

	fmt.Printf("Running Part %d\n", o.part)
	switch o.part {
	case 1:
		part1(o)
	case 2:
		part2(o)
	}
}
EOF

cat <<EOF >util.go
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type opts struct {
	fn      string
	verbose bool
	part    int
}

func run() opts {
	var fn string
	var verb bool
	var part int

	flag.StringVar(&fn, "fn", "", "Filename of the list")
	flag.BoolVar(&verb, "v", false, "Verbose output")
	flag.IntVar(&part, "part", 1, "Part of the day")
	flag.Parse()

	if fn == "" {
		fmt.Println("I need a filename")
		flag.Usage()
	}
	o := opts{}
	o.fn = fn
	o.verbose = verb
	o.part = part
	return o
}

func load(n string) ([]byte, error) {
	c, err := os.ReadFile(n)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
EOF

for part in {1..2}; do
  cat <<EOF >"part${part}.go"
package main

func part${part}(o opts) {

}
EOF
done

go mod init
