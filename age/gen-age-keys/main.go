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

// pwEncrypt - Encrypt string with a password.
func pwEncrypt(filename, plaintext, password string) {
	nsr, err := age.NewScryptRecipient(password)
	if err != nil {
		log.Fatal(err)
	}

	out, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	aw := armor.NewWriter(out)

	ae, err := age.Encrypt(aw, nsr)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := io.WriteString(ae, plaintext); err != nil {
		log.Fatal(err)
	}

	if err := ae.Close(); err != nil {
		log.Fatal(err)
	}

	if err := aw.Close(); err != nil {
		log.Fatal(err)
	}
}

// keyEncrypt - Encrypt string with a key.
func keyEncrypt(filename, plaintext, pubkey string) {
	nxr, err := age.ParseX25519Recipient(pubkey)
	if err != nil {
		log.Fatal(err)
	}

	out, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	ae, err := age.Encrypt(out, nxr)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := io.WriteString(ae, plaintext); err != nil {
		log.Fatal(err)
	}

	if err := ae.Close(); err != nil {
		log.Fatal(err)
	}
}

// fileEncrypt - Encrypt a file
func fileEncrypt(filename, pubkey string) {
	// The recipient
	nxr, err := age.ParseX25519Recipient(pubkey)
	if err != nil {
		log.Fatal(err)
	}

	// The file to encrypt
	in, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer in.Close()

	// The encrypted file
	out, err := os.Create(filename + ".age")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	ae, err := age.Encrypt(out, nxr)
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 4096)

	for {
		read, err := in.Read(buf)
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}

		if _, err := io.WriteString(ae, string(buf[:read])); err != nil {
			log.Fatal(err)
		}
	}

	if err := ae.Close(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Generate keys
	users := []string{"one", "two", "three"}

	for _, user := range users {
		genKey(user)
	}

	// Encrypt
	pwEncrypt("password-encrypted.txt.age", "the plain text", "howdy there partner")
	keyEncrypt("key-encrypted.txt.age", "the plain text", "age1x6xa2agttdw2ejldtun9fgx2xwlen45h96uc8ef2g6avtggdc3gqrzywl2")
	fileEncrypt("test.bin", "age1x6xa2agttdw2ejldtun9fgx2xwlen45h96uc8ef2g6avtggdc3gqrzywl2")
}
