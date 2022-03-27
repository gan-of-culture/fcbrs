package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gan-of-culture/fcbrs/fcbrs"
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) < 3 {
		os.Exit(1)
	}

	a, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatal(err.Error())
	}

	b, err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatal(err.Error())
	}

	c, err := strconv.Atoi(args[2])
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(fcbrs.Solve(a, b, c))
}
