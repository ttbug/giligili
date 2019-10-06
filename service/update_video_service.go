package service

import (
	"giligili/model"
	"giligili/serializer"
)

// UpdateVideoService 更新视频服务
type UpdateVideoService struct {
	Title string `form:"title" json:"title"`
	Info  string `form:"info" json:"info"`
}

// Update 更新视频
func (service *UpdateVideoService) Update(id string) serializer.Response {
	var video model.Video
	err := model.DB.First(&video, id).Error
	if err != nil {
		return serializer.Response{
			Status: 404,
			Msg:    "找不到视频",
			Error:  err.Error(),
		}
	}

	video.Title = service.Title
	video.Info = service.Info
	err = model.DB.Save(&video).Error
	if err != nil {
		return serializer.Response{
			Status: 5001,
			Msg:    "视频更新失败",
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: 200,
		Msg:    "视频更新成功",
	}
}
