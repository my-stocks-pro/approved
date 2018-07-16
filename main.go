package main

import (
	"github.com/my-stocks-pro/approved-service/client"
	"github.com/gin-gonic/gin"
	"fmt"
)

func main() {
	router := gin.Default()

	Approved := client.New()

	router.POST("/approved/redis", func(c *gin.Context) {

		data := &client.TypeRedis{}
		if c.BindJSON(data) == nil {
			fmt.Println("Post fron scheduler -> ", data)

			go Approved.MasterService()
			go Approved.WorkerService()
			go Approved.PSQLService()

			Approved.Base.ListIDS = data.ListIDS
			Approved.Base.Date = data.Date

			Approved.Base.GetCurrDate()
			Approved.GetCountIDS()
			Approved.GetCurrentIDS()

		}

	})

	router.Run(":8002")
}
