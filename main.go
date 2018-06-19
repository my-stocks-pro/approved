package main

import (
	"fmt"
	"github.com/my-stocks-pro/approved/env"
	"github.com/my-stocks-pro/approved/run"
)





//func baseRequest(page string, curr_date time.Time) TestType {
//	year := strconv.Itoa(curr_date.Year())
//	month := checkLen(fmt.Sprintf("%d", curr_date.Month()))
//	day := checkLen(strconv.Itoa(curr_date.Day()))
//
//	url := fmt.Sprintf(baseUrl, year, month, "%20", year, month, day, page)
//
//	req, err := http.NewRequest("GET", url, nil)
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	cookie := http.Cookie{Name: "session", Value: "s%3AFLsDQ0KkRmbbHJSFijJz_5VxQPCQI7Ol.t5LQWhFeOPA9qV2S0fqa6JBsFB0Rq%2BrxMDPc1URXyHE"}
//	req.AddCookie(&cookie)
//
//	client := http.Client{}
//	resp, errResp := client.Do(req)
//	if errResp != nil {
//		fmt.Println(errResp)
//	}
//
//	body, errReadBody := ioutil.ReadAll(resp.Body)
//	if errReadBody != nil {
//		fmt.Println(errReadBody)
//	}
//
//	tmp := TestType{}
//
//	errUnm := json.Unmarshal(body, &tmp)
//	if errUnm != nil {
//		fmt.Println(errUnm)
//	}
//
//	return tmp
//}




func main() {

	env.Load()

	dataFromRedis := []string{"111", "222", "333"}
	fmt.Println(env.ENV)

	if (env.ENV & env.FIRSTRUN) != 0 {
		run.First()
	} else {
		run.Normal(dataFromRedis)
	}
}
