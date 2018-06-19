package run

import (
	"strconv"
	"fmt"
	"time"
	. "github.com/my-stocks-pro/approved/new"
)

func Normal(dateFromRedis []string) {

	//TODO get data from Redis (api call to server)
	//redis.GET(currDateStr)

	Approved.New()

	page := 1
	for {
		p := strconv.Itoa(page)

		Approved.BaseUrl = fmt.Sprintf(Approved.BaseUrl, p)

		res := Approved.NewRequest()

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