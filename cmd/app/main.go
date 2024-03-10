package main

import (
	"fmt"

	"github.com/jezzaho/goro/cmd/internal"
)

func main() {

	// POST REQUEST BUILD
	apiData := internal.GetApiData()
	fmt.Println(apiData)
}
