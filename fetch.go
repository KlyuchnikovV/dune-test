package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/go-echarts/go-echarts/v2/opts"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func fetchData(
	url string,
) ([][]string, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("can't fetch data: %w", err)
	}
	defer response.Body.Close()

	rows, err := csv.NewReader(response.Body).ReadAll()
	if err != nil {
		return nil, fmt.Errorf("can't read data: %w", err)
	}

	return rows, nil
}

func fetchRatioData(
	url string,
) ([]opts.BarData, []string, error) {
	rows, err := fetchData(url)
	if err != nil {
		return nil, nil, fmt.Errorf("can't fetch rows: %w", err)
	}

	sort.SliceStable(rows, func(i, j int) bool {
		return rows[i][0] < rows[j][0]
	})

	var (
		xAxis  = make([]string, len(rows))
		result = make([]opts.BarData, len(rows))
	)
	for i := 1; i < len(rows); i++ {
		xAxis[i] = rows[i][0]
		result[i] = opts.BarData{
			Name:  rows[i][0],
			Value: rows[i][1],
			Tooltip: &opts.Tooltip{
				Show:      opts.Bool(true),
				TriggerOn: "mousemove",
			},
		}
	}

	return result, xAxis, nil
}

func fetchBreakdownData(
	url string,
) ([][]string, error) {
	rows, err := fetchData(url)
	if err != nil {
		return nil, fmt.Errorf("can't fetch rows: %w", err)
	}

	if len(rows) == 0 {
		return rows, nil
	}

	rows[0] = rows[0][1:]
	for i := range rows[0] {
		rows[0][i] = cases.Title(
			language.AmericanEnglish,
		).String(strings.ReplaceAll(rows[0][i], "_", " "))
	}

	for i := 1; i < len(rows); i++ {
		rows[i] = rows[i][1:]
		for j := 2; j < len(rows[i]); j++ {
			value, err := strconv.ParseFloat(rows[i][j], 64)
			if err != nil {
				return nil, fmt.Errorf("can't parse number: %w", err)
			}

			if j > 2 {
				rows[i][j] = fmt.Sprintf("%.2f%%", value*100)
			} else {
				rows[i][j] = fmt.Sprintf("%.2f", value)
			}
		}
	}

	return rows, nil
}
