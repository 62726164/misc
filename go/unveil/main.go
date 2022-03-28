package main

import (
	"flag"
	"fmt"
	"golang.org/x/sys/unix"
	"io/ioutil"
	"log"
	"os"
)

// simple program that demonstrates how unveil works in Go on OpenBSD systems.
func main() {
	var help = flag.Bool("help", false, "unveil-example -f /path/to/file.txt")
	var file = flag.String("f", "", "the file path to read")

	flag.Parse()
	if *help || len(os.Args) == 1 {
		flag.PrintDefaults()
		return
	}

	// allow the program to read files in /home/user/tmp
	// these files must be present *before* the program runs
	// permissions may be r, w, x or c *or* any combination
	err := unix.Unveil("/home/user/tmp/", "r")
	if err != nil {
		log.Fatal(err)
	}

	// block future unveil calls
	err = unix.UnveilBlock()
	if err != nil {
		log.Fatal(err)
	}

	// attempt to read a file
	bytes, err := ioutil.ReadFile(*file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bytes))
}
