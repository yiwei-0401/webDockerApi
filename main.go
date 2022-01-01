package main

import (
	"github.com/gin-gonic/gin"
	"webDockerApi/controller"
)

func main()  {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"index": "在线Docker管理网站",
		})
	})

	v1 := r.Group("/docker")
	{
		v1.GET("/ps", controller.Ps)
	}

	v2 := r.Group("/images")
	{
		v2.GET("/list", controller.List)
	}
	r.Run()
}



