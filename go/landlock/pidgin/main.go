package main

// Wrap pidgin with landlock.
// More restrictive than firejail and no SUID.
// Requires Linux kernel 5.13 or later and .config should look something like this:
// CONFIG_SECURITY_LANDLOCK=y
// CONFIG_LSM="landlock,lockdown,yama,loadpin,safesetid,integrity,apparmor,selinux,smack,tomoyo"

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
		landlock.RWDirs("/proc"),
		landlock.RODirs("/dev",
			"/etc",
			fmt.Sprintf("%s/.config", home),
			fmt.Sprintf("%s/.purple", home),
			"/lib",
			"/run",
			"/sys",
			"/tmp",
			"/usr",
			"/var"))

	if err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command("/usr/bin/pidgin")
	log.Printf("Running command and waiting for it to finish...")
	err = cmd.Run()
	log.Printf("Command finished with error: %v", err)
}
