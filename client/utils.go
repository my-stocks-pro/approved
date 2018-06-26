package client

import (
	"fmt"
	"strconv"
	"time"
	"strings"
	"github.com/dyninc/qstring"
)

func (a *ApprovedType) GetCurrentIDS() {
	res, err := a.BaseRequest("1")
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, image := range res.Data {
		a.ChanBase <- image.Media_id
	}
}

func (a *ApprovedType) GetCountIDS() {

	res, err := a.BaseRequest("1")
	if err != nil {
		fmt.Println(err)
		return
	}

	a.Base.PerPage = res.Total
}

func (b *TypeBase) GetCurrDate() {
	currDate := time.Now()

	b.TimeStamp = currDate
	b.CurrDate = currDate.Format("2006-01-02")
	b.CurrYear = strconv.Itoa(currDate.Year())
	b.CurrMonth = CheckDate(fmt.Sprintf("%d", currDate.Month()))
	b.CurrDay = CheckDate(strconv.Itoa(currDate.Day()))

	if strings.Compare(b.CurrDate, b.Date) != 0 {
		b.ListIDS = []string{}
		b.Date = b.CurrDate
	}
}

//
//func (a *ApprovedType) StaticURL(page string) string {
//
//	year := strconv.Itoa(a.TimeStamp.Year())
//	month := a.CheckDate(fmt.Sprintf("%d", a.TimeStamp.Month()))
//	day := a.CheckDate(strconv.Itoa(a.TimeStamp.Day()))
//
//	return 	fmt.Sprintf(baseUrl, year, month, "%20", year, month, day, page)
//}

type QueryType struct {
	ID   []string
	View string
}

func MakeQuery(newIDs []string) string {
	query := &QueryType{
		ID:   newIDs,
		View: "full",
	}
	q, errQ := qstring.MarshalString(query)
	if errQ != nil {
		fmt.Println(errQ)
	}

	return q
}

func Contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}
	_, ok := set[item]
	return ok
}

func CheckDate(tmp string) string {
	var res string
	if len(tmp) == 1 {
		res = "0" + tmp
	} else {
		res = tmp
	}
	return res
}
