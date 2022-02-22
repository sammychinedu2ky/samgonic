package service

import "github.com/sammychinedu2ky/samgonic/entity"

type VideoService interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
}
type videoService struct {
	vidoes []entity.Video
}

func New() VideoService {
	return &videoService{}
}

func (service *videoService) Save(video entity.Video) entity.Video {
	service.vidoes = append(service.vidoes, video)
	return video
}

func (service *videoService) FindAll() []entity.Video {
	return service.vidoes
}
