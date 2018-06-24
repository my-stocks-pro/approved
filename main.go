package main

import (
	"github.com/my-stocks-pro/approved/client"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"fmt"
)


func main() {
	r := gin.Default()

	Approved := client.NewClient()

	//go Approved.RunWorker()

	Approved.ChanRedis <- Approved.Redis.GET()

	r.POST("/approved", func(c *gin.Context) {
		var data []string

		if c.Bind(&data) == nil {
			fmt.Println(data)
			c.JSON(200, gin.H{"status": "ok"})
		}
	})

	r.Run(":8001")
}
