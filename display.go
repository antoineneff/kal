package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/antoineneff/kal/utils"
)

func printMonth() {
	printHeaders()
	printWeeks()
}

func printHeaders() {
	fmt.Printf("┌%s┐\n", strings.Repeat("─", 34))
	fmt.Printf("│ %v", currentYear)
	fmt.Printf("%*s │\n", 28, currentMonth)
	fmt.Println("├────┬────┬────┬────┬────┬────┬────┤")
	fmt.Println("│ Su │ Mo │ Tu │ We │ Th │ Fr │ Sa │")
}

func printWeeks() {
	firstDayIndex, numberOfDays := utils.MonthDetails(currentYear, currentMonth, now.Location())

	numberOfWeeks := int(math.Ceil(float64(firstDayIndex+numberOfDays) / 7))
	for i := 0; i < numberOfWeeks; i += 1 {
		startIndex := -firstDayIndex + i*7
		printWeek(startIndex, numberOfDays)
	}
	fmt.Println("└────┴────┴────┴────┴────┴────┴────┘")
}

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
