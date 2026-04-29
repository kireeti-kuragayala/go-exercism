//go:build ignore

package main

import (
	"affinecipher"
	"fmt"
)

func main() {
	plain_text := "test 123 786 hhg"
	fmt.Printf("Original text: '%s'\n", plain_text)
	enc, err := affinecipher.Encode(plain_text, 9, 11)
	fmt.Printf("Encoded: '%s'\n", enc)
	if err != nil {
		panic(err)
	}

	dec, err := affinecipher.Decode(enc, 9, 11)
	fmt.Printf("Decoded: '%s'\n", dec)
	if err != nil {
		panic(err)
	}
}
