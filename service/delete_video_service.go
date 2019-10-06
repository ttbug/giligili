package service

import (
	"giligili/model"
	"giligili/serializer"
)

// DeleteService 视频删除服务
type DeleteService struct {
}

// Delete 视频删除
func (service *DeleteService) Delete(id string) serializer.Response {
	var video model.Video

	err := model.DB.First(&video, id).Error
	if err != nil {
		return serializer.Response{
			Status: 404,
			Msg:    "视频找不到",
			Error:  err.Error(),
		}
	}

	err = model.DB.Delete(&video).Error
	if err != nil {
		return serializer.Response{
			Status: 5001,
			Msg:    "视频删除失败",
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: 200,
		Msg:    "视频删除成功",
	}
}
