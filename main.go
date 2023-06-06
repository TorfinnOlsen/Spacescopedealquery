package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type Deal struct {
	StatDate                            string `json:"stat_date"`
	DailyProposedRegularDealCount       int    `json:"daily_proposed_regular_deal_count"`
	DailyProposedVerifiedDealCount      int    `json:"daily_proposed_verified_deal_count"`
	DailyActivatedRegularDealCount      int    `json:"daily_activated_regular_deal_count"`
	DailyActivatedVerifiedDealCount     int    `json:"daily_activated_verified_deal_count"`
	DailySlashedRegularDealCount        int    `json:"daily_slashed_regular_deal_count"`
	DailySlashedVerifiedDealCount       int    `json:"daily_slashed_verified_deal_count"`
	DailyExpiredRegularDealCount        int    `json:"daily_expired_regular_deal_count"`
	DailyExpiredVerifiedDealCount       int    `json:"daily_expired_verified_deal_count"`
	TotalRegularDealCount               int    `json:"total_regular_deal_count"`
	TotalVerifiedDealCount              int    `json:"total_verified_deal_count"`
	ActiveRegularDealCount              int    `json:"active_regular_deal_count"`
	ActiveVerifiedDealCount             int    `json:"active_verified_deal_count"`
}

type Response struct {
	Data []Deal `json:"data"`
}

func main() {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://api.spacescope.io/v2/deals/deal_count", nil)
	query := req.URL.Query()
	query.Add("start_date", time.Now().AddDate(0, 0, -7).Format("2006-01-02"))
	query.Add("end_date", time.Now().Format("2006-01-02"))
	req.URL.RawQuery = query.Encode()
	req.Header.Add("Authorization", "Bearer "+os.Getenv("SPACESCOPE_API_KEY"))
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var dealsResponse Response
	json.Unmarshal(body, &dealsResponse)

	var totalDailyProposedRegularDealCount, totalDailyProposedVerifiedDealCount int
	var totalDailyActivatedRegularDealCount, totalDailyActivatedVerifiedDealCount int
	var totalDailySlashedRegularDealCount, totalDailySlashedVerifiedDealCount int
	var totalDailyExpiredRegularDealCount, totalDailyExpiredVerifiedDealCount int

	for _, deal := range dealsResponse.Data {
		totalDailyProposedRegularDealCount += deal.DailyProposedRegularDealCount
		totalDailyProposedVerifiedDealCount += deal.DailyProposedVerifiedDealCount
		totalDailyActivatedRegularDealCount += deal.DailyActivatedRegularDealCount
		totalDailyActivatedVerifiedDealCount += deal.DailyActivatedVerifiedDealCount
		totalDailySlashedRegularDealCount += deal.DailySlashedRegularDealCount
		totalDailySlashedVerifiedDealCount += deal.DailySlashedVerifiedDealCount
		totalDailyExpiredRegularDealCount += deal.DailyExpiredRegularDealCount
		totalDailyExpiredVerifiedDealCount += deal.DailyExpiredVerifiedDealCount
	}

	fmt.Println("Total proposed regular deals:", totalDailyProposedRegularDealCount)
	fmt.Println("Total proposed verified deals:", totalDailyProposedVerifiedDealCount)
	fmt.Println("Total activated regular deals:", totalDailyActivatedRegularDealCount)
	fmt.Println("Total activated verified dealsThe message appears to have been cut off, here is the remaining part of the code:

```go
	fmt.Println("Total activated verified deals:", totalDailyActivatedVerifiedDealCount)
	fmt.Println("Total slashed regular deals:", totalDailySlashedRegularDealCount)
	fmt.Println("Total slashed verified deals:", totalDailySlashedVerifiedDealCount)
	fmt.Println("Total expired regular deals:", totalDailyExpiredRegularDealCount)
	fmt.Println("Total expired verified deals:", totalDailyExpiredVerifiedDealCount)
}
