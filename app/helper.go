package app

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// ParseArguments CLI引数を二つもらって年月を返す関数
func ParseArguments(inputArgs []string) (int, int, error) {

	if len(inputArgs) < 2 {
		return 0, 0, errors.New("errors arguments length is less then 2")
	}

	year, err := strconv.Atoi(inputArgs[0])
	if err != nil {
		fmt.Errorf("ErrAtoi: %w\n", err)
	}

	month, err := strconv.Atoi(inputArgs[1])
	if month > 12 || month <= 0 {
		return 0, 0, errors.New("errors month should be greater then 0 and less then 13")
	}
	if err != nil {
		fmt.Errorf("ErrAtoi: %w\n", err)
	}

	return year, month, nil
}

// ConvertToHoliday httpリクエストからもらった祝日リストをHolidayに変換する関数
func ConvertToHoliday(dates []string) (holidays []*Holiday, err error) {
	for _, date := range dates {
		split := strings.Split(date, "-")

		year, err := strconv.Atoi(split[0])
		if err != nil {
			return nil, err
		}
		month, err := strconv.Atoi(split[1])
		if err != nil {
			return nil, err
		}
		day, err := strconv.Atoi(split[2])
		if err != nil {
			return nil, err
		}

		holiday := &Holiday{Year: year, Month: month, Day: day}
		holidays = append(holidays, holiday)
	}
	return
}

// PickByMonth 特定の月をもらって該当するHolidayだけ返す関数
func PickByMonth(holidays []*Holiday, month int) []*Holiday {
	var pickedHoliday []*Holiday
	for _, holiday := range holidays {
		if holiday.Month == month {
			pickedHoliday = append(pickedHoliday, holiday)
		}
	}
	return pickedHoliday
}
