package main

import (
	"github.com/gin-gonic/gin"
	"webDockerApi/controller/docker"
	"webDockerApi/controller/images"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"index": "在线Docker管理接口",
		})
	})

	v1 := r.Group("/docker")
	{
		v1.GET("/ps", docker.Ps)
		v1.GET("/create", docker.Create)
	}

	v2 := r.Group("/images")
	{
		v2.GET("/list", images.List)
	}
	r.Run(":5050")
}
