package main

import (
	"github.com/jezzaho/goro/cli"
	"github.com/jezzaho/goro/cmd/internal"
	"github.com/joho/godotenv"
	"github.com/pterm/pterm"
)

// "github.com/jezzaho/goro/cmd/internal"
// 	"github.com/jezzaho/goro/cli"

func main() {
	godotenv.Load("somerandomfile")
	// input := cli.CliInput{
	// 	Line:      "LH",
	// 	Beggining: "19JUL24",
	// 	Ending:    "19JUL24",
	// }
	input := cli.RenderMainScreen()
	spinner, _ := pterm.DefaultSpinner.Start("Pobieranie rozkładu.")
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
			StartDate:       input.Beggining,
			EndDate:         input.Ending,
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
	ENQuery := []internal.ApiQuery{
		{
			Airline:         "EN",
			StartDate:       input.Beggining,
			EndDate:         input.Ending,
			DaysOfOperation: "1234567",
			TimeMode:        "LT",
			Origin:          "KRK",
			Destination:     "MUC",
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
	case "EN":
		queryList = ENQuery
	default:
		queryList = []internal.ApiQuery{}
	}
	apiAuth := internal.PostForAuth()
	apiData := internal.GetApiData(queryList, apiAuth)
	flattened := internal.FlattenJSON(apiData)
	internal.CreateCSVFromResponse(flattened, input.SeparateDays)
	spinner.Success("ROZKŁAD POBRANY!")
	cli.RenderFinal()
	//FRONTEND

}
