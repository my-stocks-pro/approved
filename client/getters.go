package client

import (
	"time"
	"strconv"
	"fmt"
	"strings"
)

func (r *TypeRedis) GetCurrDate() {
	currDate := time.Now()

	r.TimeStamp = currDate
	r.CurrDate = currDate.Format("2006-01-02")
	r.CurrYear = strconv.Itoa(currDate.Year())
	r.CurrMonth = CheckDate(fmt.Sprintf("%d", currDate.Month()))
	r.CurrDay = CheckDate(strconv.Itoa(currDate.Day()))

	if strings.Compare(r.CurrDate, r.Date) == 0 {
		r.ListIDS = []string{}
	}
}
