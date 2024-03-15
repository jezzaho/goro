package main

import (
	"github.com/jezzaho/goro/cli"
	"github.com/jezzaho/goro/cmd/internal"
)

// "github.com/jezzaho/goro/cmd/internal"
// 	"github.com/jezzaho/goro/cli"

func main() {
	// input := cli.CliInput{
	// 	Line:      "LH",
	// 	Beggining: "19JUL24",
	// 	Ending:    "19JUL24",
	// }
	input := cli.RenderMainScreen()
	LHQuery := []internal.ApiQuery{
		{
			Airline:         "LH",
			StartDate:       input.Beggining,
			EndDate:         input.Ending,
			DaysOfOperation: "1234567",
			TimeMode:        "LT",
			Origin:          "KRK",
			Destination:     "FRA",
		},
		{
			Airline:         "LH",
			StartDate:       input.Beggining,
			EndDate:         input.Ending,
			DaysOfOperation: "1234567",
			TimeMode:        "LT",
			Origin:          "KRK",
			Destination:     "MUC",
		},
	}
	OSQuery := []internal.ApiQuery{
		{
			Airline:         "OS",
			StartDate:       input.Beggining,
			EndDate:         input.Ending,
			DaysOfOperation: "1234567",
			TimeMode:        "LT",
			Origin:          "KRK",
			Destination:     "VIE",
		},
	}
	LXQuery := []internal.ApiQuery{
		{
			Airline:         "LX",
			StartDate:       string(input.Beggining),
			EndDate:         string(input.Ending),
			DaysOfOperation: "1234567",
			TimeMode:        "LT",
			Origin:          "KRK",
			Destination:     "ZRH",
		},
	}
	SNQuery := []internal.ApiQuery{
		{
			Airline:         "SN",
			StartDate:       input.Beggining,
			EndDate:         input.Ending,
			DaysOfOperation: "1234567",
			TimeMode:        "LT",
			Origin:          "KRK",
			Destination:     "BRU",
		},
	}
	var queryList []internal.ApiQuery

	switch string(input.Line[:2]) {
	case "LH":
		queryList = LHQuery
	case "OS":
		queryList = OSQuery
	case "LX":
		queryList = LXQuery
	case "SN":
		queryList = SNQuery
	default:
		queryList = []internal.ApiQuery{}
	}
	apiAuth := internal.PostForAuth()
	apiData := internal.GetApiData(queryList, apiAuth)
	flattened := internal.FlattenJSON(apiData)
	println(string(flattened))
	internal.CreateCSVFromResponse(flattened)

	cli.RenderFinal()
	//FRONTEND

}
