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
		if string(b) == "q" {
			break
		}
		fmt.Print(b)
	}

}
