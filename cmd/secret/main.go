package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

// Table to look up when generating each char in string.
const tableFull = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
const tableNumbersOnly = "0123456789"

func main() {
	// Parse flags.
	secretLen := flag.Int("l", 32, "length of secret")
	numbersOnly := flag.Bool("n", false, "numbers only")
	flag.Parse()

	// Deciding which table to use.
	var table = tableFull
	if *numbersOnly {
		table = tableNumbersOnly
	}

	// Generating secret.
	// By changing seed first, or we'll generating same string...
	rand.Seed(time.Now().Unix())
	out := make([]byte, *secretLen)
	tableLen := len(table)
	for i := 0; i < *secretLen; i++ {
		out[i] = table[rand.Intn(tableLen)]
	}

	// Just print result to stdout.
	fmt.Println(string(out))
}
