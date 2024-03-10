package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type Auth struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int32  `json:"expires_in"`
}

func GetApiData() string {
	return getApiResponse(postForAuth())
}

func getApiResponse(auth Auth) string {

	client := http.Client{}
	getUrl := "https://api.lufthansa.com/v1/flight-schedules/flightschedules/passenger"

	queryParams := url.Values{}
	queryParams.Add("airlines", "LH")
	queryParams.Add("startDate", "19JUL24")
	queryParams.Add("endDate", "25JUL24")
	queryParams.Add("daysOfOperation", "1234567")
	queryParams.Add("timeMode", "LT")
	queryParams.Add("origin", "KRK")
	queryParams.Add("destination", "FRA")

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
	return string(body)
}

func postForAuth() Auth {

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
