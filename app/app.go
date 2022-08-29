package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
)

// GetDate httpリクエストで祝日をもらう (リクエスト先:https://holidays-jp.github.io/api/v1/{year}/date.json)
func GetDate(year int) ([]string, error) {
	httpClient := &http.Client{}

	url := fmt.Sprintf("https://holidays-jp.github.io/api/v1/%d/date.json", year)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	response, err := httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var holidays map[string]string
	if err := json.Unmarshal(body, &holidays); err != nil {
		return nil, nil
	}

	var holidayDates []string
	for date, _ := range holidays {
		holidayDates = append(holidayDates, date)
	}

	sort.Slice(holidayDates, func(i, j int) bool {
		return holidayDates[i] < holidayDates[j]
	})

	return holidayDates, nil
}
