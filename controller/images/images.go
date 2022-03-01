package images

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
)

type imagesController struct {
	gin.Context
}

func List(c *gin.Context) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	imgs, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		panic(err)
	}
	c.JSON(200, imgs)
}

/*func CreateImage(c *gin.Context) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err.Error())
	}
	res, err := cli.ImageBuild()
	v := 1
	c.JSON(200, v)
}*/
