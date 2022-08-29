package app

import (
	"time"
)

type Holiday struct {
	Year  int
	Month int
	Day   int
}

// TotalWeekdays 東京時間軸を固定し、該当する月の平日数を計算するメソッド
func (h *Holiday) TotalWeekdays() (int, error) {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return 0, err
	}

	var businessDay int
	month := h.getMonth()
	start := time.Date(h.Year, month, 1, 0, 0, 0, 0, jst)
	for i := 0; i < 31; i++ {
		duration := time.Hour * 24 * time.Duration(i)
		add1Day := start.Add(duration)
		if add1Day.Month() != month {
			break
		}

		if add1Day.Weekday() == time.Sunday || add1Day.Weekday() == time.Saturday {
			continue
		}

		businessDay++
	}

	return businessDay, nil
}

// IsWeekday time.Weekday()で特定日が平日かどうかを判断するメソッド
func (h *Holiday) IsWeekday() (bool, error) {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return false, err
	}
	date := time.Date(h.Year, h.getMonth(), h.Day, 0, 0, 0, 0, jst)

	if date.Weekday() == time.Saturday || date.Weekday() == time.Sunday {
		return false, nil
	} else {
		return true, nil
	}
}

// getMonth Holidayの持つ月情報(int)からtime.Monthに変換するPrivateメソッド
func (h *Holiday) getMonth() time.Month {
	switch h.Month {
	case 1:
		return time.January
	case 2:
		return time.February
	case 3:
		return time.March
	case 4:
		return time.April
	case 5:
		return time.May
	case 6:
		return time.June
	case 7:
		return time.July
	case 8:
		return time.August
	case 9:
		return time.September
	case 10:
		return time.October
	case 11:
		return time.November
	case 12:
		return time.December
	default:
		return 0
	}
}
