package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

// Table to look up when generating each char in string.
const alphaNumerics = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
const forNumbersOnly = "0123456789"

// From https://owasp.org/www-community/password-special-characters
const specialChars = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
const specialCharsAndSpace = " " + specialChars

func main() {
	// Parse flags.
	secretLen := flag.Int("l", 20, "length of secret")
	numericsOnly := flag.Bool("n", false, "numbers only")
	withSpecialChar := flag.Bool("s", false, "with special characters")
	withSpecialCharAndSpace := flag.Bool("ss", false, "with special characters including space")
	flag.Parse()

	// Deciding which table to use.
	var table = alphaNumerics
	if *numericsOnly {
		table = forNumbersOnly
	} else if *withSpecialCharAndSpace {
		table = table + specialCharsAndSpace
	} else if *withSpecialChar {
		table = table + specialChars
	}

	// Generating secret.
	// After Go 1.20, this is suggested way to seed random generator.
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	out := make([]byte, *secretLen)
	tableLen := len(table)
	for i := 0; i < *secretLen; i++ {
		out[i] = table[rnd.Intn(tableLen)]
	}

	// Just print result to stdout.
	fmt.Println(string(out))
}
