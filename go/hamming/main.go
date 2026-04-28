//go:build ignore

package main

import (
	"fmt"
	"hamming"
)

func main() {
	d, err := hamming.Distance("GAGCCTACTAACGGGAT", "CATCGTAATGACGGCCT")
	fmt.Println(d, err)
}
