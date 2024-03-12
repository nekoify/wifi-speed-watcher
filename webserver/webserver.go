package webserver

import (
	"fmt"
	"main/config"
	"main/utils"

	"github.com/gin-gonic/gin"
)

func StartWebsever() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.GET("/", func(c *gin.Context) { utils.GenerateGraph(c) })
	go r.Run(":" + *config.Port)
	fmt.Println("Running dashboard at http://localhost:" + *config.Port)
}
