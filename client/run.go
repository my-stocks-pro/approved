package client

import (
	"strconv"
	"fmt"
	"encoding/json"
)

type TypeTest struct {
	Date string
	IDlist []string
}


func (a *ApprovedType) RunWorker () {

	page := 1
	for {
		p := strconv.Itoa(page)

		a.Config.NewURL = fmt.Sprintf(a.Config.BaseURL,
			a.Redis.CurrYear, a.Redis.CurrMonth, "%20", a.Redis.CurrYear, a.Redis.CurrMonth, a.Redis.CurrDay, p)

		res, err := a.NewRequest()
		if err != nil {
			fmt.Println(err)
		}

		baseResp := BaseRespType{}
		errUnm := json.Unmarshal(res, &baseResp)
		if errUnm != nil {
			fmt.Println(errUnm)
		}

		if len(baseResp.Data) == 0 {
			break
		}

		a.BaseChan <- &baseResp

		page += 1
	}
}

