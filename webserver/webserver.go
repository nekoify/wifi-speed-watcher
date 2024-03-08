package webserver

import (
	"fmt"
	"main/utils"

	"github.com/gin-gonic/gin"
)

func StartWebsever() {
	fmt.Println("Running dashboard at http://localhost:8080")
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.GET("/", func(c *gin.Context) { utils.GenerateGraph(c) })
	r.Run(":8080")
}
