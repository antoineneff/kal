package main

import (
	"fmt"
	"math"
	"strings"
	"time"
)

func printWeek(startIndex, limit int) {
	fmt.Println("├────┼────┼────┼────┼────┼────┼────┤")
	fmt.Print("│")
	for i := startIndex; i < startIndex+7; i++ {
		if i >= 0 && i < limit {
			fmt.Printf(" %02d │", i+1)
		} else {
			fmt.Print("    │")
		}
	}
	fmt.Println("")
}

func printWeeks(firstDayIndex, numberOfDays int) {
	numberOfWeeks := int(math.Ceil(float64(firstDayIndex+numberOfDays) / 7))
	for i := 0; i < numberOfWeeks; i += 1 {
		startIndex := -firstDayIndex + i*7
		printWeek(startIndex, numberOfDays)
	}
	fmt.Println("└────┴────┴────┴────┴────┴────┴────┘")
}

func printHeaders(year int, month time.Month) {
	fmt.Printf("┌%s┐\n", strings.Repeat("─", 34))
	fmt.Printf("│ %v", year)
	fmt.Printf("%*s │\n", 28, month)
	fmt.Println("├────┬────┬────┬────┬────┬────┬────┤")
	fmt.Println("│ Su │ Mo │ Tu │ We │ Th │ Fr │ Sa │")
}

func monthDetails(year int, month time.Month) (int, int) {
	startOfMonth := time.Date(year, month, 1, 0, 0, 0, 0, now.Location())
	endOfMonth := time.Date(year, month+1, 1, 0, 0, 0, 0, now.Location())

	weekday := startOfMonth.Weekday()
	numberOfDays := endOfMonth.Sub(startOfMonth).Hours() / 24

	return int(weekday), int(numberOfDays)
}

func printMonth() {
	firstDayIndex, numberOfDays := monthDetails(currentYear, currentMonth)

	printHeaders(currentYear, currentMonth)
	printWeeks(firstDayIndex, numberOfDays)
}
