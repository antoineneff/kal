package main

import (
	"fmt"
	"time"

	"github.com/antoineneff/kal/utils"
	"github.com/eiannone/keyboard"
)

var now = time.Now()
var currentYear = now.Year()
var currentMonth = now.Month()

func updateMonth(dir int) {
	if int(currentMonth) == 12 && dir > 0 {
		currentYear += 1
		currentMonth = 1
	} else if int(currentMonth) == 1 && dir < 0 {
		currentYear -= 1
		currentMonth = 12
	} else if dir < 0 {
		currentMonth -= 1
	} else {
		currentMonth += 1
	}
}

func main() {
	utils.Clear()

	printMonth()

	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	for {
		fmt.Println("Use arrow keys to change month and year")
		fmt.Println("Press ESC or q to quit")

		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		if key == keyboard.KeyEsc || char == 'q' {
			break
		}

		switch key {
		case keyboard.KeyArrowLeft:
			updateMonth(-1)
		case keyboard.KeyArrowRight:
			updateMonth(1)
		case keyboard.KeyArrowUp:
			currentYear += 1
		case keyboard.KeyArrowDown:
			currentYear += -1
		}
		utils.Reset()
		printMonth()
	}
}
