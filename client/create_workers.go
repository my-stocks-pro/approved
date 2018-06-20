package client

import (
	"github.com/my-stocks-pro/approved/redis"
	"github.com/my-stocks-pro/approved/env"
	"fmt"
	"encoding/json"
	"github.com/my-stocks-pro/approved/psql"
	"github.com/my-stocks-pro/approved/slack"
)

func (a *ApprovedType) GreateWorkers() {
	for i := 0; i < 10; i++ {
		a.workerRequest()
		a.workerRedis()
		a.workerPSQL()
		a.workerSlack()
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

		fmt.Println(imageData)

		a.ChanRedis <- imageData
		a.ChanPSQL <- imageData
		a.ChanSlack <- imageData
	}
}

func (a *ApprovedType) workerPSQL() {
	for resp := range a.ChanPSQL {
		psql.Post(resp)
	}
}

func (a *ApprovedType) workerRedis() {
	for resp := range a.ChanRedis {
		redis.Post(resp)
	}
}

func (a *ApprovedType) workerSlack() {
	for resp := range a.ChanSlack {
		slack.Post(resp)
	}
}
