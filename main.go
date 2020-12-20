package main

import (
	"fmt"
	"golang.org/x/term"
	"os"
)

// Editor global state. For now hold terminal size
type Editor struct {
	wx, wy int
}

func main() {
	oldState, err := term.MakeRaw(0)
	if err != nil {
		panic(err)
	}
	defer term.Restore(0, oldState)

	// Buffer to store input
	var b []byte = make([]byte, 1)

	wx, wy, err := term.GetSize(0)
	if err != nil {
		panic(err)
	}
	editor := &Editor{
		wx: wx,
		wy: wy,
	}

	for run := true; run; {
		refresh(editor)
		run = processKeyPress(b)
	}
}

// Handle keypress event
func processKeyPress(b []byte) bool {

	switch ch := readKey(b); {

	// ASCII 17 (CTRL + q) as quit -> b[0] == 17 .
	// By design, CTRL+char ASCII value can be calculated by bitwise-AND
	// binary 00011111 (0x1f) with char.
	case ch == 0x1f&'q':
		// Clear screen on exit.
		fmt.Print("\x1b[H\x1b[2J")
		return false

	// Skip control characters. ASCII codes 0–31 are all control characters.
	// 127 is also a control character. 32–126 are all printable.
	case ch < 32:
		fallthrough
	case ch == 127:
		break
	default:

	}
	return true
}

// Wait for a keypress and return its value
func readKey(b []byte) rune {
	os.Stdin.Read(b)
	return rune(b[0])
}

func refresh(editor *Editor) {
	// <esc>[2J clear the entire screen
	fmt.Print("\x1b[2J")
	// <esc>[1;1H position the cursor to the coordinate (1,1) i.e. top left.
	// row and column number starts with 1. default argument for H is 1.
	// <esc>[H is equivalent to <esc>[1;1H
	fmt.Print("\x1b[H")

	drawRows(editor)
	// Reposition cursor after draw
	fmt.Print("\x1b[H")
}

// Handle drawing each row of the buffer of text being edited.
// Draws a tilde in each row, which means that row is not part of the file
// and can’t contain any text.
func drawRows(editor *Editor) {
	// the screen buffer string
	var screen string
	for y := 0; y < editor.wy; y++ {
		screen += "~"
		if y < editor.wy-1 {
			screen += "\r\n"
		}
	}
	fmt.Print(screen)
}
