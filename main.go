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
const specialCharsExcludeSpace = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
const specialChars = " " + specialCharsExcludeSpace

func main() {
	// Parse flags.
	secretLen := flag.Int("l", 20, "length of secret")
	numericsOnly := flag.Bool("n", false, "numbers only")
	withSpecialChar := flag.Bool("s", false, "with special characters")
	excludeSpace := flag.Bool("s-no-space", false, "with special characters but excludes space")
	flag.Parse()

	// Deciding which table to use.
	var table = alphaNumerics
	if *numericsOnly {
		table = forNumbersOnly
	}
	if *withSpecialChar {
		if *excludeSpace {
			table = table + specialCharsExcludeSpace
		} else {
			table = table + specialChars
		}
	}

	// Generating secret.
	// By changing seed first, or we'll generating same string...
	rand.Seed(time.Now().UnixNano())
	out := make([]byte, *secretLen)
	tableLen := len(table)
	for i := 0; i < *secretLen; i++ {
		out[i] = table[rand.Intn(tableLen)]
	}

	// Just print result to stdout.
	fmt.Println(string(out))
}
