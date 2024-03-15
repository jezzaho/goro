package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Auth struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int32  `json:"expires_in"`
}
type ApiQuery struct {
	Airline         string
	StartDate       string
	EndDate         string
	DaysOfOperation string
	TimeMode        string
	Origin          string
	Destination     string
}

func (a *ApiQuery) Swap() {
	a.Origin, a.Destination = a.Destination, a.Origin
}

func GetApiData(queryList []ApiQuery, apiAuth Auth) []byte {
	queryResult := ""
	
	time.Sleep(2000 * time.Millisecond)
	for _, query := range queryList {
		time.Sleep(2000 * time.Millisecond)
		queryResult += getApiResponse(apiAuth, query)
		// Swap query fields Origin and Destination for full result
		query.Swap()
		// Has to sleep - otherwise QPS is exceeded for Api Call
		time.Sleep(2000 * time.Millisecond)
		queryResult += getApiResponse(apiAuth, query)
	}

	return []byte(queryResult)

}

func getApiResponse(auth Auth, query ApiQuery) string {
	
	client := http.Client{}
	getUrl := "https://api.lufthansa.com/v1/flight-schedules/flightschedules/passenger"

	queryParams := url.Values{}
	queryParams.Add("airlines", query.Airline)
	queryParams.Add("startDate", query.StartDate)
	queryParams.Add("endDate", query.EndDate)
	queryParams.Add("daysOfOperation", query.DaysOfOperation)
	queryParams.Add("timeMode", query.TimeMode)
	queryParams.Add("origin", query.Origin)
	queryParams.Add("destination", query.Destination)

	fullURL := fmt.Sprintf("%s?%s", getUrl, queryParams.Encode())

	// Perform the GET request
	request, err := http.NewRequest("GET", fullURL, nil)
	request.Header.Add("Accept", "application/json")
	authStr := "Bearer " + auth.AccessToken
	request.Header.Add("Authorization", authStr)

	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	defer response.Body.Close()

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return ""
	}

	// Print the response body
	body = bytes.Replace(body, []byte("]["), []byte(","), -1)
	return string(body)
}

func PostForAuth() Auth {

	postString := "https://api.lufthansa.com/v1/oauth/token"

	client := http.Client{}

	form := url.Values{}
	form.Add("client_id", "uwbazeekjaagq3zdayamjp4y3")
	form.Add("client_secret", "EH4mZgk9Xj")
	form.Add("grant_type", "client_credentials")

	req, err := http.NewRequest("POST", postString, strings.NewReader(form.Encode()))
	if err != nil {
		fmt.Println("Error")
		return Auth{}
	}
	req.PostForm = form
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error")
		return Auth{}
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error")
		return Auth{}
	}

	// GET REQUEST BUILDER

	var auth Auth
	err = json.Unmarshal([]byte(body), &auth)
	if err != nil {
		fmt.Println("Error")
		return Auth{}
	}

	return auth
}
