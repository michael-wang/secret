package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

// Table to look up when generating each char in string.
const table = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

func main() {
	l := flag.Int("l", 32, "length of string")
	flag.Parse()

	out := make([]byte, *l)
	n := len(table)
	// Change seed or we'll generating same string...
	rand.Seed(time.Now().Unix())
	for i := 0; i < *l; i++ {
		out[i] = table[rand.Intn(n)]
	}

	// Just print result to stdout.
	fmt.Println(string(out))
}
