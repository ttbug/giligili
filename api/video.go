package api

import (
	"giligili/service"

	"github.com/gin-gonic/gin"
)

// CreateVideo 视频投稿
func CreateVideo(c *gin.Context) {
	ser := service.CreateVideoService{}
	if err := c.ShouldBind(&ser); err == nil {
		res := ser.Create()
		c.JSON(200, res)
	} else {
		c.JSON(500, ErrorResponse(err))
	}
}

// ShowVideo 视频展示
func ShowVideo(c *gin.Context) {
	ser := service.ShowVideoService{}

	res := ser.Show(c.Param("id"))
	c.JSON(200, res)

}

// ListVideo 视频列表
func ListVideo(c *gin.Context) {
	ser := service.ListVideoService{}
	if err := c.ShouldBind(&ser); err == nil {
		res := ser.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UpdateVideo 更新视频
func UpdateVideo(c *gin.Context) {
	ser := service.UpdateVideoService{}
	if err := c.ShouldBind(&ser); err == nil {
		res := ser.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// DeleteVideo 删除视频
func DeleteVideo(c *gin.Context) {
	ser := service.DeleteService{}
	res := ser.Delete(c.Param("id"))
	c.JSON(200, res)

}
