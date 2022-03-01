package docker

import (
	"context"
	"github.com/docker/docker/api/types"

	//docker "github.com/docker/docker/client"
	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
)

type dockerController struct {
	gin.Context
}

func Ps(c *gin.Context) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	container, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}
	c.JSON(200, container)
}

//运行一个容器
func Run(c *gin.Context) {

}
