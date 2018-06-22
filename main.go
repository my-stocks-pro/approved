package main

import (
	"github.com/my-stocks-pro/approved/env"
	"github.com/my-stocks-pro/approved/client"

)


func main() {

	env.Load()

	Approved := client.NewClient()

	if (env.FLAG & env.FIRSTRUN) != 0 {
		Approved.FirstRUN()
	} else {
		Approved.NormalRUN()
	}

	<-Approved.RespDone
	close(Approved.ChanResp)
	<-Approved.RedisDone
	close(Approved.RedisDone)

	//<- Approved.PSQLDone
	//<- Approved.SlackDone
}
