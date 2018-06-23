package main

import (
	"github.com/my-stocks-pro/approved/env"
	"github.com/my-stocks-pro/approved/client"
	"github.com/jasonlvhit/gocron"

)


func main() {

	env.Load()

	Approved := client.NewClient()

	//if (env.FLAG & env.FIRSTRUN) != 0 {
	//	//	Approved.FirstRUN()
	//	//} else {
	//	//	Approved.NormalRUN()
	//	//}

	go Approved.RunWorker()

	//TODO run in
	Approved.ChanRedis <- Approved.Redis.GET()

	gocron.Every(2).Hours().Do(task)
}
