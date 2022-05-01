package main

// Wrap firefox with landlock.
// More restrictive than firejail and no SUID.
// Requires Linux kernel 5.13 or later.

import (
	"fmt"
	"github.com/landlock-lsm/go-landlock/landlock"
	"log"
	"os"
	"os/exec"
)

func main() {
	home, _ := os.UserHomeDir()

	err := landlock.V1.RestrictPaths(
		landlock.RWFiles("/dev/null"),
		landlock.RWDirs(fmt.Sprintf("%s/.cache/mozilla/firefox", home),
			fmt.Sprintf("%s/.mozilla", home),
			fmt.Sprintf("%s/Downloads", home),
			"/tmp",
			"/proc"),
		landlock.ROFiles(fmt.Sprintf("%s/.config/mimeapps.list", home)),
		landlock.RODirs("/dev",
			"/etc",
			fmt.Sprintf("%s/.config/dconf", home),
			fmt.Sprintf("%s/.config/ibus", home),
			fmt.Sprintf("%s/.config/gtk-3.0", home),
			fmt.Sprintf("%s/.local", home),
			fmt.Sprintf("%s/.pki", home),
			"/lib",
			"/run",
			"/sys",
			"/usr",
			"/var"))

	if err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command("/usr/bin/firefox-esr")
	log.Printf("Running command and waiting for it to finish...")
	err = cmd.Run()
	log.Printf("Command finished with error: %v", err)
}
