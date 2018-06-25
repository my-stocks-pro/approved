package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"github.com/my-stocks-pro/approved/client"
)




func main() {
	router := gin.Default()

	Approved := client.New()

	router.POST("/approved/redis", func(c *gin.Context) {

		data := &client.TypeTest{}
		if c.BindJSON(data) == nil {
			fmt.Println(data)

			Approved.Redis.ListIDS = data.IDlist
			Approved.Redis.Date = data.Date
			Approved.Redis.GetCurrDate()

			go Approved.RunWorker()
		}

	})

	router.Run(":8002")
}
