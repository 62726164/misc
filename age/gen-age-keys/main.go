package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"filippo.io/age"
	"filippo.io/age/armor"
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

	f, err := os.Create("penc.age")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	aw := armor.NewWriter(f)

	w, err := age.Encrypt(aw, nsr)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := io.WriteString(w, "The plain text.\n"); err != nil {
		log.Fatal(err)
	}

	if err := w.Close(); err != nil {
		log.Fatal(err)
	}

	if err := aw.Close(); err != nil {
		log.Fatalf("Failed to close armor: %v", err)
	}
}

// keyEncrypt - Encrypt with a key.
func keyEncrypt(pubkey string) {
	nxr, err := age.ParseX25519Recipient(pubkey)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create("kenc.age")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	w, err := age.Encrypt(f, nxr)
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
	keyEncrypt("age1x6xa2agttdw2ejldtun9fgx2xwlen45h96uc8ef2g6avtggdc3gqrzywl2")
}
