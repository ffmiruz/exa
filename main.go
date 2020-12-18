package main

import (
	"fmt"
	"os"
)

func main() {

	// Buffer to store input
	var b []byte = make([]byte, 1)

	for {
		os.Stdin.Read(b)
		fmt.Println(b)
	}

}
