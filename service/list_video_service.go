package service

import (
	"giligili/model"
	"giligili/serializer"
)

// ListVideoService 视频列表展示
type ListVideoService struct {
}

// List 视频列表展示
func (service *ListVideoService) List() serializer.Response {
	var videos []model.Video

	err := model.DB.Find(&videos).Error
	if err != nil {
		return serializer.Response{
			Status: 404,
			Msg:    "视频不存在",
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildVideos(videos),
	}
}
