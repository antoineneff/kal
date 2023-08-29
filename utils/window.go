package utils

import (
	"fmt"
	"os"
	"os/exec"
)

func Clear() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func Reset() {
	// Erase content from cursor to beginning of screen
	fmt.Printf("\033[1J")
	// Move cursor to home position (0, 0)
	fmt.Printf("\033[H")
}
