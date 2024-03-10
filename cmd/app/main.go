package main

import (
	"fmt"

	"github.com/jezzaho/goro/cmd/internal"
)

func main() {
	queryFirst := internal.ApiQuery{
		Airline: "LH",
		StartDate: "19JUL24",
		EndDate: "19JUL24",
		DaysOfOperation: "1234567",
		TimeMode: "LT",
		Origin: "KRK",
		Destination: "FRA",
	}
	querySecond := internal.ApiQuery{
		Airline: "LH",
		StartDate: "22JUL24",
		EndDate: "22JUL24",
		DaysOfOperation: "1234567",
		TimeMode: "LT",
		Origin: "KRK",
		Destination: "MUC",
	}

	queryList := []internal.ApiQuery{queryFirst, querySecond}

	// apiData stores the string from a single api request
	apiData := internal.GetApiData(queryList)
	fmt.Println(apiData)
}
