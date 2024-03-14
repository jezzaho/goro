package main

import (
	"github.com/jezzaho/goro/cli"
	"github.com/jezzaho/goro/cmd/internal"
)

// "github.com/jezzaho/goro/cmd/internal"
// 	"github.com/jezzaho/goro/cli"

func main() {

	input := cli.RenderMainScreen()
	// input := cli.RenderMainScreen()
	LHQuery := []internal.ApiQuery{
		{
			Airline:         "LH",
			StartDate:       "19JUL24",
			EndDate:         "19JUL24",
			DaysOfOperation: "1234567",
			TimeMode:        "LT",
			Origin:          "KRA",
			Destination:     "FRA",
		},
		{
			Airline:         "LH",
			StartDate:       "18JUL24",
			EndDate:         "18JUL24",
			DaysOfOperation: "1234567",
			TimeMode:        "LT",
			Origin:          "KRA",
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
			Origin:          "KRA",
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
			Origin:          "KRA",
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
			Origin:          "KRA",
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

	apiData := internal.GetApiData(queryList)
	flattened := internal.FlattenJSON(apiData)
	println(string(flattened))
	internal.CreateCSVFromResponse(flattened)

	cli.RenderFinal()
	//FRONTEND

}
