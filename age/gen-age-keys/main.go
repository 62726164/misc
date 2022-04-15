package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"filippo.io/age"
)

// genKey - This is equivalent to age-keygen from the cli
func genKey(user string) {
	identity, err := age.GenerateX25519Identity()
	if err != nil {
		log.Fatalf("Failed to generate key pair: %v", err)
	}

	fmt.Printf("# created: %s\n", time.Now().Format(time.RFC3339))
	fmt.Printf("# public key: %s\n", identity.Recipient().String())
	fmt.Printf("%s\n", identity.String())
}

// pwEncrypt - Encrypt with a password.
func pwEncrypt(password string) {
	nsr, err := age.NewScryptRecipient(password)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create("enc.age")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	w, err := age.Encrypt(f, nsr)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := io.WriteString(w, "The plain text.\n"); err != nil {
		log.Fatal(err)
	}

	if err := w.Close(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	users := []string{"one", "two", "three"}

	for _, user := range users {
		genKey(user)
	}

	pwEncrypt("howdy there partner")
}