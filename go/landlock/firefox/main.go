package main

// Wrap firefox with landlock.
// More restrictive than firejail and no SUID.
// Requires Linux kernel 5.13 or later.

import (
	"github.com/landlock-lsm/go-landlock/landlock"
	"log"
	"os/exec"
)

func main() {
	err := landlock.V1.RestrictPaths(
		landlock.RWFiles("/dev/null"),
		landlock.RWDirs("/home/user/.cache/mozilla/firefox",
			"/home/user/.mozilla",
			"/home/user/Downloads",
			"/tmp",
			"/proc"),
		landlock.ROFiles("/home/user/.config/mimeapps.list"),
		landlock.RODirs("/dev",
			"/etc",
			"/home/user/.config/dconf",
			"/home/user/.config/ibus",
			"/home/user/.config/gtk-3.0",
			"/home/user/.local",
			"/home/user/.pki",
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
