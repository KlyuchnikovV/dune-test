package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

type BreakdownData struct {
	Day        string      `json:"day"`
	Symbol     string      `json:"symbol"`
	Amount     json.Number `json:"amount"`
	AmountUSD  json.Number `json:"usd_amount"`
	Persentage json.Number `json:"percentage"`
	Change1D   json.Number `json:"change_1d"`
	Change7D   json.Number `json:"change_7d"`
	Change30D  json.Number `json:"change_30d"`
}

type RatioData struct {
	Day             string      `json:"day"`
	TotalRestaked   json.Number `json:"total_restaked"`
	RestakedRatio   json.Number `json:"restaked_ratio"`
	RetakedBillion  json.Number `json:"retaked_billion"`
	RestakedPercent json.Number `json:"restaked_percent"`
}

const (
	breakdownURL = "https://api.dune.com/api/v1/query/3592795/results/csv?api_key=%s"
	ratioURL     = "https://api.dune.com/api/v1/query/3592784/results/csv?api_key=%s"
)

func main() {
	if len(os.Args) < 3 || len(os.Args[1]) == 0 || len(os.Args[2]) == 0 {
		fmt.Printf("not enough arguments passed: expected token and result file path")
		os.Exit(1)
	}

	var (
		bar        = charts.NewBar()
		token      = os.Args[1]
		resultFile = os.Args[2]
	)
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{Title: "Dune Test"}))

	breakdownData, err := fetchBreakdownData(fmt.Sprintf(breakdownURL, token))
	if err != nil {
		fmt.Printf("can't fetch data: %s", err)
		os.Exit(1)
	}

	ratioData, xAxis, err := fetchRatioData(fmt.Sprintf(ratioURL, token))
	if err != nil {
		fmt.Printf("can't fetch data: %s", err)
		os.Exit(1)
	}

	bar.Renderer = NewRenderer(bar, breakdownData...)
	bar.SetXAxis(xAxis).
		AddSeries("Total Restaked", ratioData)

	file, err := os.Create(resultFile)
	if err != nil {
		fmt.Printf("can't create file: %s", err)
		os.Exit(1)
	}

	if err := bar.Render(file); err != nil {
		fmt.Printf("can't render charts: %s", err)
		os.Exit(1)
	}
}
