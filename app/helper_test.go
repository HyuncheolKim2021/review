package app_test

import (
	"github.com/google/go-cmp/cmp"
	"review/app"
	"testing"
)

func TestParseArguments(t *testing.T) {
	cases := []struct {
		name    string
		input   []string
		expect1 int
		expect2 int
	}{
		{
			name:    "CLIからもらった引数を年月に変換成功",
			input:   []string{"2022", "1"},
			expect1: 2022,
			expect2: 1,
		},
		{
			name:    "CLIからもらった引数が一つしかない場合初期値である0だけ返す",
			input:   []string{"2022"},
			expect1: 0,
			expect2: 0,
		},
		{
			name:    "CLIからもらった引数が三つ以上でも正常的に年月を返す",
			input:   []string{"2022", "10", "31"},
			expect1: 2022,
			expect2: 10,
		},
		{
			name:    "第2引数でもらった数値が0だった場合両方0を返して終了する",
			input:   []string{"2022", "0"},
			expect1: 0,
			expect2: 0,
		},
		{
			name:    "第2引数でもらった数値が13以上だった場合両方0を返して終了する",
			input:   []string{"2022", "13"},
			expect1: 0,
			expect2: 0,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			year, month, _ := app.ParseArguments(c.input)
			if diff := cmp.Diff(year, c.expect1); diff != "" {
				t.Errorf("ParseArguments year diff (-got +want:\n%v", diff)
			}
			if diff := cmp.Diff(month, c.expect2); diff != "" {
				t.Errorf("ParseArguments month diff (-got +want:\n%v", diff)
			}
		})
	}
}

func TestConvertToHoliday(t *testing.T) {
	cases := []struct {
		name    string
		input   []string
		expect1 []*app.Holiday
	}{
		{
			name:    "もらったの日時データを元にHoliday構造体に変換成功",
			input:   []string{"2022-01-01"},
			expect1: []*app.Holiday{{Year: 2022, Month: 1, Day: 1}},
		},
		{
			name:    "祝日データがない場合nilになる",
			input:   []string{},
			expect1: nil,
		},
		{
			name:    "当月に祝日が二日以上ある場合長さ2の配列を返す",
			input:   []string{"2022-01-01", "2022-01-10"},
			expect1: []*app.Holiday{{Year: 2022, Month: 1, Day: 1}, {Year: 2022, Month: 1, Day: 10}},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			holidays, _ := app.ConvertToHoliday(c.input)
			if diff := cmp.Diff(holidays, c.expect1); diff != "" {
				t.Errorf("ConvertToHoliday diff (-got +want:\n%v", diff)
			}
		})
	}
}

func TestPickByMonth(t *testing.T) {
	holidaysIn2022 := []*app.Holiday{
		{Year: 2022, Month: 1, Day: 1},
		{Year: 2022, Month: 1, Day: 10},
		{Year: 2022, Month: 2, Day: 11},
		{Year: 2022, Month: 3, Day: 21},
		{Year: 2022, Month: 4, Day: 29},
		{Year: 2022, Month: 5, Day: 3},
		{Year: 2022, Month: 5, Day: 4},
		{Year: 2022, Month: 5, Day: 5},
		{Year: 2022, Month: 7, Day: 18},
		{Year: 2022, Month: 8, Day: 11},
		{Year: 2022, Month: 9, Day: 19},
		{Year: 2022, Month: 9, Day: 23},
		{Year: 2022, Month: 10, Day: 10},
		{Year: 2022, Month: 11, Day: 3},
		{Year: 2022, Month: 11, Day: 23},
	}

	cases := []struct {
		name   string
		input1 []*app.Holiday
		input2 int
		expect []*app.Holiday
	}{
		{
			name:   "引数で1年全ての祝日と1月を指定して1月の祝日だけ返す",
			input1: holidaysIn2022,
			input2: 1,
			expect: []*app.Holiday{{Year: 2022, Month: 1, Day: 1}, {Year: 2022, Month: 1, Day: 10}},
		}, {
			name:   "引数で1年全ての祝日と祝日のない6月を指定した場合nilが返ってくる",
			input1: holidaysIn2022,
			input2: 6,
			expect: nil,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			pickedHolidays := app.PickByMonth(c.input1, c.input2)
			if diff := cmp.Diff(pickedHolidays, c.expect); diff != "" {
				t.Errorf("PickByMonth diff (-got +want:\n%v", diff)
			}
		})
	}
}
