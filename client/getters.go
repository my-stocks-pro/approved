package client

import (
	"time"
	"strconv"
	"fmt"
)

func (r *TypeRedis) GET() []string {

	currDate := time.Now()

	r.TimeStamp = currDate
	r.CurrDate = currDate.Format("2006-01-02")
	r.CurrYear = strconv.Itoa(currDate.Year())
	r.CurrMonth = CheckDate(fmt.Sprintf("%d", currDate.Month()))
	r.CurrDay = CheckDate(strconv.Itoa(currDate.Day()))

	//TODO get data from Redis (api call to server)
	//redis.GET(currDateStr)
	dataFromRedis := []string{"111", "222", "333"}

	return dataFromRedis
}
