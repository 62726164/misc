package main

import (
	"fmt"
	"log"
	"time"

	"filippo.io/age"
)

func genKey(user string) {
	// This is equivalent to age-keygen
	identity, err := age.GenerateX25519Identity()
	if err != nil {
		log.Fatalf("Failed to generate key pair: %v", err)
	}

	fmt.Printf("# created: %s\n", time.Now().Format(time.RFC3339))
	fmt.Printf("# public key: %s\n", identity.Recipient().String())
	fmt.Printf("%s\n", identity.String())
}

func main() {
	users := []string{"one", "two", "three"}

	for _, user := range users {
		genKey(user)
	}
}
