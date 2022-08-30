package main

import (
	"fmt"
	"os"
	app "review/app"
)

func main() {
	inputArgs := os.Args[1:]
	year, month, err := app.ParseArguments(inputArgs)
	if year == 0 && month == 0 {
		fmt.Print("年月を正しく入力してください。\n")
		os.Exit(1)
	}
	if err != nil {
		fmt.Errorf("ErrParseArguments: %w\n", err)
	}

	holidays, err := app.GetAllHolidaysInYear(year)
	if err != nil {
		fmt.Errorf("ErrGetDate: %w\n", err)
	}

	convertedHolidays, err := app.ConvertToHoliday(holidays)
	if err != nil {
		fmt.Errorf("ErrConvertToDate: %w\n", err)
	}

	pickedHolidays := app.PickByMonth(convertedHolidays, month)

	var countBusinessDays, countHolidayIsWeekday int
	for _, holiday := range pickedHolidays {
		isWeekday, err := holiday.IsWeekday()
		if err != nil {
			fmt.Errorf("ErrIsWeekday: %w\n", err)
		}

		if isWeekday {
			countHolidayIsWeekday++
		}

		businessDays, err := holiday.TotalWeekdays()
		if err != nil {
			fmt.Errorf("ErrTotalWeekdays: %w\n", err)
		}

		countBusinessDays = businessDays
	}

	if len(pickedHolidays) == 0 {
		tempHoliday := &app.Holiday{Year: year, Month: month}
		countBusinessDays, err = tempHoliday.TotalWeekdays()
		if err != nil {
			fmt.Errorf("ErrTotalWeekdays: %w\n", err)
		}
	}

	fmt.Printf("平日: %d日\n平日の祝日: %d日\n", countBusinessDays, countHolidayIsWeekday)
}
