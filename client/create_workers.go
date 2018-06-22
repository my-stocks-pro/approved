package client

import (
	"fmt"
	"encoding/json"
	"github.com/my-stocks-pro/approved/redis"
)
//
//func (a *ApprovedType) GreateWorkers() {
//
//	go a.workerRedis()
//
//	//for i := 0; i < 10; i++ {
//	//	go a.workerRequest()
//		//go a.workerPSQL()
//		//go a.workerSlack()
//
//	//}
//}

func (a *ApprovedType) workerRequest() {

//	for resp := range a.ChanResp {
//
//		newIDs := []string{}
//		for _, image := range resp.Data {
//			if Contains(a.RedisResp, image.Media_id) == false {
//				newIDs = append(newIDs, image.Media_id)
//			}
//		}
//
//		a.NewURL = a.ApiURL + MakeQuery(newIDs)
//
//		env.FLAG |= env.HEADER
//
//		res, err := a.NewRequest()
//		if err != nil {
//			fmt.Println(err)
//		}
//
//		imageData := DataImageType{}
//		errUnm := json.Unmarshal(res, &imageData)
//		if errUnm != nil {
//			fmt.Println(errUnm)
//		}
//
//		//done := <- a.RespDone
//		//if done == true {
//		//	fmt.Println("close(a.ChanResp)")
//		//	close(a.ChanResp)
//		//	a.RespDone <- true
//		//}
//
//		//a.ChanPSQL <- imageData
//		//a.ChanSlack <- imageData
//	}
//
//	//close(a.ChanPSQL)
//	//close(a.ChanSlack)
//}
//
//func (a *ApprovedType) workerPSQL() {
//	for resp := range a.ChanPSQL {
//		d, e := json.Marshal(resp)
//		if e != nil {
//			panic(e)
//		}
//		psql.Post(d)
//	}
}

func (a *ApprovedType) workerRedis() {
	newIDs := []string{}
	currIDs := []string{}

	listIDs, err := a.Redis.GET()
	if err != nil {
		panic(err)
	}

	for resp := range a.ChanBaseResp {
		for _, image := range resp.Res.Data {
			if Contains(listIDs, image.Media_id) == false {
				newIDs = append(newIDs, image.Media_id)
			}
		}
		if resp.Done == true {

			fmt.Println(newIDs)

			if len(newIDs) == 0 {
				return
			}

			currIDs = append(currIDs, newIDs...)

			data, errMarshal := json.Marshal(newIDs)
			if errMarshal != nil {
				panic(errMarshal)
			}

			_, err := redis.POST(data)
			if err != nil {
				fmt.Println(err)
			}

			close(a.ChanBaseResp)
			return
		}
	}

}

func (a *ApprovedType) workerSlack() {
	//for resp := range a.ChanSlack {
	//	d, e := json.Marshal(resp)
	//	if e != nil {
	//		panic(e)
	//	}
	//	slack.Post(d)
	//}
}
