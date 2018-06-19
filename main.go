package main

import (
	"fmt"
	"time"
	"net/http"
	"io/ioutil"
	"strconv"
	"encoding/json"
	"github.com/my-stocks-pro/approved/redis"
	"github.com/my-stocks-pro/approved/env"
)

const (
	token  = "v2/NjYxMWYtZjk4NjItNWI0ZDEtYjc3ODktNmIyYWEtZWU3NDUvMTEyNDM5ODQ1L2N1c3RvbWVyLzMveEZSaF85amREWHFaOF" +
	"ZuNjZfcjZVaUhWYzlYZmhVVl9kZURtU1EybHBSbDJ0eDFxUXh5czEzUXdRdkV6Wjl6dFBYcXpGbHk4RWhWZUpxU1U4TlFoWll3RjF1YkZKMFhs" +
	"S0FndDQtSTY5Y0k0TV9nYy0yVFEzLXdzeC02TXdXMlloVFQyQ2kwZzBjZmtsVmNNVE5OUjZCSHdNY1kzSUQ4SW1CZHlwT1dYNFQ5enNIUkFPdUh" +
	"3VElPRmxaZ214a003dDc1UHE0Tmpta0tIN3ZYM3g4V2xlUQ"

	baseUrl = "https://submit.shutterstock.com/api/catalog_manager/media_types/all/items?filter_type=date_uploaded&filter_value=%s-%s-01%s%s-%s-%s&page_number=%s&per_page=100&sort=popular"
)

var Approved *ApprovedType


func baseRequest(page string, curr_date time.Time) TestType {
	year := strconv.Itoa(curr_date.Year())
	month := checkLen(fmt.Sprintf("%d", curr_date.Month()))
	day := checkLen(strconv.Itoa(curr_date.Day()))

	url := fmt.Sprintf(baseUrl, year, month, "%20", year, month, day, page)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}

	cookie := http.Cookie{Name: "session", Value: "s%3AFLsDQ0KkRmbbHJSFijJz_5VxQPCQI7Ol.t5LQWhFeOPA9qV2S0fqa6JBsFB0Rq%2BrxMDPc1URXyHE"}
	req.AddCookie(&cookie)

	client := http.Client{}
	resp, errResp := client.Do(req)
	if errResp != nil {
		fmt.Println(errResp)
	}

	body, errReadBody := ioutil.ReadAll(resp.Body)
	if errReadBody != nil {
		fmt.Println(errReadBody)
	}

	tmp := TestType{}

	errUnm := json.Unmarshal(body, &tmp)
	if errUnm != nil {
		fmt.Println(errUnm)
	}

	return tmp
}

func (a *ApprovedType) New() {
	a = &ApprovedType{
		BaseUrl: baseUrl,
		Token: token,
	}
}

func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

func FullRequest(approved TestType, dataFromRedis []string) {
	newIDs := []string{}

	for _, image := range approved.Data {
		if contains(newIDs, image.Media_id) == false {
			newIDs = append(newIDs, image.Media_id)
		}
	}

	Req(newIDs)

}


func main() {

	env.Load()

	currDate := time.Now()
	currDateStr := currDate.Format("2006-01-02")
	fmt.Println(currDateStr)

	//TODO get data from Redis (api call to server)
	redis.GET(currDateStr)

	dataFromRedis := []string{"111", "222", "333"}
	fmt.Println(env.ENV)

	//Approved.New()

	page := 1
	for {
		p := strconv.Itoa(page)

		res := baseRequest(p, currDate)
		if len(res.Data) == 0 {
			break
		}
		page += 1
		fmt.Println(res.Data)

		FullRequest(res, dataFromRedis)

	}


	fmt.Println()

}
