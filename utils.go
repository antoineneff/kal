package main

import (
	"fmt"
	"os"
	"os/exec"
)

func clear() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func reset() {
	// Erase content from cursor to beginning of screen
	fmt.Printf("\033[1J")
	// Move cursor to home position (0, 0)
	fmt.Printf("\033[H")
}
