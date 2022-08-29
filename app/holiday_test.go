package app_test

import (
	"github.com/google/go-cmp/cmp"
	"review/app"
	"testing"
)

func TestHoliday_IsWeekday(t *testing.T) {
	cases := []struct {
		name    string
		holiday *app.Holiday
		expect  bool
	}{
		{
			name:    "2022年1月1日のHolidayは平日ではない",
			holiday: &app.Holiday{Year: 2022, Month: 1, Day: 1},
			expect:  false,
		},
		{
			name:    "2022年1月10日のHolidayは平日である",
			holiday: &app.Holiday{Year: 2022, Month: 1, Day: 10},
			expect:  true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			isWeekday, _ := c.holiday.IsWeekday()
			if diff := cmp.Diff(isWeekday, c.expect); diff != "" {
				t.Errorf("*app.Holiday.IsWeekday diff (-got +want:\n%v", diff)
			}
		})
	}
}

func TestHoliday_TotalWeekdays(t *testing.T) {
	cases := []struct {
		name    string
		holiday *app.Holiday
		expect  int
	}{
		{
			name:    "2022年1月のHolidayから平日数を計算すると21日になる",
			holiday: &app.Holiday{Year: 2022, Month: 1},
			expect:  21,
		},
		{
			name:    "2022年2月のHolidayから平日数を計算すると20日になる",
			holiday: &app.Holiday{Year: 2022, Month: 2},
			expect:  20,
		},
		{
			name:    "2022年3月のHolidayから平日数を計算すると23日になる",
			holiday: &app.Holiday{Year: 2022, Month: 3},
			expect:  23,
		},
		{
			name:    "2022年4月のHolidayから平日数を計算すると21日になる",
			holiday: &app.Holiday{Year: 2022, Month: 4},
			expect:  21,
		},
		{
			name:    "2022年5月のHolidayから平日数を計算すると22日になる",
			holiday: &app.Holiday{Year: 2022, Month: 5},
			expect:  22,
		},
		{
			name:    "2022年6月のHolidayから平日数を計算すると22日になる",
			holiday: &app.Holiday{Year: 2022, Month: 6},
			expect:  22,
		},
		{
			name:    "2022年7月のHolidayから平日数を計算すると21日になる",
			holiday: &app.Holiday{Year: 2022, Month: 7},
			expect:  21,
		},
		{
			name:    "2022年8月のHolidayから平日数を計算すると23日になる",
			holiday: &app.Holiday{Year: 2022, Month: 8},
			expect:  23,
		},
		{
			name:    "2022年9月のHolidayから平日数を計算すると22日になる",
			holiday: &app.Holiday{Year: 2022, Month: 9},
			expect:  22,
		},
		{
			name:    "2022年10月のHolidayから平日数を計算すると21日になる",
			holiday: &app.Holiday{Year: 2022, Month: 10},
			expect:  21,
		},
		{
			name:    "2022年11月のHolidayから平日数を計算すると22日になる",
			holiday: &app.Holiday{Year: 2022, Month: 11},
			expect:  22,
		},
		{
			name:    "2022年12月のHolidayから平日数を計算すると22日になる",
			holiday: &app.Holiday{Year: 2022, Month: 12},
			expect:  22,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			totalWeekdays, _ := c.holiday.TotalWeekdays()
			if diff := cmp.Diff(totalWeekdays, c.expect); diff != "" {
				t.Errorf("*app.Holiday.TotalWeekdays diff (-got +want:\n%v", diff)
			}
		})
	}
}
