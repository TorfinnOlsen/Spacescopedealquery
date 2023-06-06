package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type Response struct {
	RequestID string `json:"request_id"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Data      []Deal `json:"data"`
}

type Deal struct {
	StatDate string `json:"stat_date"`
	// add other fields if you need them
}

func main() {
	spacescopeAPIKey := os.Getenv("SPACESCOPE_API_KEY")
	if spacescopeAPIKey == "" {
		log.Fatal("SPACESCOPE_API_KEY environment variable is not set")
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.spacescope.io/v2/deals/deal_count", nil)
	if err != nil {
		log.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("start_date", time.Now().AddDate(0, 0, -7).Format("2006-01-02"))
	q.Add("end_date", time.Now().Format("2006-01-02"))
	req.URL.RawQuery = q.Encode()

	req.Header.Add("Authorization", "Bearer "+spacescopeAPIKey)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var response Response
	if err := json.Unmarshal(body, &response); err != nil {
		log.Fatal(err)
	}

	prettyJSON, err := json.MarshalIndent(response, "", "    ")
	if err != nil {
		log.Fatal("Failed to generate json", err)
	}

	fmt.Println(string(prettyJSON))
}
