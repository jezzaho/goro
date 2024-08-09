package main

import (
	"log"
	"os"
	"time"

	"github.com/jezzaho/goro/cli"
	"github.com/jezzaho/goro/cmd/internal"
	"github.com/joho/godotenv"
	"github.com/pterm/pterm"
)

// "github.com/jezzaho/goro/cmd/internal"
// 	"github.com/jezzaho/goro/cli"

type application struct {
	logFile *os.File
}

func main() {

	godotenv.Load()

	// Create logging to output to file
	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file: ", err)
	}
	app := application{
		logFile: logFile,
	}

	log.SetOutput(app.logFile)

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
		log.Println("Set QueryList to Lufthansa")
	case "OS":
		queryList = OSQuery
		log.Println("Set QueryList to Austrian Airlines")
	case "LX":
		queryList = LXQuery
		log.Println("Set QueryList to Lux Air")
	case "SN":
		queryList = SNQuery
		log.Println("Set QueryList to Brussels Air")
	case "EN":
		queryList = ENQuery
		log.Println("Set QueryList to Air Dolomiti")
	default:
		queryList = []internal.ApiQuery{}
		log.Println("Set QueryList to Default [Empty list]")
	}
	apiAuth := internal.PostForAuth()
	apiData := internal.GetApiData(queryList, apiAuth)
	flattened := internal.FlattenJSON(apiData)
	date := time.Now().Local().Format("20060102")
	season := input.SelectedSeason
	if season == "Własny" {
		season = input.Beggining + "_" + input.Ending
	}
	fileName := input.Line[:2] + "_" + season + "_" + date
	internal.CreateCSVFromResponse(fileName, flattened, input.SeparateDays)
	spinner.Success("ROZKŁAD POBRANY!")
	cli.RenderFinal(fileName)
	//FRONTEND

}
