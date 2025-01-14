package main

import (
	"encoding/base64"
	"errors"
	"flag"
	"fmt"
	"log"

	"github.com/groob/plist"
	"github.com/micromdm/micromdm/pkg/crypto/password"
)

func main() {
	var (
		flB64      = flag.Bool("b64", false, "Output base64-encoded Plist")
		flPassword = flag.String("password", "", "Password to hash")
	)
	flag.Parse()
	if *flPassword == "" {
		log.Fatal(errors.New("no password supplied"))
	}
	p, err := password.SaltedSHA512PBKDF2(*flPassword)
	if err != nil {
		log.Fatal(err)
	}
	x, err := plist.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}
	if *flB64 {
		fmt.Print(base64.StdEncoding.EncodeToString(x))
	} else {
		fmt.Print(string(x))
	}
}
