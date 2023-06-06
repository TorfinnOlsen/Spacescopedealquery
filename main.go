package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

const BaseURL = "https://api.spacescope.io/v2/deals/deal_count"

func main() {
	APIKey := os.Getenv("SPACESCOPE_API_KEY") // Retrieve the API key from an environment variable

	startDate := time.Now().AddDate(0, 0, -7).Format("2006-01-02")
	endDate := time.Now().Format("2006-01-02")

	// Construct the request URL with query parameters
	u, err := url.Parse(BaseURL)
	if err != nil {
		log.Fatal(err)
	}

	q := u.Query()
	q.Set("start_date", startDate)
	q.Set("end_date", endDate)
	u.RawQuery = q.Encode()

	// Create a new request using http
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// Add authorization header to the req
	req.Header.Add("Authorization", "Bearer "+APIKey)

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}
