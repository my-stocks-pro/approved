package client

import (
	"strconv"
	"fmt"
	"github.com/my-stocks-pro/approved/env"
	"encoding/json"
)

func (a *ApprovedType) RunWorker () {

	for dataFromRedis := range a.ChanRedis {
		a.GetImageIDs(dataFromRedis)
	}

}


func (a *ApprovedType) GetImageIDs (currIDs []string) {

	page := 1
	for {
		p := strconv.Itoa(page)

		a.NewURL = fmt.Sprintf(baseUrl,
			a.Redis.CurrYear, a.Redis.CurrMonth, "%20", a.Redis.CurrYear, a.Redis.CurrMonth, a.Redis.CurrDay, p)

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
			break
		}

		page += 1
	}

}


//func (a *ApprovedType) NormalRUN () {
//
//	page := 1
//	for {
//		p := strconv.Itoa(page)
//
//		a.NewURL = fmt.Sprintf(baseUrl, a.CurrYear, a.CurrMonth, "%20", a.CurrYear, a.CurrMonth, a.CurrDay, p)
//
//		env.FLAG |= env.COOKIE
//
//		res, err := a.NewRequest()
//		if err != nil {
//			fmt.Println(err)
//		}
//
//		baseResp := BaseRespType{}
//		errUnm := json.Unmarshal(res, &baseResp)
//		if errUnm != nil {
//			fmt.Println(errUnm)
//		}
//
//		if len(baseResp.Data) == 0 {
//			a.ChanBaseResp <- TypeChanResp{nil, true}
//			break
//		}
//
//		a.ChanBaseResp <- TypeChanResp{baseResp, false}
//
//		page += 1
//
//	}
//}