package client

import (
	"github.com/my-stocks-pro/approved/env"
	"fmt"
	"encoding/json"
	"github.com/my-stocks-pro/approved/psql"
	"github.com/my-stocks-pro/approved/slack"
	"github.com/my-stocks-pro/approved/redis"
)

func (a *ApprovedType) GreateWorkers() {

	go a.workerRedis()

	for i := 0; i < 10; i++ {
		go a.workerRequest()
		//go a.workerPSQL()
		//go a.workerSlack()

	}
}

func (a *ApprovedType) workerRequest() {

	for resp := range a.ChanResp {

		newIDs := []string{}
		for _, image := range resp.Data {
			if Contains(a.RedisResp, image.Media_id) == false {
				newIDs = append(newIDs, image.Media_id)
			}
		}

		a.NewURL = a.ApiURL + MakeQuery(newIDs)

		env.FLAG |= env.HEADER

		res, err := a.NewRequest()
		if err != nil {
			fmt.Println(err)
		}

		imageData := DataImageType{}
		errUnm := json.Unmarshal(res, &imageData)
		if errUnm != nil {
			fmt.Println(errUnm)
		}

		//a.ChanPSQL <- imageData
		//a.ChanSlack <- imageData
	}

	//close(a.ChanPSQL)
	//close(a.ChanSlack)
}

func (a *ApprovedType) workerPSQL() {
	for resp := range a.ChanPSQL {
		d, e := json.Marshal(resp)
		if e != nil {
			panic(e)
		}
		psql.Post(d)
	}
}

func (a *ApprovedType) workerRedis() {
	dataToRedis := []string{}

	for resp := range a.ChanRedis {
		for _, image := range resp.Data {
			dataToRedis = append(dataToRedis, image.Media_id)
		}
	}

	d, e := json.Marshal(dataToRedis)
	if e != nil {
		panic(e)
	}

	go redis.Post(d, a.RedisDone)
}

func (a *ApprovedType) workerSlack() {
	for resp := range a.ChanSlack {
		d, e := json.Marshal(resp)
		if e != nil {
			panic(e)
		}
		slack.Post(d)
	}
}
