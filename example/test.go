package main

import (
	"fmt"

	"github.com/aflnk/HolidayArrangement"
)

func main() {
	fmt.Println("Weekeed Holidays:")
	for day, dayType := range HolidayArrangement.Weekend.GetHolidays() {
		fmt.Println(day, dayType)
	}

	fmt.Println("China State Holidays:")
	for day, dayType := range HolidayArrangement.ChinaState.GetHolidays() {
		fmt.Println(day, dayType)
	}
}
