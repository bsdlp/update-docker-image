package main

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"os"
)

func main() {
	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(os.Stdout, "pubkey:")
	fmt.Fprintln(os.Stdout, base64.StdEncoding.EncodeToString(pub))

	fmt.Fprintln(os.Stdout, "privkey:")
	fmt.Fprintln(os.Stdout, base64.StdEncoding.EncodeToString(priv))
}
