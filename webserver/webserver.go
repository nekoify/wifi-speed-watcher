package webserver

import (
	"fmt"
	"main/utils"

	"github.com/gin-gonic/gin"
)

func StartWebsever() {
	fmt.Println("starting..")
	r := gin.Default()
	r.GET("/", func(c *gin.Context) { utils.GenerateGraph(c) })
	r.Run(":8080")
}
