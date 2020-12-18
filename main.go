package main

import (
	"fmt"
	"golang.org/x/term"
	"os"
)

func main() {

	// Buffer to store input
	var b []byte = make([]byte, 1)

	oldState, err := term.MakeRaw(0)
	if err != nil {
		panic(err)
	}
	defer term.Restore(0, oldState)

	for {
		os.Stdin.Read(b)

		// ASCII 17 (CTRL + q) as quit -> b[0] == 17 .
		// By design, CTRL+char ASCII value can be calculated by bitwise-AND
		// binary 00011111 (0x1f) with char.
		if b[0] == 0x1f&'q' {
			break
		}
		// Skip control characters. ASCII codes 0–31 are all control characters.
		// 127 is also a control character. 32–126 are all printable.
		if b[0] < 32 || b[0] > 126 {
			continue
		}
		fmt.Print(string(b))
	}

}
