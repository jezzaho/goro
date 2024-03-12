package main

import (
	// "github.com/jezzaho/goro/cmd/internal"
	"github.com/jezzaho/goro/cli"
)

func main() {
	// PART BACKEND
	// queryFirst := internal.ApiQuery{
	// 	Airline:         "LH",
	// 	StartDate:       "19JUL24",
	// 	EndDate:         "19JUL24",
	// 	DaysOfOperation: "1234567",
	// 	TimeMode:        "LT",
	// 	Origin:          "KRK",
	// 	Destination:     "FRA",
	// }
	// querySecond := internal.ApiQuery{
	// 	Airline:         "LH",
	// 	StartDate:       "22JUL24",
	// 	EndDate:         "22JUL24",
	// 	DaysOfOperation: "1234567",
	// 	TimeMode:        "LT",
	// 	Origin:          "KRK",
	// 	Destination:     "MUC",
	// }

	// queryList := []internal.ApiQuery{queryFirst, querySecond}

	// // apiData stores the string from a single api request
	// // It cannot take querylist and process them all because we need to merga data into csv 1b1
	// var apiData []byte
	// apiData = internal.GetApiData(queryList)
	// flattened := internal.FlattenJSON(apiData)
	// println(string(flattened))
	// internal.CreateCSVFromResponse(flattened)

	//FRONTEND

	cli.InitalizeApp()
}
