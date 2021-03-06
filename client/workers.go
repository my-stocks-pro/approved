package client

import (
	"fmt"
	"encoding/json"
)

func (a *ApprovedType) MasterService() {

	count := 1

	for id := range a.ChanBase {
		if count == a.Base.PerPage {
			break
		}

		if Contains(a.Base.ListIDS, id) == false {
			a.ChanFull <- id
			a.Base.ListIDS = append(a.Base.ListIDS, id)
		}

		count++
	}

	a.RedisService()

}

func (a *ApprovedType) WorkerService() {

	count := 0
	for id := range a.ChanFull {

		res, err := a.FullRequest(id)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(res)

		//a.ChanSlack <- res
		a.ChanPSQL <- res

		count++
	}
}


func (a *ApprovedType) RedisService() {

	fmt.Println("RedisService")
	fmt.Println("a.Base.CurrDate", a.Base.CurrDate)
	fmt.Println("a.Base.ListIDS", a.Base.ListIDS)

	b, err := json.Marshal(TypeRedis{a.Base.Date, a.Base.ListIDS})
	if err != nil {
		fmt.Println(err)
	}

	a.POSTURL = a.Config.ApiRedisPostURL
	resp, errPOST := a.POST(b)
	if errPOST != nil {
		fmt.Println(errPOST)
	}

	fmt.Println(resp)
}

func (a *ApprovedType) PSQLService() {
	for dataImage := range a.ChanPSQL {

		fmt.Println("SlackService", dataImage)

		b, err := json.Marshal(*dataImage)
		if err != nil {
			fmt.Println(err)
		}

		a.POSTURL = a.Config.ApiPSQLPostURL
		resp, errPOST := a.POST(b)
		if errPOST != nil {
			fmt.Println(errPOST)
		}

		fmt.Println(resp)
	}
}

func (a *ApprovedType) SlackService() {
	for dataImage := range a.ChanSlack {

		fmt.Println("SlackService", dataImage)

		b, err := json.Marshal(dataImage)
		if err != nil {
			fmt.Println(err)
		}

		resp, errPOST := a.POST(b)
		if errPOST != nil {
			fmt.Println(errPOST)
		}

		fmt.Println(resp)
	}

}