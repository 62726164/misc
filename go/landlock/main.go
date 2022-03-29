package main

import (
	"flag"
	"fmt"
	"github.com/landlock-lsm/go-landlock/landlock"
	"io/ioutil"
	"log"
	"os"
)

// simple program that demonstrates how landlock works in Go on Linux systems.
// Requires 5.13 or newer kernel and .config should look something like this:
// CONFIG_SECURITY_LANDLOCK=y
// CONFIG_LSM="landlock,lockdown,yama,loadpin,safesetid,integrity,apparmor,selinux,smack,tomoyo"
func main() {
	var help = flag.Bool("help", false, "landlock-example -f /path/to/file.txt")
	var file = flag.String("f", "", "the file path to read")

	flag.Parse()
	if *help || len(os.Args) == 1 {
		flag.PrintDefaults()
		return
	}

	// allow the program to read files in /home/user/tmp
	err := landlock.V1.RestrictPaths(landlock.RODirs("/home/user/tmp"))
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
