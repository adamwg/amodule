package main

import (
	"os"
	"strconv"

	"github.com/adamwg/amodule/blah"
)

func main() {
	times, _ := strconv.Atoi(os.Args[1])

	f := blah.NewFoo("hello!")
	f.Dance(times)
}
