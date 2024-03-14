package internal

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

// ERROR MESSAGES

type ErrorMessage struct {
	Text  string `json:"text"`
	Level string `json:"level"`
}
type ErrorResponse struct {
	HttpStatus        uint32         `json:"httpStatus"`
	Message           []ErrorMessage `json:"messages"`
	TechnicalMessages []struct {
		Text string `json:"text"`
	} `json:"technicalMessage"`
}

// FLIGHT RESPONSE

type Leg struct {
	SequenceNumber                   int    `json:"sequenceNumber"`
	Origin                           string `json:"origin"`
	Destination                      string `json:"destination"`
	ServiceType                      string `json:"serviceType"`
	AircraftOwner                    string `json:"aircraftOwner"`
	AircraftType                     string `json:"aircraftType"`
	AircraftConfigurationVersion     string `json:"aircraftConfigurationVersion"`
	Registration                     string `json:"registration"`
	Op                               bool   `json:"op"`
	AircraftDepartureTimeUTC         int64  `json:"aircraftDepartureTimeUTC"`
	AircraftDepartureTimeDateDiffUTC int64  `json:"aircraftDepartureTimeDateDiffUTC"`
	AircraftDepartureTimeLT          int64  `json:"aircraftDepartureTimeLT"`
	AircraftDepartureTimeDateDiffLT  int64  `json:"aircraftDepartureTimeDateDiffLT"`
	AircraftDepartureTimeVariation   int64  `json:"aircraftDepartureTimeVariation"`
	AircraftArrivalTimeUTC           int64  `json:"aircraftArrivalTimeUTC"`
	AircraftArrivalTimeDateDiffUTC   int64  `json:"aircraftArrivalTimeDateDiffUTC"`
	AircraftArrivalTimeLT            int64  `json:"aircraftArrivalTimeLT"`
	AircraftArrivalTimeDateDiffLT    int64  `json:"aircraftArrivalTimeDateDiffLT"`
	AircraftArrivalTimeVariation     int64  `json:"aircraftArrivalTimeVariation"`
}

type DataElement struct {
	StartLegSequenceNumber int    `json:"startLegSequenceNumber"`
	EndLegSequenceNumber   int    `json:"endLegSequenceNumber"`
	ID                     int    `json:"id"`
	Value                  string `json:"value"`
}
type PeriodOfOperation struct {
	StartDate       string `json:"startDate"`
	EndDate         string `json:"endDate"`
	DaysOfOperation string `json:"daysOfOperation"`
}

type FlightResponse struct {
	Airline              string            `json:"airline"`
	FlightNumber         int               `json:"flightNumber"`
	Suffix               string            `json:"suffix"`
	PeriodOfOperationUTC PeriodOfOperation `json:"periodOfOperationUTC"`
	PeriodOfOperationLT  PeriodOfOperation `json:"periodOfOperationLT"`
	Legs                 []Leg             `json:"legs"`
	DataElements         []DataElement     `json:"dataElements"`
}

// Define interface to handle ErrorResponse and FlightResponse

type ApiResponse interface {
	isApiResponse()
}

// Define isApiResponse() for Flight and Error Response

func (e ErrorResponse) isApiResponse()  {}
func (f FlightResponse) isApiResponse() {}

// Process Response

func CreateCSVFromResponse(jsonData []byte) {
	var errResponse ErrorResponse
	err := json.Unmarshal(jsonData, &errResponse)
	if err != nil {
		fmt.Errorf(err.Error())
	}
	var flightResponses []FlightResponse
	err = json.Unmarshal(jsonData, &flightResponses)
	if err != nil {
		fmt.Errorf(err.Error())
	}
	// Open CSV file for writing
	file, err := os.Create("output.csv")
	if err != nil {
		fmt.Println("Error creating CSV file:", err)
		return
	}
	defer file.Close()

	// Create a CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write CSV header
	header := []string{"Z", "Do", "Linia", "Numer", "Odlot", "Przylot", "Od", "Do", "Dni", "Samolot", "Operator", "Typ"}
	err = writer.Write(header)
	if err != nil {
		fmt.Println("Error writing CSV header:", err)
		return
	}

	// Write data to CSV
	for _, d := range flightResponses {

		flightNumberWrite := strconv.Itoa(d.FlightNumber)
		startDateWrite := SSIMtoDate(d.PeriodOfOperationLT.StartDate)
		endDateWrite := SSIMtoDate(d.PeriodOfOperationLT.EndDate)
		startTimeWrite := NumberToTime(d.Legs[0].AircraftDepartureTimeLT)
		endTimeWrite := NumberToTime(d.Legs[0].AircraftArrivalTimeLT)
		daysOfOperationWrite := DaysOfOperation(d.PeriodOfOperationLT.DaysOfOperation)
		// "Z", "Do", "Linia", "Numer", "Odlot", "Przylot", "Od", "Do", "Dni", "Samolot", "Operator", "Typ"
		row := []string{d.Legs[0].Origin, d.Legs[0].Destination, d.Airline, flightNumberWrite, startTimeWrite, endTimeWrite, startDateWrite, endDateWrite, daysOfOperationWrite, d.Legs[0].AircraftType, d.Legs[0].AircraftOwner, d.Legs[0].ServiceType}
		err := writer.Write(row)
		if err != nil {
			fmt.Println("Error writing CSV row:", err)
			return
		}
	}

}
