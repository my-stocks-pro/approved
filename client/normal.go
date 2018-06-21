package client

import (
	"strconv"
	"fmt"
	"github.com/my-stocks-pro/approved/env"
	"encoding/json"
)

func (a *ApprovedType) NormalRUN () {
	page := 1
	for {
		p := strconv.Itoa(page)

		a.NewURL = fmt.Sprintf(baseUrl, a.CurrYear, a.CurrMonth, "%20", a.CurrYear, a.CurrMonth, a.CurrDay, p)

		env.FLAG |= env.COOKIE

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
			close(a.ChanResp)
			close(a.ChanRedis)
			break
		}
		page += 1

		a.ChanResp <- baseResp
		a.ChanRedis <- baseResp
	}

	a.RespDone <- true
}