package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	// strings are immutable sequences of bytes
	greek := "some greek: Τη γλώσσα μου έδωσαν"

	// the index operation returns a byte
	fmt.Println(greek[0])

	// the substring operation returns a string
	fmt.Println(greek[5:10])

	// strings can be decoded as bytes
	greekBytes := []byte(greek)
	fmt.Printf("greek decoded as bytes: %v\n", greekBytes)

	// or as utf8 unicode code points
	// these are named runes and are int32
	greekRunes := []rune(greek)
	fmt.Printf("greek decoded as runes: %v\n", greekRunes)

	// fancy decoding is required to
	// index the runes inside a string
	rune, bytesCount := utf8.DecodeRuneInString(greek[14:])
	runesCount := utf8.RuneCountInString(greek)
	fmt.Printf("found rune %c spanning %v bytes\n", rune, bytesCount)
	fmt.Printf("found %v runes\n", runesCount)

	// iterating is done over runes
	for range greek {
	}

}
